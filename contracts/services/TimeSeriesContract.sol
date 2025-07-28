// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../common/DataStructs.sol";
import "../common/AccessControl.sol";

contract TimeSeriesContract is AccessControl {
    
    mapping(string => DataStructs.DataPoint[]) public timeSeriesData;

    // --- EVENTS ---
    event DataPointAdded(string key, DataStructs.DataPoint newPoint);
    event TimeSeriesReplaced(string key, uint256 newSize);

    constructor() AccessControl() {
        _initializeData();
    }

    // ====================================================================
    // === CÁC HÀM GHI DỮ LIỆU (WRITE FUNCTIONS) ===
    // ====================================================================

    /**
     * @notice Thêm một điểm dữ liệu mới vào cuối một chuỗi.
     * @dev Dùng cho các biểu đồ cần lưu trữ toàn bộ lịch sử như Analytics & Profile Overview.
     */
    function pushDataPoint(string memory _key, uint256 _value, int256 _change) external onlyOwner {
        DataStructs.DataPoint memory newPoint = DataStructs.DataPoint(block.timestamp, _value, _change);
        timeSeriesData[_key].push(newPoint);
        emit DataPointAdded(_key, newPoint);
    }

    /**
     * @notice Thay thế toàn bộ dữ liệu của một chuỗi bằng một bộ dữ liệu mới.
     * @dev Dùng cho các biểu đồ chỉ cần hiển thị dữ liệu gần nhất (cửa sổ trượt) như Heartbeat Monitor.
     */
    // File: TimeSeriesContract.sol

/**
 * @notice Thay thế toàn bộ dữ liệu của một chuỗi bằng một bộ dữ liệu mới.
 * @dev Dùng cho các biểu đồ chỉ cần hiển thị dữ liệu gần nhất (cửa sổ trượt) như Heartbeat Monitor.
 */
function replaceTimeSeries(string memory _key, DataStructs.DataPoint[] memory _newData) external onlyOwner {
    // ✨ BƯỚC 1: Xóa toàn bộ dữ liệu cũ trong mảng storage ✨
    delete timeSeriesData[_key];

    // ✨ BƯỚC 2: Lặp qua mảng mới và push từng phần tử vào storage ✨
    for (uint i = 0; i < _newData.length; i++) {
        timeSeriesData[_key].push(_newData[i]);
    }
    
    emit TimeSeriesReplaced(_key, _newData.length);
}

    // ===================================================================
    // === CÁC HÀM ĐỌC DỮ LIỆU (VIEW FUNCTIONS) ===
    // ===================================================================

    /**
     * @notice Lấy toàn bộ dữ liệu của một chuỗi.
     * @dev Phù hợp cho các chuỗi dữ liệu nhỏ, có kích thước cố định như Heartbeat Monitor.
     */
    function getTimeSeriesData(string memory _key) external view returns (DataStructs.DataPoint[] memory) {
        return timeSeriesData[_key];
    }

    /**
     * @notice Lấy tổng số điểm dữ liệu của một chuỗi.
     * @dev Dùng cho frontend để tính toán logic phân trang cho các chuỗi dữ liệu lớn.
     */
    function getTimeSeriesLength(string memory _key) external view returns (uint256) {
        return timeSeriesData[_key].length;
    }

    /**
     * @notice Lấy một "lát cắt" (slice) của dữ liệu từ một chuỗi.
     * @dev Hàm chính cho việc phân trang, dùng cho Analytics & Profile Overview.
     */
    function getTimeSeriesSlice(string memory _key, uint256 _startIndex, uint256 _count) external view returns (DataStructs.DataPoint[] memory) {
        uint256 len = timeSeriesData[_key].length;

        if (_startIndex >= len) {
            return new DataStructs.DataPoint[](0);
        }
        
        uint256 endIndex = _startIndex + _count;
        if (endIndex > len) {
            endIndex = len;
        }

        uint256 sliceLength = endIndex - _startIndex;
        DataStructs.DataPoint[] memory slice = new DataStructs.DataPoint[](sliceLength);

        for (uint256 i = 0; i < sliceLength; i++) {
            slice[i] = timeSeriesData[_key][_startIndex + i];
        }

        return slice;
    }


    // ===================================================================
    // === HÀM KHỞI TẠO DỮ LIỆU MẪU (GIỮ NGUYÊN) ===
    // ===================================================================
    function _initializeData() private {
        // === DỮ LIỆU CHO PROFILE OVERVIEW ===
        string memory key1D = "profileOverview_1D";
        timeSeriesData[key1D].push(DataStructs.DataPoint(block.timestamp - 22 hours, 145000, 0));
        timeSeriesData[key1D].push(DataStructs.DataPoint(block.timestamp - 10 hours, 147090, -8));
        timeSeriesData[key1D].push(DataStructs.DataPoint(block.timestamp - 2 hours, 152000, 20));
        
        string memory key1W = "profileOverview_1W";
        timeSeriesData[key1W].push(DataStructs.DataPoint(block.timestamp - 6 days, 150000, 0));
        timeSeriesData[key1W].push(DataStructs.DataPoint(block.timestamp - 3 days, 154408, 12));
        timeSeriesData[key1W].push(DataStructs.DataPoint(block.timestamp, 156000, 6));

        string memory key1M = "profileOverview_1M";
        timeSeriesData[key1M].push(DataStructs.DataPoint(block.timestamp - 25 days, 140000, 0));
        timeSeriesData[key1M].push(DataStructs.DataPoint(block.timestamp - 15 days, 160000, 143));
        timeSeriesData[key1M].push(DataStructs.DataPoint(block.timestamp - 5 days, 155000, -31));

        string memory key6M = "profileOverview_6M";
        timeSeriesData[key6M].push(DataStructs.DataPoint(block.timestamp - 150 days, 150000, 0));
        timeSeriesData[key6M].push(DataStructs.DataPoint(block.timestamp - 90 days, 145000, -171));
        timeSeriesData[key6M].push(DataStructs.DataPoint(block.timestamp, 145000, -20));

        string memory key1Y = "profileOverview_1Y";
        timeSeriesData[key1Y].push(DataStructs.DataPoint(block.timestamp - 11 * 30 days, 120000, 0));
        timeSeriesData[key1Y].push(DataStructs.DataPoint(block.timestamp - 5 * 30 days, 160000, -111));
        timeSeriesData[key1Y].push(DataStructs.DataPoint(block.timestamp - 2 * 30 days, 220000, 375));

        // === DỮ LIỆU CHO ANALYTICS CHART ===
        timeSeriesData["analytics_Weekly"].push(DataStructs.DataPoint(block.timestamp - 6 days, 40, 0));
        timeSeriesData["analytics_Weekly"].push(DataStructs.DataPoint(block.timestamp - 3 days, 70, 750));
        timeSeriesData["analytics_Weekly"].push(DataStructs.DataPoint(block.timestamp, 80, 143));

        timeSeriesData["analytics_Monthly"].push(DataStructs.DataPoint(block.timestamp - 25 days, 50, 0));
        timeSeriesData["analytics_Monthly"].push(DataStructs.DataPoint(block.timestamp - 15 days, 30, -400));
        timeSeriesData["analytics_Monthly"].push(DataStructs.DataPoint(block.timestamp, 55, -83));
        
        timeSeriesData["analytics_Yearly"].push(DataStructs.DataPoint(block.timestamp - 10 * 30 days, 20, 0));
        timeSeriesData["analytics_Yearly"].push(DataStructs.DataPoint(block.timestamp - 6 * 30 days, 45, 1250));
        timeSeriesData["analytics_Yearly"].push(DataStructs.DataPoint(block.timestamp, 50, 428));
        
        // === DỮ LIỆU CHO HEARTBEAT MONITOR ===
        // Khởi tạo với một vài điểm dữ liệu để trông tự nhiên hơn
        for (uint i = 0; i < 10; i++) {
            timeSeriesData["bloodStatus"].push(DataStructs.DataPoint(block.timestamp - (10-i)*5 seconds, uint256(90 + i), 0));
            timeSeriesData["temperature"].push(DataStructs.DataPoint(block.timestamp - (10-i)*5 seconds, 37, 0));
            timeSeriesData["heartRate"].push(DataStructs.DataPoint(block.timestamp - (10-i)*5 seconds, uint256(70 + i), 0));
        }
    }
}