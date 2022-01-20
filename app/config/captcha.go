package config

import (
	"github.com/boshangad/v1/app/helpers"
)

type Captcha struct {
	// 验证码类型
	// image 图片、audio 音频、chinese 中文、math 数学、digit 点位、language 语言
	Type string `json:"type,omitempty" yaml:"type"`
	// 验证码宽度
	Width helpers.Int `json:"width,omitempty" yaml:"width"`
	// 验证码最低宽度
	MinWidth helpers.Int `json:"minWidth,omitempty" yaml:"minWidth"`
	// 验证码最大宽度
	MaxWidth helpers.Int `json:"maxWidth,omitempty" yaml:"maxWidth"`
	// 验证码高度
	Height helpers.Int `json:"height,omitempty" yaml:"height"`
	// 验证码最底高度
	MinHeight helpers.Int `json:"minHeight,omitempty" yaml:"minHeight"`
	// 验证码最大高度
	MaxHeight helpers.Int `json:"maxHeight,omitempty" yaml:"maxHeight"`
	// 验证码最大长度
	Length helpers.Int `json:"length,omitempty" yaml:"length"`
	// 验证码噪点数
	NoiseCount helpers.Int `json:"nosiseCount,omitempty" yaml:"nosiseCount"`
	// 验证码最小噪点数
	MinNoiseCount helpers.Int `json:"minNosiseCount,omitempty" yaml:"minNosiseCount"`
	// 验证码最大噪点数
	MaxNoiseCount helpers.Int `json:"maxNosiseCount,omitempty" yaml:"maxNosiseCount"`
	// 语音验证码语言、或语言验证码有效
	Language string `json:"language,omitempty" yaml:"language"`
	// 生成验证码的字体
	Fonts []string `json:"fonts,omitempty" yaml:"fonts"`
	// 背景颜色
	Color string `json:"color,omitempty" yaml:"color"`
	// 缓存,支持 memory, file, redis, 默认 memory
	Store string `json:"cache,omitempty" yaml:"cache"`
	// 失效时间，单位秒, 默认 30
	ExpireTime helpers.Int `json:"expireTime,omitempty" yaml:"expireTime"`
}
