package grpczap

import (
	"fmt"

	"go.uber.org/zap"
)

type logger struct {
	log *zap.Logger
}

func New(log *zap.Logger) logger {
	return logger{log: log}
}

func (l logger) Fatal(args ...interface{}) {
	l.log.Fatal(fmt.Sprint(args...))
}

func (l logger) Fatalf(format string, args ...interface{}) {
	l.log.Fatal(fmt.Sprintf(format, args...))
}

func (l logger) Fatalln(args ...interface{}) {
	l.Fatal(args...)
}

func (l logger) Print(args ...interface{}) {
	l.log.Info(fmt.Sprint(args...))
}

func (l logger) Printf(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l logger) Println(args ...interface{}) {
	l.Print(args...)
}
