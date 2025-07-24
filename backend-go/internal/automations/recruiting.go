package automations

import (
	"fmt"
	"math/big"
	"math/rand"
	"strings" // ✨ THÊM VÀO: Cần để chuyển chuỗi thành chữ thường
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"dashboard/backend/internal/blockchain"
)

// StartRecruitingUpdates khởi chạy goroutine để tự động thêm các công việc tuyển dụng mới.
func StartRecruitingUpdates(instance *blockchain.RecruitingContract) {
	ticker := time.NewTicker(20 * time.Second)
	var jobCounter int64 = 0

	for range ticker.C {
		jobCounter++
		localCounter := jobCounter

		blockchain.TxQueue <- blockchain.TxRequest{
			ServiceName: "Recruiting",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				// --- DỮ LIỆU MẪU ĐÃ MỞ RỘNG ---
				foundations := []string{"Chainlink", "Polygon", "Uniswap", "Aave", "Nvidia", "Tesla", "Google", "Apple"}
				positions := []string{"Smart Contract Engineer", "DeFi Analyst", "Rust Developer", "AI Scientist"}
				fields := []string{"Blockchain", "AI", "IOT", "Web Dev"}
				// ✨ 2. THÊM DANH SÁCH CÁC HÌNH THỨC LÀM VIỆC ✨
				forms := []string{"Remote", "Fulltime"}

				// --- LOGIC CHỌN LỌC DỮ LIỆU ---

				// ✨ 1. CHỌN FOUNDATION TRƯỚC ĐỂ LẤY ICON TƯƠNG ỨNG ✨
				// Chọn một foundation gốc từ danh sách.
				selectedFoundation := foundations[rand.Intn(len(foundations))]

				// Tạo tên foundation đầy đủ để hiển thị, bao gồm cả số đếm.
				fullFoundationName := selectedFoundation + fmt.Sprintf(" %d", localCounter)
				
				// Tạo IconId bằng cách chuyển tên gốc thành chữ thường.
				// Điều này giúp frontend dễ dàng ánh xạ tới component icon (ví dụ: 'chainlink' -> <ChainlinkIcon />).
				iconId := strings.ToLower(selectedFoundation)


				// --- TẠO CẤU TRÚC DỮ LIỆU MỚI ---
				newJob := blockchain.DataStructsJob{
					Foundation: fullFoundationName, // Tên đầy đủ
					Position:   positions[rand.Intn(len(positions))],
					Field:      fields[rand.Intn(len(fields))],
					Salary:     big.NewInt(int64(60000 + rand.Intn(20000))),
					Form:       forms[rand.Intn(len(forms))], // Lấy ngẫu nhiên từ danh sách forms
					Trend:      uint8(rand.Intn(2)),
					IconId:     iconId, // Lấy IconId đã được tạo
				}

				fmt.Printf("[Recruiting] Queueing request to add job: %s (Icon: %s)\n", newJob.Foundation, newJob.IconId)
				return instance.AddJob(auth, newJob)
			},
		}
	}
}