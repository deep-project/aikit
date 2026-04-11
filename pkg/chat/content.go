package chat

type Content struct {
	Type     ContentType      `json:"type"` // text
	Text     string           `json:"text,omitempty"`
	ImageURL *ContentImageURL `json:"image_url,omitempty"`
}

type ContentType string

const (
	ContentTypeText     = "text"
	ContentTypeImageURL = "image_url"
)

type ContentImageURL struct {
	URL    string `json:"url"`
	Detail string `json:"detail,omitempty"` // low / high / auto
}
