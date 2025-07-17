"use client";

import React, { useState, useEffect, useCallback } from 'react';
import { ArrowUpRight, RefreshCw } from 'lucide-react';
import { SiTesla, SiGoogle, SiNvidia, SiApple } from 'react-icons/si';
import { FaMicrosoft } from 'react-icons/fa';
import WidgetCard from '../ui/WidgetCard';
import StatusCard from '../ui/StatusCard';
import { StatusCardSkeleton } from '../ui/StatusCardSkeleton'; // Import skeleton

// --- KIỂU DỮ LIỆU ---
type ServiceStatus = 'ok' | 'error' | 'degraded';
interface ServiceData {
  id: string;
  title: string;
  subtitle: string;
  icon: React.ReactNode;
  name: string;
  status: ServiceStatus;
}

// --- HÀM MÔ PHỎNG FETCH API ---
const fetchServicesData = (): Promise<ServiceData[]> => {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve([
        { id: 'wallet', title: '1.7M User', subtitle: '2.2M Tx (0.01s)', icon: <SiApple size={24} />, name: 'Wallet Blockchain', status: 'ok' },
        { id: 'chain', title: 'Volum 125M', subtitle: '100,000 TPS', icon: <SiTesla size={24} />, name: 'Chain Analytics', status: 'error' },
        { id: 'email', title: '1.7M User', subtitle: '2.2M Mail (0.01s)', icon: <FaMicrosoft size={22} />, name: 'Dapp Email', status: 'ok' },
        { id: 'learn', title: '1.7M User', subtitle: '2.2M Course', icon: <SiGoogle size={22} />, name: 'Dapp Elearning', status: 'ok' },
        { id: 'chat', title: '1.2M Messages', subtitle: '2.1M Users (0.02s)', icon: <SiNvidia size={24} />, name: 'Dapp Chat', status: 'ok' },
      ]);
    }, 500); // Giả lập độ trễ mạng 1.5 giây
  });
};

const ServiceStatusGrid = () => {
  const [services, setServices] = useState<ServiceData[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // --- LOGIC LẤY DỮ LIỆU ---
  const loadData = useCallback(async () => {
    setIsLoading(true);
    setError(null);
    try {
      const data = await fetchServicesData();
      setServices(data);
    } catch (err) {
      setError('Failed to load service status.');
    } finally {
      setIsLoading(false);
    }
  }, []);

  useEffect(() => {
    loadData();
  }, [loadData]);
  
  // --- HÀM XÁC ĐỊNH MÀU SẮC ---
  const getStatusColor = (status: ServiceStatus) => {
    switch (status) {
      case 'ok': return '#22c55e'; // Green
      case 'error': return '#ef4444'; // Red
      case 'degraded': return '#f59e0b'; // Amber
      default: return '#6b7280'; // Gray
    }
  };

  // --- RENDER ---
  return (
    <WidgetCard>
      <div className="flex justify-between items-center mb-4">
        {/* Nút Refresh */}
        <button onClick={loadData} disabled={isLoading} className="p-2 rounded-full hover:bg-zinc-800 disabled:opacity-50 disabled:cursor-not-allowed">
          <RefreshCw size={16} className={`text-gray-400 ${isLoading ? 'animate-spin' : ''}`} />
        </button>
        {/* Nút See All */}
        <button className="flex items-center overflow-hidden border border-gray-600 rounded-full text-sm text-gray-300 hover:border-gray-400 transition-all">
          <span className="px-4 py-1.5">See all</span>
          <span className="p-2.5 border-l border-gray-600">
            <ArrowUpRight size={16} />
          </span>
        </button>
      </div>

      <div className="flex-grow flex items-center justify-center">
        <div className="grid grid-cols-1 sm:grid-cols-3 lg:grid-cols-5 gap-5 w-full">
          {isLoading ? (
            // Hiển thị Skeleton khi đang tải
            Array.from({ length: 5 }).map((_, i) => <StatusCardSkeleton key={i} />)
          ) : error ? (
            // Hiển thị lỗi
            <div className="col-span-full text-center text-red-400">{error}</div>
          ) : (
            // Hiển thị dữ liệu
            services.map((service) => (
              <StatusCard
                key={service.id}
                title={service.title}
                subtitle={service.subtitle}
                icon={service.icon}
                name={service.name}
                statusColor={getStatusColor(service.status)}
              />
            ))
          )}
        </div>
      </div>
    </WidgetCard>
  );
};

export default ServiceStatusGrid;