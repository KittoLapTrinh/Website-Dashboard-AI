// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../common/DataStructs.sol";
import "../common/AccessControl.sol";

contract SupportFundContract is AccessControl {
    mapping(string => DataStructs.FundItem[]) public fundItemsByDay;
    event FundItemsUpdated(string dayKey, DataStructs.FundItem[] newItems);

    // ✨ Constructor không còn tham số
    constructor() AccessControl() {
        _initializeData();
    }

    function updateFundItemsForDay(string memory _dayKey, DataStructs.FundItem[] memory _newItems) external onlyOwner {
        delete fundItemsByDay[_dayKey];
        for (uint i = 0; i < _newItems.length; i++) {
            fundItemsByDay[_dayKey].push(_newItems[i]);
        }
        emit FundItemsUpdated(_dayKey, fundItemsByDay[_dayKey]);
    }

    function getFundItemsByDay(string memory _day) external view returns (DataStructs.FundItem[] memory) {
        return fundItemsByDay[_day];
    }

    // ✨ Hàm khởi tạo dữ liệu mẫu
    function _initializeData() private {
        fundItemsByDay["Mon 21"].push(DataStructs.FundItem("Albufin", 26, "sun", 1, "bg-white"));
        fundItemsByDay["Sun 20"].push(DataStructs.FundItem("Vitamin D", 26, "sun", 2, "bg-gray-600"));
        fundItemsByDay["Sat 19"].push(DataStructs.FundItem("Omega 3", 26, "moon", 2, "bg-gray-600"));
        fundItemsByDay["Fri 18"].push(DataStructs.FundItem("Iron", 18, "sun", 1, "bg-gray-600"));
    }
}