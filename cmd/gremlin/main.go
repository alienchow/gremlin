package main

const (
	rootPath              = "/"
	maxFoodPoop           = 1000
	hostRootIndicatorFile = "HOST_ROOT"
)

func main() {
	quitCh := make(chan struct{})
	defer close(quitCh)

	continueIfFedCh := make(chan struct{}, maxFoodPoop)
	defer close(continueIfFedCh)

	go eatAndPoop(quitCh, continueIfFedCh)

	checkRootFiles(continueIfFedCh)
	// Check /etc/passwd
	// Check processes
	// Check own process ID
	//
	select {}
}
