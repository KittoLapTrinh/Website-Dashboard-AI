import React from 'react';

interface MonitorCardProps {
  title: string;
  value: string; // Giá trị hiển thị ở header
  icon: React.ReactNode;
  children: React.ReactNode; // Dùng để render biểu đồ và giá trị cuối
  statusColor?: string;
}

const MonitorCard = ({
  title,
  value,
  icon,
  children,
  statusColor = 'green',
}: MonitorCardProps) => {
  const gradientFrom = statusColor === 'green' ? 'from-green-500/80' : 'from-red-500/80';
  const textColor = statusColor === 'green' ? 'text-green-400' : 'text-red-400';

  return (
    // Div cha tạo hiệu ứng viền gradient
    <div className={`p-px rounded-2xl bg-gradient-to-tr ${gradientFrom} to-transparent`}>
      <div className="bg-zinc-900 rounded-[15px] p-4 h-full flex flex-col">
        {/* Header */}
        <div className="flex items-center gap-x-3">
          <div className="text-white">{icon}</div>
          <div>
            <p className="font-semibold text-white">{title}</p>
            <p className={`font-semibold text-sm ${textColor}`}>{value}</p>
          </div>
        </div>
        
        {/* Chart Area - Render children ở đây */}
        <div className="relative flex-grow mt-4 h-24">
          {children}
        </div>
      </div>
    </div>
  );
};

export default MonitorCard;