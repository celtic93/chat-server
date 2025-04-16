package converter

import (
	"github.com/celtic93/chat-server/internal/model"
	desc "github.com/celtic93/chat-server/pkg/v1/message"
)

func ToMessageFromCreateRequest(req *desc.CreateRequest) *model.Message {
	return &model.Message{
		ChatID: req.ChatId,
		UserID: req.UserId,
		Text:   req.Text,
	}
}
