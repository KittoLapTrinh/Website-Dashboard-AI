// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../common/DataStructs.sol";
import "../common/AccessControl.sol";

contract ServiceStatusContract is AccessControl {
    DataStructs.Service[] public services;
    event ServiceStatusUpdated(uint256 index, DataStructs.ServiceStatus newStatus);

    // ✨ Constructor không còn tham số
    constructor() AccessControl() {
        _initializeData();
    }
    
    function updateServiceStatus(uint256 _index, DataStructs.ServiceStatus _newStatus) external onlyOwner {
        require(_index < services.length, "INVALID_INDEX");
        services[_index].status = _newStatus;
        emit ServiceStatusUpdated(_index, _newStatus);
    }
    
    function getAllServices() external view returns (DataStructs.Service[] memory) {
        return services;
    }

    // ✨ Hàm khởi tạo dữ liệu mẫu
    function _initializeData() private {
        services.push(DataStructs.Service("Wallet Blockchain", "1.7M User", "2.2M Tx (0.01s)", DataStructs.ServiceStatus.Ok, "apple"));
        services.push(DataStructs.Service("Chain Analytics", "Volum 125M", "100,000 TPS", DataStructs.ServiceStatus.Error, "tesla"));
        services.push(DataStructs.Service("Dapp Email", "1.7M User", "2.2M Mail (0.01s)", DataStructs.ServiceStatus.Ok, "microsoft"));
        services.push(DataStructs.Service("Dapp Elearning", "1.7M User", "2.2M Course", DataStructs.ServiceStatus.Ok, "google"));
        services.push(DataStructs.Service("Dapp Chat", "1.7M User", "1.2M Messages(0.01s)", DataStructs.ServiceStatus.Ok, "nvidia"));
    }
}