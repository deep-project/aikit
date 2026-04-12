package test

import (
	"encoding/json"
	"testing"

	"github.com/deep-project/aikit/pkg/chat"
	"github.com/deep-project/aikit/providers/llm/openai"
	"github.com/deep-project/aikit/test/tools"
)

func TestChat(t *testing.T) {
	p := openai.NewChat(env.BaseURL, env.APIKey, env.Model, false)
	c := chat.New(chat.WithProvider(p))
	r, err := c.Send(&chat.Request{
		Messages: []chat.RequestMessage{
			{
				Role: chat.RoleUser,
				Content: []chat.Content{
					{Type: chat.ContentTypeText, Text: "你好, AI!"},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(b))
}

func TestChatTools(t *testing.T) {
	p := openai.NewChat(env.BaseURL, env.APIKey, env.Model, false)
	c := chat.New(chat.WithProvider(p))
	r, err := c.Send(&chat.Request{
		Tools: []chat.Tool{
			tools.NewChatGetStock(), // 查询库存的tool
		},
		Messages: []chat.RequestMessage{
			{
				Role: chat.RoleUser,
				Content: []chat.Content{
					{Type: chat.ContentTypeText, Text: "你好，帮我看看270854的库存还有多少？"},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(b))
}

func TestChatRequest(t *testing.T) {
	p := openai.NewChat(env.BaseURL, env.APIKey, env.Model, false)
	c := chat.New(chat.WithProvider(p))
	r, err := c.Request(&chat.Request{
		Messages: []chat.RequestMessage{
			{
				Role: chat.RoleUser,
				Content: []chat.Content{
					{Type: chat.ContentTypeText, Text: "你好，你是？？"},
				},
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(b))
}
