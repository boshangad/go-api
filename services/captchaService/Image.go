package captchaService

type Image struct {
	WHSize
	NoiseCount uint64 `json:"noise_count,omitempty" form:"noise_count"`
	Length uint64 `json:"length,omitempty" form:"length"`
	Source string `json:"source,omitempty" form:"source"`
}
