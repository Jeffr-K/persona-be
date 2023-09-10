package event

type CommandEvent struct{}
type QueryEvent struct{}

type Event[T CommandEvent | QueryEvent] struct{}

// Process
// / The function returns registered event list with command object.
func (e Event[T]) Process(command interface{}) error {
	return nil
}

// Apply
// / The function updates the aggregate after being received event
func (e Event[T]) Apply(event Event[T]) error {
	return nil
}
