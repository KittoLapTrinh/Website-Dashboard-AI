import { useState, useEffect, useMemo } from 'react';
import { useReadContracts, useWatchContractEvent } from 'wagmi';
import { dashboardAddress, dashboardAbi } from '@/app/contracts';

type TimeFilter = '1D' | '1W' | '1M' | '6M' | '1Y';
interface ViewerData {
  viewers: number;
  return: number;
}

export function useTotalViewersData(filter: TimeFilter) {
  // ✨ 1. State để lưu trữ dữ liệu hiển thị
  const [data, setData] = useState<ViewerData | null>(null);

  // 2. Vẫn dùng useReadContracts để tải dữ liệu ban đầu một lần
  const { data: initialData, isLoading: initialLoading, error, refetch } = useReadContracts({
    contracts: [
      {
        address: dashboardAddress,
        abi: dashboardAbi,
        functionName: 'getSingleValueData',
        args: [`viewers_${filter}`],
      },
      {
        address: dashboardAddress,
        abi: dashboardAbi,
        functionName: 'getReturnData',
        args: [`return_${filter}`],
      },
    ],
  });

  // 3. Cập nhật state khi dữ liệu ban đầu được tải hoặc khi bộ lọc thay đổi
  useEffect(() => {
    if (initialData && initialData[0].status === 'success' && initialData[1].status === 'success') {
      const viewers = Number(initialData[0].result as bigint);
      const returnValue = Number(initialData[1].result as bigint) / 10;
      setData({ viewers, return: returnValue });
    }
  }, [initialData, filter]); // Phụ thuộc vào cả initialData và filter

  // 4. ✨ LẮNG NGHE EVENT `TotalViewersUpdated` ✨
  useWatchContractEvent({
    address: dashboardAddress,
    abi: dashboardAbi,
    eventName: 'TotalViewersUpdated',
    onLogs(logs) {
      console.log('New TotalViewers event received:', logs);
      
      // Lấy dữ liệu từ event đầu tiên trong logs
      const eventData = (logs[0] as any).args;

      // Chỉ cập nhật state nếu event dành cho bộ lọc hiện tại
      if (eventData.timeframe === filter) {
        setData({
          viewers: Number(eventData.viewers),
          return: Number(eventData.returnPct) / 10,
        });
      }
    },
  });
  
  // Trả về isLoading của lần tải đầu tiên
  const isLoading = initialLoading || data === null;

  return { data, isLoading, error, refetch };
}