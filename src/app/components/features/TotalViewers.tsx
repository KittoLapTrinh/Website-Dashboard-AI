"use client";

import React, { useState, useEffect, useRef } from 'react';
import { ChevronDown, TrendingUp } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';
import Image, { StaticImageData } from 'next/image';

// Giả sử bạn import ảnh nền, đảm bảo đường dẫn đúng
import TotalviewersBackground from '@/app/assets/image/Totalview.png';

// --- DỮ LIỆU MẪU CHO CÁC BỘ LỌC ---
const viewersData = {
  '1D': { viewers: 12304.11, return: 3.5 },
  '1W': { viewers: 89450.75, return: 2.1 },
  '1M': { viewers: 350123.40, return: 5.8 },
  '6M': { viewers: 1850678.90, return: 12.3 },
  '1Y': { viewers: 4500000.00, return: 25.1 },
};
type TimeFilter = keyof typeof viewersData;

const TotalViewers = () => {
  // === STATE VÀ LOGIC CHO DROPDOWN ===
  const [isOpen, setIsOpen] = useState(false);
  const [activeFilter, setActiveFilter] = useState<TimeFilter>('6M');
  const [data, setData] = useState(viewersData[activeFilter]);
  const dropdownRef = useRef<HTMLDivElement>(null);
  const [isLoading, setIsLoading] = useState(false);
  const filters: TimeFilter[] = ['1D', '1W', '6M', '1Y'];

  // Cập nhật dữ liệu khi bộ lọc thay đổi
    useEffect(() => {
    // Không cần set loading nếu là lần render đầu tiên
    if (data !== viewersData[activeFilter]) {
      setIsLoading(true);
      const timer = setTimeout(() => {
        setData(viewersData[activeFilter]);
        setIsLoading(false);
      }, 500); // 0.5 giây để mô phỏng tải dữ liệu
      return () => clearTimeout(timer);
    }
  }, [activeFilter, data]);

  // Đóng dropdown khi click ra ngoài
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, [dropdownRef]);

   let xAxisInterval: number | 'preserveStartEnd' = 'preserveStartEnd';
  if (activeFilter === '6M') xAxisInterval = 29; // ~ 1 nhãn mỗi tháng
  else if (activeFilter === '1M') xAxisInterval = 6; // ~ 1 nhãn mỗi tuần
  else if (activeFilter === '1D') xAxisInterval = 3; 
  return (
    <WidgetCard className="relative overflow-hidden" transparent>
      <Image
        src={TotalviewersBackground as StaticImageData}
        alt="Total Viewers Background"
        fill
        className="object-cover z-0"
      />
      <div className="relative z-10 h-full flex flex-col justify-between">
        <div className="flex justify-between items-start">
          <p className="text-gray-400 text-base">Total Viewers</p>
          
          {/* ✨ DROPDOWN MENU ✨ */}
          <div className="relative" ref={dropdownRef}>
            <button
              onClick={() => setIsOpen(!isOpen)}
              className="flex items-center gap-x-2 text-sm text-gray-200 border border-gray-600 
                         rounded-full px-4 py-1.5 
                         hover:border-gray-400 transition-colors"
            >
              <span>{activeFilter}</span>
              <ChevronDown size={16} className={`text-gray-400 transition-transform duration-200 ${isOpen ? 'rotate-180' : ''}`} />
            </button>

            {/* Menu xổ xuống */}
            {isOpen && (
              <div 
                className="absolute top-full right-0 mt-2 w-32 
                           origin-top-right bg-zinc-900 border border-zinc-700 
                           rounded-xl shadow-lg z-20
                           transition-all duration-200 ease-out
                           transform opacity-100 scale-100"
              >
                <div className="p-1">
                  {filters.map((filter) => (
                    <button
                      key={filter}
                      onClick={() => {
                        setActiveFilter(filter);
                        setIsOpen(false);
                      }}
                      className={`w-full text-left rounded-md px-3 py-1.5 text-sm font-medium transition-colors
                        ${activeFilter === filter 
                          ? 'bg-blue-600 text-white' 
                          : 'text-gray-300 hover:bg-zinc-800'
                        }`}
                    >
                      {filter}
                    </button>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>

         {isLoading ? (
          // Giao diện khi đang tải
          <div className="animate-pulse space-y-4">
            <div className="h-10 bg-zinc-700/50 rounded-lg w-3/4 mb-12"></div>
            <div className="h-6 bg-zinc-700/50 rounded-lg w-1/2 mb-2"></div>
          </div>
        ) : (
          // Giao diện khi đã tải xong
          <div className="h-30">
            <p className="text-4xl font-bold text-white">
              {data.viewers.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
            </p> 
            <div className="flex items-center gap-x-2 mt-12">
              <span className="text-xs text-gray-400">Return</span>
              <span className="flex items-center gap-x-1 text-xs bg-green-900/50 text-green-400 border border-green-700/60 rounded-full px-1.5 py-px">
                <TrendingUp size={14} />
                +{data.return}%
              </span>
            </div>
          </div>
        )}
      </div>
    </WidgetCard>
  );
};

export default TotalViewers;