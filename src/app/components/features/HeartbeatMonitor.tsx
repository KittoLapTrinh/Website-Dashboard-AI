"use client";

import React, { useState, useEffect } from 'react';
import WidgetCard from '../ui/WidgetCard';
import MonitorCard from '../ui/MonitorCard';
import { HeartPulse, Thermometer, Activity } from 'lucide-react';
import { MonitorCardSkeleton } from '../ui/skeletons/MonitorCardSkeleton'; 
import { useHeartbeatData } from '@/app/hooks/useHeartbeatData';
import { type DataPoint } from '@/app/lib/dashboard-types';
const HeartbeatMonitor = () => {
   const { bloodData, tempData, heartData, isLoading, error, refetch } = useHeartbeatData();
  // Lấy ra giá trị cuối cùng để hiển thị và quyết định màu sắc
  const lastBloodValue = bloodData[bloodData.length - 1]?.value || 0;
  const lastTempValue = (tempData[tempData.length - 1]?.value || 0).toFixed(1);
  const lastHeartValue = heartData[heartData.length - 1]?.value || 0;

  // const WINDOW_SIZE = 30; 

  const getLastNItems = (arr: DataPoint[], n: number) => {
    if (!arr || arr.length === 0) return [];
    return arr.slice(Math.max(arr.length - n, 0));
  };

  //  const visibleBloodData = getLastNItems(bloodData, WINDOW_SIZE);
  // const visibleTempData = getLastNItems(tempData, WINDOW_SIZE);
  // const visibleHeartData = getLastNItems(heartData, WINDOW_SIZE);

  // HÀM XÁC ĐỊNH MÀU SẮC DỰA TRÊN GIÁ TRỊ
  const getHeartStatusColor = (rate: number) => {
    if (rate > 100) return '#ef4444'; // Đỏ
    if (rate < 60) return '#8b5cf6'; // Tím
    return '#22c55e'; // Xanh
  };
  const getTempStatusColor = (temp: number) => {
    if (temp > 37.5) return '#ef4444'; // Đỏ
    if (temp < 36.0) return '#8b5cf6'; // Tím
    return '#22c55e'; // Xanh
  };
  
  return (
    <WidgetCard>
      <div className="flex flex-col h-full">
        <h2 className="text-2xl font-bold text-white mb-6">Heartbeat Monitor</h2>
        <div className="flex-grow flex items-center">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 w-full">
            {isLoading ? (
              // Hiển thị skeleton khi đang tải
              <>
                <MonitorCardSkeleton />
                <MonitorCardSkeleton />
                <MonitorCardSkeleton />
              </>
            ) : (
              // Hiển thị dữ liệu thật khi đã tải xong
              <>
            {/* Thẻ 1: Blood Status - Giờ đã có dữ liệu động và màu động */}
            <MonitorCard
              title="Blood Status"
              value={`${lastBloodValue}/70`}
              icon={<HeartPulse size={24} />}
              data={bloodData}
              statusColor={getHeartStatusColor(lastBloodValue)}
              endValueDisplay={
                <div className="absolute inset-0 flex justify-end items-center pr-4">
                  <div className="text-right">
                    <p className="text-white font-bold text-2xl">{lastBloodValue}</p>
                    <p className="text-gray-400 text-sm">/70</p>
                  </div>
                </div>
              }
            />

            {/* Thẻ 2: Temperature */}
            <MonitorCard
              title="Temperature"
              value={`${lastTempValue}`}
              icon={<Thermometer size={24} />}
              data={tempData}
              statusColor={getTempStatusColor(parseFloat(lastTempValue))}
              endValueDisplay={
                <div className="absolute inset-0 flex justify-end items-center">
                   <div className="bg-zinc-800/80 backdrop-blur-sm p-3 rounded-lg border border-zinc-700">
                     <p className="text-white font-bold text-lg">{lastTempValue}</p>
                   </div>
                </div>
              }
            />

            {/* Thẻ 3: Heart Rate */}
            <MonitorCard
              title="Heart Rate"
              value={`${lastHeartValue} bpm`}
              icon={<Activity size={24} />}
              data={heartData}
              statusColor={getHeartStatusColor(lastHeartValue)}
              endValueDisplay={
                <div className="absolute inset-0 flex justify-end items-center">
                  <div className="bg-zinc-800/80 backdrop-blur-sm p-2 rounded-lg border border-zinc-700">
                    <p className="text-white font-bold text-lg">{lastHeartValue}</p>
                    <p className="text-gray-400 text-xs">bpm</p>
                  </div>
                </div>
              }
            />
            </>
            )}
          </div>
        </div>
      </div>
    </WidgetCard>
  );
};

export default HeartbeatMonitor;