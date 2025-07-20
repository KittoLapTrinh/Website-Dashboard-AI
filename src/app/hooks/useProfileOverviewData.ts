// src/hooks/useProfileOverviewData.tsx

"use client";

import { useState, useEffect } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
import { dashboardAddress, dashboardAbi } from '@/app/contracts';
import { type DataPoint } from '@/app/lib/dashboard-types';
import { Log } from 'viem';

type TimeFilter = '1D' | '1W' | '1M' | '6M' | '1Y';

interface RawDataPoint {
  timestamp: bigint;
  value: bigint;
  change: bigint;
}

const formatDateLabel = (timestamp: bigint, timeframe: TimeFilter): string => {
  const date = new Date(Number(timestamp) * 1000);
  if (timeframe === '1D') return date.toLocaleTimeString('en-US', { hour: 'numeric', hour12: true }).replace(/\s/g, '');
  if (timeframe === '1Y') return date.toLocaleDateString('en-US', { month: 'short' });
  // Mặc định cho 1W, 1M, 6M
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
};

type DecodedLog = { args: { key?: string; newPoint?: RawDataPoint; } } & Log;

export function useProfileOverviewData(filter: TimeFilter) {
  const [data, setData] = useState<DataPoint[]>([]);
  const contractKey = `profileOverview_${filter}`;

  const { data: initialData, isLoading, error, refetch } = useReadContract({
    address: dashboardAddress,
    abi: dashboardAbi,
    functionName: 'getTimeSeriesData',
    args: [contractKey],
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
  
  useWatchContractEvent({
    address: dashboardAddress,
    abi: dashboardAbi,
    eventName: 'DataPointAdded',
    onLogs(logs) {
      const relevantLog = logs.find(log => (log as DecodedLog).args.key === contractKey);
      if (relevantLog) {
        const { newPoint } = (relevantLog as DecodedLog).args;
        if (newPoint) {
            const formattedNewPoint: DataPoint = {
                date: formatDateLabel(newPoint.timestamp, filter),
                value: Number(newPoint.value),
                change: Number(newPoint.change) / 10,
            };
            setData(prevData => [...prevData, formattedNewPoint]);
        }
      }
    },
  });
  
  return { data, isLoading: isLoading && data.length === 0, error, refetch };
}