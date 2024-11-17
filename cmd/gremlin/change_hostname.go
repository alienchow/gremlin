package main

import (
	"log"
	"math/rand"
	"os"
	"syscall"
	"time"
)

const (
	trollHostname = "lOsEr"
)

func changeHostname(continueIfFedCh chan struct{}) {
	<-time.After(time.Duration(rand.Intn(9)+2) * time.Second)
	<-continueIfFedCh

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	if hostname == trollHostname {
		log.Println("hmmmm... The host is already a loser, no fun.")
		return
	}

	if err := syscall.Sethostname([]byte(trollHostname)); err != nil {
		panic(err)
	}
	log.Println("wHo'S a LoSeR!?")
}
