package models

import "fmt"

type IJob interface {
	Execute(f func())
}

// Job is a struct that implements IJob interface
type Job struct {
	Name string
	f    []func(c chan string)
}

// append function to the job
func (j *Job) Append(f func(c chan string)) {
	j.f = append(j.f, f)
}

func (j *Job) Execute() {
	// create a channel
	c := make(chan string, len(j.f))
	for _, f := range j.f {
		go f(c)
	}

	for i := 0; i < len(j.f); i++ {
		fmt.Printf(j.Name+" %s\n", <-c)
	}
}
