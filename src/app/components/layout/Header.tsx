// src/components/Header.tsx

import Link from 'next/link';
import Image from 'next/image';
import logo from '@/app/assets/image/logo.png';

const Header = () => {
  return (
    <header className="bg-black text-white ">
      <div className="container mx-auto flex items-center h-18 px-2">
          {/* CÁC LIÊN KẾT ĐÃ ĐƯỢC CẬP NHẬT */}
          <nav className="hidden md:flex items-center gap-x-5">
             <Link href="/" aria-label="Trang chủ">
              <Image
                src={logo}
                alt="Logo của công ty ABC" 
                width={100}
                height={100}
                className="h-12 w-12 "
              />
            </Link>
            <Link
              href="/getting-started"
              // Đã thay đổi từ text-sm thành text-base
              className="text-base font-medium text-gray-300 hover:text-white transition-colors"
            >
              Getting Started
            </Link>
            <Link
              href="/career"
              // Đã thay đổi từ text-sm thành text-base
              className="text-base font-medium text-gray-300 hover:text-white transition-colors"
            >
              Career
            </Link>
            <Link
              href="/products"
              // Đã thay đổi từ text-sm thành text-base
              className="text-base font-medium text-gray-300 hover:text-white transition-colors"
            >
              Product & Applications
            </Link>
          </nav>
        </div>

    </header>
  );
};

export default Header;