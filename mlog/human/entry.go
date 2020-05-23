package human

import (
	"fmt"
	"strings"
	"time"

	"medrepo-server/mlog"
)

type LogEntry struct {
	Time time.Time
	Level string
	Message string
	Caller string
	Fields []mlog.Field
}

func (f LogEntry) String() string {
	var sb strings.Builder
	if !f.Time.IsZero() {
		sb.WriteString(f.Time.Format(time.RFC3339Nano))
		sb.WriteRune(' ')
	}
	if f.Level != "" {
		sb.WriteString(f.Level)
		sb.WriteRune(' ')
	}
	if f.Caller != "" {
		sb.WriteString(f.Caller)
		sb.WriteRune(' ')
	}
	for _, field := range f.Fields {
		sb.WriteString(field.Key)
		sb.WriteRune('=')
		sb.WriteString(fmt.Sprintf(field.Interface))
		sb.WriteRune(' ')
	}
	if f.Message != "" {
		if strings.ContainsRune(f.Message, '\n') {
			sb.WriteRune('\n')
		}
		sb.WriteString(f.Message)
	}
	return sb.String()
}
