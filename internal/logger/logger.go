package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorCyan   = "\033[36m"
)

type colorHandler struct {
	level slog.Level
	out   io.Writer
	mu    sync.Mutex
	attrs []slog.Attr
}

func (h *colorHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *colorHandler) Handle(_ context.Context, r slog.Record) error {
	color, label := levelStyle(r.Level)
	ts := r.Time
	if ts.IsZero() {
		ts = time.Now()
	}
	line := fmt.Sprintf("%s[%s]%s %s %s", color, label, colorReset, ts.Format("15:04:05"), r.Message)
	for _, a := range h.attrs {
		line += fmt.Sprintf(" %s=%v", a.Key, a.Value.Any())
	}
	r.Attrs(func(a slog.Attr) bool {
		line += fmt.Sprintf(" %s=%v", a.Key, a.Value.Any())
		return true
	})
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := fmt.Fprintln(h.out, line)
	return err
}

func (h *colorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := make([]slog.Attr, len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)
	copy(newAttrs[len(h.attrs):], attrs)
	return &colorHandler{level: h.level, out: h.out, attrs: newAttrs}
}

func (h *colorHandler) WithGroup(_ string) slog.Handler { return h }

func levelStyle(level slog.Level) (color, label string) {
	switch {
	case level >= slog.LevelError:
		return colorRed, "ERROR"
	case level >= slog.LevelWarn:
		return colorYellow, "WARN"
	case level >= slog.LevelInfo:
		return colorGreen, "INFO"
	default:
		return colorCyan, "DEBUG"
	}
}

var defaultLogger *slog.Logger

func init() {
	level := slog.LevelInfo
	if os.Getenv("DEBUG") == "1" || os.Getenv("LOG_LEVEL") == "debug" {
		level = slog.LevelDebug
	}
	defaultLogger = slog.New(&colorHandler{level: level, out: os.Stdout})
}

func Debug(msg string, args ...any) { defaultLogger.Debug(msg, args...) }
func Info(msg string, args ...any)  { defaultLogger.Info(msg, args...) }
func Warn(msg string, args ...any)  { defaultLogger.Warn(msg, args...) }
func Error(msg string, args ...any) { defaultLogger.Error(msg, args...) }

func Fatal(msg string, args ...any) {
	Error(msg, args...)
	os.Exit(1)
}
