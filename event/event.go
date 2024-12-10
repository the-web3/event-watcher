package event

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/event-watcher/common/bigint"
	"github.com/the-web3/event-watcher/common/tasks"
	"github.com/the-web3/event-watcher/config"
	"github.com/the-web3/event-watcher/database"
	"github.com/the-web3/event-watcher/database/common"
	"github.com/the-web3/event-watcher/event/dapplink"
)

var blocksLimit = 10_000

type EventProcessor struct {
	db          *database.DB
	chainConfig config.ChainConfig

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group

	LatestBlockHeader *common.BlockHeader
}

func NewEventProcessor(db *database.DB, chainConfig config.ChainConfig, shutdown context.CancelCauseFunc) (*EventProcessor, error) {
	LatestBlockHeader, err := db.Blocks.LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	resCtx, resCancel := context.WithCancel(context.Background())

	return &EventProcessor{
		db:             db,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		chainConfig:    chainConfig,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		LatestBlockHeader: LatestBlockHeader,
	}, nil
}

func (ep *EventProcessor) Start() error {
	log.Info("starting bridge processor...")
	tickerL1Worker := time.NewTicker(time.Second * 5)
	ep.tasks.Go(func() error {
		for range tickerL1Worker.C {
			err := ep.processTreasureManagerEvents()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func (ep *EventProcessor) Close() error {
	ep.resourceCancel()
	return ep.tasks.Wait()
}

func (ep *EventProcessor) processTreasureManagerEvents() error {
	lastBlockNumber := big.NewInt(int64(ep.chainConfig.StartingHeight))
	if ep.LatestBlockHeader != nil {
		lastBlockNumber = ep.LatestBlockHeader.Number
	}
	log.Info("Process treasure manager events", "lastBlockNumber", lastBlockNumber)
	latestHeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common.BlockHeader{}).Where("number > ?", lastBlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	if latestHeaderScope == nil {
		return errors.New("latest header scope is nil")
	}
	latestHeader, err := ep.db.Blocks.BlockHeaderWithScope(latestHeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized treasure manager events state: %w", err)
	} else if latestHeader == nil {
		log.Debug("no new state to process event")
		return nil
	}

	fromHeight, toHeight := new(big.Int).Add(lastBlockNumber, bigint.One), latestHeader.Number
	log.Info("handle event block range", "from", fromHeight, "end", toHeight)
	if err := ep.db.Transaction(func(tx *database.DB) error {
		log.Info("scanning for treasure manager events", "fromHeight", fromHeight, "toHeight", toHeight)
		return dapplink.ProcessDepositEvents(tx, ep.chainConfig, fromHeight, toHeight)
	}); err != nil {
		return err
	}
	ep.LatestBlockHeader = latestHeader
	return nil
}
