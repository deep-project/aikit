package aikit

import "github.com/deep-project/aikit/pkg/chat"

func NewChat(p chat.Provider) *chat.Client {
	return chat.New(chat.WithProvider(p))
}
