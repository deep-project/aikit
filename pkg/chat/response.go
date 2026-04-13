package chat

type SendResult struct {
	Output   *ResponseMessage `json:"output"`   // 最佳答案（直接选Choices第一个）
	Messages []RequestMessage `json:"messages"` // 完整消息列表(包含tool msg和AI最后输出的msg)
	Response *Response        `json:"response"` // 核心响应
}

type Response struct {
	ID      string           `json:"id"`
	Choices []ResponseChoice `json:"choices"` // 响应回答（一般单条，多条时有备选答案）
	Usage   ResponseUsage    `json:"usage"`
}

type ResponseUsage struct {
	InputTokens  float64 `json:"input_tokens"`
	OutputTokens float64 `json:"output_tokens"`
	TotalTokens  float64 `json:"total_tokens"`
}

type ResponseChoice struct {
	Index   int             `json:"index"`
	Message ResponseMessage `json:"message"`
}

type ResponseMessage struct {
	Role      Role               `json:"role"`
	Content   []Content          `json:"content"`
	ToolCalls []ResponseToolCall `json:"tool_calls"`
}

type ResponseToolCall struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}
