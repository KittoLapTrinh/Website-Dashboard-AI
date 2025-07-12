

import { ChevronDown, TrendingUp } from 'lucide-react';
import Image from 'next/image';
import WidgetCard from '../ui/WidgetCard';
import Totalview from '@/app/assets/image/Totalview.png';
const TotalViewers = () => {
  return (
    // 2. Bọc toàn bộ nội dung bằng WidgetCard
    <WidgetCard className="relative overflow-hidden" transparent>
      <Image
        src={Totalview}
            alt="Total Viewers Background"
            layout="fill" // Hoặc `fill` nếu dùng Next.js 13+
            className="object-cover z-0" // `object-cover` để ảnh lấp đầy mà không bị méo.
      />
      <div className="relative z-10 h-full flex flex-col justify-between">
        <div className="flex justify-between items-start ">
          {/* Cột bên trái */}
          <div className="flex flex-col gap-y-2">
            <p className="text-[#FFFFFF] text-xl">Total Viewers</p>
            <p className="text-4xl font-bold text-white">12,304.11</p>
            <div className="flex items-center gap-x-2 mt-18">
              <span className="text-[12px] text-gray-400">Return</span>
              <span className="flex items-center gap-x-1 text-xs 
                          bg-green-900/50 text-green-400 
                          border border-green-700/60 
                          rounded-full px-1.5 py-px">
                <TrendingUp size={14} />
                +3.5%
              </span>
            </div>
          </div>

          {/* Cột bên phải */}
          <div className="flex items-center gap-x-2">
            <button className="text-sm text-gray-300 border border-[#FFFFFF] hover:bg-gray-700/50 rounded-full px-3 py-1 transition-colors">
              6M
            </button>
            <button className="p-1.5 border border-[#FFFFFF] hover:bg-gray-700/50 rounded-full transition-colors">
              <ChevronDown size={18} className="text-gray-300" />
            </button>
          </div>
        </div>
      </div>
    </WidgetCard>
  );
};

export default TotalViewers;