package main

import (
	"fmt"

	schedular "github.com/ekinbulut/go-schedular/internals"
)

type App struct {
	Name    string
	Version string
}

func (a *App) Run() {

	fmt.Printf("%s %s\n", a.Name, a.Version)

	// create a schedular
	schedular := schedular.NewSchedular(5)

	// create a job
	job := NewJob("job1", func(c chan string) {
		c <- "executed"
	})

	schedular.AddJob(job)
	schedular.Run()
}

func main() {
	app := App{
		Name:    "Go Schedular App",
		Version: "1.0.0",
	}

	app.Run()
}

// create a job
func NewJob(name string, f func(c chan string)) *schedular.Job {
	m := &schedular.Job{
		Name: name,
	}
	m.Append(f)
	return m
}
