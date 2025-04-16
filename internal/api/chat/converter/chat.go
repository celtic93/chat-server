package converter

import (
	"github.com/celtic93/chat-server/internal/model"
	desc "github.com/celtic93/chat-server/pkg/v1/chat"
)

func ToChatFromCreateRequest(req *desc.CreateRequest) *model.Chat {
	return &model.Chat{
		UserIDs: req.GetUserIds(),
	}
}
