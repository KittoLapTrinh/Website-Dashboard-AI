// File: internal/integrations/job_scraper.go
package integrations

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

// ScrapedJob định nghĩa cấu trúc cho một công việc được quét về.
type ScrapedJob struct {
	Foundation  string
	JobPosition string
	Field       string
	Salary      string // Lấy về dưới dạng chuỗi để xử lý sau
	Form        string
}

// ScrapeCryptoJobsList là một ví dụ cụ thể về cách quét một trang web.
func ScrapeCryptoJobsList() []ScrapedJob {
	log.Println("Starting to scrape jobs from CryptoJobsList...")
	var jobs []ScrapedJob

	c := colly.NewCollector(
		// Giới hạn chỉ quét các trang trong domain này
		colly.AllowedDomains("cryptojobslist.com"),
	)

	// Logic chính: Khi Colly tìm thấy một element HTML khớp với selector này...
	// Selector này có nghĩa là "tìm tất cả các thẻ `<li>` có class `job-card`"
	c.OnHTML("li.job-card", func(e *colly.HTMLElement) {
		// Bên trong mỗi thẻ `li.job-card`, tìm các element con và lấy text của chúng
		job := ScrapedJob{
			Foundation:  strings.TrimSpace(e.ChildText("span.job-title-company")),
			JobPosition: strings.TrimSpace(e.ChildText("h2.job-title-text")),
			// Các trường khác có thể cần selector phức tạp hơn
			Field:       "Blockchain", // Tạm thời gán cứng
			Form:        "Remote", // Tạm thời gán cứng
			Salary:      strings.TrimSpace(e.ChildText("span.job-salary-item")),
		}

		// Chỉ thêm vào danh sách nếu có cả tên công ty và vị trí
		if job.Foundation != "" && job.JobPosition != "" {
			jobs = append(jobs, job)
			log.Printf("Found job: %s at %s", job.JobPosition, job.Foundation)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Bắt đầu quét từ URL này
	c.Visit("https://cryptojobslist.com/blockchain-developer-jobs")

	log.Printf("Scraping finished. Found %d jobs.", len(jobs))
	return jobs
}