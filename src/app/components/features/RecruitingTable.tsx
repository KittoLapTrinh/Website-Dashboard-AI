"use client";

import React, { useState, useMemo } from 'react';
import { ArrowUpDown } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';
import MiniTrendChart from '../ui/MiniTrendChart';
import { SiTesla, SiApple, SiGoogle, SiNvidia } from 'react-icons/si';

// --- DỮ LIỆU VÀ KIỂU DỮ LIỆU ---
type Field = 'Blockchain' | 'AI' | 'IOT' | 'Web Dev';
type Trend = 'up' | 'down';
interface RecruitData {
  icon: React.ReactNode;
  foundation: string;
  position: string;
  field: Field;
  salary: number;
  form: string;
  trend: Trend;
}

const allRecruitingData: RecruitData[] = [
  { icon: <SiNvidia size={24} className="text-gray-400"/>, foundation: 'Nvidia', position: 'Deep Learning Scientist', field: 'AI', salary: 55000.00, form: 'Remote', trend: 'up' },
  { icon: <SiTesla size={24} className="text-gray-400"/>, foundation: 'Tesla', position: 'Firmware Engineer', field: 'IOT', salary: 45000.00, form: 'Fulltime', trend: 'up' },
  { icon: <SiApple size={24} className="text-gray-400"/>, foundation: 'ConsenSys', position: 'Smart Contract Dev', field: 'Blockchain', salary: 42000.00, form: 'Fulltime', trend: 'down' },
  { icon: <SiGoogle size={22} className="text-gray-400"/>, foundation: 'Google', position: 'Frontend Developer', field: 'Web Dev', salary: 38000.50, form: 'Fulltime', trend: 'up' },
  { icon: <SiApple size={24} className="text-gray-400"/>, foundation: 'Apple', position: 'AI Research Intern', field: 'AI', salary: 32000.21, form: 'Remote', trend: 'down' },
  { icon: <SiTesla size={24} className="text-gray-400"/>, foundation: 'FPT', position: 'Blockchain Engineer', field: 'Blockchain', salary: 26000.21, form: 'Fulltime', trend: 'up' },
];

const filters: ('All' | Field)[] = ['All', 'Blockchain', 'AI', 'IOT'];

// --- COMPONENT CHÍNH ---
const RecruitingTable = () => {
  const [activeFilter, setActiveFilter] = useState<'All' | Field>('All');
  const [sortConfig, setSortConfig] = useState<{ key: keyof RecruitData; direction: 'ascending' | 'descending' } | null>({ key: 'salary', direction: 'descending' });

   const filteredData = useMemo(() => {
    let sortableItems = [...allRecruitingData];
    if (activeFilter !== 'All') {
      sortableItems = sortableItems.filter(item => item.field === activeFilter);
    }
    
    // 👇 BỎ COMMENT VÀ HOÀN THIỆN LOGIC SORT Ở ĐÂY 👇
    if (sortConfig !== null) {
      sortableItems.sort((a, b) => {
        const aValue = a[sortConfig.key];
        const bValue = b[sortConfig.key];

        // Kiểm tra nếu là số để so sánh số, ngược lại so sánh chuỗi
        if (typeof aValue === 'number' && typeof bValue === 'number') {
          if (aValue < bValue) {
            return sortConfig.direction === 'ascending' ? -1 : 1;
          }
          if (aValue > bValue) {
            return sortConfig.direction === 'ascending' ? 1 : -1;
          }
        } else {
          // So sánh chuỗi (đã chuyển về chữ thường để không phân biệt hoa/thường)
          const aStr = String(aValue).toLowerCase();
          const bStr = String(bValue).toLowerCase();
          if (aStr < bStr) {
            return sortConfig.direction === 'ascending' ? -1 : 1;
          }
          if (aStr > bStr) {
            return sortConfig.direction === 'ascending' ? 1 : -1;
          }
        }
        
        return 0; // Nếu bằng nhau
      });
    }
    // 👆 KẾT THÚC PHẦN SORT 👆

    return sortableItems;
  }, [activeFilter, sortConfig]);

  const requestSort = (key: keyof RecruitData) => {
    let direction: 'ascending' | 'descending' = 'ascending';
    if (sortConfig?.key === key && sortConfig.direction === 'ascending') {
      direction = 'descending';
    }
    setSortConfig({ key, direction });
  };
  
  const headers: { label: string, key: keyof RecruitData }[] = [
    { label: 'Foundation', key: 'foundation' },
    { label: 'Job Position', key: 'position' },
    { label: 'Field', key: 'field' },
    { label: 'Salary', key: 'salary' },
    { label: 'Form', key: 'form' },
    { label: 'Update', key: 'trend' },
  ];

  return (
    <WidgetCard>
      <div className="flex flex-col h-60">
        {/* Header Widget */}
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold text-white">Real-time recruiting</h2>
          <div className="flex items-center p-1 bg-zinc-900 rounded-full">
            {filters.map((filter) => (
              <button key={filter} onClick={() => setActiveFilter(filter)} className={`text-sm rounded-full px-5 py-1.5 transition-colors font-medium ${activeFilter === filter ? 'bg-blue-600 text-white' : 'bg-transparent text-gray-400 hover:text-white'}`}>
                {filter}
              </button>
            ))}
          </div>
        </div>

        {/* Bảng Dữ Liệu */}
        <div className="flex-grow flex flex-col overflow-hidden">
          {/* Header Bảng */}
          <div className="grid grid-cols-6 gap-x-4 px-4 pb-3 flex-shrink-0">
            {headers.map(({ label, key }) => (
              <button key={key} onClick={() => requestSort(key)} className="text-xs text-gray-500 font-semibold flex items-center gap-1 cursor-pointer hover:text-white">
                <span>{label}</span>
                {sortConfig?.key === key && <ArrowUpDown size={12} />}
              </button>
            ))}
          </div>

          {/* Body của Bảng - ✨ THÊM GIỚI HẠN CHIỀU CAO VÀ SCROLL Ở ĐÂY ✨ */}
          <div className="overflow-y-auto space-y-1 pr-2">
            {filteredData.map((row, index) => (
              <div
                key={index}
                className="grid grid-cols-6 gap-x-4 items-center p-4 border-b border-gray-800/60"
              >
                <div className="flex items-center gap-x-3 text-white font-medium">
                  {row.icon}
                  <span>{row.foundation}</span>
                </div>
                <div className="text-white font-medium">{row.position}</div>
                <div className="text-gray-400">{row.field}</div>
                <div className="text-gray-400">${row.salary.toLocaleString(undefined, {minimumFractionDigits: 2, maximumFractionDigits: 2})}</div>
                <div className="text-gray-400">{row.form}</div>
                <div className="flex justify-start">
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

export default RecruitingTable;