// File: src/app/api/chat/route.ts
import { NextResponse } from 'next/server';
// ✨ 1. IMPORT THƯ VIỆN `openai` ✨
import OpenAI from 'openai';
import { getCompanyKnowledgeBase } from './company-data';

// ✨ 2. LẤY API KEY CỦA GROQ ✨
const GROQ_API_KEY = process.env.GROQ_API_KEY;

if (!GROQ_API_KEY) {
  throw new Error("Missing Groq API Key.");
}

// ✨ 3. KHỞI TẠO CLIENT OPENAI NHƯNG TRỎ ĐẾN SERVER CỦA GROQ ✨
const groq = new OpenAI({
  apiKey: GROQ_API_KEY,
  baseURL: 'https://api.groq.com/openai/v1', // Quan trọng nhất!
});


export async function POST(request: Request) {
  try {
    const { messages } = await request.json();

    if (!messages || !Array.isArray(messages) || messages.length === 0) {
      return NextResponse.json({ error: 'Messages are required.' }, { status: 400 });
    }
    
    const knowledgeBase = await getCompanyKnowledgeBase();
    
    // System prompt (giữ nguyên)
    const systemPrompt = `
      Bạn là Metanode, một trợ lý AI thân thiện và chuyên nghiệp của Trung tâm Blockchain AI.
      - ƯU TIÊN HÀNG ĐẦU: Khi người dùng hỏi về công ty, dự án, tuyển dụng, hoặc chính sách, hãy trả lời dựa trên "DỮ LIỆU THAM KHẢO" được cung cấp.
      - Nếu câu hỏi không liên quan đến công ty và nằm ngoài dữ liệu tham khảo, hãy sử dụng kiến thức chung của bạn để trả lời như một trợ lý AI thông thường.
      - Luôn giữ thái độ lịch sự, hữu ích và trả lời bằng tiếng Việt.
    `;

    // Lấy câu hỏi cuối cùng của người dùng
    const lastUserMessage = messages[messages.length - 1].content;

    // ✨ TẠO PROMPT CUỐI CÙNG VỚI CHỈ THỊ RÕ RÀNG HƠN ✨
    const finalPrompt = `
      DỮ LIỆU THAM KHẢO:
      ---
      ${knowledgeBase}
      ---
      
      Dựa vào dữ liệu trên và kiến thức của bạn, hãy trả lời câu hỏi sau của người dùng.
      QUAN TRỌNG: LUÔN LUÔN TRẢ LỜI BẰNG TIẾNG VIỆT.

      Câu hỏi: "${lastUserMessage}"
    `;


    // Chuẩn bị payload cho API
    const apiMessages = [
      { role: "system", content: systemPrompt },
      // Lịch sử chat giờ sẽ được thêm vào đây, trước câu hỏi cuối cùng
      ...messages.slice(0, -1), 
      // Câu hỏi cuối cùng của người dùng sẽ được gói trong prompt đã được bổ sung
      { role: "user", content: finalPrompt },
    ];

    const chatCompletion = await groq.chat.completions.create({
        messages: apiMessages as any,
        model: "llama3-8b-8192", 
        temperature: 0.7,
        max_tokens: 2048,
    });

    const botResponseText = chatCompletion.choices[0]?.message?.content || "Sorry, I couldn't get a response.";
    
    const botMessage = {
      role: 'assistant',
      content: botResponseText,
    };

    return NextResponse.json(botMessage);

  } catch (error) {
    console.error('[API_GROQ_ERROR]', error);
    return NextResponse.json({ error: "Sorry, I'm having trouble connecting right now." }, { status: 500 });
  }
}