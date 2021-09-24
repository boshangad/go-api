package core

import (
	"fmt"
	coreZap "github.com/boshangad/go-api/core/zap"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var level zapcore.Level

// Zap zap配置文件
func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.G_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.G_CONFIG.Zap.Director)
		_ = os.Mkdir(global.G_CONFIG.Zap.Director, os.ModePerm)
	}
	// 初始化配置文件的Level
	switch global.G_CONFIG.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.G_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.G_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.G_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder":
		// 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.G_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder":
		// 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.G_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder":
		// 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.G_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder":
		// 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.G_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer := GetWriteSyncer()
	return zapcore.NewCore(getEncoder(), writer, level)
}

// GetWriteSyncer 获取写入接口同步器
func GetWriteSyncer() zapcore.WriteSyncer {
	filename := path.Join(global.G_CONFIG.Zap.Director, "log." + time.Now().Format("20060102") + ".log")
	fileWriter := coreZap.Log{
		Filename: filename,
		MaxSize: 500, // megabytes
		MaxBackups: 0,
		MaxAge: 0, //days
		Compress: true, // disabled by default
	}
	if global.G_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.G_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}