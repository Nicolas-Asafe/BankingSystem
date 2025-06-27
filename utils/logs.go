package utils

import (
	"fmt"
	"time"
)

type Logs struct {
	WithTime bool
}

func LogFormat(typeLog string, text string) {
	fmt.Printf("[%s] %s", typeLog, text)
}
func LogFormatWithTime(typeLog string, text string, time string) {
	fmt.Printf("[%s][%s] %s",time, typeLog, text)

}

func GetTime(l *Logs) string {
	if l.WithTime {
		time := time.Now().Format("2025-06-14 23:23:02")
		return time
	}
	return ""
}

func (l *Logs) Info(text string) {
	time := GetTime(l)
	if time == "" {
		LogFormat("INFO", text)
		return
	}
	LogFormatWithTime("INFO",text,time)
}

func (l *Logs) Error(text string) {
	time := GetTime(l)
	if time == "" {
		LogFormat("ERROR", text)
		return
	}
	LogFormatWithTime("ERROR",text,time)
}

func (l *Logs) Warn(text string) {
	time := GetTime(l)
	if time == "" {
		LogFormat("WARN", text)
		return
	}
	LogFormatWithTime("WARN",text,time)
}

func (l *Logs) Debug(text string) {
	time := GetTime(l)
	if time == "" {
		LogFormat("DEBUG", text)
		return
	}
	LogFormatWithTime("DEBUG",text,time)
}

