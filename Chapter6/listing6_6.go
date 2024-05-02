package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

type myLogger struct{}

func (l myLogger) Write(msg []byte) (int, error) {
	pc := make([]uintptr, 50)
	n := runtime.Callers(0, pc)
	pc = pc[:n]
	frames := runtime.CallersFrames(pc)
	caller := ""

	frameCount := 0
	for {
		frameCount++
		fr, hasMore := frames.Next()

		if hasMore {
			caller = fr.Function
		} else {
			break
		}
	}

	output := fmt.Sprintf("%s%s - %s%s (called from %s%s)", "\033[32m", time.Now().Format("2006/01/02 3:04:05 pm"), "\033[0m", strings.TrimSpace(string(msg)), "\033[35m", caller)
	return fmt.Println(output)
}

func main() {
	myLog := new(myLogger)
	log.SetFlags(0)
	log.SetOutput(myLog)
	go concurrentLog()
	for i := 0; i < 10; i++ {
		log.Println(fmt.Sprintf("run #%d", i+1))
		time.Sleep(1 * time.Second)
	}
}

func concurrentLog() {
	for i := 0; i < 2; i++ {
		log.Println(fmt.Sprintf("concurrent run #%d", i+1))
		time.Sleep(5 * time.Second)
	}
}
