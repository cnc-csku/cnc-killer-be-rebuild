package responses

type Message struct {
	Type     string                 `json:"type"`
	Messages map[string]interface{} `json:"messages"`
}
