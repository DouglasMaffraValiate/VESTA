package logger

import (
	"time"

	"github.com/DouglasMaffraValiate/VESTA/pkg/iff"
)

type LogEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level,omitempty"`
	Message   string    `json:"message,omitempty"`
	Service   string    `json:"service,omitempty"`
	RequestID string    `json:"request_id,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	Fields    []field   `json:"fields"`
}

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

func newLogEvent(level, message, service, requestID, userID string, fields ...field) LogEvent {

	return LogEvent{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Service:   service,
		RequestID: iff.IF(len(requestID) >= 0, requestID, "request_id"),
		UserID:    iff.IF(len(userID) >= 0, userID, "user_id"),
		Fields:    fields,
	}

}

func NewField[T value](name string, value T) field {
	if len(name) == 0 {
		return field{}
	}

	if v, ok := any(value).(string); ok && len(v) == 0 {
		return field{}
	}

	return field{
		Name:  name,
		Value: value,
	}
}

type field struct {
	Name  string `json:"name,omitempty"`
	Value any    `json:"value,omitempty"`
}

type value interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}
