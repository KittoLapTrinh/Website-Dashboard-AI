// File: internal/automations/timeseries.go
package automations

import (
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"dashboard/backend/internal/blockchain"
)

// StartTimeSeriesUpdates khởi chạy goroutine để tự động cập nhật tất cả dữ liệu chuỗi thời gian.
// ✨ ĐÃ SỬA: Thay thế *blockchain.Blockchain bằng *blockchain.TimeSeriesContract ✨
func StartTimeSeriesUpdates(instance *blockchain.TimeSeriesContract) {
	ticker := time.NewTicker(5 * time.Second)

	var lastBloodValue int64 = 105
	var lastHeartRate int64 = 80

	var lastProfileValue int64 = 250000
	var analyticsBaseValue float64 = 50.0
	var lastWeekly, lastMonthly, lastYearly int64 = -1, -1, -1

	for range ticker.C {
		// --- Cập nhật Heartbeat Monitor ---
		tempChange := (rand.Float64() - 0.5) * 0.4
		newTemperature := int64(36.8 + tempChange)
		queuePushRequest(instance, "temperature", newTemperature, 0)

		bloodChange := rand.Int63n(7) - 3
		lastBloodValue += bloodChange
		if lastBloodValue > 120 { lastBloodValue = 118 }
		if lastBloodValue < 90 { lastBloodValue = 92 }
		queuePushRequest(instance, "bloodStatus", lastBloodValue, 0)

		heartRateChange := rand.Int63n(17) - 8
		lastHeartRate += heartRateChange
		if lastHeartRate > 125 { lastHeartRate = 123 }
		if lastHeartRate < 60 { lastHeartRate = 62 }
		queuePushRequest(instance, "heartRate", lastHeartRate, 0)

		// --- Cập nhật Profile/Analytics ---
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

// ✨ ĐÃ SỬA: Thay thế *blockchain.Blockchain bằng *blockchain.TimeSeriesContract ✨
func queuePushRequest(instance *blockchain.TimeSeriesContract, key string, value int64, change int64) {
	blockchain.TxQueue <- blockchain.TxRequest{
		ServiceName: "TimeSeries-Push-" + key,
		Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
			// fmt.Printf("[TimeSeries] Queueing PUSH request for %s with value %d\n", key, value)
			return instance.PushDataPoint(auth, key, big.NewInt(value), big.NewInt(change))
		},
	}
}

// ✨ ĐÃ SỬA: Thay thế *blockchain.Blockchain bằng *blockchain.TimeSeriesContract ✨
func queueAnalyticsUpdate(instance *blockchain.TimeSeriesContract, key string, currentValue int64, lastValue *int64) {
	if currentValue != *lastValue {
		*lastValue = currentValue
		queuePushRequest(instance, key, currentValue, 0)
	}
}