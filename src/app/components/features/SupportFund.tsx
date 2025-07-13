"use client";

import React, { useState, useMemo } from 'react';
import { Sun, Moon } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';

// --- DỮ LIỆU MẪU LỚN HƠN ---
const allFundData = {
  'Mon 20': [
    { name: 'Albufin', price: 26, icon: <Sun size={18} />, count: 1, avatarColor: 'bg-white' },
    { name: 'Vitamin C', price: 15, icon: <Sun size={18} />, count: 1, avatarColor: 'bg-gray-600' },
  ],
  'Tue 21': [
    { name: 'Vitamin D', price: 26, icon: <Sun size={18} />, count: 2, avatarColor: 'bg-gray-600' },
    { name: 'Omega 3', price: 26, icon: <Moon size={18} />, count: 2, avatarColor: 'bg-gray-600' },
    { name: 'Magnesium', price: 20, icon: <Moon size={18} />, count: 1, avatarColor: 'bg-gray-600' },
  ],
  'Wed 22': [
    { name: 'Albufin', price: 26, icon: <Sun size={18} />, count: 1, avatarColor: 'bg-white' },
    { name: 'Vitamin D', price: 26, icon: <Sun size={18} />, count: 2, avatarColor: 'bg-gray-600' },
    { name: 'Omega 3', price: 26, icon: <Moon size={18} />, count: 2, avatarColor: 'bg-gray-600' },
    { name: 'Iron', price: 18, icon: <Sun size={18} />, count: 1, avatarColor: 'bg-gray-600' },
    { name: 'Zinc', price: 12, icon: <Sun size={18} />, count: 1, avatarColor: 'bg-gray-600' },
  ],
  'Thu 23': [
    { name: 'Omega 3', price: 26, icon: <Moon size={18} />, count: 2, avatarColor: 'bg-gray-600' },
  ],
};

const days = Object.keys(allFundData);

const SupportFund = () => {
  const [activeDay, setActiveDay] = useState('Wed 22');
  
  // Lọc dữ liệu dựa trên ngày được chọn
  const displayedData = useMemo(() => {
    return allFundData[activeDay as keyof typeof allFundData] || [];
  }, [activeDay]);

  // Mảng class thụt vào
  const indentationClasses = ['pl-4', 'pl-12', 'pl-20', 'pl-28', 'pl-36'];

  return (
    <WidgetCard>
      <div className="flex flex-col h-60">
        {/* Header */}
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-2xl font-bold text-white">Support Fund</h2>
        </div>
        <div className="flex justify-between items-center text-sm text-gray-400 mb-8">
          {days.map((day) => (
            <button key={day} onClick={() => setActiveDay(day)} className={`font-medium transition-colors ${activeDay === day ? 'text-white' : 'text-gray-500 hover:text-gray-300'}`}>
              {day}
            </button>
          ))}
        </div>

        {/* === DANH SÁCH CÁC MỤC (LAYOUT MỚI) === */}
        {/* ✨ Thêm class để có scroll */}
        <div className="flex-grow flex flex-col gap-y-4 overflow-y-auto pr-2">
          {displayedData.map((item, index) => (
            <div key={item.name} className={`flex items-center gap-x-4 ${indentationClasses[index % indentationClasses.length]}`}>
              <div className="bg-zinc-800 rounded-full p-2 flex items-center gap-x-3 w-fit">
                <div className={`w-10 h-10 rounded-full ${item.avatarColor}`}></div>
                <div>
                  <p className="text-white font-semibold">{item.name}</p>
                  <p className="text-gray-500 text-xs font-medium">${item.price}</p>
                </div>
                <button className="flex items-center gap-x-2 bg-blue-600 text-white rounded-full px-3 py-1.5 font-medium">
                  {item.icon}
                  <span>{item.count}</span>
                </button>
              </div>
              <div className="flex-grow h-px bg-dashed-horizontal"></div>
            </div>
          ))}
        </div>
      </div>
    </WidgetCard>
  );
};

export default SupportFund;