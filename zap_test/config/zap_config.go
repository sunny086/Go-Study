package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type ZapConfig struct {
	Level         string // 日志级别
	Format        string // 日志格式
	Prefix        string // 日志前缀
	StacktraceKey string // 堆栈跟踪键
	Director      string // 日志目录
	EncodeLevel   string // 编码级别
	MaxAge        int    // 最大保存时间
	ShowLine      bool   // 是否显示行号
	LogInConsole  bool   // 是否在控制台打印
}

var ZapConfigDefault = ZapConfig{
	Level:         "debug",
	Format:        "json",
	Prefix:        "[xujs/zap/test]",
	StacktraceKey: "stacktrace",
	Director:      "./logs",
	EncodeLevel:   "CapitalColorLevelEncoder",
	MaxAge:        7,
	ShowLine:      true,
	LogInConsole:  true,
}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
func (z *ZapConfig) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 根据字符串转化为 zapcore.Level
func (z *ZapConfig) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
