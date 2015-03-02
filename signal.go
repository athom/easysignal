package easysignal

import (
	"os"
	"os/signal"
	"syscall"
)

type quit <-chan bool

func signals() <-chan bool {
	quit := make(chan bool)
	go func() {
		signals := make(chan os.Signal)
		defer close(signals)

		signal.Notify(signals, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt, os.Kill)
		defer signal.Stop(signals)

		<-signals
		quit <- true
	}()
	return quit
}

func OnProcessKilled(f func()) {
	go func() {
		if <-signals() {
			f()
		}
	}()
}
