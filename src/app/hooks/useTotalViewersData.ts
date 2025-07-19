import { useReadContracts } from 'wagmi';
import { dashboardAddress, dashboardAbi } from '@/app/contracts/index';
import { useMemo } from 'react';

type TimeFilter = '1D' | '1W' | '1M' | '6M' | '1Y';

export function useTotalViewersData(filter: TimeFilter) {
  const { data: rawData, isLoading, error, refetch } = useReadContracts({
    // Dùng useReadContracts để gộp 2 lần gọi thành 1 request
    contracts: [
      {
        address: dashboardAddress,
        abi: dashboardAbi,
        functionName: 'getSingleValueData',
        args: [`viewers_${filter}`], // Key động, ví dụ: 'viewers_6M'
      },
      { 
        address: dashboardAddress,
        abi: dashboardAbi,
        functionName: 'getReturnData',
        args: [`return_${filter}`], // Key động, ví dụ: 'return_6M'
      },
    ],
  });

  // Định dạng lại dữ liệu trả về từ contract
  const data = useMemo(() => {
    if (!rawData || rawData.some(d => d.status === 'failure')) return null;

    const viewers = Number(rawData[0].result as bigint);
    const returnValue = Number(rawData[1].result as bigint) / 10; // Chuyển 122 -> 12.2

    return { viewers, return: returnValue };
  }, [rawData]);

  return { data, isLoading, error, refetch };
}