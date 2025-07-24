"use client";

import React, { useState, useMemo , useEffect} from 'react';
import { ArrowUpDown } from 'lucide-react';
import WidgetCard from '../ui/WidgetCard';
import MiniTrendChart from '../ui/MiniTrendChart';
import { SiTesla, SiApple, SiGoogle, SiNvidia } from 'react-icons/si';
import { RecruitingTableRowSkeleton } from '../ui/skeletons/RecruitingTableRowSkeleton';
import { useRecruitingData } from '@/app/hooks/useRecruitingData'; 
import { type JobData, type JobField } from '@/app/lib/dashboard-types';



const filters: ('All' | JobField)[] = ['All', 'Blockchain', 'AI', 'IOT', 'Web Dev'];

// --- COMPONENT CHÍNH ---
const RecruitingTable = () => {
  const [activeFilter, setActiveFilter] = useState<'All' | JobField>('All');
  // ✨ SỬA LẠI: Mặc định sort theo Foundation thay vì salary (vì salary giờ là 0) ✨
  const [sortConfig, setSortConfig] = useState<{ key: keyof JobData; direction: 'ascending' | 'descending' } | null>({ key: 'foundation', direction: 'ascending' });
  
  const { jobs, isLoading, error } = useRecruitingData();

  const filteredData = useMemo(() => {
    if (!jobs) return [];
    
    let sortableItems = [...jobs];
    if (activeFilter !== 'All') {
      sortableItems = sortableItems.filter(item => item.field === activeFilter);
    }
    
    if (sortConfig !== null) {
      sortableItems.sort((a, b) => {
        // ✨ SỬA LẠI `any` ĐỂ AN TOÀN HƠN VỚI TYPESCRIPT ✨
        const aValue = a[sortConfig.key as keyof JobData] as any;
        const bValue = b[sortConfig.key as keyof JobData] as any;

        if (typeof aValue === 'number' && typeof bValue === 'number') {
          if (aValue < bValue) return sortConfig.direction === 'ascending' ? -1 : 1;
          if (aValue > bValue) return sortConfig.direction === 'ascending' ? 1 : -1;
        } else {
          const aStr = String(aValue).toLowerCase();
          const bStr = String(bValue).toLowerCase();
          if (aStr < bStr) return sortConfig.direction === 'ascending' ? -1 : 1;
          if (aStr > bStr) return sortConfig.direction === 'ascending' ? 1 : -1;
        }
        return 0;
      });
    }
    return sortableItems;
  }, [jobs, activeFilter, sortConfig]);

  const requestSort = (key: keyof JobData) => {
    // Không cho phép sort theo các cột không có ý nghĩa
    if (key === 'trend' || key === 'icon' || key === 'id' || key === 'salaryText') return; 
    
    let direction: 'ascending' | 'descending' = 'ascending';
    if (sortConfig?.key === key && sortConfig.direction === 'ascending') {
      direction = 'descending';
    }
    setSortConfig({ key, direction });
  };
  
  const headers: { label: string, key: keyof JobData }[] = [
    { label: 'Foundation', key: 'foundation' },
    { label: 'Job Position', key: 'position' },
    { label: 'Field', key: 'field' },
    { label: 'Salary', key: 'salary' }, // Giữ nguyên key để sort, nhưng hiển thị salaryText
    { label: 'Form', key: 'form' },
    { label: 'Update', key: 'trend' },
  ];

  return (
    <WidgetCard>
      <div className="flex flex-col h-70">
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

          {/* Body của Bảng - Render có điều kiện */}
          <div className="overflow-y-auto space-y-1 pr-2">
            {isLoading ? (
              // Hiển thị skeleton khi đang tải
              <div className="animate-pulse">
                {Array.from({ length: 4 }).map((_, i) => <RecruitingTableRowSkeleton key={i} />)}
              </div>
            ) : (
              // Hiển thị dữ liệu thật
              filteredData.map((row) => (
                <div key={row.id} className="grid grid-cols-6 gap-x-4 items-center p-4 h-20 border-b border-gray-800/60 last:border-b-0">
                  
                  {/* Cột 1: Foundation */}
                  <div 
                    className="flex items-center gap-x-3 text-white font-medium text-sm overflow-hidden" 
                    title={row.foundation} // Thêm tooltip để xem tên đầy đủ khi di chuột
                  >
                    {row.icon}
                    {/* Giới hạn 2 dòng và thêm dấu ... */}
                    <span className="line-clamp-2">{row.foundation}</span> 
                  </div>
                  
                  {/* Cột 2: Job Position */}
                  <div 
                    className="text-white font-medium text-sm overflow-hidden"
                    title={row.position} // Thêm tooltip
                  >
                    {/* Giới hạn 2 dòng và thêm dấu ... */}
                    <span className="line-clamp-2">{row.position}</span>
                  </div>
                  
                  {/* Cột 3: Field */}
                  <div className="text-gray-400 text-sm">{row.field}</div>
                  
                  {/* Cột 4: Salary */}
                  {/* `truncate` sẽ cắt bớt nếu không đủ chỗ */}
                  <div className="text-gray-400 text-sm truncate" title={row.salaryText}>{row.salaryText}</div>
                  
                  {/* Cột 5: Form */}
                  <div className="text-gray-400 text-sm">{row.form}</div>
                  
                  {/* Cột 6: Update */}
                  <div className="flex justify-start">
                    <MiniTrendChart trend={row.trend} />
                  </div>
                </div>
              ))
            )}
          </div>
          </div>
      </div>
    </WidgetCard>
  );
};

export default RecruitingTable;