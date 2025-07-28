// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./common/DataStructs.sol";

// --- INTERFACES ---
interface IServiceStatusContract { function getAllServices() external view returns (DataStructs.Service[] memory); }
interface IRecruitingContract { function getAllJobs() external view returns (DataStructs.Job[] memory); }
interface ITimeSeriesContract { function getTimeSeriesData(string memory _key) external view returns (DataStructs.DataPoint[] memory); }
interface IKeyValueContract { function getSingleValueData(string memory _key) external view returns (uint256); function getReturnData(string memory _key) external view returns (int256); }
interface ISupportFundContract { function getFundItemsByDay(string memory _day) external view returns (DataStructs.FundItem[] memory); }
interface IContentContract { function getRecentContent(uint256 _count) external view returns (DataStructs.ContentItem[] memory); }
// ✨ ĐÃ XÓA ICustomerReactionContract

contract DashboardFacade {
    // --- ĐỊA CHỈ CONTRACT CON ---
    address public immutable serviceStatusContract;
    address public immutable recruitingContract;
    address public immutable timeSeriesContract;
    address public immutable keyValueContract;
    address public immutable supportFundContract;
    address public immutable contentContract;
    // ✨ ĐÃ XÓA customerReactionContract

    constructor() {
        // ✨ DÁN ĐỊA CHỈ CÁC CONTRACT CON BẠN ĐÃ DEPLOY VÀO ĐÂY
        serviceStatusContract = 0xcc48E282F0136c30835946CC907e10cf98517Af8;
        recruitingContract = 0x048793bBC551254e19A33398167Fb80a9EAec9f1;
        timeSeriesContract = 0xB25201999Dea45fFedC1031274D9DAa730CB88a0;
        supportFundContract = 0xc5AF6359A83F37CC8a10755cb6Cc79a1b605Ec50;
        keyValueContract = 0x6B6250015B2e247E0994DFd5CbB50b4486898644; // Thêm địa chỉ mới
        contentContract = 0x89DE27611644350c53172648d46a135e76FC1ecc;
    }

    // --- HÀM VIEW ĐIỀU HƯỚNG ---
    function getAllServices() external view returns (DataStructs.Service[] memory) { return IServiceStatusContract(serviceStatusContract).getAllServices(); }
    function getAllJobs() external view returns (DataStructs.Job[] memory) { return IRecruitingContract(recruitingContract).getAllJobs(); }
    function getTimeSeriesData(string memory _key) external view returns (DataStructs.DataPoint[] memory) { return ITimeSeriesContract(timeSeriesContract).getTimeSeriesData(_key); }
    function getSingleValueData(string memory _key) external view returns (uint256) { return IKeyValueContract(keyValueContract).getSingleValueData(_key); }
    function getReturnData(string memory _key) external view returns (int256) { return IKeyValueContract(keyValueContract).getReturnData(_key); }
    function getFundItemsByDay(string memory _day) external view returns (DataStructs.FundItem[] memory) { return ISupportFundContract(supportFundContract).getFundItemsByDay(_day); }
    function getRecentContent(uint256 _count) external view returns (DataStructs.ContentItem[] memory) {
        return IContentContract(contentContract).getRecentContent(_count);
    }
    // ✨ ĐÃ XÓA getAllTestimonials()
}