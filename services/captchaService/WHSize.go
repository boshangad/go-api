package captchaService

type WHSize struct {
	WSize
	HSize
}

type WSize struct {
	Width int `json:"width,omitempty" form:"width"`
}

type HSize struct {
	Height int `json:"height,omitempty" form:"height"`
}