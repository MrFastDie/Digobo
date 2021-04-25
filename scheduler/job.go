package scheduler

import "time"

type JobFunc interface {
	Execute(string) error
}

type Job struct {
	ExecutionTime time.Time
	ExecutionFunc JobFunc
	Data          string
}
