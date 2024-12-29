package printer

import (
	"context"
	"fmt"
	"os"
	"time"
)

const (
	redColor    = "\033[31m"
	greenColor  = "\033[32m"
	yellowColor = "\033[33m"
	blueColor   = "\033[34m"
	resetColor  = "\033[0m"
)

func setColor(msg string, color string) string {
	return fmt.Sprintf("%s%s%s", color, msg, resetColor)
}

func Error(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", setColor(err.Error(), redColor))
}

func Fatal(err error) {
	Error(err)
	os.Exit(1)
}

func Print(msg string, args ...any) {
	fmt.Printf(msg, args...)
}

func Success(msg string, args ...any) {
	fmt.Print(setColor(fmt.Sprintf(msg, args...), greenColor))
}

func Load(msg string, args ...any) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool, 1)

	go func() {
		flag := "/"
		for {
			select {
			case <-ctx.Done():
				fmt.Print(setColor(fmt.Sprintf("%s [done]\n", fmt.Sprintf(msg, args...)), blueColor))
				done <- true
				return
			default:
				fmt.Print(setColor(fmt.Sprintf("%s%s\r", fmt.Sprintf(msg, args...), flag), blueColor))

				if flag == "/" {
					flag = "\\"
				} else {
					flag = "/"
				}

				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	return func() {
		cancel()
		<-done
	}
}
