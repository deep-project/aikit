# Chat

> 基本的消息交互 
> 
> 内部处理tool call loop



## 基本使用
```go
import (
  "github.com/deep-project/aikit/pkg/chat"
  "github.com/deep-project/aikit/providers/llm/openai"
)

func main(){
  p := openai.NewChat("https://api.openai.com/v1", "sk-xxxxxxxxxxxx", "gpt-4.1", false)
  c := chat.New(chat.WithProvider(p))

  // 请求
  r, err := c.Send(&chat.Request{
    Messages: []chat.RequestMessage{
      {
        Role: chat.RoleUser,
        Content: []chat.Content{
          {Type: chat.ContentTypeText, Text: "你好，AI"},
        },
      },
    },
  })

  // 响应
  // {
  //   "output": {
  //     "role": "assistant",
  //     "content": [
  //       {
  //         "type": "text",
  //         "text": "你好！很高兴见到你！ 我是一个乐于助人的AI助手。无论你想聊天、寻求帮助、解答问题，还是需要创意灵感，我都很乐意为你提供支持！我都很期待与你交流！"
  //       }
  //     ],
  //     "tool_calls": null
  //   },
  //   "messages":..., // 完整消息列表
  //   "response":..., // 大模型响应内容
  // }

}
```

## 使用tools
```go
import (
  "github.com/deep-project/aikit/pkg/chat"
  "github.com/deep-project/aikit/providers/llm/openai"
  "github.com/deep-project/aikit/test/tools"
)

func main(){
  p := openai.NewChat("https://api.openai.com/v1", "sk-xxxxxxxxxxxx", "gpt-4.1", false)
  c := chat.New(chat.WithProvider(p))
  r, err := c.Send(&chat.Request{
    Tools: []chat.Tool{
			tools.NewChatGetStock(), // 查询库存的tool
		},
    Messages: []chat.RequestMessage{
      {
        Role: chat.RoleUser,
        Content: []chat.Content{
          {Type: chat.ContentTypeText, Text: "你好，帮我看看270854的库存还有多少？"},
        },
      },
    },
  })
}
```

#### tool 示例
```go

package tools

import "github.com/deep-project/aikit/pkg/chat"

type ChatGetStock struct {
}

func NewChatGetStock() *ChatGetStock {
	return &ChatGetStock{}
}

func (ChatGetStock) Info() *chat.ToolInfo {
	return &chat.ToolInfo{
		Name:        "getStock",
		Description: "获取指定编码的库存",
		Parameters: []chat.ToolParameter{
			{
				Name:        "code",
				Type:        chat.ToolTypeString,
				Description: "商品编码",
			},
		},
	}
}

func (ChatGetStock) Call(args map[string]any) (string, error) {
  code, ok := args["code"]
	if !ok {
		return "", fmt.Errorf("missing param: code")
	}
	fmt.Println("code", code)
	return "当前商品的库存是：666", nil
}

```