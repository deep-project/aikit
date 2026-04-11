package openai

import "github.com/deep-project/aikit/pkg/chat"

// ----------------
// ----------------
// Request
// ----------------
// ----------------
type Request struct {
	Model      string           `json:"model"`
	Messages   []RequestMessage `json:"messages"`
	Tools      []Tool           `json:"tools,omitempty"`
	ToolChoice any              `json:"tool_choice,omitempty"`
}

func buildRequest(model string, req *chat.Request, multimodal bool) *Request {
	return &Request{
		Model:    model,
		Messages: toRequestMessages(req.Messages, multimodal),
		Tools:    toRequestTools(req.Tools),
	}
}

func toRequestTools(tools []chat.Tool) (res []Tool) {
	for _, tool := range tools {
		res = append(res, buildTool(tool))
	}
	return
}

// ----------------
// ----------------
// RequestMessage
// ----------------
// ----------------
type RequestMessage struct {
	Role    string  `json:"role"`
	Content Content `json:"content"`

	// 当role=assistant时，需要传递toolCalls信息
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	// 当role=tool时，需要传递当前tool call的id
	ToolCallID string `json:"tool_call_id,omitempty"`
}

func toRequestMessages(msgs []chat.RequestMessage, multimodal bool) (res []RequestMessage) {
	for _, msg := range msgs {
		res = append(res, buildRequestMessage(msg, multimodal))
	}
	return
}

func buildRequestMessage(msg chat.RequestMessage, multimodal bool) RequestMessage {
	r := RequestMessage{
		Role:       string(msg.Role),
		ToolCallID: msg.ToolCallID,
		Content:    buildContent(msg.Content, multimodal),
		ToolCalls:  toRequestToolCalls(msg.ToolCalls),
	}
	return r
}

func toRequestToolCalls(toolCalls []chat.ResponseToolCall) (res []ToolCall) {
	for _, t := range toolCalls {
		res = append(res, buildToolCall(t))
	}
	return
}
