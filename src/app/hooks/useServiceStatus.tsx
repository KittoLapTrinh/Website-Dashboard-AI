// File: src/app/hooks/useServiceStatus.tsx
"use client";

import React, { useState, useEffect } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
import { 
  dashboardFacadeAddress, 
  dashboardFacadeAbi, 
  serviceStatusContractAddress,
  serviceStatusContractAbi 
} from '@/app/contracts';
import { type ServiceData, type ServiceStatus } from '@/app/lib/dashboard-types';
import { SiApple, SiTesla, SiGoogle, SiNvidia } from 'react-icons/si';
import { FaMicrosoft } from 'react-icons/fa';
import { Log } from 'viem';

// --- KIỂU DỮ LIỆU THÔ ---
interface RawService { name: string; title: string; subtitle: string; status: number; iconId: string; }
type DecodedStatusLog = { args: { index?: bigint; newStatus?: number; }; } & Log;

// --- HÀM BIẾN ĐỔI (TRANSFORMER) ---
const formatServiceFromContract = (service: RawService, index: number): ServiceData => {
  const getIconComponent = (iconId: string): React.ReactNode => {
    switch (iconId.toLowerCase()) {
      case 'apple': return <SiApple size={24} />;
      case 'tesla': return <SiTesla size={24} />;
      case 'microsoft': return <FaMicrosoft size={22} />;
      case 'google': return <SiGoogle size={22} />;
      case 'nvidia': return <SiNvidia size={24} />;
      default: return null;
    }
  };
  const formatStatus = (status: number): ServiceStatus => 
    (['ok', 'degraded', 'error'] as ServiceStatus[])[status] || 'degraded';

  return {
    id: `service-${index}`,
    serviceName: service.name,
    title: service.title,
    subtitle: service.subtitle,
    status: formatStatus(service.status),
    icon: getIconComponent(service.iconId),
  };
};

// --- CUSTOM HOOK ---
export function useServiceStatus() {
  const [services, setServices] = useState<ServiceData[]>([]);

  const { data: initialServices, isLoading, error, refetch } = useReadContract({
    address: dashboardFacadeAddress,
    abi: dashboardFacadeAbi,
    functionName: 'getAllServices',
  });

  useEffect(() => {
    if (initialServices && Array.isArray(initialServices)) {
      setServices((initialServices as RawService[]).map(formatServiceFromContract));
    }
  }, [initialServices]);

  useWatchContractEvent({
    address: serviceStatusContractAddress,
    abi: serviceStatusContractAbi,
    eventName: 'ServiceStatusUpdated',
    onLogs(logs) {
      logs.forEach(log => {
        const { index, newStatus } = (log as DecodedStatusLog).args;
        if (index !== undefined && newStatus !== undefined) {
          setServices(prevServices => 
            prevServices.map((service, i) => 
              i === Number(index) 
                ? { ...service, status: (['ok', 'degraded', 'error'] as ServiceStatus[])[newStatus] } 
                : service
            )
          );
        }
      });
    },
  });

  return { services, isLoading: isLoading && services.length === 0, error, refetch };
}