package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @see https://www.youtube.com/watch?v=YF1qSfkDGAQ&list=PLDWZ5uzn69eyM81omhIZLzvRhTOXvpeX9&index=4
// @at Golang UK Conference 2017 | Ian Kent - Production-ready Go
// 本例主要演示  ctr+c os信号在go程序中的监听 截获
func main() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGFPE)

	// 退出main函数 时会被调用  如果不监听os信号那么ctr+c中断程序时
	// 此处不会被调用
	defer func() {
		log.Print("Done")
	}()

	// time.Sleep(time.Second * 10)
	timer := time.NewTimer(time.Second * 2)
	select {
	case <-timer.C:
	case <-sigCh:
	}
}
