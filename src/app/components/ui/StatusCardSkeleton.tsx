import React from 'react';

export const StatusCardSkeleton = () => (
  <div className="bg-zinc-900 rounded-2xl p-2 animate-pulse h-31">
    <div className="h-5 bg-zinc-800 rounded w-3/4 mb-2"></div>
    <div className="h-3 bg-zinc-800 rounded w-1/2 mb-4"></div>
    <div className="flex flex-col mt-auto">
      <div className="h-6 w-6 bg-zinc-800 rounded-full"></div>
      <div className="h-4 bg-zinc-800 rounded w-2/3"></div>
    </div>
  </div>
);