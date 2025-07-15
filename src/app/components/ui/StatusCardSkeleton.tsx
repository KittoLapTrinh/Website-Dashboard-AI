import React from 'react';

export const StatusCardSkeleton = () => (
  <div className="bg-zinc-900 rounded-2xl p-4 animate-pulse">
    <div className="h-5 bg-zinc-800 rounded w-3/4 mb-2"></div>
    <div className="h-4 bg-zinc-800 rounded w-1/2 mb-6"></div>
    <div className="flex flex-col items-center gap-y-2 mt-auto">
      <div className="h-8 w-8 bg-zinc-800 rounded-full"></div>
      <div className="h-4 bg-zinc-800 rounded w-2/3"></div>
    </div>
  </div>
);