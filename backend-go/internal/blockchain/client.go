package blockchain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// NewBlockchainClient thiết lập và trả về một client Ethereum đã kết nối,
// cùng với một hàm để tạo ra các đối tượng ký giao dịch (transactor).
// Hàm này sẽ là điểm khởi đầu cho mọi tương tác blockchain trong ứng dụng.
func NewBlockchainClient() (*ethclient.Client, func() *bind.TransactOpts, error) {
	log.Println("Setting up blockchain client...")

	// 1. Tải các biến môi trường từ file .env.local
	// Đường dẫn "../.env.local" là vì file thực thi main.go nằm trong cmd/automator
	if err := godotenv.Load("../.env.local"); err != nil {
		return nil, nil, fmt.Errorf("error loading .env.local file: %w", err)
	}
	rpcURL := getEnv("MTD_RPC_URL")
	privateKeyHex := getEnv("MTD_PRIVATE_KEY")

	// 2. Kết nối đến node Ethereum thông qua RPC URL
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Ethereum client at %s: %w", rpcURL, err)
	}
	log.Println("Successfully connected to Ethereum client.")

	// 3. Phân tích private key từ chuỗi hex thành đối tượng ecdsa.PrivateKey
	// Lưu ý: HexToECDSA không xử lý tiền tố "0x", nên chúng ta cần loại bỏ nó.
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing private key: %w", err)
	}

	// 4. Lấy Chain ID từ mạng đang kết nối
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// 5. Tạo và trả về một hàm "closure" để sản xuất transactor.
	// Hàm này "bắt" các biến `privateKey`, `chainID`, và `client` từ phạm vi của nó.
	// Điều này cho phép TransactionProcessor gọi nó mà không cần truyền nhiều tham số.
	createTransactor := func() *bind.TransactOpts {
		// Lấy nonce mới nhất ngay tại thời điểm được gọi
		fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Printf("ERROR: Failed to get nonce: %v\n", err)
			return nil // Trả về nil nếu có lỗi, TransactionProcessor sẽ bỏ qua yêu cầu
		}

		// Lấy giá gas gợi ý
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Printf("ERROR: Failed to suggest gas price: %v\n", err)
			return nil
		}

		// Tạo đối tượng transactor với chain ID
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			log.Printf("ERROR: Failed to create keyed transactor: %v\n", err)
			return nil
		}

		// Thiết lập các thông số cho giao dịch
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)      // Giao dịch không gửi Ether
		auth.GasLimit = uint64(500000) // Đặt một giới hạn gas hợp lý
		auth.GasPrice = gasPrice

		return auth
	}

	// Trả về client, hàm tạo transactor, và không có lỗi
	return client, createTransactor, nil
}

// getEnv là một hàm helper nội bộ để đọc biến môi trường và báo lỗi nếu không tìm thấy.
func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		// Dùng Fatal vì đây là lỗi cấu hình nghiêm trọng, chương trình không thể tiếp tục
		log.Fatalf("FATAL: Environment variable %s is not set.", key)
	}
	return val
}