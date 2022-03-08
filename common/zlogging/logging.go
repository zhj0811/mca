package zlogging

import (
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config is used to provide dependencies to a Logging instance.
type Config struct {
	// Format is the log record format specifier for the Logging instance. If the
	// spec is the string "json", log records will be formatted as JSON. Any
	// other string will be provided to the FormatEncoder. Please see
	// fabenc.ParseFormat for details on the supported verbs.
	//
	// If Format is not provided, a default format that provides basic information will
	// be used.
	Format string

	// LogSpec determines the log levels that are enabled for the logging system. The
	// spec must be in a format that can be processed by ActivateSpec.
	//
	// If LogSpec is not provided, loggers will be enabled at the INFO level.
	LogSpec string

	// Writer is the sink for encoded and formatted log records.
	//
	// If a Writer is not provided, os.Stderr will be used as the log sink.
	Writer io.Writer

	Encoding int8
}

type Encoding int8

const (
	CONSOLE Encoding = iota
	JSON
	LOGFMT
)

type Logging struct {
	level zapcore.Level
	//loggers       map[string]*zap.Logger
	encoderConfig zapcore.EncoderConfig
	writer        zapcore.WriteSyncer
	mutex         sync.RWMutex
	encoding      Encoding
}

// Write satisfies the io.Write contract. It delegates to the writer argument
// of SetWriter or the Writer field of Config. The Core uses this when encoding
// log records.
func (l *Logging) Write(b []byte) (int, error) {
	l.mutex.RLock()
	w := l.writer
	l.mutex.RUnlock()

	return w.Write(b)
}

// Sync satisfies the zapcore.WriteSyncer interface. It is used by the Core to
// flush log records before terminating the process.
func (l *Logging) Sync() error {
	l.mutex.RLock()
	w := l.writer
	l.mutex.RUnlock()

	return w.Sync()
}

//	if format == "json" {
//		l.encoding = JSON
//		return nil
//	}
//
//	if format == "logfmt" {
//		l.encoding = LOGFMT
//		return nil
//	}
//
//	formatters, err := enc.ParseFormat(format)
//	if err != nil {
//		return err
//	}
//	l.multiFormatter.SetFormatters(formatters)
//	l.encoding = CONSOLE
//
//	return nil
//}
//
// SetWriter controls which writer formatted log records are written to.
// Writers, with the exception of an *os.File, need to be safe for concurrent
// use by multiple go routines.
func (l *Logging) SetWriter(w io.Writer) io.Writer {
	var sw zapcore.WriteSyncer
	switch t := w.(type) {
	case *os.File:
		sw = zapcore.Lock(t)
	case zapcore.WriteSyncer:
		sw = t
	default:
		sw = zapcore.AddSync(w)
	}

	l.mutex.Lock()
	ow := l.writer
	l.writer = sw
	l.mutex.Unlock()

	return ow
}

//
// Apply applies the provided configuration to the logging system.
func (l *Logging) Apply(c Config) error {
	//err := l.SetFormat(c.Format)
	//if err != nil {
	//	return err
	//}

	//if c.LogSpec == "" {
	//	c.LogSpec = os.Getenv("LOGGING_SPEC")
	//}
	//if c.LogSpec == "" {
	//	c.LogSpec = defaultLevel.String()
	//}

	//err = l.LoggerLevels.ActivateSpec(c.LogSpec)
	//if err != nil {
	//	return err
	//}

	if c.Writer == nil {
		c.Writer = os.Stderr
	}
	l.SetWriter(c.Writer)

	return nil
}

// ZapLogger instantiates a new zap.Logger with the specified name. The name is
// used to determine which log levels are enabled.
func (l *Logging) Logger(name string) *zap.SugaredLogger {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	//conf := zap.NewDevelopmentEncoderConfig()
	//zLogger,_ := zap.NewDevelopment()
	var encoder zapcore.Encoder
	switch l.encoding {
	case JSON:
		encoder = zapcore.NewJSONEncoder(l.encoderConfig)
	//case LOGFMT:
	//	encoder = zaplogfmt.NewEncoder(l.encoderConfig)
	default: //case CONSOLE:
		encoder = zapcore.NewConsoleEncoder(l.encoderConfig)
	}

	core := zapcore.NewCore(encoder, l, &l.level)
	zLogger := zap.New(core, []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}...)
	//l.loggers[name] = zLogger
	return zLogger.Named(name).Sugar()
}

//
//// Encoding satisfies the Encoding interface. It determines whether the JSON or
//// CONSOLE encoder should be used by the Core when log records are written.
//func (l *Logging) Encoding() Encoding {
//	l.mutex.RLock()
//	e := l.encoding
//	l.mutex.RUnlock()
//	return e
//}
//
//// Sync satisfies the zapcore.WriteSyncer interface. It is used by the Core to
//// flush log records before terminating the process.
//func (l *Logging) Sync() error {
//	l.mutex.RLock()
//	w := l.writer
//	l.mutex.RUnlock()
//
//	return w.Sync()
//}
//
//// Write satisfies the io.Write contract. It delegates to the writer argument
//// of SetWriter or the Writer field of Config. The Core uses this when encoding
//// log records.
//func (l *Logging) Write(b []byte) (int, error) {
//	l.mutex.RLock()
//	w := l.writer
//	l.mutex.RUnlock()
//
//	return w.Write(b)
//}
//
//// Logger instantiates a new FabricLogger with the specified name. The name is
//// used to determine which log levels are enabled.
//func (l *Logging) Logger(name string) *zap.SugaredLogger {
//	zl := l.ZapLogger(name)
//	return zl.WithOptions(zap.AddCallerSkip(1)).Sugar()
//}
//

// New creates a new logging system and initializes it with the provided
// configuration.
func New(c Config) (*Logging, error) {
	encoderConfig := DefaultEncoderConfig()
	//encoderConfig.NameKey = "name"
	l := &Logging{
		//loggers: map[string]*zap.Logger{},
		//LoggerLevels: &LoggerLevels{
		//	defaultLevel: defaultLevel,
		//},
		level:         defaultLevel,
		encoderConfig: encoderConfig,
		writer:        zapcore.AddSync(os.Stdout),
		encoding:      CONSOLE,
		//multiFormatter: fabenc.NewMultiFormatter(),
	}
	err := l.Apply(c)
	if err != nil {
		return nil, err
	}
	return l, nil
}
