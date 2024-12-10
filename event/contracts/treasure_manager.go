package contracts

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/event-watcher/bindings"
	"github.com/the-web3/event-watcher/database"
	"github.com/the-web3/event-watcher/database/event"
)

type DepositTokensEvent struct {
	TokenAddress common.Address
	Sender       common.Address
	Amount       *big.Int
	Timestamp    uint64
	Raw          types.Log
}

func DepositTokensEvents(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]DepositTokensEvent, error) {
	treasureManagerAbi, err := bindings.TreasureManagerMetaData.GetAbi()
	if err != nil {
		log.Error("treasure manager meta data", "err", err)
		return nil, err
	}
	treasureManagerFilterer, err := bindings.NewTreasureManagerFilterer(contractAddress, nil)
	if err != nil {
		log.Error("new treasure manager fail", "err", err)
		return nil, err
	}
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress}
	transactionDepositTokenEvents, err := db.ContractEvent.ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		log.Error("contract events with filter fail")
		return nil, err
	}
	log.Info("query contract events success", "event length", len(transactionDepositTokenEvents))
	var txDepositTokens []DepositTokensEvent
	for _, eventItem := range transactionDepositTokenEvents {
		log.Info("start handle deposit token event", "event Signature", eventItem.EventSignature.String(), "Abi Signature", treasureManagerAbi.Events["DepositToken"].ID.String())
		rlpLog := eventItem.RLPLog
		if eventItem.EventSignature.String() == treasureManagerAbi.Events["DepositToken"].ID.String() {
			depositTokenEvent, err := treasureManagerFilterer.ParseDepositToken(*rlpLog)
			if err != nil {
				log.Error("parse deposit event fail", "err", err)
				return nil, err
			}
			log.Info("parse deposit token event success",
				"tokenAddress", depositTokenEvent.TokenAddress.String(),
				"sender", depositTokenEvent.Sender.String(),
				"amount", depositTokenEvent.Amount.String())

			txDepositTokenItem := DepositTokensEvent{
				TokenAddress: depositTokenEvent.TokenAddress,
				Sender:       depositTokenEvent.Sender,
				Amount:       depositTokenEvent.Amount,
				Timestamp:    uint64(time.Now().Unix()),
			}
			txDepositTokens = append(txDepositTokens, txDepositTokenItem)
		}
	}
	return txDepositTokens, nil
}
