import { NextResponse } from 'next/server';
// ✨ 1. Import các thành phần cần thiết từ thư viện Gemini
import { GoogleGenerativeAI, HarmCategory, HarmBlockThreshold } from '@google/generative-ai';

// Lấy API key từ biến môi trường
const GEMINI_API_KEY = process.env.GOOGLE_API_KEY;

if (!GEMINI_API_KEY) {
  throw new Error("Missing Google API Key. Make sure to set it in your .env.local file.");
}

// Khởi tạo model
const genAI = new GoogleGenerativeAI(GEMINI_API_KEY);
const model = genAI.getGenerativeModel({
  model: "gemini-1.5-flash", // Dùng model flash cho nhanh và miễn phí
});

// Cấu hình an toàn để tránh bị chặn
const generationConfig = {
  temperature: 0.9,
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
    
    // ✨ 2. Chuyển đổi định dạng tin nhắn cho phù hợp với Gemini
    // Gemini yêu cầu vai trò (role) xen kẽ giữa 'user' và 'model'
    // và vai trò của bot phải là 'model' thay vì 'assistant'
    const history = messages.slice(0, -1).map(message => ({
      role: message.role === 'assistant' ? 'model' : message.role,
      parts: [{ text: message.content }],
    }));

    // Lấy tin nhắn cuối cùng của người dùng
    const lastUserMessage = messages[messages.length - 1].content;
    
    // Khởi tạo phiên chat với lịch sử
    const chat = model.startChat({
      generationConfig,
      safetySettings,
      history: history,
    });

    // Gửi tin nhắn mới và chờ phản hồi
    const result = await chat.sendMessage(lastUserMessage);
    const response = result.response;
    const botResponseText = response.text();
    
    // ✨ 3. Trả về tin nhắn với định dạng chuẩn
    const botMessage = {
      role: 'assistant',
      content: botResponseText,
    };

    return NextResponse.json(botMessage);

  } catch (error) {
    console.error('[API_GEMINI_ERROR]', error);
  }
}