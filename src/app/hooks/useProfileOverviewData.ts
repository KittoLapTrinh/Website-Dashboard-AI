import { useReadContract } from 'wagmi';
import { dashboardAddress, dashboardAbi } from '@/app/contracts/index';

// Kiểu dữ liệu trả về từ Smart Contract
// Lưu ý: Solidity trả về `bigint` cho `uint256`
interface DataPointFromContract {
  timestamp: bigint;
  value: bigint;
}

// Kiểu dữ liệu đã được định dạng cho biểu đồ
interface FormattedDataPoint {
  date: string;
  value: number;
}

export function useProfileOverviewData() {
  // 1. Gọi hook `useReadContract` để đọc dữ liệu từ SMC
  const { data: rawData, isLoading, error, refetch } = useReadContract({
    address: dashboardAddress,
    abi: dashboardAbi,
    functionName: 'getProfileOverviewData',
  });

  // 2. Định dạng lại dữ liệu để Recharts có thể hiểu
  const formattedData: FormattedDataPoint[] = (rawData as DataPointFromContract[] | undefined)?.map(point => {
    // Chuyển đổi timestamp (giây) sang đối tượng Date
    const date = new Date(Number(point.timestamp) * 1000);
    return {
      // Định dạng ngày tháng
      date: date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' }),
      // Chuyển đổi bigint sang number
      value: Number(point.value)
    };
  }) || [];

  return { 
    data: formattedData, 
    isLoading, 
    error,
    refetch // Trả về hàm refetch để có thể tải lại dữ liệu thủ công
  };
}