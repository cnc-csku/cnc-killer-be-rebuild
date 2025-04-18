package responses

type GoogleResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	PictureURL string `json:"picture_url"`
}
