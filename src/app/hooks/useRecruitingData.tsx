// File: src/app/hooks/useRecruitingData.tsx
"use client";

import React, { useState, useEffect } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
import { 
  dashboardFacadeAddress, 
  dashboardFacadeAbi, 
  recruitingContractAddress,
  recruitingContractAbi 
} from '@/app/contracts';
import { type JobData, type JobField, type JobTrend } from '@/app/lib/dashboard-types';
import { SiTesla, SiApple, SiGoogle, SiNvidia } from 'react-icons/si';
import { FaBuilding } from 'react-icons/fa';
import { Log } from 'viem';

// --- KIỂU DỮ LIỆU THÔ ---
interface RawJob { foundation: string; position: string; field: string; salary: bigint; form: string; trend: number; iconId: string; salaryText: string; }
type DecodedJobLog = { args: { newJob?: RawJob; }; } & Log;

// --- HÀM BIẾN ĐỔI ---
const formatJobFromContract = (job: RawJob, index: number): JobData => {
  
  // Hàm này sẽ trả về component icon đã được chuẩn hóa kích thước
  const getIcon = (iconId: string): React.ReactNode => {
    let iconComponent: React.ReactNode;
    const iconSize = 20; // Đặt kích thước mong muốn

    switch (iconId.toLowerCase()) {
      case 'nvidia':
        iconComponent = <SiNvidia size={iconSize} />;
        break;
      case 'tesla':
        iconComponent = <SiTesla size={iconSize} />;
        break;
      case 'apple':
        iconComponent = <SiApple size={iconSize} />;
        break;
      case 'google':
        iconComponent = <SiGoogle size={iconSize} />;
        break;
      // Thêm các case khác nếu bạn có thêm icon
      // ...
      default:
        // Icon tòa nhà mặc định
        iconComponent = <FaBuilding size={iconSize} />;
        break;
    }
    
    // ✨ BƯỚC QUAN TRỌNG NHẤT: BỌC ICON TRONG MỘT KHUNG CỐ ĐỊNH ✨
    // Khung này có kích thước w-6 h-6 (24px x 24px) và căn icon vào giữa.
    // Điều này đảm bảo không gian mà icon chiếm dụng luôn nhất quán.
    return (
      <div className="w-6 h-6 flex-shrink-0 flex items-center justify-center text-gray-400">
        {iconComponent}
      </div>
    );
  };
  
  // ... (phần return của hàm formatJobFromContract giữ nguyên)
  return {
    id: `job-${index}-${job.foundation}`,
    foundation: job.foundation,
    position: job.position,
    field: job.field as JobField,
    salary: Number(job.salary),
    form: job.form,
    trend: (job.trend === 0 ? 'up' : 'down') as JobTrend,
    icon: getIcon(job.iconId), // Gọi hàm getIcon đã được sửa
    salaryText: job.salaryText,
  };
};

// --- CUSTOM HOOK ---
export function useRecruitingData() {
  const [jobs, setJobs] = useState<JobData[]>([]);

  const { data: initialJobs, isLoading, error, refetch } = useReadContract({
    address: dashboardFacadeAddress,
    abi: dashboardFacadeAbi,
    functionName: 'getAllJobs',
  });

  useEffect(() => {
    if (initialJobs && Array.isArray(initialJobs)) {
      setJobs((initialJobs as RawJob[]).map(formatJobFromContract).reverse());
    }
  }, [initialJobs]);

  useWatchContractEvent({
    address: recruitingContractAddress,
    abi: recruitingContractAbi,
    eventName: 'JobAdded',
    onLogs(logs) {
      logs.forEach(log => {
        const newJobRaw = (log as DecodedJobLog).args.newJob;
        if (newJobRaw) {
          const formattedNewJob = formatJobFromContract(newJobRaw, jobs.length);
          setJobs(prevJobs => [formattedNewJob, ...prevJobs]);
        }
      });
    },
  });

  return { jobs, isLoading: isLoading && jobs.length === 0, error, refetch };
}