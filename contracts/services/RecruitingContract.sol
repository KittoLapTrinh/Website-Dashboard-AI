// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../common/DataStructs.sol";
import "../common/AccessControl.sol";

contract RecruitingContract is AccessControl {
    DataStructs.Job[] public jobs;
    event JobAdded(DataStructs.Job newJob);

    // ✨ Constructor không còn tham số
    constructor() AccessControl() {
        _initializeData();
    }

    function addJob(DataStructs.Job memory _newJob) external onlyOwner {
        jobs.push(_newJob);
        emit JobAdded(_newJob);
    }

    function getAllJobs() external view returns (DataStructs.Job[] memory) {
        return jobs;
    }

    // ✨ Hàm khởi tạo dữ liệu mẫu
    function _initializeData() private {
        // jobs.push(DataStructs.Job("Nvidia", "Deep Learning Scientist", "AI", 55000, "Remote", DataStructs.JobTrend.Up, "nvidia"));
        // jobs.push(DataStructs.Job("Tesla", "Firmware Engineer", "IOT", 45000, "Fulltime", DataStructs.JobTrend.Up, "tesla"));
        // jobs.push(DataStructs.Job("ConsenSys", "Smart Contract Dev", "Blockchain", 42000, "Fulltime", DataStructs.JobTrend.Down, "apple"));
        // jobs.push(DataStructs.Job("Google", "Frontend Developer", "Web Dev", 38000, "Fulltime", DataStructs.JobTrend.Up, "google"));
        // jobs.push(DataStructs.Job("Apple", "AI Research Intern", "AI", 32000, "Remote", DataStructs.JobTrend.Down, "apple"));
        // jobs.push(DataStructs.Job("FPT", "Blockchain Engineer", "Blockchain", 26000, "Fulltime", DataStructs.JobTrend.Up, "fpt"));
    }
}