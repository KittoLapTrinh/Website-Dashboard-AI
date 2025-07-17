import React from 'react';

interface StatusCardProps {
  title: string;
  subtitle: string;
  icon: React.ReactNode;
  name: string;
  // Prop này được truyền từ component cha để quyết định màu sắc
  statusColor?: string; 
}

const StatusCard = ({ 
  title, 
  subtitle, 
  icon, 
  name, 
  statusColor = '#22c55e'
}: StatusCardProps) => {

  return (
    <div className="p-px rounded-2xl bg-gradient-to-br"  style={{
        background: `linear-gradient(to right, ${statusColor}, #18181b)` // #18181b là mã màu của zinc-900
      }}>
      {/* ✨ 1. Giảm padding từ p-4 xuống p-3 */}
      <div className="bg-zinc-900 rounded-[15px] p-2 h-full flex flex-col justify-between">
        <div>
          {/* ✨ 2. Giảm kích thước font title và subtitle */}
          <p className="text-white font-bold text-sm">{title}</p>
          <p className="text-[10px]" style={{ color: statusColor }}>
            {subtitle}
          </p>
        </div>
        
         <div className="flex flex-col gap-y-1 mt-3">
          {/* ✨ 3. Giảm kích thước icon và font tên */}
          <div className="text-gray-400 w-5 h-5">{icon}</div>
          <p className="text-xs text-gray-400">{name}</p>
        </div>
      </div>
    </div>
  );
};

export default React.memo(StatusCard);