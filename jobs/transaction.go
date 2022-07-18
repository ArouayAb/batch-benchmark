package jobs

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/go-co-op/gocron"
)

var taskScheduler = gocron.NewScheduler(time.UTC)

func Start(finished chan bool) {
	taskScheduler.Every(5).Seconds().Do(func() {
		log.Println("job started")
		cmd := exec.Command("python", "jobs/scripts/transaction-file.py")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Println(cmd.Run())
	})

	// taskScheduler.Every(5).Seconds().Do(func() {
	// 	log.Println("job started")
	// 	cmd := exec.Command("python", "jobs/scripts/transaction-db.py")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Stderr = os.Stderr
	// 	log.Println(cmd.Run())
	// })

	// taskScheduler.Every(5).Seconds().Do(func() {
	// 	log.Println("job started")
	// 	cmd := exec.Command("go", "run", "jobs/scripts/transaction-db.go")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Stderr = os.Stderr
	// 	log.Println(cmd.Run())
	// })

	// taskScheduler.Every(5).Seconds().Do(func() {
	// 	log.Println("job started")
	// 	cmd := exec.Command("go", "run", "jobs/scripts/transaction-file.go")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Stderr = os.Stderr
	// 	log.Println(cmd.Run())
	// })

	taskScheduler.StartAsync()
}
