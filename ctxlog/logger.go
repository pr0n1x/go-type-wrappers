package ctxlog

import "context"

type OutputLogger interface {
	Output(calldepth int, s string) error
}

type PrintLogger interface {
	Print(v ...any)
	Println(v ...any)
	Printf(format string, v ...any)
}

type FatalLogger interface {
	Fatal(args ...any)
	Fatalln(args ...any)
	Fatalf(format string, args ...any)
}

type PanicLogger interface {
	Panic(v ...any)
	Panicln(v ...any)
	Panicf(format string, v ...any)
}

type Logger interface {
	OutputLogger
	PrintLogger
	FatalLogger
	PanicLogger
}

type ZeroLogger struct{}

func (l ZeroLogger) Output(_ int, _ string) error { return nil }
func (l ZeroLogger) Print(_ ...any)               {}
func (l ZeroLogger) Println(_ ...any)             {}
func (l ZeroLogger) Printf(_ string, _ ...any)    {}
func (l ZeroLogger) Fatal(_ ...any)               {}
func (l ZeroLogger) Fatalln(_ ...any)             {}
func (l ZeroLogger) Fatalf(_ string, _ ...any)    {}
func (l ZeroLogger) Panic(_ ...any)               {}
func (l ZeroLogger) Panicln(_ ...any)             {}
func (l ZeroLogger) Panicf(_ string, _ ...any)    {}

type ContextKey string

const CtxKeyLogger ContextKey = "ctxlog.logger"

type loggerCtxWrapper struct {
	l Logger
}

func CtxLogger(ctx context.Context) Logger {
	if v, ok := ctx.Value(CtxKeyLogger).(loggerCtxWrapper); ok {
		return v.l
	}
	return ZeroLogger{}
}

func CtxNonZeroLogger(ctx context.Context) (bool, Logger) {
	if v, ok := ctx.Value(CtxKeyLogger).(loggerCtxWrapper); ok {
		return true, v.l
	}
	return false, ZeroLogger{}
}

func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, CtxKeyLogger, loggerCtxWrapper{l: logger})
}
