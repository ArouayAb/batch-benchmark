package main

import (
	"fmt"

	"batch/main/database"
	"batch/main/jobs"
)

func main() {
	database.Init()
	finished := make(chan bool)

	jobs.Start(finished)

	<-finished
	fmt.Print("Execution ended")

}
