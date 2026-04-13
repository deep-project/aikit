package tools

import (
	"fmt"

	"github.com/deep-project/aikit/pkg/chat"
)

type ChatGetStock struct {
}

func NewChatGetStock() *ChatGetStock {
	return &ChatGetStock{}
}

func (ChatGetStock) Info() *chat.ToolInfo {
	return &chat.ToolInfo{
		Name:        "getStock",
		Description: "获取指定编码的库存",
		Parameters: []chat.ToolParameter{
			{
				Name:        "code",
				Type:        chat.ToolTypeString,
				Description: "商品编码",
			},
		},
	}
}

func (ChatGetStock) Call(input chat.ToolCallInput) (_ string, stop bool, _ error) {
	code, ok := input.Response.Arguments["code"]
	if !ok {
		return "", false, fmt.Errorf("missing param: code")
	}
	fmt.Println("code", code)
	return "当前商品的库存是：666", false, nil
}
