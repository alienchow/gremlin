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

	checkRootFiles(continueIfFedCh)
	bypassChroot(continueIfFedCh)
	changeHostname(continueIfFedCh)
	// spyProcesses(continueIfFedCh)
	// checkOwnPID(continueIfFedCh)
	// mountDevices(continueIfFedCh)
	select {}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
