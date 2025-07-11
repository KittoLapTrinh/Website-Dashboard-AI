// src/components/widgets/TestimonialCard.tsx
import React from 'react';
import Image from 'next/image';
import { Quote } from 'lucide-react';

interface TestimonialCardProps {
  quote: string;
  avatarSrc: string;
  avatarAlt: string;
  imagePosition?: 'left' | 'right'; // Vị trí của avatar
  variant?: 'default' | 'embedded';
  showOpeningQuote?: boolean; // ✨ Prop để điều khiển dấu ngoặc mở
  showClosingQuote?: boolean; // ✨ Prop để điều khiển dấu ngoặc đóng

}

const TestimonialCard = ({
  quote,
  avatarSrc,
  avatarAlt,
  imagePosition = 'left',
   variant = 'default',
  showOpeningQuote = true, // Mặc định là hiện
  showClosingQuote = true, // Mặc định là hiện
  
}: TestimonialCardProps) => {
  const isAvatarLeft = imagePosition === 'left';

if (variant === 'embedded') {
  return (
    <div
        className={`bg-zinc-800 rounded-[50px] p-4 flex items-center gap-x-4
          ${!isAvatarLeft && 'flex-row-reverse'}`}>
      
      {/* Avatar */}
      <div className="flex-shrink-0">
        <Image
          src={avatarSrc}
          alt={avatarAlt}
          // ✨ 1. Tăng kích thước từ 64x64 lên 80x80 (hoặc lớn hơn nếu muốn)
          width={80}
          height={80}
          // ✨ 2. Bỏ `border-radius` và đảm bảo có `rounded-full` và `object-cover`
          className="rounded-full object-cover"
        />
      </div>
      
      {/* Quote Box */}
      <div className="bg-zinc-800 p-6 rounded-2xl relative w-full">
        {/* Dấu ngoặc kép trên */}
        {showOpeningQuote && (
          <Quote size={32} className="absolute top-3 left-4 text-blue-500 opacity-70 transform -scale-x-100" />
        )}
        <p className="text-gray-300 text-sm leading-relaxed text-left mx-8">{quote}</p>
        {showClosingQuote && (
          <Quote size={32} className="absolute bottom-3 right-4 text-blue-500 opacity-70" />
        )}
      </div>
    </div>
  );
}

const bubbleShapeClasses = isAvatarLeft ? 'rounded-r-2xl rounded-l-[50px]' : 'rounded-l-2xl rounded-r-[50px]';

  return (
    <div className={`flex items-center gap-x-5 ${!isAvatarLeft && 'flex-row-reverse'}`}>
      <div className="flex-shrink-0">
        <Image
          src={avatarSrc}
          alt={avatarAlt}
          width={80}
          height={80}
          className="rounded-full object-cover"
        />
      </div>
      <div className={`bg-zinc-800 p-6 relative ${bubbleShapeClasses}`}>
        {showOpeningQuote && (
          <Quote size={32} className="absolute top-3 left-4 text-blue-500 opacity-70 transform -scale-x-100" />
        )}
        <p className="text-gray-300 text-sm leading-relaxed text-left mx-8">{quote}</p>
        {showClosingQuote && (
          <Quote size={32} className="absolute bottom-3 right-4 text-blue-500 opacity-70" />
        )}
      </div>
    </div>
  );
};



export default TestimonialCard;