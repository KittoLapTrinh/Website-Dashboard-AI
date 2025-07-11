import { Apple, ArrowUpRight } from 'lucide-react';
// ✨ 1. Import các icon thương hiệu từ thư viện react-icons
import { SiTesla, SiGoogle, SiNvidia, SiApple } from 'react-icons/si';
import StatusCard from '../ui/StatusCard';
import WidgetCard from '../ui/WidgetCard';
import { FaMicrosoft } from 'react-icons/fa';
// ✨ 2. Cập nhật mảng services để sử dụng các icon mới
const services = [
  { title: '1.7M User', subtitle: '2.2M Tx (0.01s)', icon: <SiApple size={20} />, name: 'Wallet Blockchain', status: 'ok' as const },
  { title: 'Volum 125M', subtitle: '100,000 TPS', icon: <SiTesla size={20} />, name: 'Chain Analytics', status: 'error' as const },
  { title: '1.7M User', subtitle: '2.2M Mail (0.01s)', icon: <FaMicrosoft size={20} />, name: 'Dapp Email', status: 'ok' as const },
  { title: '1.7M User', subtitle: '2.2M Course', icon: <SiGoogle size={20} />, name: 'Dapp Elearning', status: 'ok' as const },
  { title: '1.7M User', subtitle: '1.2M Messages(0.01s)', icon: <SiNvidia size={20} />, name: 'Dapp Chat', status: 'ok' as const },
];

const ServiceStatusGrid = () => {
  return (
    <WidgetCard>
      {/* ✨ 3. Sửa lại nút "See all" thành một nút liền mạch */}
      <div className="flex justify-end mb-4">
        <button 
          className="flex items-center overflow-hidden  
                     border border-gray-400 rounded-full 
                     text-sm text-gray-300 
                     hover:bg-zinc-800 hover:border-gray-200 transition-all"
        >
          <span className="px-4 py-1.5">
            See all
          </span>
        </button>
         <button 
          className="flex items-center overflow-hidden  
                     border border-gray-400 rounded-full 
                     text-sm text-gray-300 
                     hover:bg-zinc-800 hover:border-gray-200 transition-all"
        >
          
          <span className="p-2 border-l border-gray-400">
            <ArrowUpRight size={16} />
          </span>
        </button>
      </div>

      {/* Layout grid được giữ nguyên để căn giữa và giới hạn chiều rộng */}
      <div className="flex justify-center">
        <div className="grid grid-cols-5 gap-4 max-w-5xl">
          {services.map((service, index) => (
            <StatusCard
              key={index}
              title={service.title}
              subtitle={service.subtitle}
              icon={service.icon}
              name={service.name}
              status={service.status}
            />
          ))}
        </div>
      </div>
    </WidgetCard>
  );
};

export default ServiceStatusGrid;