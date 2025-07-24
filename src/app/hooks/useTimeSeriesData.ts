// File: src/app/hooks/useTimeSeriesData.ts
"use client";

import { useState, useEffect } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
import { 
  dashboardFacadeAddress, 
  dashboardFacadeAbi,
  timeSeriesContractAddress,
  timeSeriesContractAbi
} from '@/app/contracts';
import { type DataPoint } from '@/app/lib/dashboard-types';
import { Log } from 'viem';

// ✨ THÊM VÀO: Giới hạn số lượng điểm dữ liệu để tránh làm treo trình duyệt
const MAX_DATA_POINTS = 200; 

interface RawDataPoint { timestamp: bigint; value: bigint; change: bigint; }
type DecodedDataPointLog = { args: { key?: string; newPoint?: RawDataPoint; } } & Log;

const formatDataPoint = (point: RawDataPoint, formatDate: (ts: bigint) => string): DataPoint => ({
  date: formatDate(point.timestamp),
  value: Number(point.value),
  change: Number(point.change),
});

export function useTimeSeriesData(
  contractKey: string,
  formatDate: (timestamp: bigint) => string
) {
  const [data, setData] = useState<DataPoint[]>([]);

  const { data: initialData, isLoading, error, refetch } = useReadContract({
    address: dashboardFacadeAddress,
    abi: dashboardFacadeAbi,
    functionName: 'getTimeSeriesData',
    args: [contractKey],
  });

  useEffect(() => {
    if (initialData && Array.isArray(initialData)) {
      setData((initialData as RawDataPoint[]).map(p => formatDataPoint(p, formatDate)));
    }
  }, [initialData, formatDate]);
  
  useWatchContractEvent({
    address: timeSeriesContractAddress,
    abi: timeSeriesContractAbi,
    eventName: 'DataPointAdded',
    onLogs(logs) {
      const relevantLog = logs.find(log => (log as DecodedDataPointLog).args.key === contractKey);
      if (relevantLog) {
        const { newPoint } = (relevantLog as DecodedDataPointLog).args;
        if (newPoint) {
            const formattedNewPoint = formatDataPoint(newPoint, formatDate);
            
            // ✨ --- LOGIC CẬP NHẬT ĐÃ ĐƯỢC SỬA LẠI --- ✨
            setData(prevData => {
                // Thêm điểm dữ liệu mới vào cuối mảng
                const newData = [...prevData, formattedNewPoint];
                
                // Nếu mảng vượt quá giới hạn, hãy xóa đi điểm cũ nhất
                if (newData.length > MAX_DATA_POINTS) {
                    return newData.slice(newData.length - MAX_DATA_POINTS);
                }
                
                // Nếu chưa, trả về mảng đã được thêm mới
                return newData;
            });
        }
      }
    },
  });
  
  return { data, isLoading: isLoading && data.length === 0, error, refetch };
}