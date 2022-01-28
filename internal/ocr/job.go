package ocr

import (
	"bunsan-ocr/kit/bus/event"
	"bunsan-ocr/kit/identifier"
	"context"
	"errors"
	"fmt"
)

type JobStatus int

const (
	Created   JobStatus = 1
	Processed JobStatus = 2
	Finished  JobStatus = 3
)

// ToPrimitive type converts the JobStatus into int.
func (j JobStatus) ToPrimitive() int {
	return int(j)
}

// JobID represents the Job unique identifier.
type JobID struct {
	value string
}

var ErrInvalidJobID = errors.New("invalid job id")

// NewJobID instantiate the Vo for JobID.
func NewJobID(value string) (JobID, error) {
	id, err := identifier.Parse(value)
	if err != nil {
		return JobID{}, fmt.Errorf("%w: %s", ErrInvalidJobID, value)
	}
	return JobID{
		value: id,
	}, nil
}

// String type converts the JobID into string.
func (j JobID) String() string {
	return j.value
}

// Job is the data structure that represent a Job.
type Job struct {
	id                 JobID
	fileInputPath      string
	fileInputExtension string
	status             JobStatus

	events []event.Event
}

// JobRepository defines the expected behaviour a job storage.
type JobRepository interface {
	Save(ctx context.Context, job Job) error
	FindById(ctx context.Context, id JobID) (Job, error)
	Update(ctx context.Context, id Job) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../platform/storage/storagemocks --name=JobRepository

// NewJob creates a new job.
func NewJob(id, fileInputPath, fileInputExtension string) (Job, error) {
	idVO, err := NewJobID(id)
	if err != nil {
		return Job{}, err
	}

	job := Job{
		id:                 idVO,
		fileInputPath:      fileInputPath,
		fileInputExtension: fileInputExtension,
		status:             Created,
	}

	job.Record(NewJobCreatedEvent(id, fileInputPath, fileInputExtension, job.status.ToPrimitive()))

	return job, nil
}

// UnmarshalJob creates a new job from primitive values without record event.
func UnmarshalJob(id, fileInputPath, fileInputExtension string, status int) (Job, error) {
	idVO, err := NewJobID(id)
	if err != nil {
		return Job{}, err
	}

	job := Job{
		id:                 idVO,
		fileInputPath:      fileInputPath,
		fileInputExtension: fileInputExtension,
		status:             JobStatus(status),
	}

	return job, nil
}

// ID returns the job unique identifier.
func (j Job) ID() JobID {
	return j.id
}

// FileInputPath returns the file path to be processed.
func (j Job) FileInputPath() string {
	return j.fileInputPath
}

// FileInputExtension returns the file extension.
func (j Job) FileInputExtension() string {
	return j.fileInputExtension
}

// Status returns the job current status.
func (j Job) Status() JobStatus {
	return j.status
}

// Processed returns a processed job.
func (j Job) Processed() Job {
	j.status = Processed
	return j
}

// Finished returns a processed job.
func (j Job) Finished() Job {
	j.status = Finished
	return j
}

// Record records a new domain event.
func (j *Job) Record(evt event.Event) {
	j.events = append(j.events, evt)
}

// PullEvents returns all the recorded domain events.
func (j *Job) PullEvents() []event.Event {
	evt := j.events
	j.events = []event.Event{}

	return evt
}
