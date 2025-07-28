// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title DataStructs
 * @notice Thư viện này định nghĩa các cấu trúc dữ liệu chung cho toàn bộ dashboard.
 * @dev Dùng `library` thay vì `contract` vì nó không có state variable và chỉ chứa định nghĩa.
 */
library DataStructs {
    
    // --- ENUMS ---

    enum ServiceStatus { Ok, Degraded, Error }
    enum JobTrend { Up, Down }
    enum ContentType { INSPIRATION, RESEARCH, NEWS }
    // --- STRUCTS ---

    // Dùng cho ServiceStatusGrid
    struct Service {
        string name;
        string title;
        string subtitle;
        ServiceStatus status;
        string iconId;
    }

    // Dùng cho RecruitingTable
    struct Job {
        string foundation;
        string position;
        string field;
        uint256 salary;
        string form;
        JobTrend trend;
        string iconId;
        string salaryText;
    }

    // Dùng cho SupportFund
    struct FundItem {
        string name;
        uint256 price;
        string iconId; // "sun", "moon"
        uint256 count;
        string avatarColor;
    }
    
    // Dùng cho các biểu đồ (ProfileOverview, AnalyticsChart, HeartbeatMonitor)
    struct DataPoint {
        uint256 timestamp;
        uint256 value;
        int256 change;
    }

    // Dùng cho CustomerReaction
    struct Testimonial {
        string quote;
        string avatarId; // "Customer1", "Customer2"
    }

    struct ContentItem {
        uint256 id;
        ContentType contentType;
        string title;
        string description;
        string imageUrl; // Đổi tên để rõ ràng hơn
        string author;
        uint256 timestamp;
    }
}