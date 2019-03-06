package logger

import (
	"log"
	"os"
)

type ILogger interface {
	Info(v ...interface{})
	Error(v ...interface{})
	Trace(v ...interface{})
	Warning(v ...interface{})
}

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(v ...interface{}) {
	log.New(os.Stdout, "T:", log.Ldate|log.Ltime|log.Lshortfile).Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	log.New(os.Stdout, "E:", log.Ldate|log.Ltime|log.Lshortfile).Println(v...)
}
func (l *Logger) Trace(v ...interface{}) {
	log.New(os.Stdout, "W:", log.Ldate|log.Ltime|log.Lshortfile).Println(v...)
}
func (l *Logger) Warning(v ...interface{}) {
	log.New(os.Stderr, "E:", log.Ldate|log.Ltime|log.Lshortfile).Println(v...)
}

type Logger2 struct {
}

func NewLogger2() *Logger2 {
	return &Logger2{}
}

func (l *Logger2) Info(v ...interface{}) {

}
func (l *Logger2) Error(v ...interface{}) {

}

func (l *Logger2) Trace(v ...interface{}) {

}

func (l *Logger2) Warning(v ...interface{}) {

}
