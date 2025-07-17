import React from 'react';

// Thêm prop `className` và `transparent`
const WidgetCard = ({ 
  children, 
  className = '', 
  transparent = false 
}: { 
  children: React.ReactNode, 
  className?: string, 
  transparent?: boolean 
}) => {
  
  const backgroundClass = transparent ? 'bg-transparent' : 'bg-black';

  // Kết hợp các class mặc định với class được truyền vào
  const finalClassName = `${backgroundClass} rounded-2xl p-6 h-full ${className}`;

  return (
    <div className={finalClassName}>
      {children}
    </div>
  );
};

export default WidgetCard;