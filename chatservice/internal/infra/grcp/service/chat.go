package service

import (
	"github.com/virb30/fcexperience/chatservice/internal/infra/grcp/pb"
	"github.com/virb30/fcexperience/chatservice/internal/usecase/chatcompletionstream"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
	ChatCommpletionStreamUseCase chatcompletionstream.ChatCompletionUseCase
	ChatConfigStream             chatcompletionstream.ChatCompletionConfigInputDTO
	StreamChannel                chan chatcompletionstream.ChatCompletionOutputDTO
}

func NewChatService(chatCommpletionStreamUseCase chatcompletionstream.ChatCompletionUseCase, chatConfigStream chatcompletionstream.ChatCompletionConfigInputDTO, streamChannel chan chatcompletionstream.ChatCompletionOutputDTO) *ChatService {
	return &ChatService{
		ChatCommpletionStreamUseCase: chatCommpletionStreamUseCase,
		ChatConfigStream:             chatConfigStream,
		StreamChannel:                streamChannel,
	}
}

func (c *ChatService) ChatStream(req *pb.ChatRequest, stream pb.ChatService_ChatStreamServer) error {
	chatConfig := chatcompletionstream.ChatCompletionConfigInputDTO{
		Model:                c.ChatConfigStream.Model,
		ModelMaxTokens:       c.ChatConfigStream.MaxTokens,
		Temperature:          c.ChatConfigStream.Temperature,
		TopP:                 c.ChatConfigStream.TopP,
		N:                    c.ChatConfigStream.N,
		Stop:                 c.ChatConfigStream.Stop,
		MaxTokens:            c.ChatConfigStream.MaxTokens,
		InitialSystemMessage: c.ChatConfigStream.InitialSystemMessage,
	}

	input := chatcompletionstream.ChatCompletionInputDTO{
		UserMessage: req.GetUserMessage(),
		UserID:      req.GetUserId(),
		ChatID:      req.GetChatId(),
		Config:      chatConfig,
	}

	ctx := stream.Context()

	go func() {
		for msg := range c.StreamChannel {
			stream.Send(&pb.ChatResponse{
				ChatId:  msg.ChatID,
				UserId:  msg.UserID,
				Content: msg.Content,
			})
		}
	}()

	_, err := c.ChatCommpletionStreamUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
