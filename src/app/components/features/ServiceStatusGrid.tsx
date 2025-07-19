"use client";

import React, { useState, useEffect, useCallback } from 'react';
import { ArrowUpRight, RefreshCw } from 'lucide-react';
import { SiTesla, SiGoogle, SiNvidia, SiApple } from 'react-icons/si';
import { FaMicrosoft } from 'react-icons/fa';
import WidgetCard from '../ui/WidgetCard';
import StatusCard from '../ui/StatusCard';
import { StatusCardSkeleton } from '../ui/skeletons/StatusCardSkeleton'; // Import skeleton
import { type ServiceData, type ServiceStatus } from '@/app/lib/dashboard-types';
import { useServiceStatus } from '@/app/hooks/useServiceStatus';



const ServiceStatusGrid = () => {

  const { services, isLoading, error, refetch } = useServiceStatus();


  
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
        <button onClick={() => refetch()} disabled={isLoading} className="p-2 rounded-full hover:bg-zinc-800 disabled:opacity-50 disabled:cursor-not-allowed">
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
            <div className="col-span-full text-center text-red-400">
              Error: Could not fetch data from the contract.
            </div>
          ) : (
            // Hiển thị dữ liệu
            services?.map((service) => (
              <StatusCard
                key={service.id}
                title={service.title}
                subtitle={service.subtitle}
                icon={service.icon}
                name={service.serviceName}
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