# Easy Signal


A handdy os.signal wrapper to make the usage simple and easy based on real usage.


## Useage


```go
import (
	    "github.com/athom/easysignal"
)

// Some normal code

// signal captured and cleanup code here
easysignal.OnProcessKilled(func() {
		_, _ = conn.Write([]byte("QUIT\n"))
		os.Exit(1)
})
```

No need to write tons of code like these any more

```go
	go func() {
		signals := make(chan os.Signal)
		defer close(signals)

		signal.Notify(signals, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt)
		defer signalStop(signals)

		<-signals
		// Your cleanup code here
		_, _ = conn.Write([]byte("QUIT\n"))
		os.Exit(1)
	}()
```