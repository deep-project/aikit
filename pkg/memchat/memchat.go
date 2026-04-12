package memchat

import "github.com/deep-project/aikit/pkg/chat"

type Client struct {
	Chat *chat.Client
}

type Option func(*Client)

func New(options ...Option) *Client {
	cli := &Client{}
	for _, opt := range options {
		opt(cli)
	}
	return cli
}

func WithChat(ch *chat.Client) Option {
	return func(c *Client) {
		c.SetChat(ch)
	}
}

func (c *Client) SetChat(ch *chat.Client) *Client {
	c.Chat = ch
	return c
}
