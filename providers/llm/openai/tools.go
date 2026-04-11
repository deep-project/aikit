package openai

import (
	"encoding/json"

	"github.com/deep-project/aikit/pkg/chat"
)

type Tool struct {
	Type     string       `json:"type"`
	Function ToolFunction `json:"function"`
}

func buildTool(tool chat.Tool) Tool {
	info := tool.Info()
	return Tool{
		Type: "function",
		Function: ToolFunction{
			Name:        info.Name,
			Description: info.Description,
			Parameters: ToolParameters{
				Type:                 "object",
				Required:             info.Required,
				AdditionalProperties: false,
				Properties:           toToolProperties(info.Parameters),
			},
		},
	}
}

func toToolProperties(parameters []chat.ToolParameter) map[string]ToolProperty {
	res := make(map[string]ToolProperty)
	for _, p := range parameters {
		res[p.Name] = ToolProperty{
			Type:        string(p.Type),
			Description: p.Description,
			Enum:        p.Enum,
			Title:       p.Title,
			Default:     p.Default,
			MaxLength:   p.MaxLength,
			MinLength:   p.MinLength,
			Pattern:     p.Pattern,
			Maximum:     p.Maximum,
			Minimum:     p.Minimum,
			MultipleOf:  p.MultipleOf,
		}
	}
	return res
}

type ToolFunction struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  ToolParameters `json:"parameters"`
}

type ToolParameters struct {
	Type                 string                  `json:"type,omitempty"`
	Properties           map[string]ToolProperty `json:"properties,omitempty"`
	Required             []string                `json:"required,omitempty"`
	AdditionalProperties bool                    `json:"additionalProperties,omitempty"` // 控制是否允许“乱传字段” 建议：FALSE
}

type ToolProperty struct {
	Type        string   `json:"type,omitempty"`        // 类型
	Description string   `json:"description,omitempty"` // 描述
	Enum        []string `json:"enum,omitempty"`        // 枚举值
	Title       string   `json:"title,omitempty"`       // 友好属性名称，可以显示更易读的属性名称
	Default     any      `json:"default,omitempty"`     // 默认值
	MaxLength   int      `json:"maxLength,omitempty"`   // 属性值最大长度
	MinLength   int      `json:"minLength,omitempty"`   // 属性值最小长度
	Pattern     string   `json:"pattern,omitempty"`     // 属性值必须匹配正则表达式
	Maximum     float64  `json:"maximum,omitempty"`     // 属性值为数字时的最大值
	Minimum     float64  `json:"minimum,omitempty"`     // 属性值为数字时的最小值
	MultipleOf  float64  `json:"multipleOf,omitempty"`  // 属性值为数字时必须是指定倍数（数值必须能被此值整除）
}

type ToolCall struct {
	ID       string           `json:"id"`
	Type     string           `json:"type"`
	Function ToolCallFunction `json:"function"`
}

type ToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

func buildToolCall(tc chat.ResponseToolCall) ToolCall {
	b, _ := json.Marshal(tc.Arguments)
	return ToolCall{
		ID:   tc.ID,
		Type: "function",
		Function: ToolCallFunction{
			Name:      tc.Name,
			Arguments: string(b),
		},
	}
}
