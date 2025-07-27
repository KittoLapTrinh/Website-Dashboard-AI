package blockchain

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// AllContracts là một struct "container" để giữ các instance của tất cả
// các smart contract mà ứng dụng cần.
type AllContracts struct {
	ServiceStatus *ServiceStatusContract
	Recruiting    *RecruitingContract
	TimeSeries    *TimeSeriesContract  // ✨ ĐÃ SỬA: Tên đúng là "Blockchain" ✨
	KeyValue      *KeyValueContract
	SupportFund   *SupportFundContract
	Content       *ContentContract
}

// LoadAllContracts khởi tạo và trả về một struct chứa các instance của tất cả các contract.
func LoadAllContracts(client *ethclient.Client) (*AllContracts, error) {
	log.Println("Loading all smart contract instances...")

	// 1. Load ServiceStatusContract
	serviceStatusAddr := getContractEnvAddress("SERVICE_STATUS_CONTRACT_ADDRESS")
	serviceStatusInstance, err := NewServiceStatusContract(serviceStatusAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ServiceStatusContract: %w", err)
	}

	// 2. Load RecruitingContract
	recruitingAddr := getContractEnvAddress("RECRUITING_CONTRACT_ADDRESS")
	recruitingInstance, err := NewRecruitingContract(recruitingAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate RecruitingContract: %w", err)
	}

	// 3. Load TimeSeriesContract (với tên đúng là Blockchain)
	timeSeriesAddr := getContractEnvAddress("TIME_SERIES_CONTRACT_ADDRESS")
	// ✨ ĐÃ SỬA: Gọi hàm NewTimeSeriesContract ✨
	timeSeriesInstance, err := NewTimeSeriesContract(timeSeriesAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate TimeSeries contract (as Blockchain): %w", err)
	}

	// 4. Load KeyValueContract
	keyValueAddr := getContractEnvAddress("KEY_VALUE_CONTRACT_ADDRESS")
	keyValueInstance, err := NewKeyValueContract(keyValueAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate KeyValueContract: %w", err)
	}

	// 5. Load SupportFundContract
	supportFundAddr := getContractEnvAddress("SUPPORT_FUND_CONTRACT_ADDRESS")
	supportFundInstance, err := NewSupportFundContract(supportFundAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate SupportFundContract: %w", err)
	}

	contentAddr := getContractEnvAddress("CONTENT_CONTRACT_ADDRESS")
	contentInstance, err := NewContentContract(contentAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ContentContract: %w", err)
	}

	log.Println("All smart contract instances loaded successfully.")

	return &AllContracts{
		ServiceStatus: serviceStatusInstance,
		Recruiting:    recruitingInstance,
		TimeSeries:    timeSeriesInstance, // ✨ ĐÃ SỬA ✨
		KeyValue:      keyValueInstance,
		SupportFund:   supportFundInstance,
		Content:       contentInstance,
	}, nil
}

// getContractEnvAddress là hàm helper nội bộ để đọc địa chỉ contract.
func getContractEnvAddress(key string) common.Address {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("FATAL: Environment variable for contract address %s is not set.", key)
	}
	if !common.IsHexAddress(val) {
		log.Fatalf("FATAL: Environment variable %s with value %s is not a valid Ethereum address.", key, val)
	}
	return common.HexToAddress(val)
}