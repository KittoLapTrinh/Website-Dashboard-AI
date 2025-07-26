"use client";

import React, { ReactNode } from 'react';
import { motion } from 'framer-motion';

interface AnimateOnScrollProps {
  children: ReactNode;
  delay?: number; // Thêm prop `delay` để tạo hiệu ứng so le
  className?: string;
}

const AnimateOnScroll = ({ children, delay = 0, className = '' }: AnimateOnScrollProps) => {
  // Định nghĩa các trạng thái của animation
  const variants = {
    hidden: { 
      opacity: 0, 
      y: 50 // Bắt đầu ở vị trí thấp hơn 50px
    },
    visible: { 
      opacity: 1, 
      y: 0 // Di chuyển về vị trí gốc (y=0)
    },
  };

  return (
    <motion.div
      className={className}
      variants={variants}
      initial="hidden" // Trạng thái ban đầu
      whileInView="visible" // Kích hoạt animation khi lọt vào màn hình
      viewport={{ once: true, amount: 0.2 }} // `once: true` để animation chỉ chạy 1 lần
                                            // `amount: 0.2` -> chạy khi 20% widget lọt vào màn hình
      transition={{ 
        duration: 0.6, 
        ease: "easeOut",
        delay: delay // Áp dụng độ trễ
      }}
    >
      {children}
    </motion.div>
  );
};

export default AnimateOnScroll;