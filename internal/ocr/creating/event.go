package creating

import (
	"bunsan-ocr/internal/ocr"
	"bunsan-ocr/internal/ocr/processing"
	"bunsan-ocr/kit/bus/event"
	"context"
	"errors"
)

const ProcessJobOnJobCreatedType event.HandlerType = "event.handler.processJobOnJobCreated"

type ProcessJobOnJobCreated struct {
	processingService processing.JobProcessorService
}

func NewProcessJobOnJobCreated(processingService processing.JobProcessorService) ProcessJobOnJobCreated {
	return ProcessJobOnJobCreated{
		processingService: processingService,
	}
}

func (p ProcessJobOnJobCreated) Handle(_ context.Context, evt event.Event) error {
	jobCreatedEvt, ok := evt.(ocr.JobCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return p.processingService.Process(jobCreatedEvt.ID(), jobCreatedEvt.FileInputPath(), jobCreatedEvt.FileInputContentType(), jobCreatedEvt.Status())
}

func (p ProcessJobOnJobCreated) Type() event.HandlerType {
	return ProcessJobOnJobCreatedType
}