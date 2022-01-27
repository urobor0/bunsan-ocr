package inmemory

import (
	"bunsan-ocr/kit/bus/query"
	"context"
	"errors"
)

// QueryBus is an in-memory implementation of the query.Bus.
type QueryBus struct {
	handlers map[query.Type]query.Handler
}

// NewQueryBus initializes a new instance of QueryBus.
func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

// Ask implements the query.Bus interface.
func (b *QueryBus) Ask(ctx context.Context, query query.Query) (query.Response, error) {
	handler, ok := b.handlers[query.Type()]
	if !ok {
		return nil, errors.New("")
	}

	return handler.Handle(ctx, query)
}

// Register implements the query.Bus interface.
func (b *QueryBus) Register(queryType query.Type, handler query.Handler) {
	b.handlers[queryType] = handler
}
