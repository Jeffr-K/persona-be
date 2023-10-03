package personaq

import "sync"

type PersonaQ interface {
	Subscribe(topic string) <-chan interface{}
	Publish(topic string, data interface{})
	Close()
}

type personaQ struct {
	mu     sync.RWMutex
	queue  map[string][]chan interface{}
	closed bool
}

func New() *personaQ {
	personaq := &personaQ{}
	personaq.queue = make(map[string][]chan interface{})
	return personaq
}

func (q *personaQ) Publish(topic string, data interface{}) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.closed {
		return
	}

	for _, ch := range q.queue[topic] {
		go func(ch chan interface{}) {
			ch <- data
		}(ch)
	}
}

func (q *personaQ) Subscribe(topic string) <-chan interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	ch := make(chan interface{}, 1)
	q.queue[topic] = append(q.queue[topic], ch)
	return ch
}

func (q *personaQ) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed == false {
		q.closed = true
		for _, queue := range q.queue {
			for _, ch := range queue {
				close(ch)
			}
		}
	}
}
