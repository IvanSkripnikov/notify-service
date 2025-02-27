package events

type EventBus struct {
	Error chan error
}

func MakeBus() EventBus {
	return EventBus{
		Error: make(chan error),
	}
}
