package logger

import "fmt"

func Info(s string) {
	fmt.Printf("[INFO] %s\n", s)
}
