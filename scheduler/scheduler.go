package scheduler

import (
	"Digobo/log"
	"sort"
	"sync"
	"time"
)

var instance *scheduler

type scheduler struct {
	currentJobs   chan Job
	waitGroup     sync.WaitGroup
	thresholdJobs struct {
		jobs  []Job
		mutex sync.RWMutex
	}
}

func GetScheduler() *scheduler {
	if instance == nil {
		instance = &scheduler{
			currentJobs: make(chan Job),
			waitGroup: sync.WaitGroup{},
			thresholdJobs: struct {
				jobs  []Job
				mutex sync.RWMutex
			}{},
		}

		instance.waitGroup.Add(1)
	}

	return instance
}

// Insert job to be executed instandly
func (this *scheduler) AddScheduledJob(job Job) error {
	this.currentJobs <- job

	this.waitGroup.Add(1)

	return nil
}

// Insert job to be executed at the given time
func (this *scheduler) AddThresholdJob(job Job) {
	this.thresholdJobs.mutex.Lock()

	this.thresholdJobs.jobs = append(this.thresholdJobs.jobs, job)

	// TODO sorted insert
	sort.Slice(this.thresholdJobs.jobs, func(i, j int) bool {
		return this.thresholdJobs.jobs[i].ExecutionTime.Before(this.thresholdJobs.jobs[j].ExecutionTime)
	})

	this.thresholdJobs.mutex.Unlock()
}

func init() {
	run := GetScheduler()
	run.start()
}

func (this *scheduler) start() {
	go this.run()

	go func() {
		for {
			duration := this.thresholdJobsToChannel()
			if duration.Seconds() <= 2 {
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (this *scheduler) run() {
	for {
		select {
		case job := <-this.currentJobs:
			err := job.ExecutionFunc.Execute(job.Data)
			if err != nil {
				log.Warning.Println("Could not execute scheduled job", err)
			}
		}
	}
}

func (this *scheduler) thresholdJobsToChannel() time.Duration {
	start := time.Now()
	var deleted = false
	var maxIndex = 0

	this.thresholdJobs.mutex.RLock()
	for i, job := range this.thresholdJobs.jobs {
		if start.After(job.ExecutionTime) || start.Equal(job.ExecutionTime) {
			this.currentJobs <- job

			maxIndex = i
			deleted = true
		} else {
			break
		}
	}
	this.thresholdJobs.mutex.RUnlock()

	if deleted {
		this.thresholdJobs.mutex.RLock()
		this.thresholdJobs.jobs = this.thresholdJobs.jobs[maxIndex+1:]
		this.thresholdJobs.mutex.RUnlock()
	}

	return time.Now().Sub(start)
}
