package blockchain

import (
	"fmt" // ✨ THÊM VÀO: Cần cho việc in hash của giao dịch
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// TxRequest định nghĩa một yêu cầu gửi giao dịch, bao gồm tên dịch vụ và hành động cụ thể.
type TxRequest struct {
	ServiceName string
	Action      func(auth *bind.TransactOpts) (*types.Transaction, error)
}

// TxQueue là hàng đợi chung, được export (viết hoa chữ cái đầu) để các package khác
// như `automations` có thể truy cập và gửi yêu cầu vào.
var TxQueue = make(chan TxRequest, 100)

// TransactionProcessor là goroutine duy nhất xử lý các yêu cầu từ TxQueue một cách tuần tự.
// Điều này giải quyết triệt để vấn đề tranh chấp nonce khi có nhiều goroutine cùng gửi giao dịch.
func TransactionProcessor(createTransactor func() *bind.TransactOpts) {
	log.Println("Transaction processor started...")
	for request := range TxQueue {
		// Tạo một transactor mới cho mỗi yêu cầu để đảm bảo nonce luôn là mới nhất
		auth := createTransactor()
		if auth == nil {
			log.Printf("ERROR: [%s] Failed to create transactor, skipping request.", request.ServiceName)
			continue
		}

		// Thực hiện hành động (gọi hàm smart contract)
		tx, err := request.Action(auth)

		// Ghi log kết quả
		logTransaction(request.ServiceName, tx, err)

		// Thêm một khoảng nghỉ ngắn để tránh spam node RPC
		time.Sleep(200 * time.Millisecond)
	}
}

// logTransaction là một hàm helper nội bộ (un-exported, viết thường chữ cái đầu)
// chỉ dùng bên trong package `blockchain` để ghi lại kết quả của một giao dịch.
func logTransaction(serviceName string, tx *types.Transaction, err error) {
	if err != nil {
		log.Printf("ERROR: [%s] Transaction failed: %v\n", serviceName, err)
	} else {
		// Dùng fmt.Printf để có output đẹp hơn và không có timestamp mặc định của log
		fmt.Printf("SUCCESS: [%s] Tx sent! Hash: %s\n", serviceName, tx.Hash().Hex())
	}
}