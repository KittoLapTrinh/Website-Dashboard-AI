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
	

	var analyticsWeeklyValue float64 = 50.0
	var analyticsMonthlyValue float64 = 60.0
	var analyticsYearlyValue float64 = 70.0

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

		analyticsWeeklyValue += (rand.Float64() - 0.5) * 4.0 // Biến động vừa phải
		if analyticsWeeklyValue > 100 { analyticsWeeklyValue = 100 }
		if analyticsWeeklyValue < 20 { analyticsWeeklyValue = 20 }
		queueAnalyticsPush(instance, "analytics_Weekly", analyticsWeeklyValue)
		
		// Cập nhật cho Monthly
		analyticsMonthlyValue += (rand.Float64() - 0.5) * 2.0 // Biến động ít hơn
		if analyticsMonthlyValue > 90 { analyticsMonthlyValue = 90 }
		if analyticsMonthlyValue < 30 { analyticsMonthlyValue = 30 }
		queueAnalyticsPush(instance, "analytics_Monthly", analyticsMonthlyValue)

		// Cập nhật cho Yearly
		analyticsYearlyValue += (rand.Float64() - 0.5) * 1.0 // Biến động rất ít
		if analyticsYearlyValue > 80 { analyticsYearlyValue = 80 }
		if analyticsYearlyValue < 40 { analyticsYearlyValue = 40 }
		queueAnalyticsPush(instance, "analytics_Yearly", analyticsYearlyValue)
	}
}

// --- CÁC HÀM HELPER ---

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
func queueAnalyticsPush(instance *blockchain.TimeSeriesContract, key string, currentValue float64) {
	// Logic này có thể được cải thiện để lấy giá trị cuối từ contract,
	// nhưng để mô phỏng, chúng ta sẽ tạo một 'change' ngẫu nhiên nhỏ.
	change := (rand.Float64() - 0.5) * 5.0 // Thay đổi ngẫu nhiên +/- 2.5%
	
	queuePushRequest(instance, key, int64(currentValue), int64(change))
}