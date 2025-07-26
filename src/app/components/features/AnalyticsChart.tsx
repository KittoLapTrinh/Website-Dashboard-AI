"use client";
import React, { useState, useRef, useEffect,useMemo  } from 'react';
import { AreaChart, Area, Tooltip, ResponsiveContainer, YAxis, XAxis, ReferenceLine } from 'recharts';
import { ChevronDown, TrendingUp, Loader2 } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';
import { useAnalyticsChartData } from '@/app/hooks/useAnalyticsChartData';
import { type DataPoint } from '@/app/lib/dashboard-types';
// ✨ ĐÃ XÓA HÀM `generateChartData` KHÔNG CÒN CẦN THIẾT ✨

// Component Tooltip tùy chỉnh để có vị trí chính xác
const CustomTooltip = ({ active, payload }: any) => {
  if (active && payload && payload.length) {
    // Lấy dữ liệu từ payload của Recharts
    const dataPoint = payload[0].payload;
    const value = dataPoint.value; // Giả sử value là phần trăm

    return (
      <div className="tooltip-with-tail">
        <div className="bg-zinc-800/90 backdrop-blur-sm px-3 py-1.5 rounded-lg flex items-center justify-center gap-x-2 shadow-lg">
          <span className="font-bold text-white">{`${value}%`}</span>
          {/* (Tùy chọn) Có thể thêm logic hiển thị TrendingUp/Down nếu có dữ liệu 'change' */}
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
  type TimeFilterType = 'Weekly' | 'Monthly' | 'Yearly';
  const [isOpen, setIsOpen] = useState(false);
  const [activeFilter, setActiveFilter] = useState<TimeFilterType>('Weekly');
  const filters: TimeFilterType[] = ['Weekly', 'Monthly', 'Yearly'];
  const dropdownRef = useRef<HTMLDivElement>(null);

  // Hook này giờ sẽ cung cấp dữ liệu thật và cập nhật liên tục
  const { data, isLoading, error, refetch } = useAnalyticsChartData(activeFilter);

  const WINDOW_SIZE = 10;

  const visibleData = useMemo(() => {
    if (!data || data.length === 0) return [];
    // Sử dụng slice để lấy N phần tử cuối cùng của mảng `data`
    return data.slice(Math.max(data.length - WINDOW_SIZE, 0));
  }, [data]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) setIsOpen(false);
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, [dropdownRef]);

  return (
    <WidgetCard>
      <div className="flex flex-col h-full">
        {/* Header */}
        <div className="flex justify-between items-start mb-4">
          <div>
            <h2 className="text-2xl font-bold text-white">Analytics</h2>
            {/* (Tùy chọn) Có thể làm cho ngày tháng này động */}
            <p className="text-sm text-[#FFFFFF] mt-1">Real-time Data</p>
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
              <div className="absolute top-full right-0 mt-2 w-26 bg-zinc-800 border border-gray-700 rounded-lg shadow-lg z-20">
                <ul className="py-1">
                  {filters.map((filter) => (
                    <li key={filter}>
                      <button
                        onClick={() => {
                          setActiveFilter(filter);
                          setIsOpen(false);
                        }}
                        className={`w-full text-left px-4 py-2 text-sm transition-colors ${
                          activeFilter === filter 
                            ? 'text-[#409BEE]'
                            : 'text-gray-300 hover:bg-zinc-700'
                        }`}
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
        <div className="flex-grow mt-4 h-48 flex items-center justify-center">
          {isLoading && data.length === 0 ? (
            <Loader2 className="animate-spin text-gray-500" size={32} />
          ) : error ? (
            <p className="text-xs text-red-500">Error loading data.</p>
          ) : (
            <ResponsiveContainer width="100%" height="100%">
              <AreaChart
                data={visibleData} // ✨ SỬ DỤNG DỮ LIỆU THẬT TỪ HOOK ✨
                margin={{ top: 10, right: 10, left: 10, bottom: 0 }}
              >
                <defs>
                  <linearGradient id="analyticsChartGradient" x1="0" y1="0" x2="0" y2="1">
                    <stop offset="5%" stopColor="#409BEE" stopOpacity={0.8}/>
                    <stop offset="95%" stopColor="#409BEE" stopOpacity={0}/>
                  </linearGradient>
                  <filter id="analyticsGlow" x="-100%" y="-100%" width="300%" height="300%">
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
                <XAxis dataKey="date" hide={true} />
                <Area 
                  type="monotone" 
                  dataKey="value"
                  stroke="#e5e7eb" 
                  strokeWidth={2}
                  fill="url(#analyticsChartGradient)"
                  fillOpacity={1}
                  isAnimationActive={false} // Tắt animation để cập nhật real-time mượt hơn
                  activeDot={<CustomActiveDot />}
                  style={{ filter: 'url(#analyticsGlow)' }}
                />
              </AreaChart>
            </ResponsiveContainer>
          )}
        </div>
      </div>
    </WidgetCard>
  );
};

export default AnalyticsChart;