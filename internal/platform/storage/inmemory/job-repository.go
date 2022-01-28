package inmemory

import (
	"bunsan-ocr/internal/ocr"
	"context"
	"errors"
	"sync"
)

type JobRepositoryInMemory map[string]ocr.Job

var (
	once sync.Once
	jobInstance JobRepositoryInMemory
)

func NewJobRepository() JobRepositoryInMemory {
	once.Do(func() {
		jobInstance = make(JobRepositoryInMemory, 0)
	})
	return jobInstance
}

func (j JobRepositoryInMemory) Save(ctx context.Context, job ocr.Job) error {
	jobInstance[job.ID().String()] = job
	return nil
}

func (j JobRepositoryInMemory) FindById(ctx context.Context, id ocr.JobID) (ocr.Job, error) {
	if job, ok := jobInstance[id.String()]; ok {
		return job, nil
	}
	return ocr.Job{}, errors.New("job not exists")
}

func (j JobRepositoryInMemory) Update(ctx context.Context, job ocr.Job) error {
	jobInstance[job.ID().String()] = job
	return nil
}