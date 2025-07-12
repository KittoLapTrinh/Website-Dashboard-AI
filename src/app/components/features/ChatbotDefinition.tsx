import React from 'react';
import Image from 'next/image';
import WidgetCard from '../ui/WidgetCard';
import { Send } from 'lucide-react';
import Chatbot1 from '@/app/assets/image/Chatbot1.png';
import Chatbot2 from '@/app/assets/image/Chatbot2.png';

const ChatbotDefinition = () => {
  return (
    <WidgetCard className="relative overflow-hidden">
      <h2 className="text-2xl font-bold text-white mb-8">
        Chat Bot Definition
      </h2>

      {/* ✨ 1. Layout chính: 2 cột trên màn hình lớn */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-x-8 gap-y-12 items-end">
        
        {/* === CỘT TRÁI: Chứa robot và bong bóng định nghĩa === */}
        <div className="relative h-80 lg:h-96">
          {/* Ảnh nền và robot */}
          <div className="absolute inset-0 transform sm:-translate-x-[40%] sm:-translate-y-[-25%]">
            <Image
              src={Chatbot2} 
              alt="Vortex background effect"
              fill
              className="object-contain -scale-x-100"
            />
          </div>
          <div className="absolute inset-0 transform sm:-translate-x-[33%] sm:-translate-y-[5%]">
            <Image
              src={Chatbot1} 
              alt="Friendly chatbot character"
              fill
              className="object-contain"
            />
          </div>
          {/* Bong bóng chat định nghĩa được định vị tuyệt đối */}
          <div 
            className="chat-bubble-bot absolute top-8 right-0 w-[80%] sm:w-[70%] lg:w-[90%] z-20 inset-0 transform sm:-translate-y-[15%] sm:-translate-x-[-35%]"
          >
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
        </div>

        {/* === CỘT PHẢI: Giao diện tương tác của người dùng === */}
        <div className="flex flex-col gap-y-4 w-full ">
          
          <div className="chat-bubble-user self-end max-w-sm inset-0 transform sm:-translate-y-[300%]">
            <p className="text-sm text-[#FFFFFF]">
              This technology is really impressive.
            </p>
          </div>
          
          <div className="flex gap-2 pt-4 flex-wrap justify-center">
            {['Any employee?', 'What is the trend?', 'About technology?'].map((text, i) => (
              <button key={i} className="bg-zinc-700/80 hover:bg-zinc-700 text-xs text-[#717177] px-3 py-1.5 rounded-lg transition-colors">
                {text}
              </button>
            ))}
          </div>

          <div className="flex items-center gap-x-2 mt-2">
            <input
              type="text"
              placeholder="This technology is really impressive...."
              className="flex-grow bg-zinc-800 border border-transparent rounded-full pl-4 py-3 text-sm text-white placeholder-[#FFFFFF] focus:outline-none focus:border-white/50 transition-colors"
            />
            <button 
              className="w-11 h-11 rounded-full flex-shrink-0 flex items-center justify-center transition-all bg-gradient-to-b from-[#0177FB] to-white/50 hover:opacity-80"
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