// File: src/app/hooks/useAnalyticsChartData.tsx
"use client";
import { useTimeSeriesData } from './useTimeSeriesData';
import { useCallback } from 'react'; // ✨ 1. IMPORT useCallback

type TimeFilter = 'Weekly' | 'Monthly' | 'Yearly';

const formatDateLabel = (timestamp: bigint, timeframe: TimeFilter): string => {
  const date = new Date(Number(timestamp) * 1000);
  if (timeframe === 'Weekly') return date.toLocaleDateString('en-US', { weekday: 'short' });
  if (timeframe === 'Monthly') return date.toLocaleDateString('en-US', { day: 'numeric' });
  return date.toLocaleDateString('en-US', { month: 'short' });
};

export function useAnalyticsChartData(filter: TimeFilter) {
  // ✨ THAY ĐỔI DUY NHẤT Ở ĐÂY ✨
  // Luôn dùng key "recruitment_activity" bất kể filter là gì,
  // vì backend đã cập nhật cho cả 3 khung thời gian cùng một lúc.
  const contractKey = `analytics_${filter}`; // Giữ nguyên key này vì backend đã cập nhật cả 3

  // Phần còn lại của hook không cần thay đổi
  const formatDate = useCallback((ts: bigint) => {
    return formatDateLabel(ts, filter);
  }, [filter]);

  return useTimeSeriesData(contractKey, formatDate);
}