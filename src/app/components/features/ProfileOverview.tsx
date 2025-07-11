// 1. Thêm "use client" vì Recharts sử dụng hooks và các API của trình duyệt
"use client";

import React, { useState } from 'react';
import { AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';
import WidgetCard from '../ui/WidgetCard';
import { TrendingUp } from 'lucide-react';

// 2. Dữ liệu mẫu cho biểu đồ
const data = [
  { date: '1st Jan', value: 200000 }, { date: '15th Jan', value: 125000 },
  { date: '1st Feb', value: 135000 }, { date: '15th Feb', value: 110000 },
  { date: '1st Mar', value: 145000 }, { date: '15th Mar', value: 100000 },
  { date: '1st Apr', value: 130000 }, { date: '15th Apr', value: 150000 },
  { date: '1st May', value: 175000 }, { date: '15th May', value: 120000 },
  { date: '1st Jun', value: 90000 },  { date: '15th Jun', value: 55000 },
  { date: '1st Jul', value: 85000 },  { date: '15th Jul', value: 45000 },
];

// 3. Component Tooltip tùy chỉnh để giống với thiết kế
const CustomTooltip = ({ active, payload, label }: any) => {
  if (active && payload && payload.length) {
    const value = payload[0].value;
    return (
      <div className="bg-zinc-800/90 backdrop-blur-sm p-3 rounded-lg border border-zinc-700 shadow-lg">
        <p className="text-xs text-gray-400">{`${label} 2024`}</p>
        <p className="text-lg font-bold text-white mt-1">
          {`$${value.toLocaleString()}`}
        </p>
        <div className="flex items-center gap-x-1 text-xs bg-green-900/50 text-green-400 border border-green-700/60 rounded-full px-2 py-0.5 mt-2 w-fit">
          <TrendingUp size={14} />
          +3.5%
        </div>
      </div>
    );
  }
  return null;
};

// 4. Component chính đã được cập nhật
const ProfileOverview = () => {
  const timeFilters = ['1D', '1W', '1M', '6M', '1Y'];
  const [activeFilter, setActiveFilter] = useState('6M');

  // Hàm định dạng số trên trục Y (ví dụ: 150000 -> 150k)
  const formatYAxis = (tickItem: number) => {
    return tickItem > 0 ? `${tickItem / 1000}k` : '0';
  };

  return (
    <WidgetCard>
      <div className="flex flex-col h-full">
        {/* Phần Header của Widget */}
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold text-white">Profile Overview</h2>
          <div className="flex items-center gap-x-2">
            {timeFilters.map((filter) => (
              <button
                key={filter}
                onClick={() => setActiveFilter(filter)}
                className={`text-sm rounded-full w-12 h-8 transition-colors flex items-center justify-center ${
                  activeFilter === filter
                    ? 'bg-blue-600 text-white'
                    : 'border border-gray-600 text-gray-300 hover:bg-gray-700/50'
                }`}
              >
                {filter}
              </button>
            ))}
          </div>
        </div>

        {/* Phần thân - Biểu đồ Recharts */}
        <div className="flex-grow mt-4 h-72">
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart
              data={data}
              margin={{ top: 5, right: 20, left: -10, bottom: 0 }}
            >
              <defs>
                {/* Định nghĩa gradient cho vùng area */}
                <linearGradient id="colorValue" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="5%" stopColor="#2563eb" stopOpacity={0.4}/>
                  <stop offset="95%" stopColor="#2563eb" stopOpacity={0}/>
                </linearGradient>
              </defs>
              <CartesianGrid 
                vertical={false} 
                stroke="#404040" 
                strokeDasharray="0"
              />
              <XAxis 
                dataKey="date" 
                axisLine={false} 
                tickLine={false} 
                tick={{ fill: '#a3a3a3', fontSize: 12 }}
                dy={10} // Đẩy label trục X xuống dưới một chút
              />
              <YAxis 
                axisLine={false} 
                tickLine={false} 
                tick={{ fill: '#a3a3a3', fontSize: 12 }}
                tickFormatter={formatYAxis}
              />
              <Tooltip content={<CustomTooltip />} cursor={{ stroke: '#2563eb', strokeWidth: 1, strokeDasharray: '3 3' }} />
              <Area 
                type="monotone" 
                dataKey="value" 
                stroke="#2563eb" 
                strokeWidth={2}
                fillOpacity={1} 
                fill="url(#colorValue)" 
                activeDot={{ r: 6, stroke: 'white', strokeWidth: 2, fill: '#2563eb' }}
              />
            </AreaChart>
          </ResponsiveContainer>
        </div>
      </div>
    </WidgetCard>
  );
};

export default ProfileOverview;