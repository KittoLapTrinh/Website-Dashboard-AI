import React from 'react';

interface MiniTrendChartProps {
  trend: 'up' | 'down';
}

const MiniTrendChart = ({ trend }: MiniTrendChartProps) => {
  const isUp = trend === 'up';

  // ✨ 1. Xác định màu sắc và đường đi của SVG
  const color = isUp ? '#22c55e' : '#ef4444'; // Xanh lá cây hoặc Đỏ
  const path = isUp 
    ? "M 2 20 L 10 12 L 18 18 L 26 10 L 34 14 L 42 6" // Đường dích dắc đi lên
    : "M 2 8 L 10 16 L 18 10 L 26 18 L 34 14 L 42 22"; // Đường dích dắc đi xuống
  
  // Tọa độ của chấm tròn cuối cùng
  const circleCx = 42;
  const circleCy = isUp ? 6 : 22;

  return (
    <svg width="54" height="28" viewBox="0 0 44 28" fill="none" xmlns="http://www.w3.org/2000/svg">
      
      {/* ✨ 2. Thêm hiệu ứng Glow bằng SVG filter */}
      <defs>
        <filter id="glow" x="-50%" y="-50%" width="200%" height="200%">
          <feGaussianBlur stdDeviation="2.5" result="coloredBlur" />
          <feMerge>
            <feMergeNode in="coloredBlur" />
            <feMergeNode in="SourceGraphic" />
          </feMerge>
        </filter>
      </defs>

      {/* ✨ 3. Đường line dích dắc */}
      <path 
        d={path} 
        stroke={color} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round" 
      />

      {/* ✨ 4. Chấm tròn cuối cùng với hiệu ứng glow */}
      <circle 
        cx={circleCx} 
        cy={circleCy} 
        r="4" 
        fill={color}
        style={{ filter: 'url(#glow)' }} // Áp dụng filter
      />
    </svg>
  );
};

export default MiniTrendChart;