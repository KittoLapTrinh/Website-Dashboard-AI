package automations

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"dashboard/backend/internal/blockchain"
)

// StartKeyValueUpdates khởi chạy goroutine để tự động cập nhật dữ liệu key-value cho Total Viewers.
func StartKeyValueUpdates(instance *blockchain.KeyValueContract) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		randomFilter := []string{"1D", "1W", "1M", "6M", "1Y"}[rand.Intn(5)]

		newViewers := big.NewInt(int64(10000 + rand.Intn(4000000)))
		blockchain.TxQueue <- blockchain.TxRequest{
			ServiceName: "KeyValue-Viewers-" + randomFilter,
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				fmt.Printf("[KeyValue] Queueing request to update viewers for %s\n", randomFilter)
				return instance.UpdateSingleValueData(auth, "viewers_"+randomFilter, newViewers)
			},
		}

		newReturn := big.NewInt(int64(rand.Intn(500) - 100))
		blockchain.TxQueue <- blockchain.TxRequest{
			ServiceName: "KeyValue-Return-" + randomFilter,
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				fmt.Printf("[KeyValue] Queueing request to update return for %s\n", randomFilter)
				return instance.UpdateReturnData(auth, "return_"+randomFilter, newReturn)
			},
		}
	}
}