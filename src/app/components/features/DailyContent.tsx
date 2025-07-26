"use client";

import Image from 'next/image';
import { Newspaper, Lightbulb, Beaker } from 'lucide-react';
import { useContentData } from '@/app/hooks/useContentData';
import WidgetCard from '../ui/WidgetCard';
import { type ContentItem } from '@/app/lib/dashboard-types';

// Component con để render một thẻ nội dung
const ContentCard = ({ item }: { item: ContentItem }) => {
  const getIconAndType = () => {
    switch (item.contentType) {
      case 0: // INSPIRATION
        return { icon: <Lightbulb size={16} />, type: "Inspiration", color: "text-yellow-400" };
      case 1: // RESEARCH
        return { icon: <Beaker size={16} />, type: "Research", color: "text-blue-400" };
      case 2: // NEWS
        return { icon: <Newspaper size={16} />, type: "News", color: "text-green-400" };
      default:
        return { icon: <Lightbulb size={16} />, type: "Content", color: "text-gray-400" };
    }
  };

  const { icon, type, color } = getIconAndType();
  const date = new Date(item.timestamp * 1000).toLocaleDateString('en-US', { month: 'short', day: 'numeric' });

  return (
    <div className="flex gap-x-4 p-4 bg-zinc-900 rounded-lg hover:bg-zinc-800 transition-colors">
      <div className="flex-shrink-0 w-24 h-24 relative">
        <Image 
          src={item.imageUrl} 
          alt={item.title} 
          layout="fill" 
          objectFit="cover" 
          className="rounded-md"
        />
      </div>
      <div className="flex flex-col">
        <div className={`flex items-center gap-x-2 text-xs font-medium ${color}`}>
          {icon}
          <span>{type}</span>
        </div>
        <h3 className="font-bold text-white mt-1 text-base leading-tight">{item.title}</h3>
        <p className="text-sm text-gray-400 mt-2 flex-grow">{item.description}</p>
        <div className="flex justify-between items-center mt-2 text-xs text-gray-500">
          <span>By: {item.author}</span>
          <span>{date}</span>
        </div>
      </div>
    </div>
  );
};


const DailyContent = () => {
  // Lấy 5 mục nội dung gần đây nhất
  const { content, isLoading, error } = useContentData(5);

  return (
    <WidgetCard>
      <h2 className="text-xl font-bold text-white mb-4">Daily Updates</h2>
      {isLoading ? (
        <div className="text-center text-gray-500">Loading daily content...</div>
      ) : error ? (
        <div className="text-center text-red-500">Error loading content.</div>
      ) : (
        <div className="space-y-3 overflow-y-auto custom-scrollbar pr-2 h-[28rem]"> {/* Chiều cao cố định để scroll */}
          {content.map(item => (
            <ContentCard key={item.id} item={item} />
          ))}
        </div>
      )}
    </WidgetCard>
  );
};

export default DailyContent;