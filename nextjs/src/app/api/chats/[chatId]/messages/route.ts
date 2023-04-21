import { prisma } from "@/app/prisma/prisma";
import { NextRequest, NextResponse } from "next/server";

type ChatParams = {
    params: {
        chatId: string
    }
}

export async function GET(_request: NextRequest, { params }: ChatParams) {
    const messages = await prisma.message.findMany({
        where: {
            chat_id: params.chatId
        },
        orderBy: { created_at: "asc" }
    });

    return NextResponse.json(messages);
}

export async function POST(request: NextRequest, { params }: ChatParams) {
    const chat = await prisma.chat.findUniqueOrThrow({
        where: {
            id: params.chatId
        }
    });
    const body = await request.json();

    const messageCreated = await prisma.message.create({
        data: {
            content: body.message,
            chat_id: chat.id
        }
    });

    return NextResponse.json(messageCreated);
}