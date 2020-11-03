package enterprise

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var once sync.Once

func Logger() *zap.Logger {
	once.Do(func() {
		logger = InitLogger()
	})
	return logger
}

func InitLogger() *zap.Logger {
	now := time.Now()
	f := fmt.Sprintf("logs/enterprise/%04d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// Create info level writer
	infoWrite := zapcore.AddSync(&lumberjack.Logger{
		Filename:   f,   // 输出文件
		MaxSize:    500, // 日志文件最大大小（单位：MB）
		MaxBackups: 3,   // 保留的旧日志文件最大数量
		MaxAge:     30,  // 保存日期
	})

	// Create error level writer
	errWrite := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/enterprise/err", // 输出文件
		MaxSize:    500,                   // 日志文件最大大小（单位：MB）
		MaxBackups: 3,                     // 保留的旧日志文件最大数量
		MaxAge:     30,                    // 保存日期
	})

	// Definition dynamic level
	dynamicLevel := zap.NewAtomicLevel()

	// Definition info level
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel
	})

	// Definition error level
	errLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		// 输出到控制台：动态等级
		zapcore.NewCore(zapcore.NewConsoleEncoder(NewEncoderConfig()), os.Stdout, dynamicLevel),
		//// 输出到文件：信息等级、json格式
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), infoWrite, infoLevel),
		//// 输出到文件：错误等级、json格式
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), errWrite, errLevel),
	)
	return zap.New(core, zap.AddCaller())

}

// 返回信息配置
func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",                              // 时间键
		LevelKey:       "L",                              // 日志等级键
		NameKey:        "N",                              // 日志记录器名
		CallerKey:      "C",                              // 日志文件信息键
		MessageKey:     "M",                              // 日志消息键
		StacktraceKey:  "S",                              // 堆栈键
		LineEnding:     zapcore.DefaultLineEnding,        // 友好日志换行符
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 友好日志等级名大小写（info -> INFO）
		EncodeTime:     TimeEncoder,                      // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder,    // 时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,       // 日志文件信息（包/文件.go:行号）
	}
}

// 格式化时间
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}
