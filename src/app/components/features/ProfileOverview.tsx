"use client";

import React, { useState, useEffect } from 'react';
import { AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, Brush  } from 'recharts';
import WidgetCard from '../ui/WidgetCard';
import { TrendingUp, MoreHorizontal, Loader2 } from 'lucide-react';
import { ClientOnly } from '@/app/components/helpers/ClientOnly';

// --- HÀM TẠO DỮ LIỆU MẪU ---
const generateChartData = (timeframe: '1D' | '1W' | '1M' | '6M' | '1Y') => {
  const data = [];
  const now = new Date();
  let lastValue = 150000;

  switch (timeframe) {
    case '1D':
      for (let i = 23; i >= 0; i--) {
        const date = new Date();
        date.setHours(now.getHours() - i);
        const change = (Math.random() - 0.5) * 5000;
        lastValue = Math.max(140000, lastValue + change);
        data.push({
          date: date.toLocaleTimeString('en-US', { hour: 'numeric', hour12: true }).replace(' ', ''),
          value: Math.round(lastValue)
        });
      }
      break;
    
    case '6M':
      for (let i = 5; i >= 0; i--) {
        const monthDate = new Date(now.getFullYear(), now.getMonth() - i, 1);
        const monthName = monthDate.toLocaleDateString('en-US', { month: 'short' });
        
        let change1 = (Math.random() - 0.48) * (lastValue * 0.1);
        lastValue = Math.max(40000, lastValue + change1);
        data.push({ date: `1st ${monthName}`, value: Math.round(lastValue) });
        
        let change2 = (Math.random() - 0.48) * (lastValue * 0.1);
        lastValue = Math.max(40000, lastValue + change2);
        data.push({ date: `15th ${monthName}`, value: Math.round(lastValue) });
      }
      break;

    case '1Y':
      for (let i = 11; i >= 0; i--) {
        const date = new Date(now.getFullYear(), now.getMonth() - i, 1);
        const change = (Math.random() - 0.48) * (lastValue * 0.1);
        lastValue = Math.max(40000, lastValue + change);
        data.push({ date: date.toLocaleDateString('en-US', { month: 'short' }), value: Math.round(lastValue) });
      }
      break;

    default: // 1W, 1M
      let numPoints = timeframe === '1W' ? 7 : 30;
      for (let i = numPoints - 1; i >= 0; i--) {
        const date = new Date();
        date.setDate(now.getDate() - i);
        const change = (Math.random() - 0.48) * (lastValue * 0.05);
        lastValue = Math.max(100000, lastValue + change);
        data.push({ date: date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' }), value: Math.round(lastValue) });
      }
      break;
  }
  return data;
};

// --- COMPONENT TOOLTIP TÙY CHỈNH ---
const CustomTooltip = ({ active, payload, label }: any) => {
  if (active && payload && payload.length) {
    return (
      <div className="tooltip-with-tail mb-3">
        <div className="bg-zinc-800/90 backdrop-blur-sm p-3 rounded-lg border border-zinc-700 shadow-lg min-w-[150px]">
          <div className="flex justify-between items-center text-xs text-gray-400">
            <span>{`${label} 2024`}</span>
            <MoreHorizontal size={14} />
          </div>
          <p className="text-xl font-bold text-white mt-1">
            {`$${payload[0].value.toLocaleString()}`}
          </p>
          <div className="flex items-center gap-x-1 text-xs bg-green-900/50 text-green-400 border border-green-700/60 rounded-full px-2 py-0.5 mt-2 w-fit">
            <TrendingUp size={14} />
            +3.5%
          </div>
        </div>
      </div>
    );
  }
  return null;
};

// --- COMPONENT ĐIỂM ACTIVE TÙY CHỈNH ---
const CustomActiveDot = (props: any) => {
  const { cx, cy } = props;
  if (!cx || !cy) return null;
  return (
    <g>
      <line x1={cx} y1={cy} x2={cx} y2={200} stroke="#2563eb" strokeWidth={1} strokeDasharray="3 3" />
      <circle cx={cx} cy={cy} r="6" stroke="#0f172a" strokeWidth={4} fill="#2563eb" />
    </g>
  );
};

// --- COMPONENT CHÍNH ---
const ProfileOverview = () => {
  const timeFilters: ('1D' | '1W' | '1M' | '6M' | '1Y')[] = ['1D', '1W', '1M', '6M', '1Y'];
  const [activeFilter, setActiveFilter] = useState<'1D' | '1W' | '1M' | '6M' | '1Y'>('6M');
  const [data, setData] = useState(generateChartData('6M'));
   const [isLoading, setIsLoading] = useState(false); 
  // useEffect để thay đổi dữ liệu khi lọc
   useEffect(() => {
    setIsLoading(true);
    const timer = setTimeout(() => {
      setData(generateChartData(activeFilter));
      setIsLoading(false);
    }, 500); // Mô phỏng độ trễ 0.5s

    return () => clearTimeout(timer);
  }, [activeFilter]);


  // useEffect cho realtime
  useEffect(() => {
    if (activeFilter !== '1D') return;
    const interval = setInterval(() => {
      setData(prevData => {
        const lastDataPoint = prevData[prevData.length - 1];
        const change = (Math.random() - 0.5) * 5000;
        const newValue = Math.max(20000, lastDataPoint.value + change);
        const newPoint = {
          date: new Date().toLocaleTimeString('en-US', { hour: 'numeric', hour12: true }).replace(' ', ''),
          value: Math.round(newValue),
        };
        return [...prevData.slice(1), newPoint];
      });
    }, 2000);
    return () => clearInterval(interval);
  }, [activeFilter]);

  const formatYAxis = (tickItem: number) => tickItem > 0 ? `${Math.round(tickItem / 1000)}k` : '0';

  let xAxisInterval: number | 'preserveStartEnd' = 'preserveStartEnd';
  if (activeFilter === '6M') {
    xAxisInterval = 1; // Hiển thị 1st, 15th,...
  } else if (activeFilter === '1M') {
    xAxisInterval = 6; // Hiển thị mỗi 7 ngày
  }

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
              <AreaChart data={data} margin={{ top: 20, right: 20, left: -20, bottom: 0 }}>
                <defs>
                  <linearGradient id="profileOverviewGradient" x1="0" y1="0" x2="0" y2="1">
                    <stop offset="5%" stopColor="#2563eb" stopOpacity={0.4}/>
                    <stop offset="100%" stopColor="#2563eb" stopOpacity={0}/>
                  </linearGradient>
                </defs>
                <CartesianGrid vertical={false} stroke="#404040" />
                <XAxis 
                  dataKey="date" 
                  tick={{ fill: '#a3a3a3', fontSize: 12 }} 
                  dy={10} 
                  tickLine={false} 
                  axisLine={false}
                  interval={xAxisInterval} 
                />
                <YAxis tickFormatter={formatYAxis} tick={{ fill: '#a3a3a3', fontSize: 12 }} tickLine={false} axisLine={false} />
                <Tooltip 
                  content={<CustomTooltip />} 
                  cursor={false}
                  position={{ y: 0 }}
                  isAnimationActive={false}
                />
                <Area 
                  type="monotone" 
                  dataKey="value" 
                  stroke="#60a5fa" 
                  strokeWidth={2}
                  fill="url(#profileOverviewGradient)" 
                  activeDot={<CustomActiveDot />}
                />
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