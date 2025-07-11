import React from 'react';
import WidgetCard from '../ui/WidgetCard';
import MonitorCard from '../ui/MonitorCard';
import { HeartPulse, Thermometer, Activity } from 'lucide-react';

const HeartbeatMonitor = () => {
  return (
    <WidgetCard>
      {/* ✨ 1. Thêm flex và flex-col để kiểm soát layout */}
      <div className="flex flex-col h-full">
        <h2 className="text-2xl font-bold text-white mb-6">Heartbeat Monitor</h2>
        
        {/* ✨ 2. Thêm flex-grow và bọc grid trong một div flex nữa */}
        <div className="flex-grow flex items-center">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 w-full">
            
            {/* === Thẻ 1: Blood Status === */}
            <MonitorCard
              title="Blood Status"
              value="102/70"
              icon={<HeartPulse size={24} />}
              statusColor="green"
            >
              <>
                <svg width="100%" height="100%" viewBox="0 0 100 50" preserveAspectRatio="none">
                  <path d="M 0 30 C 10 10, 20 40, 30 30 C 40 20, 50 40, 60 30 C 70 20, 80 40" stroke="white" strokeWidth="2" fill="none" />
                </svg>
                <div className="absolute right-0 bottom-0 text-right">
                  <p className="text-white font-bold text-2xl">102</p>
                  <p className="text-gray-400 text-sm">/70</p>
                </div>
              </>
            </MonitorCard>

            {/* === Thẻ 2: Temperature === */}
            <MonitorCard
              title="Temperature"
              value="34.1"
              icon={<Thermometer size={24} />}
              statusColor="green"
            >
              <>
                <svg width="100%" height="100%" viewBox="0 0 100 50" preserveAspectRatio="none">
                  <path d="M 0 40 L 60 40 C 70 40, 80 20, 95 50" stroke="white" strokeWidth="2" fill="none" />
                </svg>
                <div className="absolute right-0 bottom-0 bg-zinc-800/80 backdrop-blur-sm p-2 rounded-lg border border-zinc-700">
                  <p className="text-white font-bold text-lg">34.1</p>
                </div>
              </>
            </MonitorCard>

            {/* === Thẻ 3: Heart Rate === */}
            <MonitorCard
              title="Heart Rate"
              value="124 bpm"
              icon={<Activity size={24} />}
              statusColor="green"
            >
              <>
                <svg width="100%" height="100%" viewBox="0 0 100 50" preserveAspectRatio="none">
                  <path d="M 0 30 L 10 30 L 15 20 L 25 40 L 35 10 L 45 30 L 55 25 L 65 30 L 75 20 L 85 35 L 95 30" stroke="white" strokeWidth="2" fill="none" />
                </svg>
                <div className="absolute right-0 bottom-0 bg-zinc-800/80 backdrop-blur-sm p-2 rounded-lg border border-zinc-700">
                  <p className="text-white font-bold text-lg">102</p>
                  <p className="text-gray-400 text-xs">bpm</p>
                </div>
              </>
            </MonitorCard>

          </div>
        </div>
      </div>
    </WidgetCard>
  );
};

export default HeartbeatMonitor;