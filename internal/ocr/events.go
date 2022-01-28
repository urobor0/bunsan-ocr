package ocr

import "bunsan-ocr/kit/bus/event"

const JobCreatedEventType event.Type = "events.job.created"

type JobCreatedEvent struct {
	event.BaseEvent
	id string
	fileInputPath string
	fileInputExtension string
	status int
}

func NewJobCreatedEvent(id, fileInputPath, fileInputExtension string, status int) JobCreatedEvent {
	return JobCreatedEvent{
		BaseEvent:          event.NewBaseEvent(id),
		id:                 id,
		fileInputPath:      fileInputPath,
		fileInputExtension: fileInputExtension,
		status:             status,
	}
}

func (j JobCreatedEvent) Type() event.Type {
	return JobCreatedEventType
}