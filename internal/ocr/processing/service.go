package processing

import (
	"bunsan-ocr/internal/ocr"
	"bunsan-ocr/kit/bus/event"
	"fmt"
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
	fmt.Println("Processing :D")
	return nil
}