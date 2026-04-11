# 🪼 AI Kit


## 介绍

AI工具包

## 安装

```go
go get github.com/deep-project/aikit
```


## 使用

### Chat使用

#### 1. 基本使用
```go
import (
  "github.com/deep-project/aikit/pkg/chat"
  "github.com/deep-project/aikit/providers/llm/openai"
)

func main(){
  p := openai.NewChat("https://api.openai.com/v1", "sk-xxxxxxxxxxxx", "gpt-4.1", false)
  c := chat.New(chat.WithProvider(p))
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
}
```

#### 2. 使用tools
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
    Tools: []chat.Tool{tools.NewChatGetStock()},
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