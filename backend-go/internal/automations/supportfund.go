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

// StartSupportFundUpdates khởi chạy goroutine để tự động cập nhật dữ liệu cho Support Fund.
func StartSupportFundUpdates(instance *blockchain.SupportFundContract) {
	ticker := time.NewTicker(25 * time.Second)
	for range ticker.C {
		blockchain.TxQueue <- blockchain.TxRequest{
			ServiceName: "SupportFund",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				daysArray := make([]string, 30)
				now := time.Now()
				for i := 0; i < 30; i++ {
					date := now.AddDate(0, 0, -i)
					daysArray[i] = fmt.Sprintf("%s %d", date.Format("Mon"), date.Day())
				}
				randomDayKey := daysArray[rand.Intn(len(daysArray))]
				itemCount := rand.Intn(3) + 2
				var newItems []blockchain.DataStructsFundItem
				for i := 0; i < itemCount; i++ {
					newItems = append(newItems, blockchain.DataStructsFundItem{
						Name:        fmt.Sprintf("Item %d", i+1),
						Price:       big.NewInt(int64(10 + rand.Intn(20))),
						IconId:      []string{"sun", "moon"}[rand.Intn(2)],
						Count:       big.NewInt(int64(rand.Intn(3) + 1)),
						AvatarColor: "bg-gray-600",
					})
				}
				fmt.Printf("[SupportFund] Queueing request for %s with %d items\n", randomDayKey, len(newItems))
				return instance.UpdateFundItemsForDay(auth, randomDayKey, newItems)
			},
		}
	}
}