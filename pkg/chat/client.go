package chat

import (
	"errors"
	"fmt"
)

const DefaultMaxSteps = 8

type Provider interface {
	Chat(*Request) (*Response, error)
}

type Client struct {
	MaxSteps int // 最大循环次数
	Provider Provider
}

type Option func(*Client)

func New(options ...Option) *Client {
	cli := &Client{}
	for _, opt := range options {
		opt(cli)
	}
	return cli
}

func WithProvider(p Provider) Option {
	return func(c *Client) {
		c.SetProvider(p)
	}
}

func WithMaxSteps(n int) Option {
	return func(c *Client) {
		c.SetMaxSteps(n)
	}
}

func (c *Client) SetProvider(p Provider) *Client {
	c.Provider = p
	return c
}

func (c *Client) SetMaxSteps(n int) *Client {
	c.MaxSteps = n
	return c
}

func (c *Client) Send(req *Request) (*SendResult, error) {

	if c.Provider == nil {
		return nil, errors.New("provider is not defined")
	}
	if c.MaxSteps <= 0 {
		c.MaxSteps = DefaultMaxSteps // fallback default
	}

	var (
		// 循环次数
		step = 0
		// copy messages（避免污染）
		messages = append([]RequestMessage{}, req.Messages...)
	)

	for {
		if step >= c.MaxSteps {
			return nil, errors.New("tool loop exceeded max steps")
		}
		step++

		// 1. 调用模型
		resp, err := c.Request(&Request{
			Messages: messages,
			Tools:    req.Tools,
		})
		if err != nil {
			return nil, err
		}

		if len(resp.Choices) == 0 {
			return nil, errors.New("empty choices from provider")
		}

		msg := resp.Choices[0].Message

		// 2. 加入 assistant 消息
		messages = append(messages, RequestMessage{
			Role:      RoleAssistant,
			Content:   msg.Content,
			ToolCalls: msg.ToolCalls,
		})

		// 3. 判断是否有 tool_calls
		if len(msg.ToolCalls) == 0 {
			return &SendResult{
				Output:   &msg,
				Messages: messages,
				Response: resp,
			}, nil
		}

		// 4. 执行 tools
		for _, tc := range msg.ToolCalls {

			tool := c.findTool(req.Tools, tc.Name)
			if tool == nil {
				return nil, errors.New("tool not found: " + tc.Name)
			}

			result, err := tool.Call(tc.Arguments)
			if err != nil {
				return nil, err
			}

			// 5. 加入 tool message
			messages = append(messages, RequestMessage{
				Role: RoleTool,
				Content: []Content{
					{Type: ContentTypeText, Text: fmt.Sprintf("[tool:%s result]\n%s", tc.Name, result)},
				},
				ToolCallID: tc.ID,
			})
		}

	}
}

func (c *Client) findTool(tools []Tool, name string) Tool {
	for _, tool := range tools {
		if tool.Info().Name == name {
			return tool
		}
	}
	return nil
}

func (c *Client) Request(req *Request) (*Response, error) {
	if c.Provider == nil {
		return nil, errors.New("provider is not defined")
	}
	return c.Provider.Chat(req)
}
