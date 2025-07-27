// File: internal/integrations/web3career_scraper.go
package integrations

import (
	"log"
	"strings"
	"github.com/gocolly/colly/v2"
)

// Web3CareerScraper tuân thủ "khuôn mẫu" Scraper.
type Web3CareerScraper struct{}

// Scrape thực thi logic quét cho web3.career.
func (s *Web3CareerScraper) Scrape() []ScrapedJob {
	log.Println("--- Starting to scrape jobs from web3.career ---")
	var jobs []ScrapedJob

	c := colly.NewCollector(
		colly.AllowedDomains("web3.career"),
	)

	// Selector cho mỗi hàng trong bảng việc làm
	c.OnHTML("tr.table_row", func(e *colly.HTMLElement) {
		// Bỏ qua hàng header của bảng
		if e.DOM.Find("th").Length() > 0 {
			return
		}

		job := ScrapedJob{
			// CSS Selector cho từng phần tử
			JobPosition: strings.TrimSpace(e.ChildText("h2")),
			Foundation:  strings.TrimSpace(e.ChildText("td:nth-child(1) > div > a")),
			Salary:      strings.TrimSpace(e.ChildText("td:nth-child(3) > span")),
			Form:        inferFormFromLocation(strings.TrimSpace(e.ChildText("td:nth-child(4) > a"))),
			Field:       "Blockchain", // Trang này chuyên về Blockchain
		}
		
		if job.Foundation != "" && job.JobPosition != "" {
			jobs = append(jobs, job)
			log.Printf("Found job: '%s' at '%s'", job.JobPosition, job.Foundation)
		}
	})

	c.Visit("https://web3.career/blockchain-jobs")

	log.Printf("Finished scraping web3.career. Found %d jobs.", len(jobs))
	return jobs
}

// inferFormFromLocation là hàm helper cho web3.career.
func inferFormFromLocation(location string) string {
    lowerLoc := strings.ToLower(location)
    if strings.Contains(lowerLoc, "remote") {
        return "Remote"
    }
    if lowerLoc != "" {
        return "Fulltime"
    }
    return "N/A"
}