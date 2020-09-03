package iterator

import (
	"jobs/internal/platform/datastructure/node"
)

type IIterator interface {
	HasNext() bool
	Next() *node.Node
}
