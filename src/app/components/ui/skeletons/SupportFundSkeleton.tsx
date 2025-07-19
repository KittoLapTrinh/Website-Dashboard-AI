import React from 'react';

// Một hàng item trong trạng thái chờ
const SkeletonItem = ({ indent }: { indent: string }) => (
  <div className={`flex items-center gap-x-4 ${indent}`}>
    <div className="bg-zinc-800 rounded-full p-1 flex items-center gap-x-3 w-60 h-10">
      <div className="w-8 h-8 rounded-full bg-zinc-700"></div>
      <div className="flex-grow space-y-2">
        <div className="h-4 bg-zinc-700 rounded w-3/4"></div>
      </div>
      <div className="flex-grow space-y-2">
        <div className="h-3 bg-zinc-700 rounded w-1/4"></div>
      </div>
      <div className="w-8 h-8 rounded-full bg-zinc-700"></div>
    </div>
    <div className="flex-grow h-px bg-dashed-horizontal"></div>
  </div>
);

// Component Skeleton chính
export const SupportFundSkeleton = () => {
  const indentationClasses = ['pl-4', 'pl-12', 'pl-20'];
  return (
    <div className="animate-pulse">
      {/* Skeleton cho danh sách item */}
      <div className="flex-grow flex flex-col gap-y-4">
        {indentationClasses.map((indent, index) => (
          <SkeletonItem key={index} indent={indent} />
        ))}
      </div>
    </div>
  );
};