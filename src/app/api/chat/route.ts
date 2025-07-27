import { NextResponse } from 'next/server';
// ✨ 1. Import các thành phần cần thiết từ thư viện Gemini
import { GoogleGenerativeAI, HarmCategory, HarmBlockThreshold } from '@google/generative-ai';
import { getCompanyKnowledgeBase } from './company-data';

// Lấy API key từ biến môi trường
const GEMINI_API_KEY = process.env.GOOGLE_API_KEY;

if (!GEMINI_API_KEY) {
  throw new Error("Missing Google API Key. Make sure to set it in your .env.local file.");
}

// Khởi tạo model
const genAI = new GoogleGenerativeAI(GEMINI_API_KEY);
const model = genAI.getGenerativeModel({
  model: "gemini-1.5-flash", // Dùng model flash cho nhanh và miễn phí
  systemInstruction: "Bạn là Metanode, một trợ lý AI thân thiện và chuyên nghiệp của Trung tâm Blockchain AI. Hãy trả lời các câu hỏi của người dùng một cách ngắn gọn, súc tích và chỉ dựa vào thông tin được cung cấp trong prompt. Nếu không biết, hãy nói rằng bạn không có thông tin đó.",
});

// Cấu hình an toàn để tránh bị chặn
const generationConfig = {
  temperature: 0.7,
  topK: 1,
  topP: 1,
  maxOutputTokens: 2048,
};

const safetySettings = [
  { category: HarmCategory.HARM_CATEGORY_HARASSMENT, threshold: HarmBlockThreshold.BLOCK_MEDIUM_AND_ABOVE },
  { category: HarmCategory.HARM_CATEGORY_HATE_SPEECH, threshold: HarmBlockThreshold.BLOCK_MEDIUM_AND_ABOVE },
  { category: HarmCategory.HARM_CATEGORY_SEXUALLY_EXPLICIT, threshold: HarmBlockThreshold.BLOCK_MEDIUM_AND_ABOVE },
  { category: HarmCategory.HARM_CATEGORY_DANGEROUS_CONTENT, threshold: HarmBlockThreshold.BLOCK_MEDIUM_AND_ABOVE },
];

export async function POST(request: Request) {
  try {
    const { messages } = await request.json();

    if (!messages || !Array.isArray(messages) || messages.length === 0) {
      return NextResponse.json({ error: 'Messages are required.' }, { status: 400 });
    }
    
    // Chuyển đổi lịch sử chat cho phù hợp với Gemini (giữ nguyên)
    const history = messages.slice(0, -1).map(message => ({
      role: message.role === 'assistant' ? 'model' : message.role,
      parts: [{ text: message.content }],
    }));

    // Lấy câu hỏi cuối cùng của người dùng
    const lastUserMessage = messages[messages.length - 1].content;
    
    // ✨ 2. BƯỚC LẤY DỮ LIỆU (RETRIEVAL) ✨
    // Gọi hàm để lấy toàn bộ kiến thức về công ty.
    // Trong tương lai, bước này có thể được nâng cấp để tìm kiếm thông tin liên quan nhất
    // thay vì lấy tất cả.
    const knowledgeBase = await getCompanyKnowledgeBase();

    // ✨ 3. BƯỚC BỔ SUNG DỮ LIỆU (AUGMENTATION) ✨
    // Tạo một prompt mới, "đính kèm" kiến thức đã lấy được vào câu hỏi của người dùng.
    const augmentedPrompt = `
      Dưới đây là DỮ LIỆU THAM KHẢO về công ty. Chỉ sử dụng thông tin này để trả lời câu hỏi.
      ---
      ${knowledgeBase}
      ---
      
      Câu hỏi của người dùng: "${lastUserMessage}"
    `;
    
    // Khởi tạo phiên chat với lịch sử
    const chat = model.startChat({
      generationConfig,
      safetySettings,
      history: history,
    });

    // Gửi prompt đã được bổ sung và chờ phản hồi
    const result = await chat.sendMessage(augmentedPrompt);
    const response = result.response;
    const botResponseText = response.text();
    
    // Trả về tin nhắn với định dạng chuẩn
    const botMessage = {
      role: 'assistant',
      content: botResponseText,
    };

    return NextResponse.json(botMessage);

  } catch (error) {
    console.error('[API_GEMINI_ERROR]', error);
    return NextResponse.json({ error: "Sorry, I'm having trouble connecting right now." }, { status: 500 });
  }
}