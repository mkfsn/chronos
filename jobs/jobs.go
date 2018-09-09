package jobs

import (
	"log"
	"os"
	"strconv"
	"time"
)

// Job is a crontab like job
type Job interface {
	// Each job must have an unique name, and must be fit as an environment
	// variable
	Name() string

	// How long we execute this job
	Interval() time.Duration

	// How to run this job
	Run() error
}

// Run runs all defined jobs in separate go-routine
func Run() {
	// defined jobs
	jobs := []Job{
		newAburiyaJob(),
	}

	log.Printf("Run %d jobs\n", len(jobs))

	for _, job := range jobs {
		env := os.Getenv(job.Name())
		if env != "" {
			// Get job.Name() from environment variable and check if it's false or
			// not. If it has incorect form, treat it as false.
			b, err := strconv.ParseBool(env)

			if err != nil || !b {
				log.Println("Ignore:", job.Name(), "as", err, "or", b)
				continue
			}
		}
		log.Println("Start:", job.Name())
		go run(job)
	}
}

func run(job Job) {
	for {
		err := job.Run()
		if err != nil {
			log.Printf("[%s] Error: %s\n", job.Name(), err.Error())
		}
		time.Sleep(job.Interval())
	}
}
