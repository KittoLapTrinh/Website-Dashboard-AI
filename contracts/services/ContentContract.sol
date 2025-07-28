// File: contracts/services/ContentContract.sol

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../common/AccessControl.sol";
// ✨ 1. IMPORT THƯ VIỆN DataStructs ✨
import "../common/DataStructs.sol";

contract ContentContract is AccessControl {
    
    // ✨ 2. SỬ DỤNG STRUCT TỪ DataStructs ✨
    DataStructs.ContentItem[] public allContent;
    uint256 private nextContentId;

    event ContentAdded(DataStructs.ContentItem newContent);

    function addContent(
        // ✨ 3. SỬ DỤNG ENUM TỪ DataStructs ✨
        DataStructs.ContentType _contentType, 
        string memory _title, 
        string memory _description, 
        string memory _imageUrl, 
        string memory _author
    ) external onlyOwner {
        // ✨ 4. SỬ DỤNG STRUCT TỪ DataStructs ✨
        DataStructs.ContentItem memory newContent = DataStructs.ContentItem({
            id: nextContentId,
            contentType: _contentType,
            title: _title,
            description: _description,
            imageUrl: _imageUrl, // Đã đổi tên
            author: _author,
            timestamp: block.timestamp
        });
        
        allContent.push(newContent);
        emit ContentAdded(newContent);
        nextContentId++;
    }

    function getRecentContent(uint256 _count) external view returns (DataStructs.ContentItem[] memory) {
        uint256 len = allContent.length;
        if (_count > len) {
            _count = len;
        }
        
        // ✨ 5. SỬ DỤNG STRUCT TỪ DataStructs ✨
        DataStructs.ContentItem[] memory recentItems = new DataStructs.ContentItem[](_count);
        for (uint i = 0; i < _count; i++) {
            recentItems[i] = allContent[len - 1 - i];
        }
        return recentItems;
    }
}