package ddd

import (
	"github.com/google/uuid"
	"time"
)

type AggregateRoot struct {
	AggregateRootSequence uuid.UUID
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (agr AggregateRoot) EventPublisher() {}

func (agr AggregateRoot) EventConsumer() {}
