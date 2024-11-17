package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	procDir = "/proc"
)

func checkProcesses(continueIfFedCh chan struct{}) {
	<-time.After(time.Duration(rand.Intn(4)+2) * time.Second)

	for range continueIfFedCh {
		_, err := os.Stat(procDir)
		if os.IsNotExist(err) {
			log.Println("hEy! WhErE aRe My PrOcEsSes!?")
			continue
		}

		files, err := os.ReadDir(procDir)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("kek checking processes")
		for _, file := range files {
			if pid, err := strconv.Atoi(file.Name()); err == nil {
				statFile := filepath.Join(procDir, file.Name(), "stat")
				data, err := os.ReadFile(statFile)
				if err != nil {
					continue
				}

				stats := strings.Fields(string(data))
				if pid == 1 && stats[1] == "(exe)" {
					log.Println("cAn'T fInD aNy HoSt PrOcEsSeS...")
					return
				}
				if len(stats) > 1 {
					log.Printf("PID: %s, Command: %s\n", stats[0], stats[1])
				}
			}
		}
	}
}
