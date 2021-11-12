package config

type Captcha struct {
	// 验证码类型
	Type string `json:"type,omitempty" mapstructure:"type" yaml:"type"`
	// 验证码长度
	Length    int `json:"length,omitempty" mapstructure:"length" yaml:"length"`
	MinLength int `json:"min_length,omitempty" mapstructure:"min_length" yaml:"min_length"`
	MaxLength int `json:"max_length,omitempty" mapstructure:"max_length" yaml:"max_length"`
	// 宽度
	Width int `json:"width,omitempty" mapstructure:"width" yaml:"width"`
	// 高度
	Height int `json:"height,omitempty" mapstructure:"height" yaml:"height"`
	// 线条数量
	NoiseCount    int `json:"noise_count,omitempty" mapstructure:"noise_count" yaml:"noise_count"`
	MinNoiseCount int `json:"min_noise_count,omitempty" mapstructure:"min_noise_count" yaml:"min_noise_count"`
	MaxNoiseCount int `json:"max_noise_count,omitempty" mapstructure:"max_noise_count" yaml:"max_noise_count"`
	// 字符串来源
	SourceType string `json:"source_type,omitempty" mapstructure:"source_type" yaml:"source_type"`
	SourceText string `json:"source_text,omitempty" mapstructure:"source_text" yaml:"source_text"`
	// 字体
	Fonts []string `json:"fonts,omitempty" mapstructure:"fonts" yaml:"fonts"`
	// 语言
	Language string `json:"language,omitempty" mapstructure:"language"`
}
