// File: src/app/hooks/useProfileOverviewData.tsx
"use client";
import { useTimeSeriesData } from './useTimeSeriesData';
import { useCallback } from 'react'; // ✨ 1. IMPORT useCallback

type TimeFilter = '1D' | '1W' | '1M' | '6M' | '1Y';

const formatDateLabel = (timestamp: bigint, timeframe: TimeFilter): string => {
  const date = new Date(Number(timestamp) * 1000);
  if (timeframe === '1D') return date.toLocaleTimeString('en-US', { hour: 'numeric', minute: '2-digit' });
  if (timeframe === '1Y') return date.toLocaleDateString('en-US', { month: 'short' });
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
};

export function useProfileOverviewData(filter: TimeFilter) {
  const contractKey = `profileOverview_${filter}`;
  
  // ✨ 2. BỌC HÀM formatDate BẰNG useCallback
  // Hàm này giờ sẽ chỉ được tạo lại khi `filter` thay đổi.
  const formatDate = useCallback((ts: bigint) => {
    return formatDateLabel(ts, filter);
  }, [filter]); // Dependency là `filter`

  return useTimeSeriesData(contractKey, formatDate);
}