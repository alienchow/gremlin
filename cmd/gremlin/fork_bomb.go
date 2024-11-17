package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	forkBombCountDown = 10
)

func forkBomb(continueIfFedCh chan struct{}) {
	<-time.After(time.Duration(rand.Intn(4)+2) * time.Second)
	<-continueIfFedCh

	log.Printf("Fine, since it has come to this.")
	<-time.After(time.Duration(2) * time.Second)
	log.Printf("Have you heard of :(){:|:&};: ?")
	<-time.After(time.Duration(2) * time.Second)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	currentCounter := forkBombCountDown
	for range ticker.C {
		if currentCounter == 0 {
			break
		}
		log.Println(currentCounter)
		currentCounter -= 1
	}
	bombsAway()
}

func bombsAway() {
	for {
		go bombsAway()
	}
}
