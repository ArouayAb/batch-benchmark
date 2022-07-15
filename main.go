package main

import (
	"batch/main/jobs"
	"fmt"
)

func main() {
	finished := make(chan bool)
	jobs.Start(finished)
	<-finished

	fmt.Print("Execution ended")
}
