package main

import (
	"flag"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/alphauslabs/bluectl/pkg/logger"
)

var (
	file = flag.String("file", "testcur.csv", "Sample file to process")
)

func concurrent() {
	lines, err := os.ReadFile(*file)
	if err != nil {
		logger.Error(err)
		return
	}

	numCpu := runtime.NumCPU()
	wq := make(chan string, numCpu)
	var wg sync.WaitGroup
	var wc int64

	processor := func() {
		defer wg.Done()

		for j := range wq {
			ss := strings.Split(j, " ")
			atomic.AddInt64(&wc, int64(len(ss)))
		}
	}

	for i := 0; i < numCpu; i++ {
		wg.Add(1)
		go processor()
	}

	for _, line := range lines {
		wq <- string(line)
	}

	close(wq)

	wg.Wait()

	logger.Infof("word count: %v", atomic.LoadInt64(&wc))
}

func main() {
	flag.Parse()
	if *file == "" {
		logger.Errorf("cannot find file [%v]", *file)
		return
	}

	defer func(begin time.Time) {
		logger.Info("duration:", time.Since(begin))
	}(time.Now())

	concurrent()
}
