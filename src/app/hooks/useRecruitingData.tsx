"use client";

import React, { useState, useEffect, useMemo } from 'react';
import { useReadContract, useWatchContractEvent } from 'wagmi';
import { dashboardAddress, dashboardAbi } from '@/app/contracts';
import { type JobData, type JobTrend, type JobField } from '@/app/lib/dashboard-types';
import { SiTesla, SiApple, SiGoogle, SiNvidia } from 'react-icons/si';
import { Log } from 'viem';

// --- KIỂU DỮ LIỆU THÔ TỪ CONTRACT ---
interface RawJob {
    foundation: string;
    position: string;
    field: string;
    salary: bigint;
    form: string;
    trend: number; // 0 for Up, 1 for Down
    iconId: string;
}

// --- KIỂU DỮ LIỆU CHO LOG EVENT ---
type DecodedJobLog = {
  args: {
    newJob?: RawJob;
  };
} & Log;

// --- HÀM HELPER ĐỂ ĐỊNH DẠNG DỮ LIỆU ---
const formatJobFromContract = (job: RawJob, index: number): JobData => {
  const getIcon = (iconId: string) => {
    switch (iconId.toLowerCase()) {
      case 'nvidia': return <SiNvidia size={24} className="text-gray-400"/>;
      case 'tesla': return <SiTesla size={24} className="text-gray-400"/>;
      case 'apple': return <SiApple size={24} className="text-gray-400"/>;
      case 'google': return <SiGoogle size={22} className="text-gray-400"/>;
      case 'fpt': return <SiTesla size={24} className="text-gray-400"/>; // Giả sử
      default: return null;
    }
  };

  return {
    id: `job-${index}-${job.foundation}`,
    foundation: job.foundation,
    position: job.position,
    field: job.field as JobField,
    salary: Number(job.salary),
    form: job.form,
    trend: job.trend === 0 ? 'up' : 'down',
    icon: getIcon(job.iconId),
  };
};

// --- CUSTOM HOOK ---
export function useRecruitingData() {
  const [jobs, setJobs] = useState<JobData[]>([]);

  // 1. Tải dữ liệu ban đầu
  const { data: initialJobs, isLoading, error, refetch } = useReadContract({
    address: dashboardAddress,
    abi: dashboardAbi,
    functionName: 'getAllJobs',
  });

  // 2. Xử lý dữ liệu ban đầu khi có kết quả
  useEffect(() => {
    if (initialJobs && Array.isArray(initialJobs)) {
      const formattedJobs = (initialJobs as RawJob[]).map(formatJobFromContract);
      setJobs(formattedJobs);
    }
  }, [initialJobs]);

  // 3. Lắng nghe event khi có job mới được thêm
  useWatchContractEvent({
    address: dashboardAddress,
    abi: dashboardAbi,
    eventName: 'JobAdded',
    onLogs(logs) {
      logs.forEach(log => {
        const decodedLog = log as DecodedJobLog;
        const newJobRaw = decodedLog.args.newJob;

        if (newJobRaw) {
          console.log("New job event received:", newJobRaw);
          // Giả sử job mới luôn có index cao nhất
          const formattedNewJob = formatJobFromContract(newJobRaw, jobs.length);
          // Thêm job mới vào đầu danh sách để hiển thị ngay lập tức
          setJobs(prevJobs => [formattedNewJob, ...prevJobs]);
        }
      });
    },
  });

  return { jobs, isLoading, error, refetch };
}