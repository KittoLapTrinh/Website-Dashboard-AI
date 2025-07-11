// src/components/widgets/RealTimeRecruiting.tsx
import React from 'react';
import { Apple, ArrowUpDown } from 'lucide-react'; // Chỉ import những gì cần
import WidgetCard from '../ui/WidgetCard';
import MiniTrendChart from '../ui/MiniTrendChart'; // Import component biểu đồ nhỏ

import { SiTesla, SiGoogle, SiNvidia, SiApple } from 'react-icons/si';
// Dữ liệu cho các nút lọc
const filters = ['All', 'Blockchain', 'AI', 'IOT'];
const activeFilter = 'All';

// Dữ liệu cho bảng - cấu trúc nhất quán
const recruitingData = [
  {
    icon: <SiTesla size={20} />, // Biểu tượng cho foundation
    foundation: 'FPT',
    position: 'Blockchain Engineer',
    field: 'Blockchain',
    salary: '$26,000.21',
    form: 'Fulltime',
    trend: 'up' as const,
  },
  {
    icon: <Apple size={20} />, // Bạn có thể thay bằng icon Google
    foundation: 'Google',
    position: 'AI Research Intern',
    field: 'AI',
    salary: '$32,000.21',
    form: 'Remote',
    trend: 'down' as const,
  },
];

// Component chính đã được sửa lỗi
const RealTimeRecruiting = () => {
  return (
    <WidgetCard>
      <div className="flex flex-col h-full">
        {/* === HEADER CỦA WIDGET === */}
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-xl font-bold text-white">Real-time recruiting</h2>
          <div className="flex items-center gap-x-2">
            {filters.map((filter) => (
              <button
                key={filter}
                className={`text-sm rounded-full px-4 py-1.5 transition-colors ${
                  activeFilter === filter
                    ? 'bg-blue-600 text-white'
                    : 'border border-gray-600 text-gray-300 hover:bg-gray-700/50'
                }`}
              >
                {filter}
              </button>
            ))}
          </div>
        </div>

        {/* === BẢNG DỮ LIỆU === */}
        <div className="space-y-2">
          {/* Header của Bảng */}
          <div className="grid grid-cols-6 gap-x-4 px-4 pb-3 border-b border-white-800">
            {['Foundation', 'Job Position', 'Field', 'Salary', 'Form', 'Update'].map((header) => (
              <div key={header} className="text-xs text-white font-semibold flex items-center gap-1 cursor-pointer hover:text-white">
                {header}
                <ArrowUpDown size={12} />
              </div>
            ))}
          </div>

          {/* Body của Bảng */}
          <div className="space-y-3 pt-2">
            {recruitingData.map((row, index) => (
              <div
                key={index}
                className="grid grid-cols-6 gap-x-4 items-center rounded-lg p-4 
                           border-b border-white-800/80 hover:bg-zinc-900/50 transition-colors"
              >
                {/* Cột 1: Foundation */}
                <div className="flex items-center gap-x-3 text-white text-sm">
                  {row.icon}
                  <span>{row.foundation}</span>
                </div>
                {/* Cột 2: Job Position */}
                <div className="text-white text-sm">{row.position}</div>
                {/* Cột 3: Field */}
                <div className="text-gray-300 text-sm">{row.field}</div>
                {/* Cột 4: Salary */}
                <div className="text-gray-300 text-sm">{row.salary}</div>
                {/* Cột 5: Form */}
                <div className="text-gray-300 text-sm">{row.form}</div>
                {/* Cột 6: Update Chart */}
                <div>
                  <MiniTrendChart trend={row.trend} />
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </WidgetCard>
  );
};

export default RealTimeRecruiting;