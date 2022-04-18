package svc

import (
	"fmt"
	"time"

	"github.com/ekinbulut/go-schedular/svc/models"
)

type Schedular struct {
	Jobs   []*models.Job
	ticker *time.Ticker
}

func NewSchedular(second int) *Schedular {
	return &Schedular{
		Jobs:   make([]*models.Job, 0),
		ticker: time.NewTicker(time.Duration(second) * time.Second),
	}
}

// AddJob adds a job to the schedular
func (s *Schedular) AddJob(job *models.Job) {
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
