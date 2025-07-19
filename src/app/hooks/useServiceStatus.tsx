import { useReadContract, useAccount } from 'wagmi'; 
import { dashboardAddress, dashboardAbi } from '@/app/contracts';
import { type ServiceData, type ServiceStatus } from '@/app/lib/dashboard-types';
import { SiApple, SiTesla, SiGoogle, SiNvidia } from 'react-icons/si';
import { FaMicrosoft } from 'react-icons/fa';

import React from 'react';

// Hàm helper để map iconId từ contract sang component icon
// ... imports ...

// Hàm helper để map iconId từ contract sang component icon - ĐÃ SỬA
const getIconComponent = (iconId: string): React.ReactNode => {
  const lowerCaseIconId = iconId.toLowerCase();

  switch (lowerCaseIconId) {
    case 'apple': return <SiApple size={24} />;
    case 'tesla': return <SiTesla size={24} />;
    case 'microsoft': return <FaMicrosoft size={22} />;
    case 'google': return <SiGoogle size={22} />;
    case 'nvidia': return <SiNvidia size={24} />;
    default: return null;
  }
};

// Hàm helper để map status từ contract (số) sang chuỗi
const formatStatus = (status: number): ServiceStatus => {
  if (status === 0) return 'ok';
  if (status === 1) return 'degraded';
  if (status === 2) return 'error';
  return 'degraded'; // Mặc định
};

export function useServiceStatus() {

  const { chain } = useAccount();
  
  const { data: rawData, isLoading, error, refetch, isError, isSuccess  } = useReadContract({
    address: dashboardAddress,
    abi: dashboardAbi,
    functionName: 'getAllServices',
    
  });
  
  // console.log("Raw Data from Contract:", rawData);
  const services = React.useMemo(() => {
    if (!rawData || !Array.isArray(rawData)) return undefined;
    
    // ✨ Xử lý dữ liệu một cách an toàn hơn
    return rawData.map((service: any, index: number) => {
      // Chuyển đổi bigint sang string/number một cách an toàn
      const userCount = typeof service.userCount === 'bigint' ? service.userCount.toString() : service.userCount;
      const txCount = typeof service.txCount === 'bigint' ? service.txCount.toString() : service.txCount;
      
      return {
        id: `service-${index + 1}`,
        serviceName: service.name || '',
        title: service.title || '',
        subtitle: service.subtitle || '',
        status: formatStatus(Number(service.status)), // Đảm bảo status là number
        icon: getIconComponent(service.iconId || ''),
      }
    });
  }, [rawData]);

  return { services, isLoading, error, refetch };
}