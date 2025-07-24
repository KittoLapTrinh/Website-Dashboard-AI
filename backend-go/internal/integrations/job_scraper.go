// File: internal/integrations/job_scraper.go
package integrations

import (
	"context"
	"log"
	"strings"
	"time"
	"regexp"
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

// ScrapeTopCV quét trang tìm việc làm blockchain của TopCV với cấu hình đơn giản nhất.
func ScrapeTopCV() []ScrapedJob {
	log.Println("--- Starting to scrape TopCV with SIMPLEST config ---")
	var jobs []ScrapedJob

	// Sử dụng cấu hình mặc định, chỉ tắt headless để quan sát.
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAlloc()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var nodes []*cdp.Node

	// ✨ LOGIC ĐÃ ĐƯỢC SỬA LẠI CHO ĐÚNG CÚ PHÁP ✨
	log.Println("Step 1: Navigating to URL...")
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.topcv.vn/tim-viec-lam-blockchain?sba=1`),
	)
	if err != nil {
		log.Printf("ERROR: Failed to navigate to URL: %v", err)
		return jobs
	}

	log.Println("Step 2: Waiting for 15 seconds to observe...")
	if err := chromedp.Run(ctx, chromedp.Sleep(15*time.Second)); err != nil {
		log.Printf("ERROR: Sleep interrupted: %v", err)
		return jobs
	}
		
	log.Println("Step 3: Attempting to find and extract job nodes...")
	err = chromedp.Run(ctx,
		chromedp.Nodes(`div.job-item-search-result`, &nodes, chromedp.ByQueryAll),
	)

	if err != nil {
		log.Printf("ERROR: Failed to find job nodes. The selector might be wrong or blocked. Error: %v", err)
		return jobs
	}

	if len(nodes) == 0 {
		log.Println("WARNING: No job items found with the specified selector.")
		return jobs
	}
	
	log.Printf("SUCCESS: Found %d job items. Parsing details...", len(nodes))
	
	// Lặp qua từng node và trích xuất thông tin chi tiết
	for _, node := range nodes {
		var position, foundation, salaryHTML, location string // ✨ Đổi tên biến `salary` thành `salaryHTML`

		err := chromedp.Run(ctx,
			chromedp.Text(`h3.title a span`, &position, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(`a.company span.company-name`, &foundation, chromedp.ByQuery, chromedp.FromNode(node)),
			
			// ✨ THAY ĐỔI QUAN TRỌNG NHẤT: LẤY `innerHTML` THAY VÌ `Text` ✨
			chromedp.InnerHTML(`label.title-salary`, &salaryHTML, chromedp.ByQuery, chromedp.FromNode(node)),
			
			chromedp.Text(`label.address span.city-text`, &location, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		
		if err != nil {
			log.Printf("Could not parse all details for a job item, continuing... Error: %v", err)
			continue
		}
		
		if position != "" && foundation != "" {
			// ✨ DỌN DẸP CHUỖI HTML ĐỂ LẤY TEXT SẠCH ✨
			// Dùng regex để xóa tất cả các thẻ HTML (ví dụ: <i ...></i>)
			re := regexp.MustCompile(`<[^>]*>`)
			cleanSalary := re.ReplaceAllString(salaryHTML, "")
			cleanSalary = strings.TrimSpace(cleanSalary) // Xóa khoảng trắng thừa

			job := ScrapedJob{
				Foundation:  strings.TrimSpace(foundation),
				JobPosition: strings.TrimSpace(position),
				Salary:      cleanSalary, // Sử dụng chuỗi đã được dọn dẹp
				Field:       "Blockchain",
				Form:        inferForm(location),
			}
			jobs = append(jobs, job)
			log.Printf("Parsed job: '%s' at '%s' with Salary Text: '%s'", job.JobPosition, job.Foundation, job.Salary)
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