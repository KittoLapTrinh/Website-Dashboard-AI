// File: internal/integrations/job_scraper.go
package integrations

import (
	"context"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// ScrapedJob định nghĩa cấu trúc cho một công việc được quét về.
type ScrapedJob struct {
	Foundation  string
	JobPosition string
	Field       string
	Salary      string
	Form        string
}

// ScrapeSource định nghĩa một nguồn dữ liệu cần quét.
type ScrapeSource struct {
	Field string
	URL   string
}

// ScrapeAllJobs là hàm chính, điều phối việc quét từ nhiều nguồn trên TopCV.
func ScrapeAllJobs() []ScrapedJob {
	log.Println("--- Starting to scrape jobs from ALL TopCV sources ---")
	
	// ✨ DANH SÁCH CÁC NGUỒN DỮ LIỆU ĐÃ ĐƯỢC CẬP NHẬT CHÍNH XÁC ✨
	sources := []ScrapeSource{
		{Field: "Blockchain", URL: "https://www.topcv.vn/tim-viec-lam-blockchain?sba=1"},
		{Field: "AI", URL: "https://www.topcv.vn/tim-viec-lam-ai?type_keyword=1&sba=1"},
		{Field: "IOT", URL: "https://www.topcv.vn/tim-viec-lam-iot?type_keyword=1&sba=1"},
		{Field: "Web Dev", URL: "https://www.topcv.vn/tim-viec-lam-web-developer?type_keyword=1&sba=1"},
	}

	var allJobs []ScrapedJob
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, source := range sources {
		wg.Add(1)
		
		go func(src ScrapeSource) {
			defer wg.Done()
			
			log.Printf("Scraping jobs for field: %s from %s", src.Field, src.URL)
			jobsFromSource := scrapeSingleSource(src)

			mu.Lock()
			allJobs = append(allJobs, jobsFromSource...)
			mu.Unlock()
		}(source)
	}
	
	wg.Wait()
	
	log.Printf("Finished scraping all sources. Total jobs found: %d", len(allJobs))
	return allJobs
}

// scrapeSingleSource chịu trách nhiệm quét một URL duy nhất.
func scrapeSingleSource(source ScrapeSource) []ScrapedJob {
	var jobs []ScrapedJob

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), // Chạy ẩn để không mở nhiều cửa sổ
		chromedp.Flag("disable-gpu", true),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36`),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
	)
	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAlloc()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate(source.URL),
		chromedp.WaitVisible(`div.job-item-search-result`),
		chromedp.Nodes(`div.job-item-search-result`, &nodes, chromedp.ByQueryAll),
	)

	if err != nil {
		log.Printf("Failed to scrape source '%s' (%s): %v", source.Field, source.URL, err)
		return jobs
	}

	for _, node := range nodes {
		var position, foundation, salaryHTML, location string

		chromedp.Run(ctx,
			chromedp.Text(`h3.title a span`, &position, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(`a.company span.company-name`, &foundation, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.InnerHTML(`label.title-salary`, &salaryHTML, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(`label.address span.city-text`, &location, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		
		if position != "" && foundation != "" {
			re := regexp.MustCompile(`<[^>]*>`)
			cleanSalary := strings.TrimSpace(re.ReplaceAllString(salaryHTML, ""))

			job := ScrapedJob{
				Foundation:  strings.TrimSpace(foundation),
				JobPosition: strings.TrimSpace(position),
				Salary:      cleanSalary,
				Field:       source.Field, // Gán đúng Field từ nguồn
				Form:        inferForm(location),
			}
			jobs = append(jobs, job)
		}
	}
	return jobs
}

// inferForm là hàm helper để suy ra hình thức làm việc từ địa điểm.
func inferForm(location string) string {
    lowerLoc := strings.ToLower(location)
    if strings.Contains(lowerLoc, "remote") || strings.Contains(lowerLoc, "từ xa") {
        return "Remote"
    }
    if lowerLoc != "" {
        return "Fulltime"
    }
    return "N/A"
}