package chat

type Request struct {
	Messages []RequestMessage `json:"messages"`
	Tools    []Tool           `json:"tools,omitzero"`
}

type RequestMessage struct {
	Role    Role      `json:"role"`
	Content []Content `json:"content"`

	// 当role=assistant时，需要传递toolCalls信息
	ToolCalls []ResponseToolCall `json:"tool_calls,omitempty"`
	// 当role=tool时，需要传递当前tool call的id
	ToolCallID string `json:"tool_call_id,omitempty"`
}
