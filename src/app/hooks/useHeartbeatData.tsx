"use client";
import { useTimeSeriesData } from './useTimeSeriesData';
import { useCallback } from 'react';

// Giờ đây, useHeartbeatData sẽ là một "wrapper" đơn giản
// gọi đến useTimeSeriesData cho từng chỉ số.
export function useHeartbeatData() {
    const formatDate = useCallback((ts: bigint) => {
        return new Date(Number(ts) * 1000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' });
    }, []);

    const bloodStatus = useTimeSeriesData('bloodStatus', formatDate);
    const temperature = useTimeSeriesData('temperature', formatDate);
    const heartRate = useTimeSeriesData('heartRate', formatDate);
    // ✨ BỎ SP02 NẾU BẠN KHÔNG CẦN NỮA ✨
    // const spO2 = useTimeSeriesData('spO2', formatDate);

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