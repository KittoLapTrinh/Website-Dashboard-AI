import React from 'react';

export const MonitorCardSkeleton = () => (
  <div className="bg-zinc-900 rounded-2xl p-4 h-full flex flex-col animate-pulse">
    {/* Header Skeleton */}
    <div className="flex items-center gap-x-3">
      <div className="w-6 h-6 rounded bg-zinc-800"></div>
      <div>
        <div className="h-4 bg-zinc-800 rounded w-24 mb-2"></div>
        <div className="h-4 bg-zinc-800 rounded w-16"></div>
      </div>
    </div>
    
    {/* Chart Area Skeleton */}
    <div className="flex-grow mt-4 h-24 bg-zinc-800/50 rounded-lg"></div>
  </div>
);