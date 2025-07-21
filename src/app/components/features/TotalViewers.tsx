"use client";

import React, { useState, useEffect, useRef } from 'react';
import { ChevronDown, TrendingUp } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';
import Image, { StaticImageData } from 'next/image';
// `TotalViewersData` giờ đã có kiểu đúng cho `return`
import { type TotalViewersData } from '@/app/lib/dashboard-types';
import TotalviewersBackground from '@/app/assets/image/Totalview.png';
import { useTotalViewersData } from '@/app/hooks/useTotalViewersData';

type TimeFilter = '1D' | '1W' | '1M' | '6M' | '1Y';

const TotalViewers = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [activeFilter, setActiveFilter] = useState<TimeFilter>('6M');
  const dropdownRef = useRef<HTMLDivElement>(null);
  
  // Hook vẫn được gọi như cũ
  const { data, isLoading, error, refetch } = useTotalViewersData(activeFilter);
  const filters: TimeFilter[] = ['1D', '1W', '1M', '6M', '1Y'];

  // Logic dropdown không thay đổi
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, [dropdownRef]);

  return (
    <WidgetCard className="relative overflow-hidden" transparent>
      <Image
        src={TotalviewersBackground as StaticImageData}
        alt="Total Viewers Background"
        fill
        className="object-cover z-0"
      />
      <div className="relative z-10 h-full flex flex-col justify-between">
        {/* Phần Header không đổi */}
        <div className="flex justify-between items-start">
          <p className="text-gray-400 text-base">Total Viewers</p>
          <div className="relative" ref={dropdownRef}>
            <button
              onClick={() => setIsOpen(!isOpen)}
              className="flex items-center gap-x-2 text-sm text-gray-200 border border-gray-600 rounded-full px-4 py-1.5 hover:border-gray-400 transition-colors"
            >
              <span>{activeFilter}</span>
              <ChevronDown size={16} className={`text-gray-400 transition-transform duration-200 ${isOpen ? 'rotate-180' : ''}`} />
            </button>
            {isOpen && (
              <div className="absolute top-full right-0 mt-2 w-32 origin-top-right bg-zinc-900 border border-zinc-700 rounded-xl shadow-lg z-20">
                <div className="p-1">
                  {filters.map((filter) => (
                    <button
                      key={filter}
                      onClick={() => { setActiveFilter(filter); setIsOpen(false); }}
                      className={`w-full text-left rounded-md px-3 py-1.5 text-sm font-medium transition-colors ${activeFilter === filter ? 'bg-blue-600 text-white' : 'text-gray-300 hover:bg-zinc-800'}`}
                    >
                      {filter}
                    </button>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>

        {/* Phần hiển thị dữ liệu (Đã được sửa) */}
        {isLoading ? (
          <div className="animate-pulse space-y-4">
            <div className="h-10 bg-zinc-700/50 rounded-lg w-3/4"></div>
            <div className="h-6 bg-zinc-700/50 rounded-lg w-1/2"></div>
          </div>
        ) : error || !data ? (
          <div className="text-red-400">Error loading data.</div>
        ) : (
          <div>
            <p className="text-4xl font-bold text-white">
              {/* ✨ Hiển thị viewers không đổi */}
              {data.viewers.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
            </p>
            <div className="flex items-center gap-x-2 mt-2">
              <span className="text-xs text-gray-400">Return</span>
              
              {/* ✨ SỬA LẠI Ở ĐÂY: Truy cập `data.return.status` và `data.return.value` */}
              <span className={`flex items-center gap-x-1 text-xs rounded-full px-1.5 py-px
                ${data.return.status === 'increase' ? 'bg-green-900/50 text-green-400 border border-green-700/60' : 'bg-red-900/50 text-red-400 border border-red-700/60'}`}
              >
                <TrendingUp size={14} className={data.return.status === 'increase' ? '' : '-scale-y-100'}/>
                {data.return.value >= 0 ? '+' : ''}{data.return.value}%
              </span>
            </div>
          </div>
        )}
      </div>
    </WidgetCard>
  );
};

export default TotalViewers;