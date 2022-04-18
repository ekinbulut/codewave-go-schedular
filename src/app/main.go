package main

import (
	"fmt"

	"github.com/ekinbulut/go-schedular/svc"
	"github.com/ekinbulut/go-schedular/svc/models"
)

type App struct {
	Name    string
	Version string
}

func (a *App) Run() {

	// create a schedular
	schedular := svc.NewSchedular()
	schedular.AddJob(&models.Job{
		Name: "Job 1",
	})

	schedular.AddJob(&models.Job{
		Name: "Job 2",
	})

	schedular.Run()

	fmt.Printf("%s %s\n", a.Name, a.Version)
}

func main() {
	app := App{
		Name:    "Go Schedular App",
		Version: "1.0.0",
	}

	app.Run()
}
