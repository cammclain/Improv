package actor

// Message represents a generic message structure in the actor system.
type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// Actor defines the standard interface for all actors in the system.
type Actor interface {
	Receive(ctx Context) error
	PreStart(ctx Context)
	PostStop(ctx Context)
}

// Context provides information about the environment an actor is running in,
// including message sending capabilities.
type Context interface {
	SendMessage(msg interface{})
	SendTo(address string, msg interface{})
	Self() *ActorRef
	ActorSystem() *System
}
