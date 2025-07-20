"use client";

import React, { useState, useEffect } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
import { dashboardAddress, dashboardAbi } from '@/app/contracts';
import { type FundItemData } from '@/app/lib/dashboard-types';
import { Sun, Moon } from 'lucide-react';
import { Log } from 'viem';

interface RawFundItem {
  name: string; price: bigint; iconId: string; count: bigint; avatarColor: string;
}
type DecodedFundLog = { args: { dayKey?: string; newItems?: RawFundItem[] } } & Log;

const formatFundItem = (item: RawFundItem): FundItemData => {
  const getIcon = (iconId: string) => {
    switch (iconId.toLowerCase()) {
      case 'sun': return <Sun size={18} />;
      case 'moon': return <Moon size={18} />;
      default: return null;
    }
  };
  return {
    name: item.name, price: Number(item.price), icon: getIcon(item.iconId),
    count: Number(item.count), avatarColor: item.avatarColor,
  };
};

export function useSupportFundData(day: string) {
  const [items, setItems] = useState<FundItemData[]>([]);

  // 1. Tải dữ liệu ban đầu
  const { data: initialItems, isLoading, error, refetch } = useReadContract({
    address: dashboardAddress, abi: dashboardAbi,
    functionName: 'getFundItemsByDay',
    args: [day], // Ví dụ: day = "Sun 22"
  });

  // 2. Xử lý dữ liệu ban đầu
  useEffect(() => {
    console.log(`Fetching initial data for ${day}:`, initialItems); // ✨ Thêm log
    if (initialItems && Array.isArray(initialItems)) {
      setItems((initialItems as RawFundItem[]).map(formatFundItem));
    } else {
      setItems([]);
    }
  }, [initialItems, day]);

  // 3. Lắng nghe event
  useWatchContractEvent({
    address: dashboardAddress, abi: dashboardAbi,
    eventName: 'FundItemsUpdated',
    onLogs(logs) {
      logs.forEach(log => {
        const { dayKey, newItems } = (log as DecodedFundLog).args;
        console.log(`Event received for day: ${dayKey}`); // ✨ Thêm log
        if (dayKey === day && newItems) {
          setItems(newItems.map(formatFundItem));
        }
      });
    },
  });

  return { items, isLoading: isLoading && items.length === 0, error, refetch };
}