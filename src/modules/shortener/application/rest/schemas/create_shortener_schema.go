package schemas

type CreateShortenerSchema struct {
	Url string `json:"url" validate:"required"`
}
