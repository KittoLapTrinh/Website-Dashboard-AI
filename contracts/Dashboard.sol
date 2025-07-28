// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Dashboard {
    address public immutable owner;

    enum ServiceStatus { Ok, Degraded, Error }
    enum JobTrend { Up, Down }

    struct Service {
        string name;
        string title;
        string subtitle;
        ServiceStatus status;
        string iconId;
    }
    struct Job {
        string foundation;
        string position;
        string field;
        uint256 salary;
        string form;
        JobTrend trend;
        string iconId;
    }
    struct FundItem {
        string name;
        uint256 price;
        string iconId;
        uint256 count;
        string avatarColor;
    }
    struct DataPoint {
        uint256 timestamp;
        uint256 value;
         int256 change;
    }
    struct Testimonial {
        string quote;
        string avatarId;
    }

    // --- EVENTS ---
    event SingleValueUpdated(string key, uint256 newValue);
    event ReturnValueUpdated(string key, int256 newValue);
    event ServiceStatusUpdated(uint256 index, ServiceStatus newStatus);
    event JobAdded(Job newJob);
    event FundItemsUpdated(string dayKey, FundItem[] newItems);
    event DataPointAdded(string key, DataPoint newPoint);


    
    // --- STATE VARIABLES ---
    Service[] public services;
    Job[] public jobs;
    mapping(string => FundItem[]) public fundItemsByDay;
    Testimonial[] public testimonials;
    mapping(string => DataPoint[]) public timeSeriesData;
    mapping(string => uint256) public singleValueData;
    mapping(string => int256) public returnData;

    // --- MODIFIER & CONSTRUCTOR ---
    modifier onlyOwner() { require(msg.sender == owner, "NOT_OWNER"); _; }
    constructor() { owner = msg.sender; _initializeData(); }

    // --- VIEW FUNCTIONS ---
    function getAllServices() external view returns (Service[] memory) { return services; }
    function getAllJobs() external view returns (Job[] memory) { return jobs; }
    function getAllTestimonials() external view returns (Testimonial[] memory) { return testimonials; }
    function getFundItemsByDay(string memory _day) external view returns (FundItem[] memory) { return fundItemsByDay[_day]; }
    function getTimeSeriesData(string memory _key) external view returns (DataPoint[] memory) { return timeSeriesData[_key]; }
    function getSingleValueData(string memory _key) external view returns (uint256) { return singleValueData[_key]; }
    function getReturnData(string memory _key) external view returns (int256) { return returnData[_key]; }

    // --- SETTER FUNCTIONS ---

     function updateTotalViewersData(string memory _timeframe, uint256 _viewers, int256 _returnPct) external onlyOwner {
        string memory viewersKey = string(abi.encodePacked("viewers_", _timeframe));
        string memory returnKey = string(abi.encodePacked("return_", _timeframe));
        singleValueData[viewersKey] = _viewers;
        returnData[returnKey] = _returnPct;
        emit SingleValueUpdated(viewersKey, _viewers);
        emit ReturnValueUpdated(returnKey, _returnPct);
    }
    
   function updateServiceStatusByIndex(uint256 _index, ServiceStatus _newStatus) external onlyOwner {
    require(_index < services.length, "INVALID_INDEX");
    services[_index].status = _newStatus;
    emit ServiceStatusUpdated(_index, _newStatus); // Chỉ emit index và status mới
}
    

    function addJob(Job memory _newJob) external onlyOwner {
        jobs.push(_newJob);
        emit JobAdded(_newJob);
    }

    function updateFundItemsForDay(string memory _dayKey, FundItem[] memory _newItems) external onlyOwner {
    // ✨ BƯỚC 1: Xóa toàn bộ dữ liệu cũ của ngày đó
    delete fundItemsByDay[_dayKey];
    
    // ✨ BƯỚC 2: Dùng vòng lặp để thêm lại từng phần tử mới
    for (uint i = 0; i < _newItems.length; i++) {
        fundItemsByDay[_dayKey].push(_newItems[i]);
    }

    // ✨ BƯỚC 3: Phát event (nên phát sau khi đã cập nhật xong)
    // Lưu ý: Phát event với toàn bộ mảng có thể tốn nhiều gas
    emit FundItemsUpdated(_dayKey, fundItemsByDay[_dayKey]);
}

    function pushDataPoint(string memory _key, uint256 _value, int256 _change) external onlyOwner {
        DataPoint memory newPoint = DataPoint(block.timestamp, _value, _change);
        timeSeriesData[_key].push(newPoint);
        emit DataPointAdded(_key, newPoint);
    }

    

    // --- HÀM KHỞI TẠO DỮ LIỆU MẪU ---
    // --- HÀM KHỞI TẠO DỮ LIỆU MẪU ---
    function _initializeData() private {
        // === Dữ liệu cho ServiceStatusGrid (Giữ nguyên) ===
        services.push(Service("Wallet Blockchain", "1.7M User", "2.2M Tx (0.01s)", ServiceStatus.Ok, "apple"));
        services.push(Service("Chain Analytics", "Volum 125M", "100,000 TPS", ServiceStatus.Error, "tesla"));
        services.push(Service("Dapp Email", "1.7M User", "2.2M Mail (0.01s)", ServiceStatus.Ok, "microsoft"));
        services.push(Service("Dapp Elearning", "1.7M User", "2.2M Course", ServiceStatus.Ok, "google"));
        services.push(Service("Dapp Chat", "1.7M User", "1.2M Messages(0.01s)", ServiceStatus.Ok, "nvidia"));

         // 2. RecruitingTable
        jobs.push(Job("Nvidia", "Deep Learning Scientist", "AI", 55000, "Remote", JobTrend.Up, "nvidia"));
        jobs.push(Job("Tesla", "Firmware Engineer", "IOT", 45000, "Fulltime", JobTrend.Up, "tesla"));
        jobs.push(Job("ConsenSys", "Smart Contract Dev", "Blockchain", 42000, "Fulltime", JobTrend.Down, "apple")); // Giả sử ConsenSys dùng icon Apple
        jobs.push(Job("Google", "Frontend Developer", "Web Dev", 38000, "Fulltime", JobTrend.Up, "google"));
        jobs.push(Job("Apple", "AI Research Intern", "AI", 32000, "Remote", JobTrend.Down, "apple"));
        jobs.push(Job("FPT", "Blockchain Engineer", "Blockchain", 26000, "Fulltime", JobTrend.Up, "tesla"));
        
        // 3. SupportFund
        fundItemsByDay["Sun 20"].push(FundItem("Albufin", 26, "sun", 1, "bg-white"));
        fundItemsByDay["Sun 20"].push(FundItem("Vitamin C", 15, "sun", 1, "bg-gray-600"));

        // Dữ liệu cho 'Sat 19'
        fundItemsByDay["Sat 19"].push(FundItem("Vitamin D", 26, "sun", 2, "bg-gray-600"));
        fundItemsByDay["Sat 19"].push(FundItem("Omega 3", 26, "moon", 2, "bg-gray-600"));
        fundItemsByDay["Sat 19"].push(FundItem("Magnesium", 20, "moon", 1, "bg-gray-600"));

        // Dữ liệu cho 'Fri 18'

        fundItemsByDay["Fri 18"].push(FundItem("Omega 3", 26, "moon", 2, "bg-gray-600"));
        fundItemsByDay["Fri 18"].push(FundItem("Iron", 18, "sun", 1, "bg-gray-600"));
        fundItemsByDay["Fri 18"].push(FundItem("Zinc", 12, "sun", 1, "bg-gray-600"));


        fundItemsByDay["Thu 17"].push(FundItem("Omega 3", 26, "moon", 2, "bg-gray-600"));
        
        // === Dữ liệu cho TotalViewers (Giữ nguyên) ===
        singleValueData["viewers_1D"] = 12304; returnData["return_1D"] = 35;
        singleValueData["viewers_1W"] = 89450; returnData["return_1W"] = 21;
        singleValueData["viewers_1M"] = 350123; returnData["return_1M"] = 58;
        singleValueData["viewers_6M"] = 1850678; returnData["return_6M"] = 122;
        singleValueData["viewers_1Y"] = 4500000; returnData["return_1Y"] = 251;
        
        // === Dữ liệu cho ProfileOverview (ĐÃ MỞ RỘNG) ===
        
        // Dữ liệu cho 1 Ngày (mỗi 2 giờ)
        string memory key1D = "profileOverview_1D";
        timeSeriesData[key1D].push(DataPoint(block.timestamp - 22 hours, 145000, 0));
        timeSeriesData[key1D].push(DataPoint(block.timestamp - 18 hours, 148000, 21)); // +2.1%
        timeSeriesData[key1D].push(DataPoint(block.timestamp - 14 hours, 151004, 5)); // +0.5%
        timeSeriesData[key1D].push(DataPoint(block.timestamp - 10 hours, 147090, -8)); // -0.8%
        timeSeriesData[key1D].push(DataPoint(block.timestamp - 6 hours, 149000, 13)); // +1.3%
        timeSeriesData[key1D].push(DataPoint(block.timestamp - 2 hours, 152000, 20)); // +2.0%

        // Dữ liệu cho 1 Tuần
        string memory key1W = "profileOverview_1W";
        timeSeriesData[key1W].push(DataPoint(block.timestamp - 6 days, 150000, 0));
        timeSeriesData[key1W].push(DataPoint(block.timestamp - 5 days, 152000, 13)); // +1.3%
        timeSeriesData[key1W].push(DataPoint(block.timestamp - 4 days, 151500, -3)); // -0.3%
        timeSeriesData[key1W].push(DataPoint(block.timestamp - 3 days, 154408, 12)); // +1.2%
        timeSeriesData[key1W].push(DataPoint(block.timestamp - 2 days, 153000, -9)); // -0.9%
        timeSeriesData[key1W].push(DataPoint(block.timestamp - 1 days, 155000, 13)); // +1.3%
        timeSeriesData[key1W].push(DataPoint(block.timestamp, 156000, 6)); // +0.6%

        // Dữ liệu cho 1 Tháng
        string memory key1M = "profileOverview_1M";
        timeSeriesData[key1M].push(DataPoint(block.timestamp - 25 days, 140000, 0));
        timeSeriesData[key1M].push(DataPoint(block.timestamp - 15 days, 160000, 143)); // +14.3%
        timeSeriesData[key1M].push(DataPoint(block.timestamp - 5 days, 155000, -31)); // -3.1%

        // Dữ liệu cho 6 Tháng
        string memory key6M = "profileOverview_6M";
        timeSeriesData[key6M].push(DataPoint(block.timestamp - 150 days, 150000, 0));
        timeSeriesData[key6M].push(DataPoint(block.timestamp - 120 days, 175000, 167)); // +16.7%
        timeSeriesData[key6M].push(DataPoint(block.timestamp - 90 days, 145000, -171)); // -17.1%
        timeSeriesData[key6M].push(DataPoint(block.timestamp - 60 days, 131323, -23)); // -2.3%
        timeSeriesData[key6M].push(DataPoint(block.timestamp - 30 days, 148000, 127)); // +12.7%
        timeSeriesData[key6M].push(DataPoint(block.timestamp, 145000, -20)); // -2.0%

        // Dữ liệu cho 1 Năm
        string memory key1Y = "profileOverview_1Y";
        timeSeriesData[key1Y].push(DataPoint(block.timestamp - 11 * 30 days, 120000, 0));
        timeSeriesData[key1Y].push(DataPoint(block.timestamp - 8 * 30 days, 180000, 500));
        timeSeriesData[key1Y].push(DataPoint(block.timestamp - 5 * 30 days, 160000, -111));
        timeSeriesData[key1Y].push(DataPoint(block.timestamp - 2 * 30 days, 220000, 375));

        // 6. AnalyticsChart
        timeSeriesData["analytics_Weekly"].push(DataPoint(block.timestamp - 6 days, 40, 0));
        timeSeriesData["analytics_Weekly"].push(DataPoint(block.timestamp - 3 days, 70, 750)); // +75%
        timeSeriesData["analytics_Weekly"].push(DataPoint(block.timestamp, 80, 143)); // +14.3%

        // Dữ liệu cho Monthly
        timeSeriesData["analytics_Monthly"].push(DataPoint(block.timestamp - 25 days, 50, 0));
        timeSeriesData["analytics_Monthly"].push(DataPoint(block.timestamp - 15 days, 30, -400)); // -40%
        timeSeriesData["analytics_Monthly"].push(DataPoint(block.timestamp - 5 days, 60, 1000)); // +100%
        timeSeriesData["analytics_Monthly"].push(DataPoint(block.timestamp, 55, -83)); // -8.3%
        
        // Dữ liệu cho Yearly
        timeSeriesData["analytics_Yearly"].push(DataPoint(block.timestamp - 10 * 30 days, 20, 0)); // ~10 tháng trước
        timeSeriesData["analytics_Yearly"].push(DataPoint(block.timestamp - 6 * 30 days, 45, 1250)); // +125%
        timeSeriesData["analytics_Yearly"].push(DataPoint(block.timestamp - 2 * 30 days, 35, -222)); // -22.2%
        timeSeriesData["analytics_Yearly"].push(DataPoint(block.timestamp, 50, 428)); // +42.8%

        // 7. HeartbeatMonitor
        timeSeriesData["bloodStatus"].push(DataPoint(block.timestamp, 93, 0));
        timeSeriesData["temperature"].push(DataPoint(block.timestamp, 37, 0));
        timeSeriesData["heartRate"].push(DataPoint(block.timestamp, 115, 0));

        // 8. CustomerReaction
        testimonials.push(Testimonial("Wow, this AI solution could revolutionize everything.", "Customer1"));
    }
}