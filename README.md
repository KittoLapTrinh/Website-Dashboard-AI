```
.env.local:
GOOGLE_API_KEY=AIzaSyAt9vPcAE1kf1rsdKSUAI_f2CtjardPtmY
MTD_RPC_URL="https://rpc-proxy-sequoia.iqnb.com:8446"
MTD_PRIVATE_KEY="0x848d466692aaa9d23c69771276e498b06fea7c35d08f0b461e00149714225a48"
NEXT_PUBLIC_DASHBOARD_CONTRACT_ADDRESS="0x2dF65CDb1aa81CBa65c72129d8054BF730B86946"

# Địa chỉ cho Frontend (bắt đầu bằng NEXT_PUBLIC_)
NEXT_PUBLIC_DASHBOARD_FACADE_ADDRESS="0x72321414DA50a768a40BFa118aCE850A7fDedD5a"

# Địa chỉ cho Backend Go
SERVICE_STATUS_CONTRACT_ADDRESS="0xcc48E282F0136c30835946CC907e10cf98517Af8"
RECRUITING_CONTRACT_ADDRESS="0x048793bBC551254e19A33398167Fb80a9EAec9f1"
SUPPORT_FUND_CONTRACT_ADDRESS="0xc5AF6359A83F37CC8a10755cb6Cc79a1b605Ec50"
TIME_SERIES_CONTRACT_ADDRESS="0xB25201999Dea45fFedC1031274D9DAa730CB88a0"
KEY_VALUE_CONTRACT_ADDRESS="0x6B6250015B2e247E0994DFd5CbB50b4486898644"
CONTENT_CONTRACT_ADDRESS="0x89DE27611644350c53172648d46a135e76FC1ecc"

```

This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/app/building-your-application/optimizing/fonts) to automatically optimize and load [Geist](https://vercel.com/font), a new font family for Vercel.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
