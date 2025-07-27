"use client";

import { useState, useEffect } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
// Giả sử bạn đã thêm các export này vào `src/app/contracts/index.ts`
import { contentContractAddress, contentContractAbi, dashboardFacadeAddress, dashboardFacadeAbi } from '@/app/contracts';
import { type ContentItem } from '@/app/lib/dashboard-types'; // Giả sử bạn đã thêm type này

// --- HÀM HELPER ĐỂ BIẾN ĐỔI DỮ LIỆU ---
// Dữ liệu từ contract trả về dưới dạng mảng/tuple, cần chuyển thành object dễ sử dụng.
const formatContentItem = (rawItem: any): ContentItem => ({
  id: Number(rawItem.id),
  contentType: Number(rawItem.contentType), // 0: INSPIRATION, 1: RESEARCH, 2: NEWS
  title: rawItem.title,
  description: rawItem.description,
  imageUrl: rawItem.imageUrl,
  author: rawItem.author,
  timestamp: Number(rawItem.timestamp),
});


export function useContentData(count: number) {
  const [content, setContent] = useState<ContentItem[]>([]);

  // Lấy dữ liệu ban đầu thông qua Facade
  const { data: initialContent, isLoading, error } = useReadContract({
    address: dashboardFacadeAddress,
    abi: dashboardFacadeAbi,
    functionName: 'getRecentContent',
    args: [BigInt(count)],
  });

  // Xử lý dữ liệu ban đầu
  useEffect(() => {
    if (initialContent && Array.isArray(initialContent)) {
      const formatted = (initialContent as any[]).map(formatContentItem);
      setContent(formatted);
    }
  }, [initialContent]);

  // Lắng nghe sự kiện `ContentAdded` để cập nhật real-time
  useWatchContractEvent({
    address: contentContractAddress,
    abi: contentContractAbi,
    eventName: 'ContentAdded',
    onLogs(logs) {
      // Mỗi log là một sự kiện mới
      logs.forEach(log => {
        const newContentRaw = (log as any).args.newContent;
        if (newContentRaw) {
          const formattedNewItem = formatContentItem(newContentRaw);
          // Thêm content mới vào đầu danh sách và giữ cho danh sách có độ dài `count`
          setContent(prevContent => {
            const updatedContent = [formattedNewItem, ...prevContent];
            return updatedContent.slice(0, count);
          });
        }
      });
    },
  });

  return { content, isLoading: isLoading && content.length === 0, error };
}