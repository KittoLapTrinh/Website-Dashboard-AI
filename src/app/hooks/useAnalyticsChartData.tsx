"use client";
import { useState, useEffect } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
import { timeSeriesContractAddress, timeSeriesContractAbi } from '@/app/contracts';
import { type DataPoint } from '@/app/lib/dashboard-types';
import type { Log } from 'viem';

type TimeFilter = 'Weekly' | 'Monthly' | 'Yearly';

// Các kiểu dữ liệu thô (giữ nguyên)
interface RawDataPoint { timestamp: bigint; value: bigint; change: bigint; }
type DecodedDataPointLog = { args: { key?: string; newPoint?: RawDataPoint; } } & Log;

const MAX_DATA_POINTS = 200;

// Sửa lại hàm format: giữ `date` là `number` (miligiây)
const formatDataPoint = (point: RawDataPoint): DataPoint => ({
    date: new Date(Number(point.timestamp) * 1000).toISOString(), 
    value: Number(point.value),
    change: Number(point.change),
});

export function useAnalyticsChartData(filter: TimeFilter) {
    const [data, setData] = useState<DataPoint[]>([]);
    const contractKey = `analytics_${filter}`;

    // Hook useReadContract giờ sẽ là nguồn dữ liệu duy nhất khi tải và filter
    const { 
        data: contractData, 
        isLoading, 
        isFetching,
        error, 
    } = useReadContract({
        address: timeSeriesContractAddress,
        abi: timeSeriesContractAbi,
        functionName: 'getTimeSeriesData',
        args: [contractKey],
    });

    // useEffect chính để đồng bộ state `data` với dữ liệu từ contract
    useEffect(() => {
        // Khi wagmi trả về dữ liệu mới (sau khi tải lần đầu hoặc sau khi filter),
        // hãy cập nhật toàn bộ state.
        if (contractData && Array.isArray(contractData)) {
            setData((contractData as RawDataPoint[]).map(formatDataPoint));
        }
    }, [contractData]); // Chỉ chạy lại khi `contractData` thực sự thay đổi
    
    // ✨ XÓA BỎ HOÀN TOÀN `useEffect` GÂY LỖI MẤT DỮ LIỆU ✨
    // useEffect(() => {
    //     setData([]);
    // }, [filter]);

    // Lắng nghe sự kiện để cập nhật real-time
    useWatchContractEvent({
        address: timeSeriesContractAddress,
        abi: timeSeriesContractAbi,
        eventName: 'DataPointAdded',
        onLogs(logs) {
            const relevantLogs = logs.filter(log => (log as DecodedDataPointLog).args.key === contractKey);
            
            if (relevantLogs.length > 0) {
                const newPoints = relevantLogs.map(log => {
                    const newPointRaw = (log as DecodedDataPointLog).args.newPoint;
                    return formatDataPoint(newPointRaw!);
                });

                setData(prevData => {
                    const combinedData = [...prevData, ...newPoints];
                    // Logic cửa sổ trượt
                    if (combinedData.length > MAX_DATA_POINTS) {
                        return combinedData.slice(combinedData.length - MAX_DATA_POINTS);
                    }
                    return combinedData;
                });
            }
        },
    });
    
    // Trạng thái loading sẽ là true khi đang fetch lần đầu HOẶC đang fetch lại do đổi filter
    // và chưa có dữ liệu nào trong state.
    const finalIsLoading = (isLoading || isFetching) && data.length === 0;

    return { data, isLoading: finalIsLoading, error };
}