package graph

import (
	"time"

	"github.com/google/uuid"
)

type Iterator interface {
	Next() bool
	Error() error
	Close() error
}

type LinkIterator interface {
	Iterator
	Link() *Link
}

type EdgeIterator interface {
	Iterator
	Edge() *Edge
}

type Link struct {
	ID          uuid.UUID
	URL         string
	RetrievedAt time.Time
}

type Edge struct {
	ID        uuid.UUID
	Src       uuid.UUID
	Dst       uuid.UUID
	UpdatedAt time.Time
}

type Graph interface {
	UpsertLink(link *Link) error
	FindLink(id uuid.UUID) (*Link, error)

	UpsertEdge(edge *Edge) error
	RemoveStaleEdges(FromID uuid.UUID, updatedBefore time.Time) error

	Links(fromID, toID uuid.UUID, retrievedBefore time.Time) (LinkIterator, error)
	Edges(fromId, toID uuid.UUID, updatedBefore time.Time) (EdgeIterator, error)
}
