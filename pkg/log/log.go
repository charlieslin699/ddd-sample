package log

import (
	"fmt"
)

func Debug(o ...any) {
	nowTime := nowTime()
	o = append([]any{nowTime}, o...)
	fmt.Println(o...)
}

func Info(o ...any) {
	nowTime := nowTime()
	o = append([]any{nowTime}, o...)
	fmt.Println(o...)
}

func Warn(o ...any) {
	nowTime := nowTime()
	o = append([]any{nowTime}, o...)
	fmt.Println(o...)
}

func Error(o ...any) {
	nowTime := nowTime()
	o = append([]any{nowTime}, o...)
	fmt.Println(o...)
}
