package logging

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	info   *log.Logger
	errors *log.Logger
	debug  *log.Logger
	waring *log.Logger
}

func InitLogger() Logger {
	flags := log.Ldate | log.LstdFlags | log.Lshortfile
	return Logger{
		log.New(os.Stdout, "[INFO]", flags),
		log.New(os.Stdout, "[ERROR]", flags),
		log.New(os.Stdout, "[DEBUG]", flags),
		log.New(os.Stdout, "[WARING]", flags),
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Output(2, fmt.Sprintln(v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.errors.Output(2, fmt.Sprintln(v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Output(2, fmt.Sprintln(v...))
}

func (l *Logger) Waring(v ...interface{}) {
	l.waring.Output(2, fmt.Sprintln(v...))
}
