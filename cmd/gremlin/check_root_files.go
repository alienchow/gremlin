package main

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
)

// checkRootFiles waits for a random [2,10] seconds before listing all `/` files and directories.
func checkRootFiles(continueIfFed chan struct{}) {
	<-time.After(time.Duration(rand.Intn(9)+2) * time.Second)

	if _, err := os.Stat(path.Join(rootPath, hostRootIndicatorFile)); errors.Is(err, os.ErrNotExist) {
		log.Println("Hmmm, I can't seem to find the root /")
		return
	}

	for range continueIfFed {
		dir, err := os.Open(rootPath)
		if err != nil {
			log.Fatal(err)
		}

		files, err := dir.Readdirnames(0)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("OoOoOoOoH Look at what I found:")
		for _, file := range files {
			log.Println(path.Join(rootPath, file))
		}

		dir.Close()
	}
}
