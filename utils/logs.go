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

func (l *Logs) NewLog(text string,typeLOG string){
	time := getTime(l)
	if time == "" {
		logFormat(typeLOG, text)
		return
	}
	logFormatWithTime(typeLOG, text, time)

}

func getTime(l *Logs) string {
	if l.WithTime {
		time := time.Now().Format("2025-06-14 23:23:02")
		return time
	}
	return ""
}

func (l *Logs) Info(text string) {
	l.NewLog(text,"INFO")
}

func (l *Logs) Error(text string) {
	l.NewLog(text,"ERROR")
}

func (l *Logs) Warn(text string) {
	l.NewLog(text,"WARN")
}

func (l *Logs) Debug(text string) {
	l.NewLog(text,"DEBUG")
}
