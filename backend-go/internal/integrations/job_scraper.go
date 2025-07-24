// File: internal/integrations/job_scraper.go
package integrations

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

// ScrapedJob định nghĩa cấu trúc cho một công việc được quét về.
type ScrapedJob struct {
	Foundation  string
	JobPosition string
	Field       string
	Salary      string
	Form        string
}

// ScrapeCryptoJobsListDeveloper quét trang developer của CryptoJobsList.
func ScrapeCryptoJobsListDeveloper() []ScrapedJob {
	log.Println("Starting to scrape jobs from CryptoJobsList/developer...")
	var jobs []ScrapedJob

	c := colly.NewCollector(
		colly.AllowedDomains("cryptojobslist.com"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	
	q, _ := queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)

	// ✨ SỬA LẠI SELECTOR CHÍNH VÀ CÁC SELECTOR CON ✨
	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		// Bỏ qua các hàng là quảng cáo (có class 'notAJobAd' hoặc chứa text "Ad")
		if e.DOM.HasClass("notAJobAd") || strings.Contains(e.Text, "Ad") {
			return
		}

		// Lấy tất cả các tag để xác định hình thức làm việc
		var form string = "N/A"
		e.ForEach("td.job-tags span.category", func(_ int, el *colly.HTMLElement) {
			text := strings.ToLower(el.Text)
			if text == "remote" {
				form = "Remote"
			} else if text == "full time" {
				form = "Fulltime"
			}
		})

		job := ScrapedJob{
			Foundation:  strings.TrimSpace(e.ChildText("a.job-company-name-text")),
			JobPosition: strings.TrimSpace(e.ChildText("a.job-title-text")),
			Salary:      strings.TrimSpace(e.ChildText("div span.align-middle")),
			Form:        form,
			Field:       "Blockchain", // Tạm thời gán cứng
		}

		if job.Foundation != "" && job.JobPosition != "" {
			jobs = append(jobs, job)
			log.Printf("Found job: %s at %s", job.JobPosition, job.Foundation)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response code:", r.StatusCode, "\nError:", err)
	})
	
	// ✨ SỬA LẠI URL ĐÍCH ✨
	q.AddURL("https://cryptojobslist.com/developer")
	q.Run(c)
	c.Wait()

	log.Printf("Scraping finished. Found %d jobs.", len(jobs))
	return jobs
}