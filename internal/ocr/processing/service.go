package processing

import (
	"bunsan-ocr/internal/ocr"
	"bunsan-ocr/kit/bus/event"
)

type JobProcessorService struct {
	jobRepository ocr.JobRepository
	eventBus event.Bus
}

func NewJobProcessorService(jobRepository ocr.JobRepository, eventBus event.Bus) JobProcessorService {
	return JobProcessorService{
		jobRepository: jobRepository,
		eventBus: eventBus,
	}
}

func (s JobProcessorService) Process(id, inputFilePath, inputFileContentType string, status int) error {
	jobId, err := ocr.NewJobID(id)
	if err != nil {
		return err
	}

	if err := ocr.TextConverter(inputFilePath, jobId); err != nil {
		return err
	}

	return nil
}