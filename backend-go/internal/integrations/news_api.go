// File: internal/integrations/news_api.go
package integrations

import (
	"log"
)

// Article định nghĩa cấu trúc cho một bài báo lấy về từ API.
type Article struct {
	Title       string
	Description string
	URL         string
}

// FetchBlockchainNewsFromAPI là một hàm mẫu, sẵn sàng để được triển khai.
// Nhiệm vụ của nó sẽ là gọi đến một dịch vụ tin tức (ví dụ: NewsAPI.org, GNews.io)
// để lấy các bài báo mới nhất.
func FetchBlockchainNewsFromAPI(query string) []Article {
	log.Printf("Fetching news from external API for query: %s...", query)

	// --- TODO: TRIỂN KHAI LOGIC GỌI API Ở ĐÂY ---
	// 1. Lấy API key từ biến môi trường.
	// 2. Xây dựng URL request.
	// 3. Dùng package `net/http` và `encoding/json` của Go để gửi request và parse response.
	// 4. Trả về một mảng các Article.

	// Dữ liệu mẫu để trả về trong khi chưa triển khai
	mockArticles := []Article{
		{Title: "Tin tức mẫu 1", Description: "Mô tả cho tin tức mẫu 1.", URL: "#"},
		{Title: "Tin tức mẫu 2", Description: "Mô tả cho tin tức mẫu 2.", URL: "#"},
	}
	
	log.Println("Mock news data returned.")
	return mockArticles
}