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
	// 随机性
	// 0.001:最保守 1:有创造性 1.2:非常随机(可能胡说)
	Temperature float64 `json:"temperature,omitempty"`
	// 核采样
	// 0.5:很保守 0.9:只选高概率词 1.0:不限制（最随机）
	TopP float64 `json:"top_p,omitempty"`
	// 存在惩罚
	// 0.001:多说已存在的词 1:多用新词
	PresencePenalty float64 `json:"presence_penalty,omitempty"`
	// 频率惩罚
	// 0.001:重复性加强 1:重复性减弱
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// 千问适配
	ExtraBody *RequestExtraBody `json:"extra_body,omitempty"`
}

func buildRequest(model string, req *chat.Request, multimodal bool) *Request {
	return &Request{
		Model:            model,
		Messages:         toRequestMessages(req.Messages, multimodal),
		Tools:            toRequestTools(req.Tools),
		Temperature:      req.Temperature,
		TopP:             req.TopP,
		PresencePenalty:  req.PresencePenalty,
		FrequencyPenalty: req.FrequencyPenalty,
		ExtraBody:        &RequestExtraBody{EnableThinking: req.EnableThinking},
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

type RequestExtraBody struct {
	EnableThinking bool `json:"enable_thinking"`
}
