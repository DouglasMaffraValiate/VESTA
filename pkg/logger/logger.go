package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"
)

type Logger interface {
	Debug(message, requestID, userID string, fields ...field)
	Info(message, requestID, userID string, fields ...field)
	Warn(message, requestID, userID string, fields ...field)
	Error(message, requestID, userID string, err error, fields ...field)
}

type slogLogger struct {
	service string
	logger  *slog.Logger
}

func NewLogger(service string, output io.Writer) Logger {
	if output == nil {
		output = os.Stdout
	}

	handler := slog.NewJSONHandler(output, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			if len(groups) != 0 {
				return attr
			}

			switch attr.Key {
			case slog.TimeKey:
				attr.Key = "timestamp"
			case slog.MessageKey:
				attr.Key = "message"
			}

			return attr
		},
	})

	return &slogLogger{
		service: service,
		logger:  slog.New(handler),
	}
}

func (l *slogLogger) Debug(
	message, requestID, userID string,
	fields ...field,
) {
	l.log(slog.LevelDebug, newLogEvent(DEBUG, message, l.service, requestID, userID, fields...))
}

func (l *slogLogger) Info(
	message, requestID, userID string,
	fields ...field,
) {
	l.log(slog.LevelInfo, newLogEvent(INFO, message, l.service, requestID, userID, fields...))
}

func (l *slogLogger) Warn(
	message, requestID, userID string,
	fields ...field,
) {
	l.log(slog.LevelWarn, newLogEvent(WARN, message, l.service, requestID, userID, fields...))
}

func (l *slogLogger) Error(
	message, requestID, userID string,
	err error,
	fields ...field,
) {
	if err != nil {
		// Copia os campos para não modificar o slice do chamador.
		fieldsWithError := make([]field, 0, len(fields)+1)
		fieldsWithError = append(fieldsWithError, fields...)
		fieldsWithError = append(fieldsWithError, NewField("error", err.Error()))
		fields = fieldsWithError
	}

	l.log(slog.LevelError, newLogEvent(ERROR, message, l.service, requestID, userID, fields...))
}

// log é o ponto central de emissão dos eventos via slog.
func (l *slogLogger) log(level slog.Level, event LogEvent) {
	ctx := context.Background()

	if !l.logger.Enabled(ctx, level) {
		return
	}

	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	attrs := make([]slog.Attr, 0, len(event.Fields)+3)

	if event.Service != "" {
		attrs = append(attrs, slog.String("service", event.Service))
	}
	if event.RequestID != "" {
		attrs = append(attrs, slog.String("request_id", event.RequestID))
	}
	if event.UserID != "" {
		attrs = append(attrs, slog.String("user_id", event.UserID))
	}

	fields := make([]field, 0, len(event.Fields))
	for _, f := range event.Fields {
		if f.Name != "" {
			fields = append(fields, f)
		}
	}

	attrs = append(attrs, slog.Any("fields", fields))

	record := slog.NewRecord(
		event.Timestamp,
		level,
		event.Message,
		0,
	)
	record.AddAttrs(attrs...)

	_ = l.logger.Handler().Handle(ctx, record)
}
