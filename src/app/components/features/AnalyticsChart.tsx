"use client";
import React, { useState, useRef, useEffect } from 'react';
import { AreaChart, Area, Tooltip, ResponsiveContainer, YAxis } from 'recharts';
import { ChevronDown, TrendingUp } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';

// Dữ liệu mẫu cho biểu đồ
// const data = [
//   { name: 'Jan', uv: 40 }, { name: 'Feb', uv: 30 },
//   { name: 'Mar', uv: 60 }, { name: 'Apr', uv: 50 },
//   { name: 'May', uv: 90 }, { name: 'Jun', uv: 40 },
//   { name: 'Jul', uv: 55 }, { name: 'Aug', uv: 80 },
//   { name: 'Sep', uv: 30 }, { name: 'Oct', uv: 50 },
//   { name: 'Nov', uv: 20 }, { name: 'Dec', uv: 40 },
// ];

const generateChartData = (timeframe: 'Weekly' | 'Monthly' | 'Yearly') => {
  const data = [];
  const now = new Date();
  let numPoints;
  let pointInterval: 'day' | 'month' = 'day';
  
  switch (timeframe) {
    case 'Weekly': numPoints = 7; pointInterval = 'day'; break;
    case 'Monthly': numPoints = 30; pointInterval = 'day'; break;
    case 'Yearly': numPoints = 12; pointInterval = 'month'; break;
    default: numPoints = 7;
  }

  let lastValue = 80;
  for (let i = numPoints - 1; i >= 0; i--) {
    const date = new Date(now);
    if (pointInterval === 'day') date.setDate(now.getDate() - i);
    if (pointInterval === 'month') date.setMonth(now.getMonth() - i);

    const change = (Math.random() - 0.45) * 10;
    lastValue = Math.max(20, Math.min(100, lastValue + change));
    
    let dateLabel = date.toLocaleDateString('en-US', { weekday: 'short' });
    if (timeframe === 'Monthly') dateLabel = date.toLocaleDateString('en-US', { day: 'numeric' });
    if (timeframe === 'Yearly') dateLabel = date.toLocaleDateString('en-US', { month: 'short' });

    data.push({ name: dateLabel, uv: Math.round(lastValue) });
  }
  return data;
};


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
  const [isOpen, setIsOpen] = useState(false);
  const [activeFilter, setActiveFilter] = useState<'Weekly' | 'Monthly' | 'Yearly'>('Weekly');
  const filters: ('Weekly' | 'Monthly' | 'Yearly')[] = ['Weekly', 'Monthly', 'Yearly'];
  const dropdownRef = useRef<HTMLDivElement>(null);

   const [chartData, setChartData] = useState(() => generateChartData('Weekly'));
  useEffect(() => {
    setChartData(generateChartData(activeFilter));
  }, [activeFilter]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, [dropdownRef]);

  // const [chartData, setChartData] = useState<any>(null);
  // const handleMouseMove = (state: any) => {
  //   if (state.isTooltipActive && state.activePayload?.length) {
  //     setChartData({ payload: state.activePayload, coordinate: state.activeCoordinate });
  //   } else {
  //     setChartData(null);
  //   }
  // };
  // const handleMouseLeave = () => {
  //   setChartData(null);
  // };

  return (
    <WidgetCard>
      <div className="flex flex-col h-full">
        {/* Header */}
        <div className="flex justify-between items-start mb-4">
          <div>
            <h2 className="text-2xl font-bold text-white">Analytics</h2>
            <p className="text-sm text-[#FFFFFF] mt-1">April 2025</p>
          </div>
          <div className="relative" ref={dropdownRef}>
            {/* Nút bấm chính */}
            <button
              onClick={() => setIsOpen(!isOpen)}
              className="flex items-center gap-x-2 text-sm text-gray-300 border border-[#409BEE] rounded-full px-4 py-2 transition-colors"
            >
              {activeFilter}
              <ChevronDown size={16} className={`transition-transform duration-200 ${isOpen ? 'rotate-180' : ''}`} />
            </button>

            {/* Menu xổ xuống */}
            {isOpen && (
              <div 
                className="absolute top-full right-0 mt-2 w-26 
                           bg-zinc-800 border border-gray-700 
                           rounded-lg shadow-lg z-20"
              >
                <ul className="py-1">
                  {filters.map((filter) => (
                    <li key={filter}>
                      <button
                        onClick={() => {
                          setActiveFilter(filter);
                          setIsOpen(false);
                        }}
                        className={`w-full text-left px-4 py-2 text-sm transition-colors
                          ${
                            activeFilter === filter 
                              ? 'text-[#409BEE]' // Style cho nút active
                              : 'text-gray-300 hover:bg-zinc-700' // Style cho nút thường
                          }
                        `}
                      >
                        {filter}
                      </button>
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </div>
      
        </div>

        {/* Chart Area */}
        <div className="flex-grow mt-4 h-48">
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart
              data={chartData}
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