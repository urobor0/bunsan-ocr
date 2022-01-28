package creating

import (
	"bunsan-ocr/internal/ocr"
	"bunsan-ocr/kit/bus/event"
	"context"
)

type JobService struct {
	jobRepository ocr.JobRepository
	eventBus event.Bus
}

func NewJobService(repository ocr.JobRepository, bus event.Bus) JobService {
	return JobService{
		jobRepository: repository,
		eventBus:      bus,
	}
}

func (j JobService) CreateJob(ctx context.Context, id, fileInputPath, fileInputContentType string) error {
	job, err := ocr.NewJob(id, fileInputPath, fileInputContentType)
	if err != nil {
		return err
	}

	if err := j.jobRepository.Save(ctx, job); err != nil {
		return err
	}

	return j.eventBus.Publish(ctx, job.PullEvents())
}