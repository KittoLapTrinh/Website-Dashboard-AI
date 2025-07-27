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
import DailyContent from "./components/features/DailyContent";
import AnimateOnScroll  from "./components/helpers/AnimateOnScroll";

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
      <AnimateOnScroll  className="lg:col-span-4" delay={0.1}>
        <TotalViewers />
      </AnimateOnScroll >
      <AnimateOnScroll  className="lg:col-span-8" delay={0.2}>
        <ServiceStatusGrid />
      </AnimateOnScroll >

      {/* === HÀNG 2 === */}
      <AnimateOnScroll className="lg:col-span-12" delay={0.1}>
        <ProfileOverview />
      </AnimateOnScroll>
      {/* <div className="lg:col-span-12">
        <DailyContent />
      </div> */}

      {/* === HÀNG 3 === */}
      <AnimateOnScroll className="lg:col-span-8" delay={0.1}>
        <RealTimeRecruiting />
      </AnimateOnScroll>
      <AnimateOnScroll className="lg:col-span-4" delay={0.2}>
        <SupportFund />
      </AnimateOnScroll>
      {/* 2. THÊM HÀNG 4 MỚI */}
      {/* Analytics Chart chiếm 4 cột */}
      <AnimateOnScroll className="lg:col-span-4" delay={0.1}>
        <AnalyticsChart />
      </AnimateOnScroll>
      {/* Heartbeat Monitor chiếm 8 cột */}
      <AnimateOnScroll className="lg:col-span-8" delay={0.2}>
        <HeartbeatMonitor />
      </AnimateOnScroll>
       {/* 2. THÊM HÀNG 5 MỚI */}
      {/* Widget này chiếm toàn bộ 12 cột */}
      <AnimateOnScroll className="lg:col-span-12"  delay={0.1}>
        <CustomerReaction />
      </AnimateOnScroll>
      {/* 2. THÊM HÀNG 6 & 7 */}
      <AnimateOnScroll className="lg:col-span-12"  delay={0.1}>
        <ChatbotDefinition />
      </AnimateOnScroll>
      <AnimateOnScroll className="lg:col-span-12"  delay={0.1}>
        <LetsConnect />
      </AnimateOnScroll>
    </div>
  );
}