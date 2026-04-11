package openai

import "github.com/deep-project/aikit/pkg/chat"

type Content any

func buildContent(content []chat.Content, multimodal bool) Content {
	if multimodal {
		return buildContentMultimodal(content)
	}
	return buildContentText(content)
}

// 单文本消息
type ContentText string

func buildContentText(content []chat.Content) ContentText {
	for _, c := range content {
		if c.Type == chat.ContentTypeText {
			return ContentText(c.Text)
		}
	}
	return ""
}

// 多模态
type ContentMultimodal struct {
	Type     string           `json:"type"`
	Text     string           `json:"text,omitempty"`
	ImageURL *ContentImageURL `json:"image_url,omitempty"`
}

type ContentImageURL struct {
	URL    string `json:"url"`
	Detail string `json:"detail,omitempty"` // low / high / auto
}

func buildContentMultimodal(content []chat.Content) (res []ContentMultimodal) {
	for _, item := range content {
		c := ContentMultimodal{Type: string(item.Type), Text: item.Text}

		// ImageURL
		if item.ImageURL != nil {
			c.ImageURL = &ContentImageURL{
				URL:    item.ImageURL.URL,
				Detail: item.ImageURL.Detail,
			}
		}
		// TODO more types

		res = append(res, c)
	}
	return
}
