package main

import (
	"fmt"

	_ "batch/main/database"
	"batch/main/jobs"
)

func main() {
	finished := make(chan bool)
	jobs.Start(finished)
	<-finished

	fmt.Print("Execution ended")
}
