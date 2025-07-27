// File: internal/automations/content.go
package automations

import (
	"log"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"dashboard/backend/internal/blockchain"
	"dashboard/backend/internal/integrations" // Dùng để gọi news_api
)

// StartContentUpdates khởi chạy tất cả các tác vụ tự động liên quan đến việc đăng tải nội dung hằng ngày.
func StartContentUpdates(instance *blockchain.ContentContract) {
	log.Println("Starting Daily Content services...")
	
	// Chạy vòng lặp cho từng loại nội dung ngay khi khởi động, sau đó lặp lại theo chu kỳ.
	// Sử dụng các khoảng thời gian khác nhau để tránh gửi quá nhiều giao dịch cùng lúc.
	go runDailyTask(publishDailyInspiration, instance, 24*time.Hour)
	go runDailyTask(publishDailyResearch, instance, 23*time.Hour) // Chạy sớm hơn 1 tiếng
	go runDailyTask(publishDailyNews, instance, 25*time.Hour)   // Chạy muộn hơn 1 tiếng
}

// runDailyTask là một hàm helper để chạy một tác vụ ngay lập tức và sau đó lặp lại theo chu kỳ.
func runDailyTask(task func(*blockchain.ContentContract), instance *blockchain.ContentContract, period time.Duration) {
	// Chạy lần đầu tiên khi service khởi động
	task(instance)
	
	// Sau đó lặp lại theo chu kỳ đã định
	ticker := time.NewTicker(period)
	for range ticker.C {
		task(instance)
	}
}


// --- CÁC HÀM TÁC VỤ CHO TỪNG LOẠI NỘI DUNG ---

// 1. Chia sẻ trang web đẹp, sáng tạo
func publishDailyInspiration(instance *blockchain.ContentContract) {
	// Dữ liệu mẫu (có thể được mở rộng hoặc đọc từ file JSON/database)
	inspirations := []struct{Title, Desc, URL string}{
		{"Kima Network", "Ví dụ xuất sắc về thiết kế 3D tương tác trong Web3, kết nối Web2 và Web3 một cách liền mạch.", "https://assets-global.website-files.com/6578a3358055a95333f33876/6578a3358055a95333f33926_kima-network-thumb.jpg"},
		{"Awwwards", "Nơi tôn vinh các website có thiết kế xuất sắc trên toàn thế giới.", "https://assets-global.website-files.com/5f7b442111596b69c43d939f/6304b413028d7abe5f212f0f_awwwards-og.png"},
		{"Codrops", "Các bài viết kỹ thuật và thử nghiệm sáng tạo về hiệu ứng web.", "https://i.ytimg.com/vi/F3203_2sPuc/maxresdefault.jpg"},
	}
	
	// Chọn ngẫu nhiên một mục để đăng
	item := inspirations[rand.Intn(len(inspirations))]
		
	blockchain.TxQueue <- blockchain.TxRequest{
		ServiceName: "Content-Inspiration",
		Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
			log.Println("Queueing new inspiration content:", item.Title)
			// ContentType trong smart contract: 0 = INSPIRATION
			return instance.AddContent(auth, 0, item.Title, item.Desc, item.URL, "Creative Team")
		},
	}
}

// 2. Đăng tải bài nghiên cứu của nhân viên
func publishDailyResearch(instance *blockchain.ContentContract) {
	researches := []struct{Title, Desc, Author, URL string}{
		{"Phân Tích So Sánh ZK-Rollups", "Đánh giá chi tiết về hiệu năng và chi phí giữa StarkNet và zkSync, hai giải pháp Layer 2 hàng đầu.", "Chuyên gia Blockchain A", "https://uploads-ssl.webflow.com/6426757515a49c5e3a510515/644673966504a732a39a738a_Starknet-vs-zkSync.jpg"},
		{"AI trong Quản Lý Rủi Ro DeFi", "Sử dụng mô hình học máy để dự đoán các biến động thị trường và giảm thiểu rủi ro thanh lý tài sản.", "Chuyên gia AI B", "https://miro.medium.com/v2/resize:fit:1400/0*uCg2wta0t2xK2x-t.png"},
	}
	
	item := researches[rand.Intn(len(researches))]

	blockchain.TxQueue <- blockchain.TxRequest{
		ServiceName: "Content-Research",
		Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
			log.Println("Queueing new research content:", item.Title)
			// ContentType trong smart contract: 1 = RESEARCH
			return instance.AddContent(auth, 1, item.Title, item.Desc, item.URL, item.Author)
		},
	}
}

// 3. Cập nhật tin tức Blockchain & AI
func publishDailyNews(instance *blockchain.ContentContract) {
	// Gọi đến module integrations đã được chuẩn bị
	articles := integrations.FetchBlockchainNewsFromAPI("blockchain vietnam")
	if len(articles) == 0 {
		log.Println("No news articles found to publish.")
		return
	}
	
	article := articles[0]
	
	// TODO: Tích hợp với một dịch vụ AI (như Gemini) để tóm tắt `article.Description`
	// summary := ai_service.Summarize(article.Description)

	blockchain.TxQueue <- blockchain.TxRequest{
		ServiceName: "Content-News",
		Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
			log.Println("Queueing new AI-generated news:", article.Title)
			// ContentType trong smart contract: 2 = NEWS
			// Giả sử có một ảnh thumbnail mặc định cho tin tức
			return instance.AddContent(auth, 2, article.Title, article.Description, "https://images.ctfassets.net/s5gv235f1gc3/643d9a0ddb94322d707d32c021132ec5/1e78f9f0914f6b0f2043f11467a140f7/blockchain-in-vietnam-from-a-legal-perspective.jpg", "AI tổng hợp")
		},
	}
}