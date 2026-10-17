package main

import "fmt"

// Продукт
type Logger interface {
	Log(message string)
}

// Конкретные продукты
type ConsoleLogger struct{}
func (l *ConsoleLogger) Log(msg string) { fmt.Println("CONSOLE:", msg) }

type FileLogger struct{}
func (l *FileLogger) Log(msg string) { fmt.Println("FILE:", msg) }

// Фабричный метод
type LoggerType string

const (
	CONSOLE LoggerType = "console"
	FILE    LoggerType = "file"
)

func NewLogger(t LoggerType) Logger {
	switch t {
	case FILE:
		return &FileLogger{}
	case CONSOLE:
		fallthrough
	default:
		return &ConsoleLogger{}
	}
}

func main() {
	logger := NewLogger(FILE)
	logger.Log("Application started")
}