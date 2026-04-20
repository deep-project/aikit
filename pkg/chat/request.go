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

	// 采样温度
	// 控制模型生成文本的多样性
	// temperature越高，生成的文本更多样，反之，生成的文本更确定
	// 取值范围： [0, 2)
	// 0.001:最保守 1:有创造性 1.2:非常随机(可能胡说)
	Temperature float64 `json:"temperature,omitempty"`
	// 核采样的概率阈值
	// 控制模型生成文本的多样性
	// top_p越高，生成的文本更多样。反之，生成的文本更确定
	// 取值范围：（0,1.0]
	// 0.5:很保守 0.9:只选高概率词 1.0:不限制（最随机）
	TopP float64 `json:"top_p,omitempty"`
	// 存在惩罚
	// 控制模型生成文本时的内容重复度
	// 在创意写作或头脑风暴等需要多样性、趣味性或创造力的场景中，建议调高该值；
	// 在技术文档或正式文本等强调一致性与术语准确性的场景中，建议调低该值。
	// 取值范围：[-2.0, 2.0]
	// 正值降低重复度，负值增加重复度
	// 0.001:多说已存在的词 1:多用新词
	PresencePenalty float64 `json:"presence_penalty,omitempty"`
	// 频率惩罚
	// 0.001:重复性加强 1:重复性减弱
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// 启用深度思考
	EnableThinking bool `json:"enable_thinking"`
}

type RequestMessage struct {
	Role    Role      `json:"role"`
	Content []Content `json:"content"`

	// 当role=assistant时，需要传递toolCalls信息
	ToolCalls []ResponseToolCall `json:"tool_calls,omitempty"`
	// 当role=tool时，需要传递当前tool call的id
	ToolCallID string `json:"tool_call_id,omitempty"`
}
