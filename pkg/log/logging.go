package log

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	defaultLogger *zap.Logger
	sugar         *zap.SugaredLogger
	atomicLevel   zap.AtomicLevel
)

// Init 初始化日志系统
func Init(filePath string, logLevel zapcore.Level, isDebug bool) {
	atomicLevel = zap.NewAtomicLevelAt(logLevel)

	// --- 1. 控制台编码配置 ---
	consoleConfig := zap.NewDevelopmentEncoderConfig()
	// 重点：将 TID/GID 格式化进 Level 标签，实现 INFO |tid=11,gid=1|
	consoleConfig.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("%s |tid=%d,gid=%d|", l.CapitalString(), getTID(), getGID()))
	}
	consoleConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)

	// --- 2. 文件编码配置 (JSON) ---
	fileConfig := zap.NewProductionEncoderConfig()
	fileConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileConfig)

	// --- 3. 核心输出介质 ---
	var cores []zapcore.Core

	// 控制台输出
	cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), atomicLevel))

	// 文件输出
	if filePath != "" {
		lumberjackLogger := &lumberjack.Logger{
			Filename:   filePath,
			MaxSize:    256,
			MaxBackups: 10,
			MaxAge:     30,
			Compress:   true,
		}

		// 异步刷盘，防止磁盘 I/O 阻塞游戏逻辑
		asyncWriter := &zapcore.BufferedWriteSyncer{
			WS:            zapcore.AddSync(lumberjackLogger),
			Size:          512 * 1024,
			FlushInterval: time.Second,
		}

		fileCore := zapcore.NewCore(fileEncoder, asyncWriter, atomicLevel)
		if !isDebug {
			fileCore = zapcore.NewSamplerWithOptions(fileCore, time.Second, 100, 10)
		}
		cores = append(cores, fileCore)
	}

	defaultLogger = zap.New(zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	sugar = defaultLogger.Sugar()
}

// ---------------------- 高性能 API (已移除冗余 GID) ----------------------

func Info(msg string, fields ...zap.Field) {
	defaultLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	defaultLogger.Warn(msg, fields...)
}

func Error(err error, fields ...zap.Field) {
	if err == nil {
		return
	}
	// 自动带上 error 详情，不重复带 gid
	defaultLogger.Error(err.Error(), append(fields, zap.Error(err))...)
}

// ---------------------- Sugar 模式 ----------------------

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}

// ---------------------- 业务专属 ----------------------

func Player(playerID uint64, msg string, fields ...zap.Field) {
	// 仅带上业务相关的 pid
	defaultLogger.Info(msg, append(fields, zap.Uint64("pid", playerID))...)
}

func SetLevel(lvl zapcore.Level) {
	atomicLevel.SetLevel(lvl)
}

func Sync() {
	if defaultLogger != nil {
		_ = defaultLogger.Sync()
	}
}
