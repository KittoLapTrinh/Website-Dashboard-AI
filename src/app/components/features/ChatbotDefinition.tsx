// src/components/features/ChatbotDefinition.tsx - ĐÃ SỬA ĐỂ XẾP CHỒNG ẢNH

import React from 'react';
import Image from 'next/image';
import WidgetCard from '../ui/WidgetCard';
import { Send } from 'lucide-react';
import Chatbot1 from '@/app/assets/image/Chatbot1.png';
import Chatbot2 from '@/app/assets/image/Chatbot2.png';

const ChatbotDefinition = () => {
  return (
    <WidgetCard>
      <h2 className="text-2xl font-bold text-white mb-6">
        <span className=" pb-1">
          Chat Bot Definition
        </span>
      </h2>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center mt-8">
        {/* Cột trái: Hình ảnh được xếp chồng - ĐÃ SỬA ĐỂ RESPONSIVE */}
        <div className="relative aspect-square max-w-md mx-auto">
          {/* LỚP NỀN: Hiệu ứng xoáy (z-10) */}
          <Image
            src={Chatbot2} 
            alt="Vortex background effect"
            fill
            // ✨ Dùng phần trăm thay vì pixel.
            // ✨ Thêm các class responsive của Tailwind (ví dụ: `sm:`, `md:`)
            className="object-contain z-10 
                      -translate-x-1/4 translate-y-1/4
                      sm:-translate-x-[22%] sm:translate-y-[15%]" 
          />
          {/* LỚP TRÊN: Con robot (z-20) */}
          <Image
            src={Chatbot1} 
            alt="Friendly chatbot character"
            fill
            // ✨ Dùng phần trăm và các class responsive
            className="object-contain z-20
                      -translate-x-1/3 -translate-y-1/4
                      sm:-translate-x-[28%] sm:-translate-y-[8%]"
          />
        </div>

        {/* Cột phải: Giao diện Chat (giữ nguyên) */}
        {/* Cột phải: Giao diện Chat - ĐÃ SỬA ĐỂ RESPONSIVE */}
        <div className="flex flex-col gap-y-4">
           <div className="chat-bubble-user self-end max-w-md">
            <p className="text-sm text-gray-300">This technology is really impressive.</p>
          </div>
          {/* Bong bóng chat định nghĩa của bot */}
          {/* Bỏ các class translate, thêm max-w-md để không quá dài */}
          <div className="chat-bubble-bot z-20 ">
            <h3 className="font-bold text-white">Menu Leave Rate (20%)</h3>
            <p className="text-sm text-gray-300 mt-2 leading-relaxed">
              A 20% menu leave rate indicates that a significant portion of customers
              are navigating away without taking further action. This could suggest
              that the menu might not be engaging enough or that users are not
              finding what they're looking for quickly. It's essential to investigate the
              reasons behind this drop-off and consider optimizing the menu layout,
              content, or user experience to reduce this rate.
            </p>
          </div>
          
          {/* Bong bóng chat của người dùng */}
          {/* Dùng `self-end` để nó tự động căn sang phải */}
         
          
          {/* Các nút gợi ý */}
          {/* Thêm `flex-wrap` để các nút tự xuống hàng */}
          <div className="flex gap-2 pt-4 flex-wrap">
            {['Any employee?', 'What is the trend?', 'About technology?'].map((text, i) => (
              <button key={i} className="bg-zinc-700/80 hover:bg-zinc-700 text-xs text-[#636363] px-3 py-1.5 rounded-lg transition-colors">
                {text}
              </button>
            ))}
          </div>
          
          {/* Ô nhập liệu */}
          {/* ✨ Dùng Flexbox để sắp xếp input và button trên cùng một hàng */}
          <div className="flex items-center gap-x-2">
            
            {/* Ô input sẽ tự động chiếm hết không gian còn lại */}
            <input
              type="text"
              placeholder="This technology is really impressive..."
              // ✨ Bỏ `pr-14` vì không cần chừa chỗ cho nút nữa
              className="flex-grow bg-zinc-800 border border-transparent rounded-full pl-4 py-3 text-sm text-white placeholder-white focus:outline-none focus:border-white/50 transition-colors"
            />

            {/* Nút gửi nằm riêng biệt */}
            <button 
              className="w-11 h-11 rounded-full 
                        flex-shrink-0 flex items-center justify-center 
                        transition-all
                        bg-gradient-to-b from-[#0177FB] to-white/50
                        hover:opacity-80"
            >
              <Send size={18} className="text-white rotate-45" />
            </button>

          </div>
          

        </div>
      </div>
    </WidgetCard>
  );
};

export default ChatbotDefinition;