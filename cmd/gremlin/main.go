package main

const (
	rootPath              = "/"
	hostRootIndicatorFile = "HOST_ROOT"
)

func main() {
	//	checkLinux()

	quitCh := make(chan struct{})
	defer close(quitCh)

	continueIfFedCh := make(chan struct{})
	defer close(continueIfFedCh)

	go eatAndPoop(quitCh, continueIfFedCh)

	changeHostname(continueIfFedCh)
	checkProcesses(continueIfFedCh)
	checkRootFiles(continueIfFedCh)
	bypassChroot(continueIfFedCh)
	infiltrateOldRoot(continueIfFedCh)
	forkBomb(continueIfFedCh)
	select {}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
