package chat

import "context"

type Request struct {
	Ctx context.Context `json:"-"`
	// 元数据
	// 使用者定义后，可以直接在tool Call中使用
	// 可用于鉴权key输入
	Meta     map[string]any   `json:"meta,omitzero"`
	Messages []RequestMessage `json:"messages"`
	Tools    []Tool           `json:"tools,omitzero"`
	// 随机性
	// 0:最保守 1:有创造性 1.2:非常随机(可能胡说)
	Temperature float64 `json:"temperature,omitempty"`
	// 核采样
	// 0.5:很保守 0.9:只选高概率词 1.0:不限制（最随机）
	TopP float64 `json:"top_p,omitempty"`
	// 存在惩罚
	// 0:多说已存在的词 1:多用新词
	PresencePenalty float64 `json:"presence_penalty,omitempty"`
	// 频率惩罚
	// 0:重复性加强 1:重复性减弱
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
}

type RequestMessage struct {
	Role    Role      `json:"role"`
	Content []Content `json:"content"`

	// 当role=assistant时，需要传递toolCalls信息
	ToolCalls []ResponseToolCall `json:"tool_calls,omitempty"`
	// 当role=tool时，需要传递当前tool call的id
	ToolCallID string `json:"tool_call_id,omitempty"`
}
