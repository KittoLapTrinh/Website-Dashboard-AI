// src/components/widgets/SupportFund.tsx
import React from 'react';
import { Sun, Moon, Droplets } from 'lucide-react'; // Sử dụng các icon phù hợp
import WidgetCard from '../ui/WidgetCard';

// Dữ liệu cho các ngày
const days = ['Mon 20', 'Tue 21', 'Wed 22', 'Thu 23'];
const activeDay = 'Wed 22';

// Dữ liệu cho danh sách quỹ
const fundData = [
  { name: 'Albufin', price: 26, icon: <Sun size={16} />, count: 1 },
  { name: 'Vitamin D', price: 26, icon: <Droplets size={16} />, count: 2 },
  { name: 'Omega 3', price: 26, icon: <Moon size={16} />, count: 2 },
];

const SupportFund = () => {
  return (
    <WidgetCard>
      <div className="flex flex-col h-full">
        {/* === HEADER CỦA WIDGET === */}
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold text-white">Support Fund</h2>
        </div>

        {/* === THANH CHỌN NGÀY === */}
        <div className="flex justify-between items-center text-sm text-gray-400 mb-6">
          {days.map((day) => (
            <span
              key={day}
              className={`${activeDay === day ? 'text-white font-semibold' : ''}`}
            >
              {day}
            </span>
          ))}
        </div>

        {/* === DANH SÁCH CÁC MỤC === */}
        <div className="space-y-3">
          {fundData.map((item) => (
            <div
              key={item.name}
              className="flex items-center justify-between bg-zinc-800/80 hover:bg-zinc-700/80 p-3 rounded-lg transition-colors"
            >
              <div className="flex items-center gap-x-3">
                {/* Placeholder cho ảnh sản phẩm */}
                <div className="w-8 h-8 rounded-full bg-gray-600 flex-shrink-0"></div>
                <div>
                  <p className="text-white text-sm font-medium">{item.name}</p>
                  <p className="text-gray-500 text-xs">${item.price}</p>
                </div>
              </div>
              <button className="flex items-center gap-x-1.5 bg-blue-600 text-white rounded-full px-3 py-1 text-sm">
                {item.icon}
                <span>{item.count}</span>
              </button>
            </div>
          ))}
        </div>
      </div>
    </WidgetCard>
  );
};

export default SupportFund;