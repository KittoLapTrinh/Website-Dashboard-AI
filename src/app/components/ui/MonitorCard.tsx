"use client";
import React from 'react';
import { AreaChart, Area, ResponsiveContainer } from 'recharts';

interface MonitorCardProps {
  title: string;
  value: string;
  icon: React.ReactNode;
  data: { value: number }[]; // Mảng dữ liệu cho biểu đồ
  endValueDisplay: React.ReactNode; // Component để hiển thị giá trị cuối
  statusColor?: string;
}

const MonitorCard = ({
  title, value, icon, data, endValueDisplay, statusColor = '#22c55e', // Mặc định là xanh
}: MonitorCardProps) => {


  return (
    <div 
      className="rounded-2xl p-px h-full" 
      style={{ background: statusColor }}
    >
    <div
      className="gradient-border-corner bg-zinc-900 rounded-2xl p-4 h-full flex flex-col"
      style={{ '--border-color': statusColor } as React.CSSProperties}
    >
       <div className="flex items-center gap-x-3">
        <div className="text-white">{icon}</div>
        <div>
          <p className="font-semibold text-white">{title}</p>
          {/* ✨ 3. Đặt màu cho text trực tiếp bằng `style` */}
          <p 
            className="font-semibold text-sm"
            style={{ color: statusColor }}
          >
            {value}
          </p>
        </div>
      </div>
      
        {/* Chart Area */}
        <div className="relative flex-grow mt-4 h-24">
          <ResponsiveContainer width="80%" height="100%">
            <AreaChart data={data} margin={{ top: 35, right: 5, left: 5, bottom: 5 }} >
              <Area type="monotone" dataKey="value" stroke="white" strokeWidth={2} fill="transparent" />
            </AreaChart>
          </ResponsiveContainer>
          {/* Render giá trị cuối cùng ở đây */}
          {endValueDisplay}
        </div>
      </div>
</div>
  );
};

export default MonitorCard; 