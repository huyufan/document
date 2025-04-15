package logger

import (
	"sync/atomic"
)

type Level int

var (
	Levels = struct {
		Off    Level
		Panic  Level
		Error  Level
		Warn   Level
		Info   Level
		Debug  Level
		Access Level
	}{
		Access: (-1),
		Off:    (0),
		Panic:  (1),
		Error:  (2),
		Warn:   (3),
		Info:   (4),
		Debug:  (5),
	}

	levelMap = map[Level]string{Levels.Access: "Access", Levels.Off: "Off", Levels.Panic: "Panic", Levels.Error: "Error", Levels.Warn: "Warn", Levels.Info: "Info", Levels.Debug: "Debug"}

	cfgLevels = map[string]Level{
		"access": Levels.Access,
		"off":    Levels.Off,
		"panic":  Levels.Panic,
		"error":  Levels.Error,
		"warn":   Levels.Warn,
		"info":   Levels.Info,
		"debug":  Levels.Debug,
	}

	logCount  uint64
	dropCount uint64
	errCount  uint64
)

func Stats() (logs, pending, drop, errs uint64) {
	return atomic.LoadUint64(&logCount), uint64(1), atomic.LoadUint64(&dropCount), atomic.LoadUint64(&errCount)
}

type Logger struct {
	level               Level
	sample, sampleCount uint64
}

func (level Level) String() string {
	return levelMap[level]
}

func New(level Level) (l *Logger) {
	l = new(Logger)
	l.level = level
	l.sample = 1
	return
}

func (l *Logger) Printf(level Level, prefix, format string, v ...interface{}) {
	switch {
	case level == Levels.Access:
		count := atomic.AddUint64(&l.sampleCount, 1)
		if l.sample == 0 || count%l.sample != 0 {
			return
		}
	case level > l.level, level == Levels.Off:
		return
	}
}

func (l *Logger) Debug(prefix, format string, v ...interface{}) {
	l.Printf(Levels.Debug, prefix, format, v...)
}

func (l *Logger) Info(prefix, format string, v ...interface{}) {
	l.Printf(Levels.Info, prefix, format, v...)
}

func (l *Logger) Warn(prefix, format string, v ...interface{}) {
	l.Printf(Levels.Warn, prefix, format, v...)
}

func (l *Logger) Error(prefix, format string, v ...interface{}) {
	l.Printf(Levels.Error, prefix, format, v...)
}

func (l *Logger) Panic(prefix, format string, v ...interface{}) {
	l.Printf(Levels.Panic, prefix, format, v...)
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

func (l *Logger) Level() Level {
	return l.level
}

func (l *Logger) SetAccessLogSample(sample uint64) {
	atomic.StoreUint64(&l.sample, sample)
}
