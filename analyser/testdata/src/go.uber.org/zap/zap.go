package zap

type Logger struct{}

func (l *Logger) Info(msg string, fields ...Field)  {}
func (l *Logger) Error(msg string, fields ...Field) {}
func (l *Logger) Warn(msg string, fields ...Field)  {}
func (l *Logger) Debug(msg string, fields ...Field) {}

type Field struct{}

func String(key, val string) Field {
	return Field{}
}
