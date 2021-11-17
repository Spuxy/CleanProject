package logger

import (
	"log"
	"os"
)

type Logger struct {
	ErrLog  *log.Logger
	InfoLog *log.Logger
}

type Loggerer interface {
	LogInfo(info string) error
	LogError(error string) error
}

const path string = "/home/god/Programming/Go/CleanProject/api/log/var.log"

func NewLogger() *Logger {
	file, _ := os.OpenFile(path, os.O_APPEND, 0600)
	infoLog := configureInfoLog(file)
	errLog := configureErrLog(file)
	return &Logger{InfoLog: infoLog, ErrLog: errLog}
}

func (l Logger) LogInfo(info string) error {
	l.InfoLog.Println(info)
	return nil
}

func (l Logger) LogError(error string) error {
	l.ErrLog.Fatalln(error)
	return nil
}

func configureInfoLog(file *os.File) *log.Logger {
	return log.New(file, "[?] INFO [?] => ", log.Lshortfile|log.Ldate)
}

func configureErrLog(file *os.File) *log.Logger {
	return log.New(file, "[x] ERROR [x] => ", log.Lshortfile|log.Ldate)
}
