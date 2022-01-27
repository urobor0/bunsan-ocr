package query

import "context"

// Bus defines the expected behaviour from a query bus.
type Bus interface {
	// Ask is the method used to dispatch new queries.
	Ask(context.Context, Query) (Response, error)
	// Register is the method used to register a new query handler.
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=querymocks --output=querydmocks --name=Bus

// Type represents an application bus type.
type Type string

// Query represents an application bus.
type Query interface {
	Type() Type
}

type Response interface{}

// Handler defines the expected behaviour from a bus handler.
type Handler interface {
	Handle(context.Context, Query) (Response, error)
	SubscribedTo() Query
}
