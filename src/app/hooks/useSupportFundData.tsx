// File: src/app/hooks/useSupportFundData.tsx
"use client";

import React, { useState, useEffect } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
import { 
  dashboardFacadeAddress, 
  dashboardFacadeAbi,
  supportFundContractAddress,
  supportFundContractAbi
} from '@/app/contracts';
import { type FundItemData } from '@/app/lib/dashboard-types';
import { Sun, Moon } from 'lucide-react';
import { Log } from 'viem';

// --- KIỂU DỮ LIỆU THÔ ---
interface RawFundItem { name: string; price: bigint; iconId: string; count: bigint; avatarColor: string; }
type DecodedFundLog = { args: { dayKey?: string; newItems?: RawFundItem[] } } & Log;

// --- HÀM BIẾN ĐỔI ---
const formatFundItem = (item: RawFundItem): FundItemData => ({
  name: item.name,
  price: Number(item.price),
  icon: item.iconId.toLowerCase() === 'sun' ? <Sun size={18} /> : <Moon size={18} />,
  count: Number(item.count),
  avatarColor: item.avatarColor,
});

// --- CUSTOM HOOK ---
export function useSupportFundData(day: string) {
  const [items, setItems] = useState<FundItemData[]>([]);

  const { data: initialItems, isLoading, error, refetch } = useReadContract({
    address: dashboardFacadeAddress,
    abi: dashboardFacadeAbi,
    functionName: 'getFundItemsByDay',
    args: [day],
  });

  useEffect(() => {
    if (initialItems && Array.isArray(initialItems)) {
      setItems((initialItems as RawFundItem[]).map(formatFundItem));
    } else {
      setItems([]);
    }
  }, [initialItems, day]);

  useWatchContractEvent({
    address: supportFundContractAddress,
    abi: supportFundContractAbi,
    eventName: 'FundItemsUpdated',
    onLogs(logs) {
      logs.forEach(log => {
        const { dayKey, newItems } = (log as DecodedFundLog).args;
        if (dayKey === day && newItems) {
          setItems(newItems.map(formatFundItem));
        }
      });
    },
  });

  return { items, isLoading, error, refetch };
}