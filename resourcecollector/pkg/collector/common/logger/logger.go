package logger

import (
	"fmt"
	"log"
)

func ReadInfof(format string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("\033[00;36m%s\033[0m\n", format), a...)
}

func Infof(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

func Info(str string) {
	fmt.Println(str + "\n")
}

func Fatalf(format string, a ...interface{}) {
	log.Fatalf(format+"\n", a...)
}

func Errorf(format string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("\033[00;36m%s\033[0m\n", format), a...)
}

func Error(format string) {
	Errorf(format)
}

func Warnf(format string, a ...interface{}) {
	Errorf(format, a...)
}