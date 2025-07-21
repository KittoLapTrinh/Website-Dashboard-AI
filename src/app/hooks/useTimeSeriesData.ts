// File: src/app/hooks/useTimeSeriesData.ts (File mới hoặc đặt trong lib)
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
            setData(prevData => [...prevData.slice(1), formattedNewPoint]);
        }
      }
    },
  });
  
  return { data, isLoading: isLoading && data.length === 0, error, refetch };
}