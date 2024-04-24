package actor

type BaseActor struct {
	// Add fields necessary for actor state
}

func (a *BaseActor) Receive(ctx Context) error {
	// Handle messages
	return nil
}

func (a *BaseActor) PreStart(ctx Context) {
	// Initialization logic
}

func (a *BaseActor) PostStop(ctx Context) {
	// Cleanup logic
}
