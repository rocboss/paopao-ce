// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/robfig/cron/v3"
	"github.com/rocboss/paopao-ce/pkg/types"
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

type emptyJobManager types.Empty

type simpleJobManager struct {
	m *cron.Cron
}

func (emptyJobManager) Start() {
	// nothing
}

func (emptyJobManager) Stop() {
	// nothing
}

func (emptyJobManager) Remove(id EntryID) {
	// nothing
}

func (emptyJobManager) Schedule(job Job) EntryID {
	return 0
}

func (j *simpleJobManager) Start() {
	j.m.Start()
}

func (j *simpleJobManager) Stop() {
	j.m.Stop()
}

// Remove an entry from being run in the future.
func (j *simpleJobManager) Remove(id EntryID) {
	j.m.Remove(id)
}

// Schedule adds a Job to the Cron to be run on the given schedule.
// The job is wrapped with the configured Chain.
func (j *simpleJobManager) Schedule(job Job) EntryID {
	return j.m.Schedule(job, job)
}

func NewJobManager() JobManager {
	return &simpleJobManager{
		m: cron.New(),
	}
}
