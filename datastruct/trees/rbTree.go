package trees

const (
	RED   = true
	BLACK = false
)

type Node struct {
	Parent      *Node
	Left, Right *Node
	color       bool
}
