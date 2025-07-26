// File: internal/automations/timeseries.go
package automations

import (
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"dashboard/backend/internal/blockchain"
)

// StartTimeSeriesUpdates khởi chạy goroutine để tự động cập nhật tất cả dữ liệu chuỗi thời gian.
func StartTimeSeriesUpdates(instance *blockchain.TimeSeriesContract) {
	// Chu kỳ cập nhật chính, mỗi chu kỳ sẽ tạo ra một "đoạn" sóng mới cho Heartbeat Monitor
	ticker := time.NewTicker(5 * time.Second)

	var lastProfileValue int64 = 250000
	var analyticsBaseValue float64 = 50.0
	var lastWeekly, lastMonthly, lastYearly int64 = -1, -1, -1

	for range ticker.C {
		// =====================================================================
		// === 1. CẬP NHẬT CHO HEARTBEAT MONITOR (VỚI LOGIC SÓNG MỚI) ===
		// =====================================================================

		// --- A. Temperature: Vẫn giữ logic biến động nhẹ, push 1 điểm duy nhất ---
		tempChange := (rand.Float64() - 0.5) * 0.4
		newTemperature := int64(36.8 + tempChange)
		queuePushRequest(instance, "temperature", newTemperature, 0)
		
		// --- B & C. Blood Status & Heart Rate: Tạo ra một "đoạn" sóng mới ---
		// Chạy việc tạo sóng trong một goroutine riêng để không làm block vòng lặp chính.
		go generateEcgWaveform(instance, "bloodStatus", 15, 105, 15) // 15 điểm, giá trị nền 105, biên độ 15
		go generateEcgWaveform(instance, "heartRate", 20, 80, 25)  // 20 điểm, giá trị nền 80, biên độ 25 (biến động nhiều hơn)


		// =====================================================================
		// === 2. CẬP NHẬT CHO PROFILE/ANALYTICS (GIỮ NGUYÊN) ===
		// =====================================================================
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

// --- HÀM HELPER MỚI ĐỂ TẠO SÓNG ĐIỆN TÂM ĐỒ ---

// generateEcgWaveform tạo ra một chuỗi các điểm dữ liệu mô phỏng một nhịp đập và gửi chúng lên blockchain.
func generateEcgWaveform(instance *blockchain.TimeSeriesContract, key string, numPoints int, baseline int, amplitude int) {
	// Tạo ra một "nhịp đập" chính (QRS complex) ở một vị trí ngẫu nhiên trong chuỗi
	spikePosition := 2 + rand.Intn(numPoints-4) // Đảm bảo spike không ở đầu hoặc cuối

	for i := 0; i < numPoints; i++ {
		var value int64
		// Dùng công thức toán học đơn giản để tạo hình dạng sóng
		if i > 0 && i < spikePosition-1 {
			// Sóng P nhỏ trước nhịp đập chính
			pWave := float64(amplitude/4) * math.Sin(float64(i)*math.Pi/float64(spikePosition))
			value = int64(baseline) + int64(pWave)
		} else if i == spikePosition {
			// Đỉnh của nhịp đập chính (spike)
			value = int64(baseline + amplitude)
		} else if i == spikePosition-1 {
			// Điểm lõm ngay trước spike
			value = int64(baseline - amplitude/2)
		} else {
			// Các điểm nhiễu nền ở các vị trí khác
			noise := rand.Int63n(int64(amplitude/4)) - int64(amplitude/8)
			value = int64(baseline) + noise
		}
		
		// Gửi từng điểm dữ liệu vào hàng đợi
		queuePushRequest(instance, key, value, 0)
		
		// Thêm một khoảng nghỉ rất ngắn giữa mỗi điểm để tạo hiệu ứng real-time
		// Tổng thời gian cho một sóng sẽ là numPoints * sleepDuration
		time.Sleep(100 * time.Millisecond)
	}
}


// --- CÁC HÀM HELPER KHÁC (GIỮ NGUYÊN) ---

func queuePushRequest(instance *blockchain.TimeSeriesContract, key string, value int64, change int64) {
	blockchain.TxQueue <- blockchain.TxRequest{
		ServiceName: "TimeSeries-Push-" + key,
		Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
			// fmt.Printf("[TimeSeries] Queueing PUSH request for %s with value %d\n", key, value)
			return instance.PushDataPoint(auth, key, big.NewInt(value), big.NewInt(change))
		},
	}
}

func queueAnalyticsUpdate(instance *blockchain.TimeSeriesContract, key string, currentValue int64, lastValue *int64) {
	if currentValue != *lastValue {
		*lastValue = currentValue
		queuePushRequest(instance, key, currentValue, 0)
	}
}