"use client";

import React, { useState } from 'react';
import { Sun, Moon } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';

const days = ['Mon 20', 'Tue 21', 'Wed 22', 'Thu 23'];
const fundData = [
  { name: 'Albufin', price: 26, icon: <Sun size={18} />, count: 1, avatarColor: 'bg-white' },
  { name: 'Vitamin D', price: 26, icon: <Sun size={18} />, count: 2, avatarColor: 'bg-gray-600' },
  { name: 'Omega 3', price: 26, icon: <Moon size={18} />, count: 2, avatarColor: 'bg-gray-600' },
];

const SupportFund = () => {
  const [activeDay, setActiveDay] = useState('Wed 22');

  return (
    <WidgetCard>
      <div className="flex flex-col h-full">
        {/* Header */}
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold text-white">Support Fund</h2>
        </div>
        <div className="flex justify-between items-center text-sm text-gray-400 mb-8">
          {days.map((day) => (
            <button key={day} onClick={() => setActiveDay(day)} className={`font-medium transition-colors ${activeDay === day ? 'text-white' : 'text-gray-500 hover:text-gray-300'}`}>
              {day}
            </button>
          ))}
        </div>

        {/* Danh sách các mục */}
        <div className="relative flex-grow min-h-[250px]">
          
          {/* === CÁC ĐƯỜNG KẺ === */}

          {/* ✨ Đường kẻ ngang 1 */}
          <div className="absolute top-7 left-6 w-[90%] h-px bg-dashed-horizontal"></div>
          {/* ✨ Đường kẻ ngang 2 */}
          <div className="absolute top-1/2 left-6 w-[90%] h-px bg-dashed-horizontal"></div>
          {/* ✨ Đường kẻ ngang 3 */}
          <div className="absolute bottom-7 left-6 w-[90%] h-px bg-dashed-horizontal"></div>


          {/* === CÁC ITEM === */}
          {/* Item 1 */}
          <div className="absolute flex items-center">
      
            <div className="ml-4 bg-zinc-800 rounded-full p-2 flex items-center gap-x-3 w-[100%]">
              <div className={`w-10 h-10 rounded-full ${fundData[0].avatarColor}`}></div>
              <div>
                <p className="text-white font-semibold">{fundData[0].name}</p>
                <p className="text-gray-500 text-xs font-medium">${fundData[0].price}</p>
              </div>
              <button className="flex items-center gap-x-2 bg-blue-600 text-white rounded-full px-3 py-1.5 font-medium">
                {fundData[0].icon}<span>{fundData[0].count}</span>
              </button>
            </div>
          </div>

          {/* Item 2 */}
          <div className="absolute top-1/2 -translate-y-1/2 left-10 flex items-center">  
             <div className="ml-4 bg-zinc-800 rounded-full p-2 flex items-center gap-x-3 w-[100%]">
              <div className={`w-10 h-10 rounded-full ${fundData[1].avatarColor}`}></div>
              <div>
                <p className="text-white font-semibold">{fundData[1].name}</p>
                <p className="text-gray-500 text-xs font-medium">${fundData[1].price}</p>
              </div>
              <button className="flex items-center gap-x-2 bg-blue-600 text-white rounded-full px-3 py-1.5 font-medium">
                {fundData[1].icon}<span>{fundData[1].count}</span>
              </button>
            </div>
          </div>

          {/* Item 3 */}
          <div className="absolute bottom-0 left-25 flex items-center">
             <div className="ml-4 bg-zinc-800 rounded-full p-2 flex items-center gap-x-3 w-[100%]">
              <div className={`w-10 h-10 rounded-full ${fundData[2].avatarColor}`}></div>
              <div>
                <p className="text-white font-semibold">{fundData[2].name}</p>
                <p className="text-gray-500 text-xs font-medium">${fundData[2].price}</p>
              </div>
              <button className="flex items-center gap-x-2 bg-blue-600 text-white rounded-full px-3 py-1.5 font-medium">
                {fundData[2].icon}<span>{fundData[2].count}</span>
              </button>
            </div>
          </div>
          
        </div>
      </div>
    </WidgetCard>
  );
};

export default SupportFund;