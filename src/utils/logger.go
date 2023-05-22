package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	LogFile io.Writer
	Time    *time.Time
	FLerr   error
}

func generateLogFileName() string {
	now := time.Now().UTC()
	return fmt.Sprintf("./log/%s.log", now.Format("2006-01-02"))
}

func (l *Logger) NewLogPath() error {
	file, err := os.OpenFile(generateLogFileName(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.FLerr = err
		return err
	}
	l.LogFile = io.MultiWriter(file, os.Stdout)
	return err
}

func (l *Logger) LogMaker() {
	l.Info = log.New(l.LogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Warning = log.New(l.LogFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Error = log.New(l.LogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewLogger() *Logger {
	logger := &Logger{}
	err := logger.NewLogPath()
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	logger.LogMaker()
	return logger
}
