package main

import (
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	foodPath = "/food"
	poopPath = "/poop"

	foodPrefix = "food_"
	poopPrefix = "poop_"
)

// eat tries to find food to consume. Returns true if fed.
func eat() bool {
	foodPathInfo, err := os.Stat(foodPath)
	if err != nil {
		log.Printf("wHeRe'S mY fOod DiR %s!?\n", foodPath)
		return false
	}
	if !foodPathInfo.IsDir() {
		log.Printf("I cAn'T oPeN mY fOoD dIr: %s!\n", foodPath)
		return false
	}

	food, err := os.ReadDir(foodPath)
	if err != nil {
		log.Printf("I cAn'T aCcEsS mY fOoD in %s!\n", foodPath)
		return false
	}

	foundFood := false
	for _, f := range food {
		if f.IsDir() {
			continue
		}
		if strings.HasPrefix(f.Name(), foodPrefix) {
			foundFood = true
			fullFoodPath := path.Join(foodPath, f.Name())
			err := os.Remove(fullFoodPath)
			if err != nil {
				log.Printf("I cAn'T eAt %s!\n", fullFoodPath)
			} else {
				log.Printf("Om nomNOM nOMnoMNOm. Ate %s\n", fullFoodPath)
			}
			break
		}
	}

	if !foundFood {
		log.Printf("ThErE iS nO fOoD iN tHe FoOd DiR %s!\n", foodPath)
		return false
	}
	return true
}

// poop tries to take a dump. Returns true if successfully defecated.
func poop() bool {
	poopPathInfo, err := os.Stat(poopPath)
	if err != nil {
		log.Printf("wHeRe'S mY pOoP DiR %s!?\n", poopPath)
		return false
	}

	if !poopPathInfo.IsDir() {
		log.Printf("I cAn'T oPeN mY pOoP dIr: %s!\n", poopPath)
		return false
	}

	poopName := poopPrefix + uuid.New().String()
	fullPoopPath := path.Join(poopPath, poopName)
	poop, err := os.Create(fullPoopPath)
	if err != nil {
		log.Printf("I cAn'T pOoP %s!\n", fullPoopPath)
		return false
	}
	defer poop.Close()

	log.Printf("pOoPeD oUt %s!\n", fullPoopPath)

	return true
}

func eatAndPoop(quitCh chan struct{}, continueIfFed chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-quitCh:
			return
		case <-ticker.C:
			if eat() && poop() {
				continueIfFed <- struct{}{}
			}
		}
	}
}
