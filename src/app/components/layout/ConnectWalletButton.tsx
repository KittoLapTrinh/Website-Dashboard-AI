"use client"

import React, { useState, useEffect } from 'react';
import { useAccount, useConnect, useDisconnect } from 'wagmi'
import { injected } from 'wagmi/connectors'

export function ConnectWalletButton() {
  const [hasMounted, setHasMounted] = useState(false);
  const { address, isConnected } = useAccount()
  const { connect } = useConnect()
  const { disconnect } = useDisconnect()
  
  useEffect(() => {
    setHasMounted(true);
  }, []);

  // Nếu chưa mount ở client, render một nút placeholder để tránh lỗi hydration
  if (!hasMounted) {
    return (
        <button
            className="bg-zinc-800 text-transparent px-4 py-1.5 rounded-md text-sm font-semibold animate-pulse"
            disabled
        >
            Connect Wallet
        </button>
    );
  }

  if (isConnected) {
    return (
      <div className="flex items-center gap-x-2">
        <span className="text-sm font-mono bg-zinc-800 px-3 py-1 rounded-md">
          {`${address?.slice(0, 6)}...${address?.slice(-4)}`}
        </span>
        <button
          onClick={() => disconnect()}
          className="bg-red-600 text-white px-4 py-1.5 rounded-md text-sm font-semibold hover:bg-red-700"
        >
          Disconnect
        </button>
      </div>
    )
  }

  return (
    <button
      onClick={() => connect({ connector: injected() })}
      className="bg-blue-600 text-white px-4 py-1.5 rounded-md text-sm font-semibold hover:bg-blue-700"
    >
      Connect Wallet
    </button>
  )
}