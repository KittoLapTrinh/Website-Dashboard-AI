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
  const contractKey = `analytics_${filter}`;

  // ✨ 2. BỌC HÀM formatDate BẰNG useCallback
  const formatDate = useCallback((ts: bigint) => {
    return formatDateLabel(ts, filter);
  }, [filter]); // Dependency là `filter`

  return useTimeSeriesData(contractKey, formatDate);
}