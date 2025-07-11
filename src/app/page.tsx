// src/app/page.tsx

// Import các component widget của bạn
import TotalViewers from "./components/features/TotalViewers";
import ServiceStatusGrid from "./components/features/ServiceStatusGrid";
import ProfileOverview from "./components/features/ProfileOverview";
import RealTimeRecruiting from "./components/features/RecruitingTable"; // Tên đã sửa ở bước trước
import SupportFund from "./components/features/SupportFund";
import AnalyticsChart from "./components/features/AnalyticsChart";
import HeartbeatMonitor from "./components/features/HeartbeatMonitor";
import CustomerReaction from "./components/features/CustomerReaction";
import ChatbotDefinition from "./components/features/ChatbotDefinition";
import LetsConnect from "./components/layout/Footer";
export default function HomePage() {
  return (
    // Đây là layout grid chính, chứa TẤT CẢ các widget
     <div 
      // 👇 THAY ĐỔI MÀU Ở ĐÂY 👇
      className="bg-[#16191d] rounded-2xl p-8 
                 shadow-2xl shadow-black/40 
                  border border-white/5
                 grid grid-cols-1 lg:grid-cols-12 gap-6"
    >
      
      {/* === HÀNG 1 === */}
      <div className="lg:col-span-4">
        <TotalViewers />
      </div>
      <div className="lg:col-span-8">
        <ServiceStatusGrid />
      </div>

      {/* === HÀNG 2 === */}
      <div className="lg:col-span-12">
        <ProfileOverview />
      </div>

      {/* === HÀNG 3 === */}
      <div className="lg:col-span-8">
        <RealTimeRecruiting />
      </div>
      <div className="lg:col-span-4">
        <SupportFund />
      </div>
      {/* 2. THÊM HÀNG 4 MỚI */}
      {/* Analytics Chart chiếm 4 cột */}
      <div className="lg:col-span-4">
        <AnalyticsChart />
      </div>
      {/* Heartbeat Monitor chiếm 8 cột */}
      <div className="lg:col-span-8">
        <HeartbeatMonitor />
      </div>
       {/* 2. THÊM HÀNG 5 MỚI */}
      {/* Widget này chiếm toàn bộ 12 cột */}
      <div className="lg:col-span-12">
        <CustomerReaction />
      </div>
      {/* 2. THÊM HÀNG 6 & 7 */}
      <div className="lg:col-span-12">
        <ChatbotDefinition />
      </div>
      <div className="lg:col-span-12">
        <LetsConnect />
      </div>
    </div>
  );
}