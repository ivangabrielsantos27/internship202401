package main

import (
	"flag"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/alphauslabs/bluectl/pkg/logger"
)

var (
	file = flag.String("file", "", "Sample file to process")
	cc   = flag.Bool("concurrent", false, "If true, run the concurrent function")
)

func sequential() {
	lines, err := os.ReadFile(*file)
	if err != nil {
		logger.Error(err)
		return
	}

	var wc int64
	for _, line := range lines {
		ss := strings.Split(string(line), " ")
		atomic.AddInt64(&wc, int64(len(ss)))
		for i := 0; i < 10000; i++ {
			dummy := i % 10
			_ = dummy
		}
	}

	logger.Infof("word count: %v", atomic.LoadInt64(&wc))
}

func concurrent() {
	lines, err := os.ReadFile(*file)
	if err != nil {
		logger.Error(err)
		return
	}

	numCpu := runtime.NumCPU()
	wq := make(chan string, numCpu)
	wqdone := make(chan bool, numCpu)

	var wc int64
	processor := func(id int) {
		defer func() {
			wqdone <- true
		}()

		for j := range wq {
			ss := strings.Split(j, " ")
			atomic.AddInt64(&wc, int64(len(ss)))
			for i := 0; i < 10000; i++ {
				dummy := i % 10
				_ = dummy
			}
		}
	}

	for i := 0; i < numCpu; i++ {
		go processor(i)
	}

	for _, line := range lines {
		wq <- string(line)
	}

	logger.Infof("word count: %v", atomic.LoadInt64(&wc))

	close(wq)
	for i := 0; i < numCpu; i++ {
		<-wqdone
	}
}

func main() {
	flag.Parse()
	if *file == "" {
		logger.Errorf("cannot find file [%v]", *file)
		return
	}

	// Log how long did it take.
	defer func(begin time.Time) {
		logger.Info("duration:", time.Since(begin))
	}(time.Now())

	if !*cc {
		sequential()
	} else {
		concurrent()
	}
}
