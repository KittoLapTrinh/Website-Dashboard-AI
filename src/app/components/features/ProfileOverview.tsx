"use client";

import React, { useState, useEffect } from 'react';
import { AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, Brush  } from 'recharts';
import WidgetCard from '../ui/WidgetCard';
import { TrendingUp, TrendingDown, MoreHorizontal, Loader2 } from 'lucide-react';
import { ClientOnly } from '@/app/components/helpers/ClientOnly';
import { type DataPoint } from '@/app/lib/dashboard-types';
import { useProfileOverviewData } from '@/app/hooks/useProfileOverviewData';


const CustomTooltip = ({ active, payload, label }: any) => {
  if (active && payload && payload.length) {
    const dataPoint: DataPoint = payload[0].payload;
    const { value, change } = dataPoint;
    const isPositive = change !== undefined && change >= 0;

    return (
      <div className="tooltip-with-tail mb-3">
        <div className="bg-zinc-800/90 backdrop-blur-sm p-3 rounded-lg border border-zinc-700 shadow-lg min-w-[150px]">
          <div className="flex justify-between items-center text-xs text-gray-400">
            <span>{label}</span>
            <MoreHorizontal size={14} />
          </div>
          <p className="text-xl font-bold text-white mt-1">
            {`$${value.toLocaleString()}`}
          </p>
          {change !== undefined && (
            <div className={`flex items-center gap-x-1 text-xs rounded-full px-2 py-0.5 mt-2 w-fit ${isPositive ? 'bg-green-900/50 text-green-400 border border-green-700/60' : 'bg-red-900/50 text-red-400 border border-red-700/60'}`}>
              {isPositive ? <TrendingUp size={14} /> : <TrendingDown size={14} />}
              <span>{isPositive ? '+' : ''}{change?.toFixed(1)}%</span>
            </div>
          )}
        </div>
      </div>
    );
  }
  return null;
};

// --- COMPONENT ĐIỂM ACTIVE ---
const CustomActiveDot = (props: any) => {
  const { cx, cy } = props;
  if (typeof cx !== 'number' || typeof cy !== 'number') return null;
  return (
    <g>
      <line x1={cx} y1={cy} x2={cx} y2={200} stroke="#2563eb" strokeWidth={1} strokeDasharray="3 3" />
      <circle cx={cx} cy={cy} r="6" stroke="#0f172a" strokeWidth={4} fill="#2563eb" />
    </g>
  );
};

// --- COMPONENT CHÍNH ---
const ProfileOverview = () => {
  type TimeFilterType = '1D' | '1W' | '1M' | '6M' | '1Y';
  const timeFilters: TimeFilterType[] = ['1D', '1W', '1M', '6M', '1Y'];
  const [activeFilter, setActiveFilter] = useState<TimeFilterType>('6M');
  
  // ✨ GỌI HOOK ĐỂ LẤY DỮ LIỆU TỪ BLOCKCHAIN
  const { data, isLoading, error, refetch } = useProfileOverviewData(activeFilter);

  const formatYAxis = (tickItem: number) => tickItem > 0 ? `${Math.round(tickItem / 1000)}k` : '0';
  
  let xAxisInterval: number | 'preserveStartEnd' = 'preserveStartEnd';
  if (activeFilter === '6M') xAxisInterval = 2;
  else if (activeFilter === '1M') xAxisInterval = 6;
  else if (activeFilter === '1D') xAxisInterval = 3;

  return (
    <WidgetCard>
      <div className="flex flex-col h-full">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold text-white">Profile Overview</h2>
          <div className="flex items-center p-1 bg-zinc-900 rounded-full">
            {timeFilters.map((filter) => (
              <button
                key={filter}
                onClick={() => setActiveFilter(filter)}
                className={`text-sm rounded-full w-12 h-8 transition-colors flex items-center justify-center font-medium
                  ${activeFilter === filter ? 'bg-blue-600 text-white' : 'bg-transparent text-gray-400 hover:text-white'}`}
              >
                {filter}
              </button>
            ))}
          </div>
        </div>
          <div className="flex-grow mt-4 h-72 flex items-center justify-center">
          {isLoading ? (
            <Loader2 className="animate-spin text-gray-500" size={32} />
          ) : (
            <ClientOnly>
              <ResponsiveContainer width="100%" height="100%">
                <AreaChart data={data} margin={{ top: 20, right: 20, left: -20, bottom: 20  }}>
                  <defs>
                    <linearGradient id="profileOverviewGradient" x1="0" y1="0" x2="0" y2="1">
                      <stop offset="5%" stopColor="#2563eb" stopOpacity={0.4}/>
                      <stop offset="100%" stopColor="#2563eb" stopOpacity={0}/>
                    </linearGradient>
                  </defs>
                  <CartesianGrid vertical={false} stroke="#404040" />
                  <XAxis dataKey="date" tick={{ fill: '#a3a3a3', fontSize: 12 }} dy={10} tickLine={false} axisLine={false} interval={xAxisInterval} />
                  <YAxis tickFormatter={formatYAxis} tick={{ fill: '#a3a3a3', fontSize: 12 }} tickLine={false} axisLine={false} />
                  <Tooltip content={<CustomTooltip />} cursor={false} position={{ y: 0 }} isAnimationActive={false} />
                  <Area type="monotone" dataKey="value" stroke="#60a5fa" strokeWidth={2} fill="url(#profileOverviewGradient)" activeDot={<CustomActiveDot />} />
                  {data.length > 20 && ( // Chỉ hiển thị Brush khi có nhiều dữ liệu
                  <Brush 
                    dataKey="date" 
                    height={30} 
                    stroke="#60a5fa"
                    fill="rgba(39, 39, 42, 0.5)"
                    startIndex={Math.max(0, data.length - 20)} // Mặc định hiển thị 20 điểm cuối
                    endIndex={data.length - 1}
                  >
                    {/* Biểu đồ nhỏ bên trong Brush */}
                    <AreaChart>
                       <Area dataKey="value" stroke="#818cf8" fill="#818cf8" fillOpacity={0.3} />
                    </AreaChart>
                  </Brush>
                )}
                </AreaChart>
              </ResponsiveContainer>
            </ClientOnly>
          )}
        </div>
      </div>
    </WidgetCard>
  );
};

export default ProfileOverview;