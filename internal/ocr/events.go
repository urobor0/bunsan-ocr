package ocr

import "bunsan-ocr/kit/bus/event"

const JobCreatedEventType event.Type = "events.job.created"

type JobCreatedEvent struct {
	event.BaseEvent
	id string
	fileInputPath        string
	fileInputContentType string
	status               int
}

func NewJobCreatedEvent(id, fileInputPath, fileInputExtension string, status int) JobCreatedEvent {
	return JobCreatedEvent{
		BaseEvent:            event.NewBaseEvent(id),
		id:                   id,
		fileInputPath:        fileInputPath,
		fileInputContentType: fileInputExtension,
		status:               status,
	}
}

func (j JobCreatedEvent) FileInputPath() string {
	return j.fileInputPath
}

func (j JobCreatedEvent) FileInputContentType() string {
	return j.fileInputContentType
}

func (j JobCreatedEvent) Status() int {
	return j.status
}

func (j JobCreatedEvent) Type() event.Type {
	return JobCreatedEventType
}