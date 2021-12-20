package algors

type Color bool
type Key interface {
	CompareTo(Key) int
}
type Value interface{}

const (
	RED   = true
	BLACK = false
)

type Node struct {
	Key         Key
	Value       Value
	Left, Right *Node // subtress
	N           int   // nodes in this subtree
	Color       bool
}

func NewNode(key Key, value Value, N int, color bool) *Node {
	return &Node{
		Key:   key,
		Value: value,
		N:     N,
		Color: color,
	}
}

func isRed(x *Node) bool {
	if x == nil {
		return false
	}

	return x.Color
}

func size(x *Node) int {
	if x == nil {
		return 0
	}

	return x.N
}

func rotateLeft(h *Node) *Node {
	x := h.Right
	h.Right = x.Left
	x.Left = h
	x.Color = h.Color
	h.Color = RED
	x.N = h.N
	h.N = 1 + h.Left.N + h.Right.N
	return x
}

func rotateRight(h *Node) *Node {
	x := h.Left
	h.Left = x.Right
	x.Right = h
	x.Color = h.Color
	h.Color = RED
	x.N = h.N
	h.N = 1 + size(h.Left) + size(h.Right)
	return x
}

func flipColors(h *Node) {
	h.Color = RED
	h.Left.Color = BLACK
	h.Right.Color = BLACK
}

type RedBlackBST struct {
	root *Node
}

func (rb *RedBlackBST) Put(key Key, value Value) {
	root := rb.put(rb.root, key, value)
	root.Color = BLACK
	rb.root = root
}

func (rb *RedBlackBST) put(h *Node, key Key, value Value) *Node {
	if h == nil {
		return NewNode(key, value, 1, RED)
	}

	cmp := key.CompareTo(h.Key)
	if cmp < 0 {
		h.Left = rb.put(h.Left, key, value)
	} else if cmp > 0 {
		h.Right = rb.put(h.Right, key, value)
	} else {
		h.Value = value
	}

	if isRed(h.Right) && !isRed(h.Left) {
		h = rotateLeft(h)
	}
	if isRed(h.Left) && isRed(h.Left.Left) {
		h = rotateRight(h)
	}
	if isRed(h.Left) && isRed(h.Right) {
		flipColors(h)
	}

	h.N = size(h.Left) + size(h.Right) + 1
	return h
}
