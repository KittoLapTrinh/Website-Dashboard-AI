// File: src/app/hooks/useTotalViewersData.tsx
"use client";

import { useState, useEffect } from 'react';
import { useReadContracts, useWatchContractEvent } from 'wagmi';
import { 
  dashboardFacadeAddress, 
  dashboardFacadeAbi,
  keyValueContractAddress,
  keyValueContractAbi
} from '@/app/contracts';
import { type TotalViewersData, type ReturnStatus } from '@/app/lib/dashboard-types';
import { Log } from 'viem';

// --- KIỂU DỮ LIỆU THÔ ---
type TimeFilter = '1D' | '1W' | '1M' | '6M' | '1Y';
type DecodedValueLog = { args: { key?: string; newValue?: bigint; } } & Log;

// --- HÀM BIẾN ĐỔI ---
const getReturnStatus = (value: number): ReturnStatus => {
  if (value > 0) return 'increase';
  if (value < 0) return 'decrease';
  return 'neutral';
};

// --- CUSTOM HOOK ---
export function useTotalViewersData(filter: TimeFilter) {
  const [data, setData] = useState<TotalViewersData | null>(null);
  const viewersKey = `viewers_${filter}`;
  const returnKey = `return_${filter}`;

  const { data: initialData, isLoading: initialLoading, error, refetch } = useReadContracts({
    contracts: [
      { address: dashboardFacadeAddress, abi: dashboardFacadeAbi, functionName: 'getSingleValueData', args: [viewersKey] },
      { address: dashboardFacadeAddress, abi: dashboardFacadeAbi, functionName: 'getReturnData', args: [returnKey] },
    ],
  });

  useEffect(() => {
    if (initialData && initialData[0].status === 'success' && initialData[1].status === 'success') {
      const viewers = Number(initialData[0].result as bigint);
      const returnPct = Number(initialData[1].result as bigint);
      setData({ 
        viewers, 
        return: { value: returnPct, status: getReturnStatus(returnPct) }
      });
    }
  }, [initialData, filter]);

  useWatchContractEvent({
    address: keyValueContractAddress,
    abi: keyValueContractAbi,
    eventName: 'SingleValueUpdated',
    onLogs(logs) {
      logs.forEach(log => {
        const { key, newValue } = (log as DecodedValueLog).args;
        if (key === viewersKey && newValue !== undefined) {
          setData(prevData => ({ ...prevData!, viewers: Number(newValue) }));
        }
      });
    },
  });

  useWatchContractEvent({
    address: keyValueContractAddress,
    abi: keyValueContractAbi,
    eventName: 'ReturnValueUpdated',
    onLogs(logs) {
      logs.forEach(log => {
        const { key, newValue } = (log as DecodedValueLog).args;
        if (key === returnKey && newValue !== undefined) {
          const newReturnPct = Number(newValue);
          setData(prevData => ({ 
            ...prevData!, 
            return: { value: newReturnPct, status: getReturnStatus(newReturnPct) } 
          }));
        }
      });
    },
  });

  return { data, isLoading: initialLoading || data === null, error, refetch };
}