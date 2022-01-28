package creating

import (
	"bunsan-ocr/kit/bus/command"
	"context"
	"errors"
)

const JobCommandType command.Type = "command.creating.job"

// JobCommand is the command dispatched to create a new job.
type JobCommand struct {
	id   string
	fileInputPath string
	fileInputContentType string
}

// NewJobCommand creates a new JobCommand.
func NewJobCommand(id, fileInputPath, fileInputContentType string) JobCommand {
	return JobCommand{
		id:                   id,
		fileInputPath:        fileInputPath,
		fileInputContentType: fileInputContentType,
	}
}

// Type returns a command.Type.
func (j JobCommand) Type() command.Type {
	return JobCommandType
}

// JobCommandHandler is the command handler responsible for creating job.
type JobCommandHandler struct {
	service JobService
}

// NewJobCommandHandler initializes a new JobCommandHandler.
func NewJobCommandHandler(service JobService) JobCommandHandler {
	return JobCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (j JobCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createJobCmd, ok := cmd.(JobCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return j.service.CreateJob(
		ctx,
		createJobCmd.id,
		createJobCmd.fileInputPath,
		createJobCmd.fileInputContentType,
	)
}

// SubscribedTo returns the command to which the handler is subscribed.
func (j JobCommandHandler) SubscribedTo() command.Command {
	return JobCommand{}
}