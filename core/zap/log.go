package zap

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

type Log struct {
	Filename string `json:"filename" yaml:"filename"`
	MaxSize int `json:"maxsize" yaml:"maxsize"`
	MaxAge int `json:"maxage" yaml:"maxage"`
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`
	LocalTime bool `json:"localtime" yaml:"localtime"`
	Compress bool `json:"compress" yaml:"compress"`
}

func (that Log) Write(p []byte) (n int, err error) {
	fileWriter := lumberjack.Logger{
		Filename: that.Filename,
		MaxSize: that.MaxSize, // megabytes
		MaxBackups: that.MaxBackups,
		MaxAge: that.MaxAge, //days
		Compress: that.Compress, // disabled by default
	}
	return fileWriter.Write(p)
}

func (that Log) Sync() error {
	return nil
}