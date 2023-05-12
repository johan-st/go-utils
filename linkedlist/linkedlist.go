package linkedlist

type node struct {
	val any
}

// LinkedList is a doubly linked list
type LinkedList struct {
	head *node
	tail *node
}

func (l *LinkedList) Add(val any)
func (l *LinkedList) AddBack(val any)
func (l *LinkedList) AddAfter(index int, val any) error
func (l *LinkedList) Get(index int) (any, error)
func (l *LinkedList) Remove(index int) error
func (l *LinkedList) RemoveValue(val any) error

func (l *LinkedList) nodeAt(index int) (*node, error)
func (l *LinkedList) detatch(node *node)
