"use client";

import React, { useState, useEffect } from 'react';
import WidgetCard from '../ui/WidgetCard';
import MonitorCard from '../ui/MonitorCard';
import { HeartPulse, Thermometer, Activity } from 'lucide-react';
import { MonitorCardSkeleton } from '../ui/MonitorCardSkeleton'; 
// --- HÀM HELPER ĐỂ TẠO DỮ LIỆU BAN ĐẦU ---
const createInitialData = (length: number, base: number, variance: number) => 
  Array.from({ length }, () => ({ value: base + (Math.random() - 0.5) * variance }));

const HeartbeatMonitor = () => {
  // === STATE ĐỂ LƯU DỮ LIỆU ĐỘNG ===
  const [bloodData, setBloodData] = useState(() => createInitialData(20, 30, 20));
  const [tempData, setTempData] = useState(() => createInitialData(5, 38, 5));
  const [heartData, setHeartData] = useState(() => createInitialData(20, 30, 20));
 const [isLoading, setIsLoading] = useState(true);
  // Lấy ra giá trị cuối cùng để hiển thị và quyết định màu sắc
  const lastBloodValue = Math.round(bloodData[bloodData.length - 1]?.value * 3 + 40); // Giả lập giá trị thực
  const lastTempValue = (tempData[tempData.length - 1]?.value + 15).toFixed(1);
  const lastHeartValue = Math.round(heartData[heartData.length - 1]?.value * 3 + 60);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
    }, 500); // Tải trong 1.5 giây
    return () => clearTimeout(timer);
  }, []);
  // === useEffect ĐỂ MÔ PHỎNG DỮ LIỆU REALTIME ===
  useEffect(() => {
    const interval = setInterval(() => {
      // Cập nhật Blood Status
      setBloodData(prev => {
        const change = (Math.random() - 0.5) * 5;
        const newValue = prev[prev.length - 1].value + change;
        return [...prev.slice(1), { value: Math.max(10, Math.min(50, newValue)) }];
      });
      // Cập nhật Temperature
      setTempData(prev => {
        const change = (Math.random() - 0.5) * 2;
        const newValue = prev[prev.length - 1].value + change;
        return [...prev.slice(1), { value: Math.max(15, Math.min(25, newValue)) }];
      });
      // Cập nhật Heart Rate
      setHeartData(prev => {
        const change = (Math.random() - 0.5) * 8;
        const newValue = prev[prev.length - 1].value + change;
        return [...prev.slice(1), { value: Math.max(10, Math.min(50, newValue)) }];
      });
    }, 2000); // Cập nhật mỗi 2 giây

    return () => clearInterval(interval); // Dọn dẹp khi component unmount
  }, []);

  // === HÀM XÁC ĐỊNH MÀU SẮC DỰA TRÊN GIÁ TRỊ ===
  const getHeartStatusColor = (rate: number) => {
    if (rate > 100) return '#ef4444'; // Đỏ - Báo động cao
    if (rate < 60) return '#8b5cf6'; // Tím - Báo động thấp
    return '#22c55e'; // Xanh - An toàn
  };
  const getTempStatusColor = (temp: number) => {
    if (temp > 37.5) return '#ef4444'; // Đỏ - Sốt
    if (temp < 36.0) return '#8b5cf6'; // Tím - Hạ thân nhiệt
    return '#22c55e'; // Xanh - An toàn
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