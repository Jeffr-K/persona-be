package bus

import "sync"

type CommandBus struct {
	mutex  sync.Mutex
	bus    map[string][]chan interface{}
	quit   chan struct{}
	closed bool
}

func New() *CommandBus {
	return &CommandBus{
		bus:  make(map[string][]chan interface{}),
		quit: make(chan struct{}),
	}
}

func (commandBus *CommandBus) Close() {
	commandBus.mutex.Lock()
	defer commandBus.mutex.Unlock()

	if commandBus.closed {
		return
	}

	for _, ch := range commandBus.bus {
		for _, bus := range ch {
			close(bus)
		}
	}
}

func (commandBus *CommandBus) Publish(topic string, event interface{}) {
	commandBus.mutex.Lock()
	defer commandBus.mutex.Unlock()

	if commandBus.closed {
		return
	}

	for _, ch := range commandBus.bus[topic] {
		ch <- event
	}
}

func (commandBus *CommandBus) Subscribe(topic string) <-chan interface{} {
	commandBus.mutex.Lock()
	defer commandBus.mutex.Unlock()

	if commandBus.closed {
		return nil
	}

	ch := make(chan interface{})
	commandBus.bus[topic] = append(commandBus.bus[topic], ch)
	return ch
}
