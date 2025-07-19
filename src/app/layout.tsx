// src/app/layout.tsx
// "use client";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from "./components/layout/Header";
import { Web3Provider } from "@/app/components/providers/Web3Provider";
const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "My Dashboard",
  description: "Analytics Dashboard",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      {/* THAY ĐỔI TỪ bg-black THÀNH bg-zinc-900 */}
      <body className={`${inter.className} bg-black`}>  
        <Web3Provider>   
          <Header />
          <main className="p-4 sm:p-6 lg:p-8">
            {children}
          </main>
        </Web3Provider>  
      </body>
    </html>
  );
}