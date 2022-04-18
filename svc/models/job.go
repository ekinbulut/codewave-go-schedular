package models

type IJob interface {
	Execute(f func())
}

// Job is a struct that implements IJob interface
type Job struct {
	Name string
}

func (j *Job) Execute(f func()) {
	f()
}
