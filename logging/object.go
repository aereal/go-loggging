package logging

import (
	"go.uber.org/zap/zapcore"
)

type user struct {
	name string
	age  int
}

var _ zapcore.ObjectMarshaler = &user{}

func (u *user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", u.name)
	enc.AddInt("age", u.age)
	return nil
}
