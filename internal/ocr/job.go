package ocr

import (
	"bunsan-ocr/kit/bus/event"
	"bunsan-ocr/kit/identifier"
	"context"
	"errors"
	"fmt"
)

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
	id JobID
	fileInputPath string
	fileInputExtension string
	status int

	events []event.Event
}

// JobRepository defines the expected behaviour a job storage.
type JobRepository interface {
	Save(ctx context.Context, job Job) error
	FindById(ctx context.Context, id JobID) (Job, error)
	Update(ctx context.Context, id Job) error
}


//go:generate mockery --case=snake --outpkg=storagempcks --output=../../platform/stprage/storagemocks --name=JobRepository
