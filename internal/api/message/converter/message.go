package converter

import (
	"github.com/celtic93/chat-server/internal/model"
	desc "github.com/celtic93/chat-server/pkg/v1/chat"
)

func ToMessageFromSendMessageRequest(req *desc.SendMessageRequest) *model.Message {
	return &model.Message{
		ChatID: req.ChatId,
		UserID: req.UserId,
		Text:   req.Text,
	}
}
