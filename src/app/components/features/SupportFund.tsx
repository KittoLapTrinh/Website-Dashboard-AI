"use client";

import React, { useState, useMemo , useEffect} from 'react';
import { Sun, Moon } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';
import { SupportFundSkeleton } from '../ui/SupportFundSkeleton';

// --- HÀM TẠO DỮ LIỆU MẪU CHO 30 NGÀY ---
const generateAllFundData = () => {
  const data: { [key: string]: any[] } = {};
  const now = new Date();
  const today = now.getDate();
  const dayOfWeek = now.toLocaleDateString('en-US', { weekday: 'short' });

  for (let i = 0; i < 30; i++) {
    const date = new Date();
    date.setDate(now.getDate() - i);
    const dayKey = `${date.toLocaleDateString('en-US', { weekday: 'short' })} ${date.getDate()}`;
    
    const itemCount = Math.floor(Math.random() * 4) + 1; // 1 đến 4 item mỗi ngày
    data[dayKey] = Array.from({ length: itemCount }, () => ({
      name: ['Albufin', 'Vitamin D', 'Omega 3', 'Iron', 'Zinc'][Math.floor(Math.random() * 5)],
      price: Math.floor(Math.random() * 20) + 10,
      icon: Math.random() > 0.5 ? <Sun size={18} /> : <Moon size={18} />,
      count: Math.floor(Math.random() * 3) + 1,
      avatarColor: Math.random() > 0.3 ? 'bg-gray-600' : 'bg-white',
    }));
  }
  return { data, initialDay: `${dayOfWeek} ${today}` };
};

const { data: allFundData, initialDay } = generateAllFundData();
const days = Object.keys(allFundData).reverse();

const SupportFund = () => {
  const [activeDay, setActiveDay] = useState(initialDay);
  const [isLoading, setIsLoading] = useState(true);
  const displayedData = useMemo(() => {
    return allFundData[activeDay as keyof typeof allFundData] || [];
  }, [activeDay]);

     useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
    }, 500); // Tải trong 1 giây
    return () => clearTimeout(timer);
  }, []);

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
          <div className="flex items-center space-x-6 whitespace-nowrap px-1">
            {days.map((day) => (
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