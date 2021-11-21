package config

// 跨域请求配置
type Cors struct {
	// 是否允许跨域
	AllowCors bool `mapstructure:"allowCors" json:"allowCors,omitempty" yaml:"allowCors"`
	// 允许跨域的地址
	AllowOrigins []string `mapstructure:"allowOrigins" json:"allowOrigins,omitempty" yaml:"allowOrigins"`
	// 允许跨域的请求方法
	AllowMethods []string `mapstructure:"allowMethods" json:"allowMethods,omitempty" yaml:"allowMethods"`
	// 响应报头指示的请求的响应是否可以暴露于该页面
	AllowCredentials bool `mapstructure:"allowCredentials" json:"allowCredentials,omitempty" yaml:"allowCredentials"`
	// 暴露的请求头
	ExposeHeaders []string `mapstructure:"exposeHeaders" json:"exposeHeaders,omitempty" yaml:"exposeHeaders"`
	// 允许缓存的时长
	MaxAge int64 `mapstructure:"maxAge" json:"maxAge,omitempty" yaml:"maxAge"`
}
