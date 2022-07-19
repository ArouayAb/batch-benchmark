package main

import (
	"fmt"

	_ "batch/main/database"
	"batch/main/jobs"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	finished := make(chan bool)
	jobs.Start(finished)
	<-finished

	fmt.Print("Execution ended")
}
