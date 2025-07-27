package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"dashboard/backend/internal/automations"
	"dashboard/backend/internal/blockchain"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Backend Automator Service...")

	// --- BƯỚC 1: THIẾT LẬP CÁC THÀNH PHẦN CỐT LÕI ---

	client, createTransactorFunc, err := blockchain.NewBlockchainClient()
	if err != nil {
		log.Fatalf("CRITICAL: Could not initialize blockchain client: %v", err)
	}

	allContracts, err := blockchain.LoadAllContracts(client)
	if err != nil {
		log.Fatalf("CRITICAL: Could not load smart contracts: %v", err)
	}

	// --- BƯỚC 2: KHỞI CHẠY CÁC DỊCH VỤ NỀN (BACKGROUND SERVICES) ---
	log.Println("All components initialized. Starting background services...")

	// 2.1. Khởi chạy bộ xử lý giao dịch trung tâm
	go blockchain.TransactionProcessor(createTransactorFunc)

	// ✨ 2.2. KHỞI CHẠY CÁC DỊCH VỤ TỪ PACKAGE `automations` ✨
	// Đây là phần đã được cập nhật để gọi các hàm mới của bạn.
	go automations.StartRecruitingUpdates(allContracts.Recruiting)
	go automations.StartServiceStatusUpdates(allContracts.ServiceStatus)
	go automations.StartTimeSeriesUpdates(allContracts.TimeSeries)
	go automations.StartKeyValueUpdates(allContracts.KeyValue)
	go automations.StartSupportFundUpdates(allContracts.SupportFund)
	go automations.StartContentUpdates(allContracts.Content)
	
	log.Println("All services are now running in the background.")

	// --- BƯỚC 3: CHỜ TÍN HIỆU DỪNG ĐỂ SHUTDOWN MỘT CÁCH AN TOÀN ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("Shutdown signal received. Exiting gracefully.")
}