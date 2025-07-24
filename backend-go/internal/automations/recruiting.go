// File: internal/automations/recruiting.go
package automations

import (
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"dashboard/backend/internal/blockchain"
	"dashboard/backend/internal/integrations" // ✨ 1. IMPORT MODULE SCRAPER ✨
)

// StartRecruitingUpdates khởi chạy goroutine để quét và thêm việc làm thật.
func StartRecruitingUpdates(instance *blockchain.RecruitingContract) {
	// Hàm để chạy logic quét và cập nhật
	runScrapingCycle := func() {
		log.Println("Running scraping and updating cycle...")
		
		// ✨ 2. GỌI HÀM SCRAPER ĐỂ LẤY DỮ LIỆU THẬT ✨
		scrapedJobs := integrations.ScrapeCryptoJobsListDeveloper()

		if len(scrapedJobs) == 0 {
			log.Println("No new jobs found. Waiting for next cycle.")
			return
		}

		// 3. Lặp qua các công việc đã quét được và gửi lên blockchain
		for _, job := range scrapedJobs {
			// Tạo bản sao để closure sử dụng đúng
			localJob := job 
			
			blockchain.TxQueue <- blockchain.TxRequest{
				ServiceName: "Recruiting-Scraper",
				Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
					// Chuyển đổi dữ liệu từ ScrapedJob sang DataStructsJob
					newJob := blockchain.DataStructsJob{
						Foundation:  localJob.Foundation,
						Position:    localJob.JobPosition,
						Field:       mapField(localJob.JobPosition), // Tự động đoán Field
						Salary:      big.NewInt(70000), // Tạm thời gán cứng
						Form:        localJob.Form,
						Trend:       0, // 0 = Up
						IconId:      strings.ToLower(strings.Split(localJob.Foundation, " ")[0]),
					}
					
					log.Printf("Queueing REAL job to add: %s at %s", newJob.Position, newJob.Foundation)
					return instance.AddJob(auth, newJob)
				},
			}
			// Thêm một khoảng nghỉ nhỏ giữa các giao dịch để không làm quá tải hàng đợi
			time.Sleep(1 * time.Second) 
		}
	}

	// Chạy lần đầu tiên ngay khi khởi động
	go runScrapingCycle()

	// ✨ 4. THAY ĐỔI TẦN SUẤT QUÉT ✨
	// Quét lại sau mỗi 6 giờ, không nên quét quá thường xuyên
	ticker := time.NewTicker(6 * time.Hour)
	for range ticker.C {
		runScrapingCycle()
	}
}

// mapField là một hàm helper đơn giản để tự động đoán lĩnh vực dựa trên chức danh.
func mapField(position string) string {
	lowerPos := strings.ToLower(position)
	if strings.Contains(lowerPos, "ai") || strings.Contains(lowerPos, "scientist") {
		return "AI"
	}
	if strings.Contains(lowerPos, "iot") || strings.Contains(lowerPos, "firmware") {
		return "IOT"
	}
	// Mặc định là Blockchain
	return "Blockchain"
}