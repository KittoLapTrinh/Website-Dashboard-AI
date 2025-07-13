"use client";

import React, { useState } from 'react';
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

  const handleSendMessage = async (e: React.FormEvent) => {
    e.preventDefault();
    const currentInput = input.trim();
    if (!currentInput) return;

    const userMessage: Message = { role: 'user', content: currentInput };
    const newMessages = [...messages, userMessage];
    setMessages(newMessages);
    setInput('');
    setIsLoading(true);

    try {
      // Gọi đến API Route Gemini
      const response = await fetch('/api/chat', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ messages: newMessages }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`API error: ${response.status} ${errorText}`);
      }

      const botMessage = await response.json();
      setMessages(prev => [...prev, botMessage]);

    } catch (error) {
      console.error("Failed to send message:", error);
      const errorMessage: Message = {
        role: 'assistant',
        content: 'Sorry, I encountered an error. Please try again.',
      };
      setMessages(prev => [...prev, errorMessage]);
    } finally {
      setIsLoading(false);
    }
  };

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
            <h3 className="font-bold text-white">Menu Leave Rate (20%)</h3>
            <p className="text-sm text-gray-300 mt-2 leading-relaxed">
              A 20% menu leave rate indicates that a significant portion of customers...
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
                className={msg.role === 'user' ? 'chat-bubble-user self-end max-w-sm' : 'chat-bubble-bot max-w-md'}
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
              <input
                type="text"
                value={input}
                onChange={(e) => setInput(e.target.value)}
                placeholder="This technology is really impressive...."
                className="flex-grow bg-zinc-800 border border-transparent rounded-full pl-4 py-3 text-sm text-white placeholder-gray-500 focus:outline-none focus:border-white/50 transition-colors"
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