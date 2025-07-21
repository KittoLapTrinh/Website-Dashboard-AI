package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"time"
	"sync"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// 1. Định nghĩa một "Yêu Cầu Giao Dịch"
type TxRequest struct {
	ServiceName string
	Action      func(auth *bind.TransactOpts) (*types.Transaction, error)
}

// 2. Tạo một "Hàng Đợi" (Channel) chung
var txQueue = make(chan TxRequest, 100) // Buffer 100 yêu cầu

var txMutex = &sync.Mutex{}

// Helper để đọc địa chỉ từ .env
func getEnvAddress(key string) common.Address {
	return common.HexToAddress(getEnv(key))
}
func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("FATAL: Environment variable %s is not set.", key)
	}
	return val
}

func main() {
	// --- SETUP ---
	if err := godotenv.Load("../.env.local"); err != nil {
		log.Fatalf("Error loading .env.local file: %v", err)
	}
	rpcURL := getEnv("MTD_RPC_URL")
	privateKeyHex := getEnv("MTD_PRIVATE_KEY")
	
	client, err := ethclient.Dial(rpcURL)
	if err != nil { log.Fatalf("Failed to connect to client: %v", err) }

	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil { log.Fatalf("Error parsing private key: %v", err) }

	chainID, err := client.ChainID(context.Background())
	if err != nil { log.Fatalf("Failed to get chain ID: %v", err) }
	
	serviceStatusInstance, err := NewServiceStatusContract(getEnvAddress("SERVICE_STATUS_CONTRACT_ADDRESS"), client)
	if err != nil { log.Fatalf("Failed to instantiate ServiceStatusContract: %v", err) }
	
	recruitingInstance, err := NewRecruitingContract(getEnvAddress("RECRUITING_CONTRACT_ADDRESS"), client)
	if err != nil { log.Fatalf("Failed to instantiate RecruitingContract: %v", err) }
	
	timeSeriesInstance, err := NewTimeSeriesContract(getEnvAddress("TIME_SERIES_CONTRACT_ADDRESS"), client)
	if err != nil { log.Fatalf("Failed to instantiate TimeSeriesContract: %v", err) }

	keyValueInstance, err := NewKeyValueContract(getEnvAddress("KEY_VALUE_CONTRACT_ADDRESS"), client)
	if err != nil { log.Fatalf("Failed to instantiate KeyValueContract: %v", err) }

	supportFundInstance, err := NewSupportFundContract(getEnvAddress("SUPPORT_FUND_CONTRACT_ADDRESS"), client)
	if err != nil { log.Fatalf("Failed to instantiate SupportFundContract: %v", err) }
	
	fmt.Println("Backend service started. All services are running concurrently...")

	// 3. Khởi chạy "Người Xử Lý Giao Dịch" - chỉ một goroutine duy nhất
	go transactionProcessor(privateKey, chainID, client)

	// 4. Các service giờ chỉ tạo và GỬI YÊU CẦU vào hàng đợi
	go updateServiceStatus(serviceStatusInstance)
	go addRecruitingJob(recruitingInstance)
	go updateTimeSeries(timeSeriesInstance)
	go updateKeyValue(keyValueInstance)
	go updateSupportFund(supportFundInstance)

	// Giữ cho chương trình chính chạy mãi mãi
	select {}
}

// 5. "Người Xử Lý Giao Dịch" Trung Tâm
func transactionProcessor(pk *ecdsa.PrivateKey, chainID *big.Int, client *ethclient.Client) {
	fmt.Println("Transaction processor started...")
	for request := range txQueue {
		// Lấy nonce mới nhất CHO MỖI YÊU CẦU
		auth := createTransactor(pk, chainID, client)
		if auth == nil {
			log.Printf("[%s] Failed to create transactor, skipping request.", request.ServiceName)
			continue
		}

		tx, err := request.Action(auth)
		logTransaction(request.ServiceName, tx, err)
		
		// Khoảng nghỉ ngắn giữa các giao dịch để tránh spam node
		time.Sleep(200 * time.Millisecond)
	}
}



// --- CÁC HÀM DỊCH VỤ CON (Sửa lại để chỉ gửi yêu cầu) ---

func updateServiceStatus(instance *ServiceStatusContract) {
	ticker := time.NewTicker(15 * time.Second)
	for range ticker.C {
		txQueue <- TxRequest{
			ServiceName: "ServiceStatus",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				serviceIndex := big.NewInt(int64(rand.Intn(5)))
				newStatus := uint8(rand.Intn(3))
				fmt.Printf("[ServiceStatus] Queueing request to update index %d to status %d\n", serviceIndex, newStatus)
				return instance.UpdateServiceStatus(auth, serviceIndex, newStatus)
			},
		}
	}
}

func addRecruitingJob(instance *RecruitingContract) {
	ticker := time.NewTicker(20 * time.Second)
	var jobCounter int64 = 0
	for range ticker.C {
		jobCounter++
		// Tạo biến cục bộ để goroutine không bị ảnh hưởng bởi các lần lặp sau
		localCounter := jobCounter 
		txQueue <- TxRequest{
			ServiceName: "Recruiting",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				foundations := []string{"Chainlink", "Polygon", "Uniswap", "Aave"}
				positions := []string{"Smart Contract Engineer", "DeFi Analyst", "Rust Developer"}
				fields := []string{"Blockchain", "AI", "IOT"}
				newJob := DataStructsJob{
					Foundation: foundations[rand.Intn(len(foundations))] + fmt.Sprintf(" %d", localCounter),
					Position:   positions[rand.Intn(len(positions))],
					Field:      fields[rand.Intn(len(fields))],
					Salary:     big.NewInt(int64(60000 + rand.Intn(20000))),
					Form:       "Remote",
					Trend:      uint8(rand.Intn(2)),
					IconId:     "default",
				}
				fmt.Printf("[Recruiting] Queueing request to add job: %s\n", newJob.Foundation)
				return instance.AddJob(auth, newJob)
			},
		}
	}
}

func updateTimeSeries(instance *TimeSeriesContract) {
	ticker := time.NewTicker(5 * time.Second)
	var lastValue int64 = 250000
	for range ticker.C {
		// 1. Gửi yêu cầu cho Blood Status
		txQueue <- TxRequest{
			ServiceName: "TimeSeries-Blood",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				return instance.PushDataPoint(auth, "bloodStatus", big.NewInt(int64(90+rand.Intn(20))), big.NewInt(0))
			},
		}

		// 2. Gửi yêu cầu cho Temperature
		txQueue <- TxRequest{
			ServiceName: "TimeSeries-Temp",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				return instance.PushDataPoint(auth, "temperature", big.NewInt(int64(36+rand.Intn(2))), big.NewInt(0))
			},
		}

		// 3. Gửi yêu cầu cho Heart Rate
		txQueue <- TxRequest{
			ServiceName: "TimeSeries-Heart",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				return instance.PushDataPoint(auth, "heartRate", big.NewInt(int64(70+rand.Intn(40))), big.NewInt(0))
			},
		}

		// 4. Tính toán và gửi yêu cầu cho Profile Overview
		changePercent := (rand.Float64() - 0.5) * 2.0
		newValueFloat := float64(lastValue) * (1 + changePercent/100)
		lastValue = int64(newValueFloat)
		
		// Tạo bản sao của giá trị để closure sử dụng đúng
		valueToSend := lastValue
		changeToSend := int64(changePercent * 10)

		txQueue <- TxRequest{
			ServiceName: "TimeSeries-Profile",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				fmt.Printf("[TimeSeries] Queueing request to push profile data: %d\n", valueToSend)
				return instance.PushDataPoint(auth, "profileOverview_1D", big.NewInt(valueToSend), big.NewInt(changeToSend))
			},
		}
	}
}



func updateKeyValue(instance *KeyValueContract) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		randomFilter := []string{"1D", "1W", "1M", "6M", "1Y"}[rand.Intn(5)]
		
		newViewers := big.NewInt(int64(10000 + rand.Intn(4000000)))
		txQueue <- TxRequest{
			ServiceName: "KeyValue-Viewers-" + randomFilter,
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				fmt.Printf("[KeyValue] Queueing request to update viewers for %s\n", randomFilter)
				return instance.UpdateSingleValueData(auth, "viewers_"+randomFilter, newViewers)
			},
		}
		
		newReturn := big.NewInt(int64(rand.Intn(500) - 100))
		txQueue <- TxRequest{
			ServiceName: "KeyValue-Return-" + randomFilter,
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				fmt.Printf("[KeyValue] Queueing request to update return for %s\n", randomFilter)
				return instance.UpdateReturnData(auth, "return_"+randomFilter, newReturn)
			},
		}
	}
}

func updateSupportFund(instance *SupportFundContract) {
	ticker := time.NewTicker(25 * time.Second)
	for range ticker.C {
		txQueue <- TxRequest{
			ServiceName: "SupportFund",
			Action: func(auth *bind.TransactOpts) (*types.Transaction, error) {
				daysArray := make([]string, 30)
				// Lấy thời gian hiện tại một lần
				now := time.Now() 
				for i := 0; i < 30; i++ {
					// ✨ SỬA LẠI Ở ĐÂY:
                    // Tính toán ngày bằng cách trừ đi `i` ngày từ `now`
					date := now.AddDate(0, 0, -i) 
					daysArray[i] = fmt.Sprintf("%s %d", date.Format("Mon"), date.Day())
				}
				randomDayKey := daysArray[rand.Intn(len(daysArray))]
				itemCount := rand.Intn(3) + 2
				var newItems []DataStructsFundItem
				for i := 0; i < itemCount; i++ {
					newItems = append(newItems, DataStructsFundItem{
						Name:        fmt.Sprintf("Item %d", i+1),
						Price:       big.NewInt(int64(10 + rand.Intn(20))),
						IconId:      []string{"sun", "moon"}[rand.Intn(2)],
						Count:       big.NewInt(int64(rand.Intn(3) + 1)),
						AvatarColor: "bg-gray-600",
					})
				}
				fmt.Printf("[SupportFund] Queueing request for %s with %d items\n", randomDayKey, len(newItems))
				return instance.UpdateFundItemsForDay(auth, randomDayKey, newItems)
			},
		}
	}
}


// --- HÀM HELPER (Không đổi) ---
func createTransactor(pk *ecdsa.PrivateKey, chainID *big.Int, client *ethclient.Client) *bind.TransactOpts {
	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil { log.Printf("Nonce error: %v\n", err); return nil }
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil { log.Printf("Gas price error: %v\n", err); return nil }

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil { log.Printf("Tx opts error: %v\n", err); return nil }
	
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(500000)
	auth.GasPrice = gasPrice
	return auth
}

func logTransaction(serviceName string, tx *types.Transaction, err error) {
	if err != nil {
		log.Printf("[%s] Transaction failed: %v\n", serviceName, err)
	} else {
		fmt.Printf("[%s] Tx sent! Hash: %s\n", serviceName, tx.Hash().Hex())
	}
}