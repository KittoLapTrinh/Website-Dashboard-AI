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

// StartServiceStatusUpdates khởi chạy goroutine để tự động cập nhật trạng thái của các dịch vụ.
func StartServiceStatusUpdates(instance *blockchain.ServiceStatusContract) {
	ticker := time.NewTicker(15 * time.Second)
	for range ticker.C {
		blockchain.TxQueue <- blockchain.TxRequest{
			ServiceName: "ServiceStatus",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				serviceIndex := big.NewInt(int64(rand.Intn(5)))
				newStatus := uint8(rand.Intn(3))
				fmt.Printf("[ServiceStatus] Queueing request to update index %d to status %d\n", serviceIndex, newStatus)
				return instance.UpdateServiceStatus(auth, serviceIndex, newStatus)
			},
		}
	}
}