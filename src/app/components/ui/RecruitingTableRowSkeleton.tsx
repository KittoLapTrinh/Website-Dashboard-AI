import React from 'react';

export const RecruitingTableRowSkeleton = () => (
  <div className="grid grid-cols-6 gap-x-4 items-center p-4 border-b border-gray-800/60">
    {/* Cột 1: Foundation */}
    <div className="flex items-center gap-x-3">
      <div className="w-6 h-6 rounded-full bg-zinc-800"></div>
      <div className="h-4 bg-zinc-800 rounded w-20"></div>
    </div>
    {/* Cột 2: Job Position */}
    <div className="h-4 bg-zinc-800 rounded w-3/4"></div>
    {/* Cột 3: Field */}
    <div className="h-4 bg-zinc-800 rounded w-1/2"></div>
    {/* Cột 4: Salary */}
    <div className="h-4 bg-zinc-800 rounded w-2/3"></div>
    {/* Cột 5: Form */}
    <div className="h-4 bg-zinc-800 rounded w-1/2"></div>
    {/* Cột 6: Update Chart */}
    <div className="h-7 w-14 bg-zinc-800 rounded"></div>
  </div>
);