package openai

import (
	"encoding/json"

	"github.com/deep-project/aikit/pkg/chat"
)

// ----------------
// ----------------
// openai Response
// ----------------
// ----------------
type Response struct {
	ID      string           `json:"id"`
	Object  string           `json:"object"`
	Created int              `json:"created"` // 创建的时间戳
	Model   string           `json:"model"`   // 模型名称
	Choices []ResponseChoice `json:"choices"`
	Usage   ResponseUsage    `json:"usage"`
}

type ResponseUsage struct {
	PromptTokens     float64 `json:"prompt_tokens"`     // 输入token用量
	CompletionTokens float64 `json:"completion_tokens"` // 输出token用量
	TotalTokens      float64 `json:"total_tokens"`      // 总用量
}

type ResponseChoice struct {
	Index   int             `json:"index"`
	Message ResponseMessage `json:"message"`

	// FinishReason 停止原因
	// stop: 正常结束
	// length: 长度限制
	// tool_calls: 函数调用
	// content_filter: 安全拦截
	FinishReason string `json:"finish_reason"`
}

type ResponseMessage struct {
	Role      string     `json:"role"`
	Content   string     `json:"content"`
	ToolCalls []ToolCall `json:"tool_calls"`
}

// ----------------
// ----------------
// to chat Response
// ----------------
// ----------------

func toChatResponse(resp *Response) *chat.Response {
	return &chat.Response{
		ID:      resp.ID,
		Choices: toChatResponseChoices(resp.Choices),
		Usage: chat.ResponseUsage{
			InputTokens:  resp.Usage.PromptTokens,
			OutputTokens: resp.Usage.CompletionTokens,
			TotalTokens:  resp.Usage.TotalTokens,
		},
	}
}

func toChatResponseChoices(items []ResponseChoice) (res []chat.ResponseChoice) {
	for _, item := range items {
		res = append(res, toChatResponseChoice(item))
	}
	return res
}

func toChatResponseChoice(item ResponseChoice) chat.ResponseChoice {
	return chat.ResponseChoice{
		Index: item.Index,
		Message: chat.ResponseMessage{
			Role: chat.Role(item.Message.Role),
			Content: []chat.Content{
				{Type: chat.ContentTypeText, Text: item.Message.Content},
			},
			ToolCalls: toChatResponseToolCalls(item.Message.ToolCalls),
		},
	}
}

func toChatResponseToolCalls(items []ToolCall) (res []chat.ResponseToolCall) {
	for _, item := range items {
		res = append(res, toChatResponseToolCall(item))
	}
	return
}

func toChatResponseToolCall(item ToolCall) chat.ResponseToolCall {
	var args map[string]any
	json.Unmarshal([]byte(item.Function.Arguments), &args)
	return chat.ResponseToolCall{
		ID:        item.ID,
		Name:      item.Function.Name,
		Arguments: args,
	}
}
