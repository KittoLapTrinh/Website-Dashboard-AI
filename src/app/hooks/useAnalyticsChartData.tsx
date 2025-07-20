"use client";

import { useState, useEffect } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
import { dashboardAddress, dashboardAbi } from '@/app/contracts'; // Sửa lại đường dẫn
import { type DataPoint } from '@/app/lib/dashboard-types'; // Sửa lại đường dẫn

type TimeFilter = 'Weekly' | 'Monthly' | 'Yearly';

interface RawDataPoint {
  timestamp: bigint;
  value: bigint;
  change: bigint;
}

// Định nghĩa kiểu dữ liệu cụ thể cho log event mà chúng ta mong đợi
type DataPointAddedLog = {
  eventName: 'DataPointAdded';
  args: {
    key?: string;
    newPoint?: RawDataPoint;
  };
}

const formatDateLabel = (timestamp: bigint, timeframe: TimeFilter): string => {
  const date = new Date(Number(timestamp) * 1000);
  if (timeframe === 'Weekly') return date.toLocaleDateString('en-US', { weekday: 'short' });
  if (timeframe === 'Monthly') return date.toLocaleDateString('en-US', { day: 'numeric' });
  return date.toLocaleDateString('en-US', { month: 'short' });
};

export function useAnalyticsChartData(filter: TimeFilter) {
  const [data, setData] = useState<DataPoint[]>([]);
  const contractKey = `analytics_${filter}`;

  const { data: initialData, isLoading, error, refetch } = useReadContract({
    address: dashboardAddress, abi: dashboardAbi,
    functionName: 'getTimeSeriesData', args: [contractKey],
  });

  useEffect(() => {
    if (initialData && Array.isArray(initialData)) {
      const formatted = (initialData as RawDataPoint[]).map(point => ({
        date: formatDateLabel(point.timestamp, filter),
        value: Number(point.value),
        change: Number(point.change) / 10,
      }));
      setData(formatted);
    }
  }, [initialData, filter]);
  
  // ✨ SỬA LẠI KHỐI NÀY ✨
  // Lắng nghe event
    useWatchContractEvent({
    address: dashboardAddress,
    abi: dashboardAbi,
    eventName: 'DataPointAdded',
    onLogs(logs) {
        logs.forEach(log => {
        // ✨ 1. Ép kiểu `log` thành `any` TRƯỚC KHI truy cập `args`
        const eventArgs = (log as any).args;

        // ✨ 2. Kiểm tra sự tồn tại của `eventArgs` và các thuộc tính bên trong
        if (eventArgs && eventArgs.key === contractKey && eventArgs.newPoint) {
            // ✨ 3. Gán `newPoint` vào một biến để sử dụng
            const newPoint = eventArgs.newPoint as RawDataPoint;

            console.log(`Event received for ${eventArgs.key}:`, newPoint);
            
            const formattedNewPoint: DataPoint = {
            date: formatDateLabel(newPoint.timestamp, filter),
            value: Number(newPoint.value),
            change: Number(newPoint.change) / 10,
            };

            setData(prevData => [...(prevData || []).slice(1), formattedNewPoint]);
        }
        });
    },
    });
  
  return { data, isLoading: isLoading && data.length === 0, error, refetch };
}