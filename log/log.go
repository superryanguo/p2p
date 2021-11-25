package log

import (
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	p       *zap.Logger
	sp      *zap.SugaredLogger
	inited  bool
	logFile string
	level   zap.AtomicLevel
)

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func init() {
	if !inited {
		logFile = "./peer2peer.log"
		file := zapcore.Lock(getLogWriter())
		console := zapcore.Lock(os.Stdout)
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
		cfg := zap.NewDevelopmentEncoderConfig()
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder := zapcore.NewConsoleEncoder(cfg)

		core := zapcore.NewTee(
			zapcore.NewCore(encoder, console, level),
			zapcore.NewCore(encoder, file, level))

		p = zap.New(core, zap.Development(), zap.AddCaller(), zap.AddCallerSkip(1))
		sp = p.Sugar()

		inited = true
	}
}

//LogServer make the log leverl could be changed on the fly
//curl -X PUT localhost:8080/log/level?level=debug
//curl -X PUT localhost:8080/log/level -d level=debug
//curl -X PUT localhost:8080/log/level -H "Content-Type: application/json" -d '{"level":"debug"}'
func LogServer(addr string) error {
	http.HandleFunc("/log/level", level.ServeHTTP)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}

func Sync() error {
	return p.Sync()
}
func Debug(msg string, fields ...zap.Field) {
	p.Debug(msg, fields...)
}
func Info(msg string, fields ...zap.Field) {
	p.Info(msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {
	p.Warn(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	p.Error(msg, fields...)
}
func Fatal(msg string, fields ...zap.Field) {
	p.Fatal(msg, fields...)
}
func Panic(msg string, fields ...zap.Field) {
	p.Panic(msg, fields...)
}
func Debugf(template string, args ...interface{}) {
	sp.Debugf(template, args...)
}
func Infof(template string, args ...interface{}) {
	sp.Infof(template, args...)
}
func Warnf(template string, args ...interface{}) {
	sp.Warnf(template, args...)
}
func Errorf(template string, args ...interface{}) {
	sp.Errorf(template, args...)
}
func Fatalf(template string, args ...interface{}) {
	sp.Fatalf(template, args...)
}
func Panicf(template string, args ...interface{}) {
	sp.Panicf(template, args...)
}
