import React from 'react';
import { Linkedin, Mail, Phone, Instagram, Facebook } from 'lucide-react';

// Dữ liệu có thể giữ nguyên
const services = [
  { title: 'Chat Blockchain', description: 'A 20% menu leave rate indicates that a significant' },
  { title: 'Email Blockchain', description: 'A 20% menu leave rate indicates that a significant' },
  { title: 'Wallet', description: 'A 20% menu leave rate indicates that a significant' },
  { title: 'AI Camera', description: 'A 20% menu leave rate indicates that a significant' },
  { title: 'Chain Analytics', description: 'A 20% menu leave rate indicates that a significant' },
  { title: 'AI Speed To Text', description: 'A 20% menu leave rate indicates that a significant' },
];

const socialIcons = [
  { icon: <Linkedin size={20} />, href: '#' },
  { icon: <Mail size={20} />, href: '#' },
  { icon: <Phone size={20} />, href: '#' },
  { icon: <Instagram size={20} />, href: '#' },
  { icon: <Facebook size={20} />, href: '#' },
];

// 1. Đổi tên component thành "Footer"
const Footer = () => {
  return (
    // 2. Thay thế WidgetCard bằng thẻ `<footer>` với styling riêng
    <footer className="bg-black text-white p-8 lg:p-12 mt-12 rounded-2xl">
      {/* 3. Bỏ class chiều cao cố định h-[450px] để nó tự co giãn */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Cột trái: Thông tin liên hệ */}
        <div className="lg:col-span-1 flex flex-col justify-center">
          <p className="text-5xl font-light text-[#919191]">Let's</p>
          <p className="text-6xl font-bold text-white mt-2">Connect</p>
          <p className="text-lg text-[#919191] mt-6">Example@vn.ai</p>
          <div className="flex items-center gap-x-4 mt-8">
            {socialIcons.map((social, index) => (
              <a key={index} href={social.href} className="text-white hover:text-white transition-colors">
                {social.icon}
              </a>
            ))}
          </div>
        </div>

        {/* Cột phải: Lưới dịch vụ */}
        <div className="lg:col-span-2 grid grid-cols-1 sm:grid-cols-2 gap-6 place-content-center">
          {services.map((service, index) => (
            <div key={index}>
              <h3 className="text-lg font-bold text-white">{service.title}</h3>
              <p className="text-sm text-[#FFFFFF] mt-2">{service.description}</p>
            </div>
          ))}
        </div>
      </div>
    </footer>
  );
};

// 4. Export component đã đổi tên
export default Footer;