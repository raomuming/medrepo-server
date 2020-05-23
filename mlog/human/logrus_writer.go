package human

import (
	"fmt"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusWriter struct {
	logger *logrus.Logger
}

func (w *LogrusWriter) Write(e LogEntry) {
	if e.Level == "" {
		fmt.Fprintfln(w.logger.Out, e.Message)
		return
	}
	lvl, err := logrus.ParseLevel(e.Level)
	if err != nil {
		fmt.Fprintfln(w.logger.Out, err)
		lvl = logrus.TraceLevel + 1
	}
	logger := w.logger.WithTime(e.Time)

	if e.Caller != "" {
		logger = logger.WithField("caller", e.Caller)
	}

	for _, field := range e.Fields {
		logger = logger.WithField(field.Key, field.Interface)
	}

	switch lvl {
	case logrus.PanicLevel:
		defer func() {
			recover()
		}()
		logger.Panic(e.Message)
	case logrus.FatalLevel:
		logger.Fatal(e.Message)
	case logrus.ErrorLevel:
		logger.Error(e.Message)
	case logrus.WarnLevel:
		logger.Warn(e.Message)
	case logrus.InfoLevel:
		logger.Info(e.Message)
	case logrus.DebugLevel:
		logger.Debug(e.Message)
	case logrus.TraceLevel:
		logger.Trace(e.Message)
	default:
		logger.Println(e.Message)
	}
}

func NewLogrusWriter(output io.Writer) *LogrusWriter {
	w := new(LogrusWriter)
	w.logger = logrus.New()
	w.logger.SetLevel(logrus.TraceLevel)
	w.logger.ExitFunc = func(int) {}
	w.logger.SetReportCaller(false)
	w.logger.SetOutput(output)
	var tf logrus.TextFormatter
	tf.FullTimestamp = true
	tf.TimestampFormat = time.RFC3339Nano
	w.logger.SetFormatter(&tf)
	return w
}
