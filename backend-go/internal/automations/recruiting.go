package automations

import (
	"log"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"dashboard/backend/internal/blockchain"
	"dashboard/backend/internal/integrations"
)

// StartRecruitingUpdates khởi chạy goroutine để quét và thêm việc làm thật.
func StartRecruitingUpdates(instance *blockchain.RecruitingContract) {
	runScrapingCycle := func() {
		log.Println("Running scraping and updating cycle for ALL fields...")
		
		scrapedJobs := integrations.ScrapeAllJobs()

		if len(scrapedJobs) == 0 {
			log.Println("No new jobs found. Waiting for next cycle.")
			return
		}

		for _, job := range scrapedJobs {
			localJob := job
			
			blockchain.TxQueue <- blockchain.TxRequest{
				ServiceName: "Recruiting-Scraper",
				Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
					// ✨ LOGIC ĐÃ ĐƯỢC ĐƠN GIẢN HÓA ✨

					trend := uint8(rand.Intn(2))
					iconId := strings.ToLower(strings.Split(localJob.Foundation, " ")[0])

					newJob := blockchain.DataStructsJob{
						Foundation:  localJob.Foundation,
						Position:    localJob.JobPosition,
						Field:       localJob.Field,
						Salary:      big.NewInt(0),     // Gán cứng trường số là 0
						Form:        localJob.Form,
						Trend:       trend,
						IconId:      iconId,
						SalaryText:  localJob.Salary, // Lấy thẳng chuỗi lương đã quét được
					}
					
					log.Printf("Queueing REAL job to add: [%s] %s at %s | Salary Text: '%s'", newJob.Field, newJob.Position, newJob.Foundation, newJob.SalaryText)
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