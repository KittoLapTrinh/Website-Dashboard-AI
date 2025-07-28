// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../common/AccessControl.sol";

/**
 * @title KeyValueContract
 * @notice Lưu trữ các cặp key-value đơn giản.
 */
contract KeyValueContract is AccessControl {
    
    mapping(string => uint256) public singleValueData;
    mapping(string => int256) public returnData;

    event SingleValueUpdated(string key, uint256 newValue);
    event ReturnValueUpdated(string key, int256 newValue);

    constructor() AccessControl() {
        _initializeData();
    }

    // --- VIEW FUNCTIONS ---
    function getSingleValueData(string memory _key) external view returns (uint256) {
        return singleValueData[_key];
    }
    function getReturnData(string memory _key) external view returns (int256) {
        return returnData[_key];
    }

    // --- SETTER FUNCTIONS ---
    function updateSingleValueData(string memory _key, uint256 _newValue) external onlyOwner {
        singleValueData[_key] = _newValue;
        emit SingleValueUpdated(_key, _newValue);
    }
    function updateReturnData(string memory _key, int256 _newValue) external onlyOwner {
        returnData[_key] = _newValue;
        emit ReturnValueUpdated(_key, _newValue);
    }

    // --- HÀM KHỞI TẠO DỮ LIỆU MẪU ---
    function _initializeData() private {
        // Dữ liệu cho TotalViewers
        singleValueData["viewers_1D"] = 12304; returnData["return_1D"] = 35;
        singleValueData["viewers_1W"] = 89450; returnData["return_1W"] = 21;
        singleValueData["viewers_1M"] = 350123; returnData["return_1M"] = 58;
        singleValueData["viewers_6M"] = 1850678; returnData["return_6M"] = 122;
        singleValueData["viewers_1Y"] = 4500000; returnData["return_1Y"] = 251;
    }
}