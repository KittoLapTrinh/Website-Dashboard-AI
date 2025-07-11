"use client";
import React, { useState } from 'react';
import { AreaChart, Area, Tooltip, ResponsiveContainer, YAxis } from 'recharts';
import { ChevronDown, TrendingUp } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';

// Dữ liệu mẫu cho biểu đồ
const data = [
  { name: 'Jan', uv: 40 }, { name: 'Feb', uv: 30 },
  { name: 'Mar', uv: 60 }, { name: 'Apr', uv: 50 },
  { name: 'May', uv: 90 }, { name: 'Jun', uv: 40 },
  { name: 'Jul', uv: 55 }, { name: 'Aug', uv: 80 },
  { name: 'Sep', uv: 30 }, { name: 'Oct', uv: 50 },
  { name: 'Nov', uv: 20 }, { name: 'Dec', uv: 40 },
];

// Component Tooltip tùy chỉnh để có vị trí chính xác
const CustomTooltip = ({ active, payload }: any) => {
  if (active && payload && payload.length) {
    return (
      <div className="tooltip-with-tail">
        <div className="bg-zinc-800/90 backdrop-blur-sm px-3 py-1.5 rounded-lg flex items-center justify-center gap-x-2 shadow-lg">
          <span className="font-bold text-white">{`${payload[0].value}%`}</span>
          <TrendingUp size={16} className="text-white" />
        </div>
      </div>
    );
  }
  return null;
};

// Component điểm active tùy chỉnh
const CustomActiveDot = (props: any) => {
  const { cx, cy } = props;
  if (!cx || !cy) return null;

  return (
    <g>
      <line x1={cx} y1={cy} x2={cx} y2={200} stroke="#a7a7a7" strokeWidth={1} strokeDasharray="3 3" />
      <circle cx={cx} cy={cy} r="10" stroke="#111827" strokeWidth={4} fill="white" />
    </g>
  );
};

const AnalyticsChart = () => {
  const [activeFilter, setActiveFilter] = useState('Weekly');
  const filters = ['Weekly', 'Monthly', 'Yearly'];

  return (
    <WidgetCard>
      <div className="flex flex-col h-full">
        {/* Header */}
        <div className="flex justify-between items-start mb-4">
          <div>
            <h2 className="text-2xl font-bold text-white">Analytics</h2>
            <p className="text-sm text-gray-400 mt-1">April 2025</p>
          </div>
          <button className="flex items-center gap-x-2 text-sm text-gray-300 border border-gray-600 hover:bg-zinc-800 rounded-full px-4 py-1.5 transition-colors">
            Weekly
            <ChevronDown size={16} />
          </button>
        </div>

        {/* Chart Area */}
        <div className="flex-grow mt-4 h-48">
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart
              data={data}
              margin={{ top: 10, right: 10, left: 10, bottom: 0 }}
            >
              <defs>
                <linearGradient id="analyticsChartGradient" x1="0" y1="0" x2="0" y2="1">
                  {/* ✨ 1. Thay đổi màu và opacity */}
                  <stop offset="5%" stopColor="#409BEE" stopOpacity={10}/>
                  <stop offset="95%" stopColor="#409BEE" stopOpacity={0}/>
                </linearGradient>
                <filter id="analyticsGlow" x="-100%" y="-100%" width="300%" height="300%">
                  {/* ✨ 2. Tăng độ mờ để vầng sáng rộng hơn */}
                  <feGaussianBlur in="SourceGraphic" stdDeviation="12" result="blur" />
                  <feMerge>
                    <feMergeNode in="blur" />
                    <feMergeNode in="SourceGraphic" />
                  </feMerge>
                </filter>
              </defs>
              
              <YAxis domain={[0, 110]} hide={true} />

              <Tooltip 
                content={<CustomTooltip />} 
                cursor={false}
                isAnimationActive={false}
              />

              <Area 
                type="monotone" 
                dataKey="uv" 
                stroke="#e5e7eb" 
                strokeWidth={2}
                fill="url(#analyticsChartGradient)"
                fillOpacity={1}
                activeDot={<CustomActiveDot />}
                style={{ filter: 'url(#analyticsGlow)' }}
              />
            </AreaChart>
          </ResponsiveContainer>
        </div>
      </div>
    </WidgetCard>
  );
};

export default AnalyticsChart;