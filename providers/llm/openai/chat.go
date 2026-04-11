package openai

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/deep-project/aikit/pkg/chat"
	"github.com/go-resty/resty/v2"
)

// Chat Completions 标准
type Chat struct {
	BaseURL string
	APIKey  string
	Model   string

	// 多模态
	// content是多条
	Multimodal bool
}

func NewChat(baseURL, apiKey, model string, multimodal bool) *Chat {
	return &Chat{BaseURL: baseURL, APIKey: apiKey, Model: model, Multimodal: multimodal}
}

func (c *Chat) Chat(req *chat.Request) (*chat.Response, error) {

	if c.BaseURL == "" {
		return nil, errors.New("baseURL is not defined")
	}
	if c.APIKey == "" {
		return nil, errors.New("apiKey is not defined")
	}
	if c.Model == "" {
		return nil, errors.New("model is not defined")
	}

	var chatResp Response
	var chatReq = buildRequest(c.Model, req, c.Multimodal)

	r, err := resty.New().R().
		SetHeader("Authorization", c.getAuthorization()).
		SetHeader("Content-Type", "application/json").
		SetBody(chatReq).
		SetResult(&chatResp).
		Post(c.getFullURL())

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	if r.StatusCode() >= 400 {
		return nil, fmt.Errorf("status: %s \n body: %s", r.Status(), r.String())
	}
	//fmt.Printf("Body: %#v\n", r.StatusCode())
	return toChatResponse(&chatResp), nil
}

func (c *Chat) getFullURL() string {
	fullURL, _ := url.JoinPath(c.BaseURL, "chat/completions")
	return fullURL
}

func (c *Chat) getAuthorization() string {
	return fmt.Sprintf("Bearer %s", c.APIKey)
}
