// src/components/widgets/CustomerReaction.tsx
import React from 'react';
import Image from 'next/image';
import WidgetCard from '../ui/WidgetCard';
import TestimonialCard from '../ui/TestimonialCard';
import avatar1 from '@/app/assets/image/Customer1.png';
import avatar2 from '@/app/assets/image/Customer2.jpg';
import avatar3 from '@/app/assets/image/Customer3.png';
import avatar4 from '@/app/assets/image/Customer4.png';

const testimonials = [
  {
    quote: 'Wow, this AI solution could revolutionize how we understand our industry and customers.',
    avatarSrc: avatar2, // Đặt ảnh avatar trong /public/avatars
    imagePosition: 'left' as const,
  },
  {
    quote: 'This technology is really impressive.',
    avatarSrc: avatar3,
    imagePosition: 'right' as const,
  },
  {
    quote: 'This is truly a major leap forward. An AI tool that provides better customer insights.',
    avatarSrc: avatar4,
    imagePosition: 'left' as const,
  },
];

const CustomerReaction = () => {
  return (
    <WidgetCard>
      <h2 className="text-2xl font-bold text-white -mb-[45px] pt-[30px]">
        Customer's Reaction
      </h2>
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Cột trái: Hình ảnh lớn */}
        <div className="lg:col-span-1 w-full h-full min-h-[300px] relative">
          <Image
            src={avatar1} // Đặt ảnh này trong /public/images
            alt="Team discussing project on a large screen"
            fill // fill sẽ làm ảnh lấp đầy div cha
            className="rounded-2xl object-cover"
          />
        </div>

        {/* Cột phải: Danh sách testimonials */}
        <div className="flex flex-col justify-center gap-y-2">
          {testimonials.map((testimonial, index) => (
            <TestimonialCard
              key={index}
              quote={testimonial.quote}
              avatarSrc={testimonial.avatarSrc}
              avatarAlt={`Avatar of customer ${index + 1}`}
              imagePosition={testimonial.imagePosition}
               avatarPosition={index === 1 ? 'right' : 'left'}
              variant={index === 1 ? 'embedded' : 'default'}
              // ✨ Logic điều khiển dấu ngoặc kép theo yêu cầu của bạn
              showOpeningQuote={index !== 0 && index !== 2} // Hiện khi index không phải 0 và 2
              showClosingQuote={index !== 1}
            />
          ))}
        </div>
      </div>
    </WidgetCard>
  );
};

export default CustomerReaction;