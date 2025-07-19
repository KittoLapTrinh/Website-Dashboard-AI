import React from 'react';
import { Sun, Moon } from 'lucide-react';
import { SiTesla, SiApple, SiGoogle, SiNvidia } from 'react-icons/si';
import { FaMicrosoft } from 'react-icons/fa';

// ================================================================
// --- 1. ĐỊNH NGHĨA CÁC KIỂU DỮ LIỆU ---
// ================================================================



export type ServiceStatus = 'ok' | 'degraded' | 'error';
export interface ServiceData {
  id: string;
  serviceName: string; // Đã đổi tên từ 'name'
  title: string;
  subtitle: string;
  status: ServiceStatus;
  icon: React.ReactNode;
}

export type JobTrend = 'up' | 'down';
export type JobField = 'Blockchain' | 'AI' | 'IOT' | 'Web Dev';
export interface JobData {
  id: string;
  foundation: string;
  position: string;
  field: JobField;
  salary: number;
  form: string;
  trend: JobTrend;
  icon: React.ReactNode;
}

export interface FundItemData {
  name: string;
  price: number;
  icon: React.ReactNode;
  count: number;
  avatarColor: string;
}

export interface DataPoint {
  date: string;
  value: number;
  change?: number; 
}

export interface TestimonialData {
  id: number;
  quote: string;
  avatarSrc: string;
}

export interface TotalViewersData {
  viewers: number;
  // ✨ `return` giờ là một object
  return: {
    value: number;
    status: 'increase' | 'decrease' | 'neutral';
  }
}
