package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mkfsn/chronos/jobs"
)

func main() {
	jobs.Run()

	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
