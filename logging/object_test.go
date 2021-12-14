package logging

import (
	"bytes"
	"io"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createLogger(out io.Writer) *zap.Logger {
	w := zapcore.AddSync(out)
	zap.NewExample()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			NameKey:        "logger",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		}),
		w,
		zap.DebugLevel,
	)
	return zap.New(core)
}

func Test_userLog(t *testing.T) {
	buf := new(bytes.Buffer)
	l := createLogger(buf)
	l.Info("found user", zap.Any("user", &user{name: "aereal", age: 17}))
	got := buf.String()
	expected := `{"level":"info","msg":"found user","user":{"name":"aereal","age":17}}` + "\n"
	if got != expected {
		t.Errorf("log entry mismatch\n\texpected:%q\n\tgot:%q", expected, got)
	}
}
