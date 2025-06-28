package utils

import (
	"fmt"
	"time"
)

type Logs struct {
	WithTime bool
}

func logFormat(typeLog string, text string) {
	fmt.Printf("[%s] %s", typeLog, text)
}
func logFormatWithTime(typeLog string, text string, time string) {
	fmt.Printf("[%s][%s] %s", time, typeLog, text)

}

func getTime(l *Logs) string {
	if l.WithTime {
		time := time.Now().Format("2025-06-14 23:23:02")
		return time
	}
	return ""
}

func (l *Logs) Info(text string) {
	time := getTime(l)
	if time == "" {
		logFormat("INFO", text)
		return
	}
	logFormatWithTime("INFO", text, time)
}

func (l *Logs) Error(text string) {
	time := getTime(l)
	if time == "" {
		logFormat("ERROR", text)
		return
	}
	logFormatWithTime("ERROR", text, time)
}

func (l *Logs) Warn(text string) {
	time := getTime(l)
	if time == "" {
		logFormat("WARN", text)
		return
	}
	logFormatWithTime("WARN", text, time)
}

func (l *Logs) Debug(text string) {
	time := getTime(l)
	if time == "" {
		logFormat("DEBUG", text)
		return
	}
	logFormatWithTime("DEBUG", text, time)
}
