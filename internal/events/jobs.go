// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/robfig/cron/v3"
)

var (
	_defaultJobManager JobManager = (*jobManager)(nil)
)

type (
	EntryID = cron.EntryID
)

// JobFn job help function that implement cron.Job interface
type JobFn func()

func (fn JobFn) Run() {
	fn()
}

// Job job interface
type Job interface {
	cron.Schedule
	cron.Job
}

type simpleJob struct {
	cron.Schedule
	cron.Job
}

// JobManager job manger interface
type JobManager interface {
	Start()
	Stop()
	Remove(id EntryID)
	Schedule(Job) EntryID
}

type jobManager struct {
	m *cron.Cron
}

func (j *jobManager) Start() {
	j.m.Start()
}

func (j *jobManager) Stop() {
	j.m.Stop()
}

// Remove an entry from being run in the future.
func (j *jobManager) Remove(id EntryID) {
	j.m.Remove(id)
}

// Schedule adds a Job to the Cron to be run on the given schedule.
// The job is wrapped with the configured Chain.
func (j *jobManager) Schedule(job Job) EntryID {
	return j.m.Schedule(job, job)
}

func initJobManager() {
	_defaultJobManager = &jobManager{
		m: cron.New(),
	}
	StartJobManager()
}

func StartJobManager() {
	_defaultJobManager.Start()
}

func StopJobManager() {
	_defaultJobManager.Stop()
}

// NewJob create new Job instance
func NewJob(s cron.Schedule, fn JobFn) Job {
	return &simpleJob{
		Schedule: s,
		Job:      fn,
	}
}

// RemoveJob an entry from being run in the future.
func RemoveJob(id EntryID) {
	_defaultJobManager.Remove(id)
}

// ScheduleJob adds a Job to the Cron to be run on the given schedule.
// The job is wrapped with the configured Chain.
func ScheduleJob(job Job) EntryID {
	return _defaultJobManager.Schedule(job)
}

// Schedule adds a Job to the Cron to be run on the given schedule.
// The job is wrapped with the configured Chain.
func Schedule(s cron.Schedule, fn JobFn) EntryID {
	job := &simpleJob{
		Schedule: s,
		Job:      fn,
	}
	return _defaultJobManager.Schedule(job)
}
