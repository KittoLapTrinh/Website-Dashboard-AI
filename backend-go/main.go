package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("FATAL: Environment variable %s is not set.", key)
	}
	return val
}

func main() {
	if err := godotenv.Load("../.env.local"); err != nil {
		log.Fatalf("Error loading .env.local file: %v", err)
	}

	rpcURL := getEnv("MTD_RPC_URL")
	privateKeyHex := getEnv("MTD_PRIVATE_KEY")

	addressBytes, err := os.ReadFile("../src/app/contracts/dashboard-contract-address.json")
	if err != nil { log.Fatalf("Failed to read address file: %v", err) }
	var addrInfo struct{ Address string `json:"address"` }
	if err := json.Unmarshal(addressBytes, &addrInfo); err != nil { log.Fatalf("Failed to parse address JSON: %v", err) }
	contractAddress := common.HexToAddress(addrInfo.Address)

	client, err := ethclient.Dial(rpcURL)
	if err != nil { log.Fatalf("Failed to connect to client: %v", err) }

	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil { log.Fatalf("Error parsing private key: %v", err) }

	chainID, err := client.ChainID(context.Background())
	if err != nil { log.Fatalf("Failed to get chain ID: %v", err) }
	
	instance, err := NewDashboard(contractAddress, client)
	if err != nil { log.Fatalf("Failed to instantiate contract: %v", err) }

	fmt.Println("Backend service started. Sending realtime data to Smart Contract...")
	ticker := time.NewTicker(8 * time.Second) // Tăng thời gian giữa các giao dịch
	defer ticker.Stop()

	// Khởi tạo các giá trị cuối cùng cho dữ liệu realtime
	lastProfileValue := 250000
	lastHeartRate := 75
	lastBloodValue := 95      
	lastTempValue := 37.0 
	jobCounter := 0
 	lastAnalyticsValue := 80
	
	for range ticker.C {
		fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil { log.Printf("Nonce error: %v\n", err); continue }

		gasPrice := big.NewInt(1000000000) // 1 Gwei

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil { log.Printf("Tx opts error: %v\n", err); continue }
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(500000)
		auth.GasPrice = gasPrice
		
		// Chọn ngẫu nhiên một hành động để thực hiện
		action := rand.Intn(7) // 0, 1, 2, hoặc 3

		switch action {
		case 0:
			// Cập nhật Profile Overview
			changePercent := (rand.Float64() - 0.5) * 5.0
			newValueFloat := float64(lastProfileValue) * (1 + changePercent/100)
			lastProfileValue = int(newValueFloat)
			changeInt := int64(changePercent * 10)
			fmt.Printf("Pushing data for profileOverview_1D: %d, change: %d\n", lastProfileValue, changeInt)
			
			tx, err := instance.PushDataPoint(auth, "profileOverview_1D", big.NewInt(int64(lastProfileValue)), big.NewInt(changeInt))
			if err != nil { log.Printf("Tx failed (profileOverview): %v\n", err); continue }
			fmt.Printf("Tx sent! Hash: %s\n", tx.Hash().Hex())
		
		case 1:
			// Cập nhật Total Viewers
			newViewers := big.NewInt(int64(1800000 + rand.Intn(100000)))
			newReturn := big.NewInt(int64(100 + rand.Intn(50)))
			fmt.Printf("Updating TotalViewers for 6M: %s viewers, %s return\n", newViewers.String(), newReturn.String())
			
			tx, err := instance.UpdateTotalViewersData(auth, "6M", newViewers, newReturn)
			if err != nil { log.Printf("Tx failed (totalViewers): %v\n", err); continue }
			fmt.Printf("Tx sent! Hash: %s\n", tx.Hash().Hex())

		case 2:
			// Cập nhật Heartbeat Monitor
			 metricToUpdate := rand.Intn(3) // 0: Blood, 1: Temp, 2: Heart Rate

            var tx *types.Transaction
            var err error

            switch metricToUpdate {
            case 0:
                // Cập nhật Blood Status
                change := rand.Intn(6) - 3 // +/- 3
                newBloodValue := int64(lastBloodValue + change)
                if newBloodValue < 80 { newBloodValue = 80 }
                if newBloodValue > 140 { newBloodValue = 140 }
                lastBloodValue = int(newBloodValue)
                fmt.Printf("Pushing data for bloodStatus: %d\n", lastBloodValue)
                tx, err = instance.PushDataPoint(auth, "bloodStatus", big.NewInt(newBloodValue), big.NewInt(0))

            case 1:
                // Cập nhật Temperature
                change := (rand.Float64() - 0.5) * 0.4 // +/- 0.2
                newTempValue := lastTempValue + change
                if newTempValue < 36.0 { newTempValue = 36.0 }
                if newTempValue > 39.0 { newTempValue = 39.0 }
                lastTempValue = newTempValue
                tempInt := int64(newTempValue) 
                fmt.Printf("Pushing data for temperature: %.1f\n", newTempValue)
                tx, err = instance.PushDataPoint(auth, "temperature", big.NewInt(tempInt), big.NewInt(0))

            case 2:
                // Cập nhật Heart Rate
                change := rand.Intn(10) - 5
                newHeartRate := int64(lastHeartRate + change)
                if newHeartRate < 50 { newHeartRate = 50 }
                if newHeartRate > 120 { newHeartRate = 120 }
                lastHeartRate = int(newHeartRate)
                fmt.Printf("Pushing data for heartRate: %d\n", lastHeartRate)
                tx, err = instance.PushDataPoint(auth, "heartRate", big.NewInt(newHeartRate), big.NewInt(0))
            }
            
            if err != nil { 
                log.Printf("Tx failed (Heartbeat): %v\n", err)
                continue 
            }
            fmt.Printf("Tx sent! Hash: %s\n", tx.Hash().Hex())
		
		case 3:
			// Cập nhật Service Status
			serviceIndex := big.NewInt(int64(rand.Intn(5)))
			newStatus := uint8(rand.Intn(3))
			fmt.Printf("Updating service at index %d to status %d\n", serviceIndex, newStatus)
			
			tx, err := instance.UpdateServiceStatusByIndex(auth, serviceIndex, newStatus)
			if err != nil { log.Printf("Tx failed (updateService): %v\n", err); continue }
			fmt.Printf("Tx sent! Hash: %s\n", tx.Hash().Hex())
		
		case 4:
            // Thêm một Job mới vào RecruitingTable
            jobCounter++

			 foundations := []string{"Chainlink", "Polygon", "Uniswap", "Aave", "The Graph"}
				positions := []string{"Smart Contract Engineer", "Blockchain Developer", "DeFi Analyst", "Rust Developer", "Core Protocol Engineer"}
				fields := []string{"Blockchain", "AI", "IOT", "Web Dev"}
				forms := []string{"Remote", "Fulltime", "Hybrid"}
				iconIds := []string{"apple", "tesla", "google", "nvidia", "fpt"}

				randomFoundation := foundations[rand.Intn(len(foundations))]
				randomPosition := positions[rand.Intn(len(positions))]
				randomField := fields[rand.Intn(len(fields))]
				randomForm := forms[rand.Intn(len(forms))]
				randomIconId := iconIds[rand.Intn(len(iconIds))]
			
				newJob := DashboardJob{
                Foundation: randomFoundation + " " + fmt.Sprintf("%d", jobCounter), // Thêm số đếm để không bị trùng
                Position:   randomPosition,
                Field:      randomField,
                Salary:     big.NewInt(int64(40000 + rand.Intn(50000))), // Lương từ 40k đến 90k
                Form:       randomForm,
                Trend:      uint8(rand.Intn(2)), // 0 for Up, 1 for Down
                IconId:     randomIconId,
            }

            fmt.Printf(
                "Adding new job: %s - %s (%s) - Icon: %s\n", 
                newJob.Foundation, 
                newJob.Position,
                newJob.Field,
                newJob.IconId,
            )
            
            tx, err := instance.AddJob(auth, newJob)
            if err != nil { log.Printf("Tx failed (addJob): %v\n", err); continue }
            fmt.Printf("Tx sent! Hash: %s\n", tx.Hash().Hex())

		case 5:
            // Cập nhật danh sách cho SupportFund
            daysArray := make([]string, 30)
            now := time.Now()
            for i := 0; i < 30; i++ {
                date := now.AddDate(0, 0, -i)
                daysArray[i] = fmt.Sprintf("%s %d", date.Format("Mon"), date.Day())
            }
            
            // ✨ 2. Chọn ngẫu nhiên một ngày từ danh sách 30 ngày đó
            randomDayKey := daysArray[rand.Intn(len(daysArray))]

            // ✨ 3. Tạo dữ liệu ngẫu nhiên cho ngày đã chọn
            itemCount := rand.Intn(3) + 2 // 2 đến 4 item
            var newItems []DashboardFundItem
            for i := 0; i < itemCount; i++ {
                newItems = append(newItems, DashboardFundItem{
                    Name:        fmt.Sprintf("Item %d", i+1),
                    Price:       big.NewInt(int64(10 + rand.Intn(20))),
                    IconId:      []string{"sun", "moon"}[rand.Intn(2)],
                    Count:       big.NewInt(int64(rand.Intn(3) + 1)),
                    AvatarColor: "bg-gray-600",
                })
            }
            fmt.Printf("Updating SupportFund for %s with %d items\n", randomDayKey, len(newItems))

            // ✨ 4. Gọi hàm trong Smart Contract với `dayKey` ngẫu nhiên
            tx, err := instance.UpdateFundItemsForDay(auth, randomDayKey, newItems)
            if err != nil { 
                log.Printf("Tx failed (updateFundItems): %v\n", err)
                continue 
            }
            fmt.Printf("Tx sent! Hash: %s\n", tx.Hash().Hex())
		case 6:
            // ✨ 1. Chọn ngẫu nhiên một trong ba bộ lọc
            filters := []string{"Weekly", "Monthly", "Yearly"}
            randomFilter := filters[rand.Intn(len(filters))]
            dataKey := "analytics_" + randomFilter

            // ✨ 2. Tạo dữ liệu mới
            change := rand.Intn(20) - 10 // Thay đổi từ -10 đến +9
            newValue := int64(lastAnalyticsValue + change)
            if newValue < 20 { newValue = 20 }
            if newValue > 100 { newValue = 100 }
            lastAnalyticsValue = int(newValue)
            
            // Giả sử Analytics Chart không cần % change
            changeInt := int64(0)

            fmt.Printf("Pushing data for %s: %d\n", dataKey, lastAnalyticsValue)
            
            // ✨ 3. Gọi hàm `pushDataPoint` với key động
            tx, err := instance.PushDataPoint(auth, dataKey, big.NewInt(newValue), big.NewInt(changeInt))
            if err != nil { 
                log.Printf("Tx failed (Analytics): %v\n", err)
                continue 
            }
            fmt.Printf("Tx sent! Hash: %s\n", tx.Hash().Hex())
		} // Dấu `}` của switch phải nằm ở đây
	} // Dấu `}` của for phải nằm ở đây
}