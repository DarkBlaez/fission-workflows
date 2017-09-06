package fes

// Fast, minimal Event Sourcing
type Aggregator interface {
	// Entity-specific
	ApplyEvent(event *Event) error

	// The aggregate provides type information about the entity, such as the aggregate id and the aggregate type.
	//
	// Implemented by AggregatorMixin
	Aggregate() Aggregate

	// Update entity to the provided target state
	//
	// This is implemented by AggregatorMixin, can be overridden for performance approach
	UpdateState(targetState Aggregator) error
}

type EventHandler interface {
	HandleEvent(event *Event) error
}

// EventStore is a persistent store for events
type EventStore interface {
	EventHandler
	Get(subject string) ([]*Event, error)
	List(matcher StringMatcher) ([]string, error)
}

// EventBus is the volatile reactive store that processes, stores events, and notifies subscribers
//type Dispatcher interface {
//	EventHandler
//	//pubsub.Publisher
//}

// Projector projects events into an entity
type Projector interface {
	Project(target Aggregator, events ...*Event) error
}

type CacheReader interface {
	Get(entity Aggregator) error
}

type CacheWriter interface {
	Put(entity Aggregator) error
}

type CacheReaderWriter interface {
	CacheReader
	CacheWriter
}

type StringMatcher interface {
	Match(target string) bool
}
