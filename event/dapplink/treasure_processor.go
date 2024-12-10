package dapplink

import (
	"math/big"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/event-watcher/config"
	"github.com/the-web3/event-watcher/database"
	"github.com/the-web3/event-watcher/database/worker"
	"github.com/the-web3/event-watcher/event/contracts"
)

func ProcessDepositEvents(db *database.DB, chainCfg config.ChainConfig, fromHeight, toHeight *big.Int) error {
	txDepositTokens, err := contracts.DepositTokensEvents(chainCfg.Contracts[0], db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	log.Info("detected transaction deposit events", "size", len(txDepositTokens))

	depositTk := make([]worker.DepositTokens, len(txDepositTokens))
	for i := range txDepositTokens {
		depositTk[i] = worker.DepositTokens{
			GUID:         uuid.New(),
			TokenAddress: depositTk[i].TokenAddress,
			Sender:       txDepositTokens[i].Sender,
			Amount:       big.NewInt(0),
			BlockNumber:  big.NewInt(1),
			Timestamp:    txDepositTokens[i].Timestamp,
		}
	}
	if len(depositTk) > 0 {
		if err := db.DepositTokens.StoreDepositTokens(depositTk); err != nil {
			return err
		}
	}
	return nil
}
