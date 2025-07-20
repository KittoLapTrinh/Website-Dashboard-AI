"use client";

import { useState, useEffect, useMemo } from 'react';
import { useReadContracts, useWatchContractEvent } from 'wagmi';
import { dashboardAddress, dashboardAbi } from '@/app/contracts';
import { type DataPoint } from '@/app/lib/dashboard-types';
import { Log } from 'viem';

// --- KIỂU DỮ LIỆU THÔ TỪ CONTRACT ---
interface RawDataPoint {
  timestamp: bigint;
  value: bigint;
  change: bigint;
}

// --- KIỂU DỮ LIỆU CHO LOG EVENT ---
type DecodedDataPointLog = {
  args: {
    key?: string;
    newPoint?: RawDataPoint;
  };
} & Log;

// --- HÀM HELPER ĐỂ ĐỊNH DẠNG DỮ LIỆU ---
const formatDataPoints = (rawData: RawDataPoint[] | undefined): DataPoint[] => {
  if (!rawData || !Array.isArray(rawData)) return [];
  return rawData.map(point => ({
    date: new Date(Number(point.timestamp) * 1000).toLocaleTimeString(),
    value: Number(point.value),
    change: Number(point.change) / 10,
  }));
};


// --- CUSTOM HOOK ---
export function useHeartbeatData() {
  // 1. Dùng state riêng cho mỗi loại dữ liệu
  const [bloodData, setBloodData] = useState<DataPoint[]>([]);
  const [tempData, setTempData] = useState<DataPoint[]>([]);
  const [heartData, setHeartData] = useState<DataPoint[]>([]);

  // 2. Tải dữ liệu ban đầu cho cả 3 biểu đồ trong 1 lần gọi
  const { 
    data: initialData, 
    isLoading, 
    error, 
    refetch 
  } = useReadContracts({
    contracts: [
      { address: dashboardAddress, abi: dashboardAbi, functionName: 'getTimeSeriesData', args: ['bloodStatus'] },
      { address: dashboardAddress, abi: dashboardAbi, functionName: 'getTimeSeriesData', args: ['temperature'] },
      { address: dashboardAddress, abi: dashboardAbi, functionName: 'getTimeSeriesData', args: ['heartRate'] },
    ],
  });

  // 3. Xử lý dữ liệu ban đầu khi có kết quả
  useEffect(() => {
    if (initialData && initialData.every(d => d.status === 'success')) {
      setBloodData(formatDataPoints(initialData[0].result as RawDataPoint[]));
      setTempData(formatDataPoints(initialData[1].result as RawDataPoint[]));
      setHeartData(formatDataPoints(initialData[2].result as RawDataPoint[]));
    }
  }, [initialData]);

  // 4. Lắng nghe event và cập nhật đúng state
  useWatchContractEvent({
    address: dashboardAddress,
    abi: dashboardAbi,
    eventName: 'DataPointAdded',
    onLogs(logs) {
      logs.forEach(log => {
        const decodedLog = log as DecodedDataPointLog;
        const { key, newPoint } = decodedLog.args;

        if (key && newPoint) {
          const formattedNewPoint = {
            date: new Date(Number(newPoint.timestamp) * 1000).toLocaleTimeString(),
            value: Number(newPoint.value),
            change: Number(newPoint.change) / 10,
          };

          // Dùng switch để cập nhật đúng state dựa trên key từ event
          switch (key) {
            case 'bloodStatus':
              setBloodData(prev => [...prev.slice(1), formattedNewPoint]);
              break;
            case 'temperature':
              setTempData(prev => [...prev.slice(1), formattedNewPoint]);
              break;
            case 'heartRate':
              setHeartData(prev => [...prev.slice(1), formattedNewPoint]);
              break;
            default:
              break;
          }
        }
      });
    },
  });

  return { bloodData, tempData, heartData, isLoading, error, refetch };
}