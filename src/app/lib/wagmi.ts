import { http, createConfig } from 'wagmi'
import { type Chain } from 'viem' // ✨ Lấy kiểu Chain từ viem

// ✨ 1. ĐỊNH NGHĨA MẠNG MTD MAINNET
export const mtdMainnet = {
  id: 991, // Chain ID
  name: 'MTD Mainnet',
  nativeCurrency: { name: 'MTD Token', symbol: 'MTD', decimals: 18 },
  rpcUrls: {
    default: { http: ['https://rpc-proxy-sequoia.iqnb.com:8446'] },
  },
  // blockExplorers: { (Tùy chọn)
  //   default: { name: 'MTDScan', url: 'https://scan.mtd.com' },
  // },
} as const satisfies Chain

// ✨ 2. SỬ DỤNG MẠNG MTD TRONG CẤU HÌNH
export const config = createConfig({
  chains: [mtdMainnet], // Sử dụng chain của công ty bạn
  transports: {
    [mtdMainnet.id]: http(),
  },
})