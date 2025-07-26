// File: src/app/api/chat/company-data.ts

// ====================================================================
// === CƠ SỞ DỮ LIỆU "BỘ NÃO" CHO CHATBOT ===
// ====================================================================
// Đây là nơi bạn định nghĩa tất cả kiến thức về công ty.
// Càng chi tiết, chatbot sẽ càng thông minh.

// --- Dữ liệu về Công ty ---
const companyInfo = {
    name: "Trung tâm Blockchain AI",
    mission: "Sứ mệnh của chúng tôi là phát triển các sản phẩm công nghệ đột phá trên nền tảng Blockchain và AI, đồng thời xây dựng một môi trường làm việc sáng tạo và chuyên nghiệp.",
    employeeCount: 50,
    mainTechnologies: "Solidity, Go (Golang), React, Next.js cho các ứng dụng blockchain, và Python, TensorFlow, PyTorch cho các mô hình AI.",
    dataTrend: "Dữ liệu trên dashboard được cập nhật liên tục từ các smart contract, phản ánh các hoạt động và hiệu suất của trung tâm theo thời gian thực.",
    locations: "Chúng tôi có văn phòng tại Bát Nàn (TP. Hồ Chí Minh), Điện Biên Phủ (TP. Hồ Chí Minh), và D'Capitale - 119 Trần Duy Hưng (Hà Nội)."
};

// --- Dữ liệu về các Dự án ---
const projects = [
    { name: "Chat Blockchain", description: "Hệ thống chat phi tập trung, đảm bảo tin nhắn được mã hóa end-to-end và không thể bị kiểm duyệt." },
    { name: "Email Blockchain", description: "Dịch vụ email an toàn trên nền tảng blockchain, bảo vệ danh tính và quyền riêng tư của người dùng." },
    { name: "Ví tiền điện tử (Wallet)", description: "Ví non-custodial an toàn để người dùng tự quản lý và giao dịch các loại tài sản số." },
    { name: "Phân tích dữ liệu Blockchain (Chain Analytics)", description: "Công cụ phân tích dữ liệu on-chain mạnh mẽ để theo dõi các giao dịch, phát hiện gian lận và dự báo xu hướng thị trường." },
    { name: "Nghiên cứu phần cứng Blockchain", description: "Phát triển các thiết bị phần cứng chuyên dụng, tối ưu cho việc vận hành node và tăng hiệu suất staking." },
    { name: "Ứng dụng AI trong Bất động sản", description: "Sử dụng AI để phân tích thị trường, định giá tài sản và gợi ý các cơ hội đầu tư bất động sản tiềm năng." }
];

// --- Dữ liệu về Tuyển dụng ---
const recruiting = {
    generalPolicy: "Chúng tôi luôn tìm kiếm những tài năng đam mê công nghệ blockchain và AI. Quy trình tuyển dụng của chúng tôi tập trung vào việc đánh giá năng lực thực tế, tư duy giải quyết vấn đề và sự phù hợp với văn hóa công ty.",
    applicationEmail: "hr@ten_cong_ty.com", // Thay bằng email thật
    positions: [
        // Blockchain
        { position: "Senior Blockchain Engineer", field: "Blockchain", salary: "Thỏa thuận (lên đến $7,000)", requirements: "Trên 3 năm kinh nghiệm với Solidity & Go. Có kinh nghiệm sâu về kiến trúc hệ thống phi tập trung, DeFi, và các giải pháp Layer 2." },
        { position: "Junior Blockchain Developer", field: "Blockchain", salary: "$1,500 - $2,500", requirements: "Hiểu biết cơ bản về blockchain, đã hoàn thành các dự án cá nhân hoặc khóa học về Solidity. Có tư duy logic tốt và ham học hỏi." },
        { position: "Smart Contract Auditor", field: "Blockchain", salary: "Thỏa thuận (rất cạnh tranh)", requirements: "Có kinh nghiệm kiểm thử và phát hiện lỗ hổng bảo mật trong smart contract. Thành thạo các công cụ như Slither, Mythril." },
        
        // AI
        { position: "AI Scientist (NLP)", field: "AI", salary: "Lên đến $5,000", requirements: "Có kinh nghiệm xây dựng và triển khai các mô hình xử lý ngôn ngữ tự nhiên (NLP) như Transformer, BERT. Có khả năng làm việc với các bộ dữ liệu lớn." },
        { position: "Data Engineer", field: "AI", salary: "$2,000 - $4,000", requirements: "Có kinh nghiệm xây dựng và tối ưu các pipeline dữ liệu (ETL). Thành thạo SQL, Python và các công cụ như Spark, Airflow." },
        { position: "Machine Learning Intern", field: "AI", salary: "Hỗ trợ 5-8 triệu VND/tháng", requirements: "Sinh viên năm cuối hoặc mới tốt nghiệp, có kiến thức nền tảng vững chắc về Machine Learning và mong muốn phát triển trong lĩnh vực AI." },
        
        // Vị trí khác
        { position: "Product Manager (DeFi)", field: "Product", salary: "Thỏa thuận", requirements: "Có kinh nghiệm quản lý sản phẩm trong lĩnh vực tài chính hoặc blockchain. Hiểu biết sâu sắc về các giao thức DeFi." },
        { position: "Community Manager (Web3)", field: "Marketing", salary: "Thỏa thuận", requirements: "Có kinh nghiệm xây dựng và quản lý cộng đồng trên các nền tảng như Discord, Telegram, Twitter trong lĩnh vực blockchain/crypto." },
        { position: "DevOps Engineer", field: "Operations", salary: "$2,500 - $4,500", requirements: "Có kinh nghiệm với AWS/GCP, Docker, Kubernetes, và các công cụ CI/CD. Ưu tiên có kinh nghiệm vận hành các node blockchain." }
    ],
    // ✨ BỔ SUNG MỤC CÂU HỎI THƯỜNG GẶP VỀ TUYỂN DỤNG ✨
    faq: [
        { q: "Quy trình tuyển dụng gồm mấy vòng?", a: "Quy trình tuyển dụng của chúng tôi thường bao gồm 3 vòng: Sàng lọc CV, Phỏng vấn kỹ thuật với team lead, và Phỏng vấn văn hóa với quản lý." },
        { q: "Tôi có thể ứng tuyển vào nhiều vị trí cùng lúc không?", a: "Có, bạn có thể ứng tuyển vào các vị trí mà bạn cảm thấy phù hợp. Tuy nhiên, chúng tôi khuyến khích bạn tập trung vào vị trí phù hợp nhất với năng lực của mình." },
        { q: "Thời gian làm bài test kỹ thuật là bao lâu?", a: "Tùy thuộc vào vị trí, bài test kỹ thuật có thể là một bài tập nhỏ làm ở nhà (take-home) trong vòng 2-3 ngày, hoặc một phiên giải quyết vấn đề trực tiếp trong buổi phỏng vấn." },
        { q: "Tôi cần chuẩn bị gì cho buổi phỏng vấn?", a: "Bạn nên xem lại các kiến thức chuyên môn liên quan đến vị trí ứng tuyển và chuẩn bị các ví dụ về những dự án bạn đã thực hiện. Quan trọng nhất là thể hiện được tư duy logic và đam mê của bạn." },
        { q: "Sau bao lâu thì tôi sẽ nhận được kết quả?", a: "Chúng tôi sẽ cố gắng phản hồi cho bạn trong vòng 7 ngày làm việc sau mỗi vòng phỏng vấn." },
        { q: "Công ty có tuyển thực tập sinh không?", a: "Có, chúng tôi thường xuyên có các đợt tuyển thực tập sinh cho các mảng Blockchain và AI. Hiện tại đang có vị trí Machine Learning Intern." }
    ]
};

// --- Dữ liệu về Chính sách & Phúc lợi ---
const policies = {
    workingHours: "Giờ làm việc linh hoạt (flex-time), từ thứ 2 đến thứ 7. Nhân viên có thể tự sắp xếp thời gian miễn là đảm bảo hiệu quả công việc và tham gia các cuộc họp quan trọng.",
    remotePolicy: "Chúng tôi hỗ trợ chính sách làm việc hybrid, kết hợp làm việc tại văn phòng và từ xa để tạo sự thuận tiện cho nhân viên.",
    paidLeave: "15 ngày nghỉ phép có lương mỗi năm, cộng thêm các ngày nghỉ lễ theo quy định.",
    healthBenefits: "Cung cấp bữa ăn trưa miễn phí hàng ngày tại văn phòng, có quầy bar đồ ăn nhẹ và trái cây. Có quỹ hỗ trợ sức khỏe để chi trả các chi phí y tế phát sinh.",
    insurance: "Toàn bộ nhân viên được đóng bảo hiểm xã hội, bảo hiểm y tế, và bảo hiểm thất nghiệp theo đúng quy định của pháp luật."
};

// --- Dữ liệu về Hoạt động nội bộ ---
const internalActivities = {
    training: "Tổ chức các buổi workshop và seminar hàng tuần về các công nghệ mới. Khuyến khích và tài trợ cho nhân viên tham gia các khóa học và lấy chứng chỉ quốc tế.",
    teamBuilding: "Các hoạt động team building được tổ chức hàng quý. Ngoài ra còn có các câu lạc bộ thể thao và giải đấu game nội bộ.",
    hackathons: "Tổ chức các cuộc thi hackathon nội bộ 2 lần/năm với những giải thưởng hấp dẫn để khuyến khích sự sáng tạo và đổi mới."
};


// ====================================================================
// === HÀM TẠO RA CHUỖI VĂN BẢN KIẾN THỨC ===
// ====================================================================
export const getCompanyKnowledgeBase = (): string => {
    let knowledge = "Đây là thông tin về Trung tâm Blockchain AI:\n\n";

    knowledge += "== THÔNG TIN CHUNG VỀ CÔNG TY ==\n";
    knowledge += `- Tên công ty: ${companyInfo.name}\n`;
    knowledge += `- Sứ mệnh: ${companyInfo.mission}\n`;
    knowledge += `- Số lượng nhân viên: Khoảng ${companyInfo.employeeCount} người.\n`;
    knowledge += `- Địa chỉ văn phòng: ${companyInfo.locations}\n`;
    knowledge += `- Công nghệ sử dụng: ${companyInfo.mainTechnologies}\n`;
    knowledge += `- Về xu hướng dữ liệu (trend) trên dashboard: ${companyInfo.dataTrend}\n\n`;

    knowledge += "== CÁC DỰ ÁN NỔI BẬT ==\n";
    projects.forEach(p => {
        knowledge += `- Dự án "${p.name}": ${p.description}\n`;
    });
    knowledge += "\n";

    knowledge += "== THÔNG TIN TUYỂN DỤNG ==\n";
    knowledge += `- Chính sách chung: ${recruiting.generalPolicy}\n`;
    knowledge += `- Ứng viên có thể gửi CV về địa chỉ email: ${recruiting.applicationEmail}\n`;
    
    knowledge += "\n--- CÁC VỊ TRÍ ĐANG MỞ ---\n";
    recruiting.positions.forEach(r => {
        knowledge += `- Vị trí: ${r.position} | Lĩnh vực: ${r.field} | Mức lương: ${r.salary} | Yêu cầu chính: ${r.requirements}\n`;
    });
    
    knowledge += "\n--- CÂU HỎI THƯỜNG GẶP VỀ TUYỂN DỤNG ---\n";
    recruiting.faq.forEach(item => {
        knowledge += `- Câu hỏi: ${item.q}\n  Trả lời: ${item.a}\n`;
    });
    knowledge += "\n";
    
    knowledge += "== CHÍNH SÁCH & PHÚC LỢI DÀNH CHO NHÂN VIÊN ==\n";
    knowledge += `- Giờ làm việc: ${policies.workingHours}\n`;
    knowledge += `- Làm việc từ xa: ${policies.remotePolicy}\n`;
    knowledge += `- Ngày nghỉ phép: ${policies.paidLeave}\n`;
    knowledge += `- Phúc lợi sức khỏe: ${policies.healthBenefits}\n`;
    knowledge += `- Bảo hiểm: ${policies.insurance}\n\n`;

    knowledge += "== HOẠT ĐỘNG NỘI BỘ ==\n";
    knowledge += `- Đào tạo & Phát triển: ${internalActivities.training}\n`;
    knowledge += `- Gắn kết đội ngũ: ${internalActivities.teamBuilding}\n`;
    knowledge += `- Khuyến khích sáng tạo: ${internalActivities.hackathons}\n`;

    return knowledge;
}