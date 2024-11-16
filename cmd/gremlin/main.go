package main

import (
	"sync"
)

const (
	rootPath = "/"
)

func checkRootFiles(wg *sync.WaitGroup) {
	defer wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	quitCh := make(chan struct{})
	defer close(quitCh)

	go eatAndPoop(quitCh)

	wg.Add(1)
	go checkRootFiles(wg)
	wg.Wait()
	select {}

	// Check /etc/passwd

	// Check processes

	// Check own process ID
	//
	//
}
