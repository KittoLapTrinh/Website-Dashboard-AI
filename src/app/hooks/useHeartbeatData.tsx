// File: src/app/hooks/useHeartbeatData.tsx
"use client";
import { useTimeSeriesData } from './useTimeSeriesData';
import { useCallback } from 'react';


export function useHeartbeatData() {
    // Định nghĩa hàm formatDate bên trong hook và bọc bằng useCallback.
    // Thêm tham số `ts: bigint` vào định nghĩa hàm.
    const formatDate = useCallback((ts: bigint) => {
        return new Date(Number(ts) * 1000).toLocaleTimeString('en-US', {
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit'
        });
    }, []); // Mảng dependency rỗng vì hàm này không phụ thuộc vào giá trị nào từ bên ngoài.

  const bloodStatus = useTimeSeriesData('bloodStatus', formatDate);
  const temperature = useTimeSeriesData('temperature', formatDate);
  const heartRate = useTimeSeriesData('heartRate', formatDate);

  return {
    bloodData: bloodStatus.data,
    tempData: temperature.data,
    heartData: heartRate.data,
    isLoading: bloodStatus.isLoading || temperature.isLoading || heartRate.isLoading,
    error: bloodStatus.error || temperature.error || heartRate.error,
    refetch: () => {
      bloodStatus.refetch();
      temperature.refetch();
      heartRate.refetch();
    }
  };
}