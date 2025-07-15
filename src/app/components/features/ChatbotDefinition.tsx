"use client";

import React, { useState, useEffect, useRef } from 'react';
import Image from 'next/image';
import WidgetCard from '../ui/WidgetCard';
import { Send, Loader2 } from 'lucide-react';

// Giữ nguyên đường dẫn import của bạn, nhưng đảm bảo nó đúng
import Chatbot1 from '@/app/assets/image/Chatbot1.png';
import Chatbot2 from '@/app/assets/image/Chatbot2.png';

interface Message {
  role: 'user' | 'assistant';
  content: string;
}

const ChatbotDefinition = () => {
  // === STATE VÀ LOGIC CHO CHAT ===
  const [messages, setMessages] = useState<Message[]>([]);
  const [botResponse, setBotResponse] = useState<Message | null>(null);
  const [input, setInput] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const chatContainerRef = useRef<HTMLDivElement>(null);
  const textareaRef = useRef<HTMLTextAreaElement>(null);
    const initialBotMessages = [
  "How can I help you today?",
  "Ask me anything about this data.",
  "I'm here to assist you with your analytics.",
  "What insights are you looking for?",
];
  const [displayMessage, setDisplayMessage] = useState<string>(
    "What insights are you looking for?"
  );
  const [displayTitle, setDisplayTitle] = useState<string>("Menu Leave Rate (20%)");

  useEffect(() => {
    if (chatContainerRef.current) {
      chatContainerRef.current.scrollTop = chatContainerRef.current.scrollHeight;
    }
  }, [messages, isLoading]);

  useEffect(() => {
    if (textareaRef.current) {
      const textarea = textareaRef.current;
      // 1. Reset chiều cao về auto để tính toán lại scrollHeight
      textarea.style.height = 'auto'; 
      
      // 2. Lấy chiều cao thực của nội dung
      const scrollHeight = textarea.scrollHeight;
      
      // 3. Đặt chiều cao tối đa
      const maxHeight = 140; // Giới hạn 140px

      // 4. Áp dụng chiều cao mới, không vượt quá giới hạn
      textarea.style.height = `${Math.min(scrollHeight, maxHeight)}px`;

      // 5. Điều khiển việc hiển thị thanh cuộn
      // Nếu scrollHeight lớn hơn maxHeight, hiển thị thanh cuộn, ngược lại thì ẩn đi.
      textarea.style.overflowY = scrollHeight > maxHeight ? 'auto' : 'hidden';
    }
  }, [input]);


 const handleSendMessage = async (e?: React.FormEvent) => {
  if (e) e.preventDefault();
  
  const currentInput = input.trim();
  if (!currentInput || isLoading) return;

  const userMessage: Message = { role: 'user', content: currentInput };

  // 1. Cập nhật state tin nhắn của cả hai bên
  // Thêm tin nhắn của user vào lịch sử chung
  setMessages(prev => [...prev, userMessage]); 
  
  // Xóa input và bật trạng thái loading
  setInput('');
  setIsLoading(true);
  
  // Cập nhật bong bóng bên trái để hiển thị loading
  setDisplayTitle("Assistant is thinking...");
  setDisplayMessage(""); // Xóa nội dung cũ

  try {
    // 2. Tạo lịch sử hội thoại để gửi đi
    // Gửi đi toàn bộ mảng messages hiện tại CỘNG với tin nhắn mới của người dùng
    const conversationHistory = [...messages, userMessage];

    const response = await fetch('/api/chat', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ messages: conversationHistory }),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`API error: ${response.status} ${errorText}`);
    }

    const botReply: Message = await response.json();
    
    // 3. Cập nhật câu trả lời của bot vào cả lịch sử chung và bong bóng hiển thị
    setMessages(prev => [...prev, botReply]);
    setDisplayMessage(botReply.content);

  } catch (error) {
    console.error("Failed to send message:", error);
    const errorMessage: Message = {
      role: 'assistant',
      content: 'Sorry, I encountered an error. Please try again.',
    };
    // Cập nhật cả hai nơi
    setMessages(prev => [...prev, errorMessage]);
    setDisplayMessage(errorMessage.content);
  } finally {
    setIsLoading(false);
    // Sau khi có kết quả, có thể giữ nguyên title "Assistant" hoặc đổi lại
    setDisplayTitle("Assistant");
  }
};

   const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleSendMessage();
    }
    // Shift + Enter sẽ tự động xuống dòng, không cần xử lý
  };



   useEffect(() => {
    // Chỉ chạy hiệu ứng này khi chưa có cuộc trò chuyện nào
    if (messages.length === 0 && !isLoading) {
      const interval = setInterval(() => {
        const randomIndex = Math.floor(Math.random() * initialBotMessages.length);
        setDisplayTitle("Assistant"); // Đổi tiêu đề thành "Assistant"
        setDisplayMessage(initialBotMessages[randomIndex]);
      }, 4000); // Đổi sau mỗi 4 giây

      return () => clearInterval(interval);
    }
  }, [messages, isLoading]);

  return (
    <WidgetCard className="relative overflow-hidden">
      <h2 className="text-2xl font-bold text-white mb-8">
        Chat Bot Definition
      </h2>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-x-8 gap-y-12 items-end">
        
        {/* === CỘT TRÁI: Giữ nguyên === */}
        <div className="relative h-80 lg:h-96">
          <div className="absolute inset-0 transform sm:-translate-x-[40%] sm:-translate-y-[-25%]">
            <Image src={Chatbot2} alt="Vortex" fill className="object-contain -scale-x-100" />
          </div>
          <div className="absolute inset-0 transform sm:-translate-x-[33%] sm:-translate-y-[5%]">
            <Image src={Chatbot1} alt="Chatbot" fill className="object-contain" />
          </div>
          <div className="chat-bubble-bot absolute top-8 right-0 w-[80%] sm:w-[70%] lg:w-[70%] z-20 inset-0 transform sm:-translate-y-[15%] sm:-translate-x-[-35%]">
            <h3 className="font-bold text-white">Chatbot Metanode</h3>
            <p className="text-sm text-gray-300 mt-2 leading-relaxed">
              {displayMessage}
            </p>
          </div>
        </div>

        {/* === CỘT PHẢI: Đã được kết nối với state và logic === */}
        <div className="flex flex-col h-full">
          
          {/* ✨ 1. Khu vực hiển thị tin nhắn */}
          <div className="flex-grow space-y-4 overflow-y-auto pr-2 h-64">
            {/* Render các tin nhắn từ state */}
            {messages.map((msg, index) => (
              <div 
                key={index} 
                className={msg.role === 'user' ? 'chat-bubble-user self-end min-w-[630] max-w-sm' : 'chat-bubble-bot max-w-md'}
              >
                <p className="text-sm text-white">{msg.content}</p>
              </div>
            ))}
            
            {/* Hiển thị "Bot is typing..." */}
            {isLoading && (
              <div className="chat-bubble-bot max-w-xs">
                 <Loader2 className="animate-spin" size={20} />
              </div>
            )}
          </div>
          
          {/* ✨ 2. Khu vực nhập liệu */}
          <div className="mt-auto pt-4 flex-shrink-0">
            <div className="flex gap-2 flex-wrap justify-center mb-4">
              {['Any employee?', 'What is the trend?', 'About technology?'].map((text, i) => (
                <button key={i} onClick={() => setInput(text)} className="bg-zinc-700/80 hover:bg-zinc-700 text-xs text-gray-400 px-3 py-1.5 rounded-lg transition-colors">
                  {text}
                </button>
              ))}
            </div>
            
            <form onSubmit={handleSendMessage} className="flex items-center gap-x-2">
              <textarea
                ref={textareaRef}
                value={input}
                onChange={(e) => setInput(e.target.value)}
                onKeyDown={handleKeyDown}
                placeholder="This technology is really impressive...."
                // ✨ Bỏ `overflow-hidden` và sửa `rounded`
                className="flex-grow bg-zinc-800 border border-transparent rounded-2xl px-4 py-2 text-sm text-white placeholder-gray-500 focus:outline-none focus:border-white/50 transition-colors resize-none"
                rows={1}
                style={{ maxHeight: '140px' }} // Thêm style nội tuyến để CSS biết giới hạn
                disabled={isLoading}
              />
              <button 
                type="submit"
                className="w-11 h-11 rounded-full flex-shrink-0 flex items-center justify-center transition-all bg-gradient-to-b from-[#0177FB] to-white/50 hover:opacity-80 disabled:opacity-50 disabled:cursor-not-allowed"
                disabled={isLoading}
              >
                {isLoading ? <Loader2 className="animate-spin" /> : <Send size={18} className="text-white rotate-45" />}
              </button>
            </form>
          </div>
        </div>
      </div>
    </WidgetCard>
  );
};

export default ChatbotDefinition;