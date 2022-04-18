package svc

import (
	"fmt"

	"github.com/ekinbulut/go-schedular/svc/models"
)

type Schedular struct {
	Jobs []*models.Job
}

func NewSchedular() *Schedular {
	return &Schedular{}
}

// AddJob adds a job to the schedular
func (s *Schedular) AddJob(job *models.Job) {
	s.Jobs = append(s.Jobs, job)
}

// Stop stops the schedular
func (s *Schedular) Stop() {
}

// Run starts the schedular
func (s *Schedular) Run() {

	// create output channel for jobs execution in string
	out := make(chan string)
	for _, job := range s.Jobs {
		go job.Execute(func() {

			// send output to channel
			out <- fmt.Sprintf("%s executed", job.Name)

		})
	}

	// wait for all jobs to finish
	for i := 0; i < len(s.Jobs); i++ {
		fmt.Println(<-out)
	}
}
