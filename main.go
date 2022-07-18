package main

import (
	"batch/main/jobs"
	"fmt"

	_ "batch/main/database"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	finished := make(chan bool)
	jobs.Start(finished)
	<-finished

	fmt.Print("Execution ended")
}
