"use client";

import React, { useState, useEffect } from 'react';

// Component này sẽ chỉ render `children` ở phía client
export const ClientOnly = ({ children }: { children: React.ReactNode }) => {
  const [hasMounted, setHasMounted] = useState(false);

  useEffect(() => {
    setHasMounted(true);
  }, []);

  if (!hasMounted) {
    // Trả về null hoặc một placeholder/spinner ở phía server
    return null; 
  }

  return <>{children}</>;
};