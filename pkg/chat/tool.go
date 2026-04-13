package chat

type Tool interface {
	Info() *ToolInfo
	Call(ToolCallInput) (_ string, stop bool, _ error)
}

type ToolCallInput struct {
	Request  *Request
	Response ResponseToolCall
}

type ToolInfo struct {
	Name        string          `json:"name"`        // 工具名称(例：getStock)
	Description string          `json:"description"` // 工具表述
	Parameters  []ToolParameter `json:"parameters"`  // 工具参数
}

type ToolParameter struct {
	Name        string   `json:"name"`        // 属性名
	Type        ToolType `json:"type"`        // 类型
	Description string   `json:"description"` // 描述
	Enum        []string `json:"enum"`        // 枚举值（约束限定可选范围）
	Default     any      `json:"default"`     // 默认值
	Required    bool     `json:"required"`    // 必填的

	// 以下字段优先级较低，可选填
	Examples   []string `json:"examples"`   // 举例
	Title      string   `json:"title"`      // 友好属性名称，可以显示更易读的属性名称
	MaxLength  int      `json:"maxLength"`  // 属性值最大长度
	MinLength  int      `json:"minLength"`  // 属性值最小长度
	Pattern    string   `json:"pattern"`    // 属性值必须匹配正则表达式
	Maximum    float64  `json:"maximum"`    // 属性值为数字时的最大值
	Minimum    float64  `json:"minimum"`    // 属性值为数字时的最小值
	MultipleOf float64  `json:"multipleOf"` // 属性值为数字时必须是指定倍数（数值必须能被此值整除）
}

// 工具类型
type ToolType string

// openai标准要求符合JSON Schema基本字段
const (
	ToolTypeString  ToolType = "string"
	ToolTypeNumber  ToolType = "number"
	ToolTypeInteger ToolType = "integer"
	ToolTypeBoolean ToolType = "boolean"
	ToolTypeArray   ToolType = "array"
	ToolTypeObject  ToolType = "object"
)
