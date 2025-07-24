// File: internal/automations/timeseries.go
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

// StartTimeSeriesUpdates khởi chạy goroutine để tự động cập nhật dữ liệu chuỗi thời gian.
// ✨ ĐÃ SỬA: Thay thế *blockchain.TimeSeries bằng *blockchain.Blockchain ✨
func StartTimeSeriesUpdates(instance *blockchain.Blockchain) { 
	ticker := time.NewTicker(5 * time.Second)

	const windowSize = 30 
	var bloodData, tempData, heartData []blockchain.DataStructsDataPoint


	var lastProfileValue int64 = 250000

	var analyticsBaseValue float64 = 50.0
	var lastWeekly, lastMonthly, lastYearly int64 = -1, -1, -1

	for range ticker.C {
		// --- Cập nhật Heartbeat Monitor (dùng REPLACE) ---
		now := big.NewInt(time.Now().Unix())
		newBloodPoint := blockchain.DataStructsDataPoint{Timestamp: now, Value: big.NewInt(int64(90 + rand.Intn(30))), Change: big.NewInt(0)}
		newTempPoint := blockchain.DataStructsDataPoint{Timestamp: now, Value: big.NewInt(int64(36 + rand.Intn(3))), Change: big.NewInt(0)}
		newHeartPoint := blockchain.DataStructsDataPoint{Timestamp: now, Value: big.NewInt(int64(70 + rand.Intn(40))), Change: big.NewInt(0)}
		// newSpO2Point := blockchain.DataStructsDataPoint{Timestamp: now, Value: big.NewInt(int64(95 + rand.Intn(6))), Change: big.NewInt(0)}

		bloodData = updateWindow(bloodData, newBloodPoint, windowSize)
		tempData = updateWindow(tempData, newTempPoint, windowSize)
		heartData = updateWindow(heartData, newHeartPoint, windowSize)
		// spO2Data = updateWindow(spO2Data, newSpO2Point, windowSize)

		queuePushRequest(instance, "bloodStatus", int64(90 + rand.Intn(30)), 0)
		queuePushRequest(instance, "temperature", int64(36 + rand.Intn(2)), 0)
		queuePushRequest(instance, "heartRate", int64(70 + rand.Intn(40)), 0)
		// queuePushRequest(instance, "spO2", int64(95 + rand.Intn(6)), 0)

		// --- Cập nhật Profile/Analytics (dùng PUSH) ---
		changePercent := (rand.Float64() - 0.5) * 2.0
		newValueFloat := float64(lastProfileValue) * (1 + changePercent/100)
		lastProfileValue = int64(newValueFloat)
		queuePushRequest(instance, "profileOverview_1D", lastProfileValue, int64(changePercent*10))

		change := (rand.Float64() - 0.5) * 4.0
		analyticsBaseValue += change
		if analyticsBaseValue > 100 { analyticsBaseValue = 100 }
		if analyticsBaseValue < 0 { analyticsBaseValue = 0 }
		weeklyValue := int64(analyticsBaseValue + (rand.Float64()-0.5)*2)
		monthlyValue := int64(analyticsBaseValue + (rand.Float64()-0.5)*4)
		yearlyValue := int64(analyticsBaseValue)
		queueAnalyticsUpdate(instance, "analytics_Weekly", weeklyValue, &lastWeekly)
		queueAnalyticsUpdate(instance, "analytics_Monthly", monthlyValue, &lastMonthly)
		queueAnalyticsUpdate(instance, "analytics_Yearly", yearlyValue, &lastYearly)
	}
}

// --- CÁC HÀM HELPER ---

func updateWindow(data []blockchain.DataStructsDataPoint, newPoint blockchain.DataStructsDataPoint, maxSize int) []blockchain.DataStructsDataPoint {
	if len(data) >= maxSize {
		return append(data[1:], newPoint)
	}
	return append(data, newPoint)
}

func queuePushRequest(instance *blockchain.Blockchain, key string, value int64, change int64) {
	blockchain.TxQueue <- blockchain.TxRequest{
		ServiceName: "TimeSeries-Push-" + key,
		Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
			fmt.Printf("[TimeSeries] Queueing PUSH request for %s with value %d\n", key, value)
			return instance.PushDataPoint(auth, key, big.NewInt(value), big.NewInt(change))
		},
	}
}


// ✨ ĐÃ SỬA: Thay thế *blockchain.TimeSeries bằng *blockchain.Blockchain ✨
func queueReplaceRequest(instance *blockchain.Blockchain, key string, data []blockchain.DataStructsDataPoint) {
	dataToSend := make([]blockchain.DataStructsDataPoint, len(data))
	copy(dataToSend, data)

	blockchain.TxQueue <- blockchain.TxRequest{
		ServiceName: "TimeSeries-Replace-" + key,
		Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
			// ✨ SỬA LẠI THÔNG BÁO LOG (TÙY CHỌN NHƯNG NÊN LÀM) ✨
			fmt.Printf("[TimeSeries] Queueing REPLACE request for %s with %d data points\n", key, len(dataToSend))
			return instance.ReplaceTimeSeries(auth, key, dataToSend)
		},
	}
}


// ✨ ĐÃ SỬA: Thay thế *blockchain.TimeSeries bằng *blockchain.Blockchain ✨
func queueAnalyticsUpdate(instance *blockchain.Blockchain, key string, currentValue int64, lastValue *int64) {
	if currentValue != *lastValue {
		*lastValue = currentValue
		queuePushRequest(instance, key, currentValue, 0)
	}
}