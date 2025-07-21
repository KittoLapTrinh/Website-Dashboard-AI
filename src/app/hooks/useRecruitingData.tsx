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
interface RawJob { foundation: string; position: string; field: string; salary: bigint; form: string; trend: number; iconId: string; }
type DecodedJobLog = { args: { newJob?: RawJob; }; } & Log;

// --- HÀM BIẾN ĐỔI ---
const formatJobFromContract = (job: RawJob, index: number): JobData => {
  const getIcon = (iconId: string) => {
    switch (iconId.toLowerCase()) {
      case 'nvidia': return <SiNvidia size={24} className="text-gray-400"/>;
      case 'tesla': return <SiTesla size={24} className="text-gray-400"/>;
      case 'apple': return <SiApple size={24} className="text-gray-400"/>;
      case 'google': return <SiGoogle size={22} className="text-gray-400"/>;
      case 'fpt': // Fallthrough
      case 'consensys': // Fallthrough
      default: return <FaBuilding size={24} className="text-gray-400"/>;
    }
  };

  return {
    id: `job-${index}-${job.foundation}`,
    foundation: job.foundation,
    position: job.position,
    field: job.field as JobField,
    salary: Number(job.salary),
    form: job.form,
    trend: (job.trend === 0 ? 'up' : 'down') as JobTrend,
    icon: getIcon(job.iconId),
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