package log

import (
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/boshangad/v1/app/helpers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Zap struct {
	// 日志配置等级
	level zapcore.Level
	// 日志存放日志
	Director string `json:"director,omitempty" yaml:"director,omitempty"`
	// 日志等级
	Level string `json:"level,omitempty" yaml:"level,omitempty"`
	// 是否显示行数
	ShowLine bool `json:"showLine,omitempty" yaml:"showLine,omitempty"`
	// 堆栈跟踪密钥
	StacktraceKey string `json:"stacktraceKey,omitempty" yaml:"stacktraceKey,omitempty"`
	// 格式化等级
	// LowercaseLevelEncoder 小写编码器(默认)
	// LowercaseColorLevelEncoder 小写编码器带颜色
	// CapitalLevelEncoder 大写编码器
	// CapitalColorLevelEncoder 大写编码器带颜色
	EncodeLevel string `json:"encodeLevel,omitempty" yaml:"encodeLevel,omitempty"`
	// 是否将日志输出到控制台
	LogInConsole bool `json:"logInConsole,omitempty" yaml:"logInConsole,omitempty"`
	// 日志内容前缀
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	// 内容格式化
	// json 表示格式为json
	// row 行列表示（默认）
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
	// 日志文件最大可用大小,小于等于0表示不限制，单位MB
	MaxSize int `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`
	// 最大可备份日志文件数量,小于等于0表示不限制
	MaxBackups int `json:"maxBackups,omitempty" yaml:"maxBackups,omitempty"`
	// 保留旧日志文件的最大天数,小于等于0表示不限制
	MaxAge int `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`
	// 是否使用gzip压缩轮换的日志文件，默认为true
	Compress bool `json:"compress,omitempty" yaml:"compress,omitempty"`
}

// 获取日志编码配置
func (that Zap) getEncoderConfig() (config zapcore.EncoderConfig) {
	// customTimeEncoder 自定义日志输出时间格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(that.Prefix + "2006/01/02 - 15:04:05.000"))
	}
	config = zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: that.StacktraceKey,
		// 默认行尾字符
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case that.EncodeLevel == "LowercaseLevelEncoder":
		// 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case that.EncodeLevel == "LowercaseColorLevelEncoder":
		// 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case that.EncodeLevel == "CapitalLevelEncoder":
		// 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case that.EncodeLevel == "CapitalColorLevelEncoder":
		// 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// 获取日志编码器
func (that Zap) getEncoder() zapcore.Encoder {
	encoderConfig := that.getEncoderConfig()
	switch strings.ToLower(that.Format) {
	case "json":
		return zapcore.NewJSONEncoder(encoderConfig)
	default:
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}

// 将文本等级转为数值型等级
func (that Zap) GetZapLevel(level string) (zapLevel zapcore.Level) {
	switch strings.ToLower(level) {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	case "dpanic":
		zapLevel = zap.DPanicLevel
	case "panic":
		zapLevel = zap.PanicLevel
	case "fatal":
		zapLevel = zap.FatalLevel
	default:
		zapLevel = zap.InfoLevel
	}
	return
}

// 获取Encoder的zapcore.Core
func (that Zap) getEncoderCore() (core zapcore.Core) {
	var writer zapcore.WriteSyncer
	writerSyncer := lumberjack.Logger{
		Filename:   path.Join(that.Director, "log."+time.Now().Format("2006010215")+".log"),
		MaxSize:    that.MaxSize,
		MaxBackups: that.MaxBackups,
		MaxAge:     that.MaxAge,
		Compress:   that.Compress,
	}
	if that.LogInConsole {
		writer = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(&writerSyncer),
			zapcore.AddSync(os.Stdout),
		)
	} else {
		writer = zapcore.AddSync(&writerSyncer)
	}
	// 文件拆分
	// core := zapcore.NewTee(
	// 	zapcore.NewCore(that.getEncoder(), writer, that.level),
	// 	zapcore.NewCore(that.getEncoder(), writer, that.level),
	// )
	return zapcore.NewCore(that.getEncoder(), writer, that.level)
}

// 默认zap配置
func DefaultZapConfig() Zap {
	return Zap{
		Director:      "",
		Level:         "debug",
		ShowLine:      true,
		StacktraceKey: "",
		EncodeLevel:   "LowercaseLevelEncoder",
		LogInConsole:  true,
		Prefix:        "",
		Format:        "row",
		MaxSize:       0,
		MaxBackups:    0,
		MaxAge:        0,
		Compress:      true,
	}
}

// 实例化配置
func NewLogger(zapConfig Zap) (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if zapConfig.Director == "" {
		zapConfig.Director = path.Join(helpers.GetCurrentDirectory(), "runtime/log")
	}
	if !helpers.IsDir(zapConfig.Director) {
		err := os.Mkdir(zapConfig.Director, os.ModePerm)
		if err != nil {
			log.Println("", err)
		}
	}
	// 初始化配置文件的Level
	level := zapConfig.GetZapLevel(zapConfig.Level)
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(zapConfig.getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(zapConfig.getEncoderCore())
	}
	if zapConfig.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
