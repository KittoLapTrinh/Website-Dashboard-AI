"use client";

import React, { useState, useMemo , useEffect} from 'react';
import { Sun, Moon } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';
import { SupportFundSkeleton } from '../ui/skeletons/SupportFundSkeleton';
import { type FundItemData } from '@/app/lib/dashboard-types';
import { useSupportFundData } from '@/app/hooks/useSupportFundData';
// --- HÀM TẠO DỮ LIỆU MẪU CHO 30 NGÀY ---
const generateDays = (count: number) => {
  const daysArray = [];
  const now = new Date();
  for (let i = 0; i < count; i++) {
    const date = new Date();
    date.setDate(now.getDate() - i);
    // ✨ TẠO KEY THEO ĐỊNH DẠNG "Mon 23", "Sun 22",...
    daysArray.push(
      `${date.toLocaleDateString('en-US', { weekday: 'short' })} ${date.getDate()}`
    );
  }
  // Đảo ngược để ngày gần nhất ở cuối, sau đó lại đảo ngược để hiển thị
  // Logic này đảm bảo ngày hôm nay là lựa chọn cuối cùng trong mảng
  return { days: daysArray.reverse(), initialDay: daysArray[0] };
};

// ✨ Tạo ra 30 ngày và ngày khởi tạo là ngày hôm nay
const { days, initialDay } = generateDays(30);

const SupportFund = () => {
  const [activeDay, setActiveDay] = useState(initialDay);
  const { items: displayedData, isLoading, error, refetch } = useSupportFundData(activeDay);

  const indentationClasses = ['pl-4', 'pl-12', 'pl-20', 'pl-28', 'pl-36'];

  return (
    <WidgetCard>
      
      <div className="flex flex-col h-60"> {/* Thay đổi h-60 thành h-full */}
        {/* Header */}
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-2xl font-bold text-white">Support Fund</h2>
        </div>
        
        {/* === THANH CHỌN NGÀY CÓ SCROLL NGANG === */}
        <div className="overflow-x-auto hide-scrollbar mb-8">
          <div className="flex items-center space-x-6 whitespace-nowrap px-1 ">
            {[...days].reverse().map((day) => (
              <button 
                key={day} 
                onClick={() => setActiveDay(day)} 
                className={`font-medium transition-colors py-1
                  ${activeDay === day 
                    ? 'text-white border-b-2 border-white' 
                    : 'text-gray-500 hover:text-gray-300'
                  }`}
              >
                {day}
              </button>
            ))}
          </div>
        </div>

        {/* === DANH SÁCH CÁC MỤC VỚI SCROLL DỌC === */}
        <div className="flex-grow flex flex-col gap-y-4 overflow-y-auto custom-scrollbar pr-4">
          {isLoading ? (
            <SupportFundSkeleton />
          ) : displayedData.length > 0 ? (
            displayedData.map((item, index) => (
              <div key={`${item.name}-${index}`} className={`flex items-center gap-x-4 ${indentationClasses[index % indentationClasses.length]}`}>
                <div className="bg-zinc-800 rounded-full p-1 flex items-center gap-x-3 w-fit">
                  <div className={`w-8 h-8 rounded-full ${item.avatarColor}`}></div>
                  <div>
                    <p className="text-white font-semibold">{item.name}</p>
                  </div>
                  <div>
                    <p className="text-gray-500 text-xs font-medium">${item.price}</p>
                  </div>
                  <button className="flex items-center gap-x-2 bg-blue-600 text-white rounded-full px-3 py-1 font-medium">
                    {item.icon}
                    <span>{item.count}</span>
                  </button>
                </div>
                <div className="flex-grow h-px bg-dashed-horizontal"></div>
              </div>
            ))
          ) : (
            <div className="flex items-center justify-center h-full text-gray-500">
              No data for this day.
            </div>
          )}
        </div>
      </div>
    </WidgetCard>
  );
};

export default SupportFund;