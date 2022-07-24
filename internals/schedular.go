package schedular

import (
	"fmt"
	"time"
)

type Schedular struct {
	Jobs   []*Job
	ticker *time.Ticker
}

func NewSchedular(second int) *Schedular {
	return &Schedular{
		Jobs:   make([]*Job, 0),
		ticker: time.NewTicker(time.Duration(second) * time.Second),
	}
}

// AddJob adds a job to the schedular
func (s *Schedular) AddJob(job *Job) {
	s.Jobs = append(s.Jobs, job)
}

// Stop stops the schedular
func (s *Schedular) Stop() {
	s.ticker.Stop()
}

// Run starts the schedular
func (s *Schedular) Run() {
	go s.BackgroundTicker()
	select {}
}

func (s *Schedular) BackgroundTicker() {

	for range s.ticker.C {

		fmt.Printf("%s\n", time.Now().String())

		for _, job := range s.Jobs {
			go job.Execute()
		}
	}
}

// return schedular size
func (s *Schedular) Size() int {
	return len(s.Jobs)
}
