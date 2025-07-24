// File: internal/automations/recruiting.go
package automations

import (
	"log"
	"math/big"
	"strings"
	"time"
	"math/rand"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"dashboard/backend/internal/blockchain"
	"dashboard/backend/internal/integrations" // ✨ 1. IMPORT MODULE SCRAPER ✨
)

// StartRecruitingUpdates khởi chạy goroutine để quét và thêm việc làm thật.
func StartRecruitingUpdates(instance *blockchain.RecruitingContract) {
	runScrapingCycle := func() {
		log.Println("Running scraping and updating cycle...")
		scrapedJobs := integrations.ScrapeTopCV()

		if len(scrapedJobs) == 0 {
			log.Println("No new jobs found. Waiting for next cycle.")
			return
		}

		for _, job := range scrapedJobs {
			localJob := job
			blockchain.TxQueue <- blockchain.TxRequest{
				ServiceName: "Recruiting-Scraper",
				Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
					trend := uint8(rand.Intn(2))
					iconId := strings.ToLower(strings.Split(localJob.Foundation, " ")[0])

					// ✨ THAY ĐỔI LOGIC TẠI ĐÂY ✨
					newJob := blockchain.DataStructsJob{
						Foundation:  localJob.Foundation,
						Position:    localJob.JobPosition,
						Field:       mapField(localJob.JobPosition),
						Salary:      big.NewInt(0), // Gán cứng là 0
						Form:        localJob.Form,
						Trend:       trend,
						IconId:      iconId,
						SalaryText:  localJob.Salary, // Lấy thẳng chuỗi lương đã crawl
					}
					
					log.Printf("Queueing REAL job to add: %s at %s | Salary Text: '%s'", newJob.Position, newJob.Foundation, newJob.SalaryText)
					return instance.AddJob(auth, newJob)
				},
			}
			time.Sleep(1 * time.Second)
		}
	}
	go runScrapingCycle()
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