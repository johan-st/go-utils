package lru

import (
	"sync"

	ll "github.com/johan-st/go-utils/linkedlist"
)

type node struct {
	value any
	prev  *node
	next  *node
}

// generic comparable check?

// Thread safe LRU cache
type Lru struct {
	list         *ll.LinkedList
	cap          int
	lookup       map[any]*node
	reverse      map[*node]any
	mutexLookup  sync.RWMutex
	mutexReverse sync.RWMutex
}

func New(capacity int) *Lru {
	return &Lru{
		cap:     capacity,
		lookup:  make(map[any]*node),
		reverse: make(map[*node]any),
	}
}

// Update visits the value, moving it to the front of the list or creates a new if not found
func (l *Lru) Update(key string, value any) (any, bool) {

}

func (l *Lru) add(key string, value any) (any, bool)
func (l *Lru) moveToFront(n *node)
func (l *Lru) trim()
func (l *Lru) setLookup(key string, n *node)
func (l *Lru) rmLookup(key string, n *node)
