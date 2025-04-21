package responses

type Message struct {
	MessageType string                 `json:"type"`
	Contents    map[string]interface{} `json:"messages"`
}
