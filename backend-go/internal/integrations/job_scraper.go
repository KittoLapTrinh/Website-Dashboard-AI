// File: internal/integrations/job_scraper.go
package integrations

import (
	"context"
	"log"
	"os"
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
	
	sources := []ScrapeSource{
		{Field: "Blockchain", URL: "https://www.topcv.vn/tim-viec-lam-blockchain?sba=1"},
		{Field: "AI", URL: "https://www.topcv.vn/tim-viec-lam-ai?type_keyword=1&sba=1"},
		{Field: "IOT", URL: "https://www.topcv.vn/tim-viec-lam-iot?type_keyword=1&sba=1"},
		{Field: "Web Dev", URL: "https://www.topcv.vn/tim-viec-lam-web-developer?type_keyword=1&sba=1"},
	}

	var allJobs []ScrapedJob
	var wg sync.WaitGroup
	var mu sync.Mutex

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36`),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
	)
	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAlloc()

	browserCtx, cancelBrowser := chromedp.NewContext(allocCtx)
	defer cancelBrowser()

	for _, source := range sources {
		wg.Add(1)
		
		go func(src ScrapeSource) {
			defer wg.Done()
			
			log.Printf("Scraping jobs for field: %s from %s", src.Field, src.URL)
			jobsFromSource := scrapeSingleSource(browserCtx, src)

			mu.Lock()
			if jobsFromSource != nil {
				allJobs = append(allJobs, jobsFromSource...)
			}
			mu.Unlock()
		}(source)

		time.Sleep(2 * time.Second)
	}
	
	wg.Wait()
	
	log.Printf("Finished scraping all sources. Total jobs found: %d", len(allJobs))
	return allJobs
}

// scrapeSingleSource chịu trách nhiệm quét một URL duy nhất trong một tab mới.
func scrapeSingleSource(browserCtx context.Context, source ScrapeSource) []ScrapedJob {
	var jobs []ScrapedJob

	ctx, cancel := chromedp.NewContext(browserCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 90*time.Second)
	defer cancel()

	var nodes []*cdp.Node
	var screenshotBuf []byte

	err := chromedp.Run(ctx,
		chromedp.Navigate(source.URL),
		chromedp.Sleep(8*time.Second),
		chromedp.FullScreenshot(&screenshotBuf, 90),
		chromedp.WaitVisible(`div.job-item-search-result`),
		chromedp.Nodes(`div.job-item-search-result`, &nodes, chromedp.ByQueryAll),
	)
	
	if err != nil {
		fileName := "error_screenshot_" + strings.ReplaceAll(source.Field, " ", "_") + ".png"
		if err_ss := os.WriteFile(fileName, screenshotBuf, 0o644); err_ss != nil {
			log.Printf("ERROR: Failed to save screenshot for %s: %v", source.Field, err_ss)
		} else {
			log.Printf("SUCCESS: Error screenshot for %s saved to %s", source.Field, fileName)
		}
		log.Printf("Failed to scrape source '%s' (%s): %v", source.Field, source.URL, err)
		return nil
	}
	
	// ✨ BẮT ĐẦU PHẦN CODE ĐÃ ĐƯỢC ĐƯA VÀO ĐÚNG CHỖ ✨
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
				Field:       source.Field,
				Form:        inferForm(location),
			}
			jobs = append(jobs, job)
		}
	}
	return jobs
	// ✨ KẾT THÚC PHẦN CODE ĐÃ ĐƯỢC ĐƯA VÀO ĐÚNG CHỖ ✨
}


// ✨ HÀM INFERFORM ĐÃ ĐƯỢC ĐƯA VÀO ĐÚNG CHỖ ✨
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