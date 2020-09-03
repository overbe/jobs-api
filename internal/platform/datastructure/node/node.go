package node

type Node struct {
	item    int
	status  string
	jobType string
	next    *Node
}

// Init Node
func New(item int, status, jobType string) *Node {
	return &Node{item, status, jobType, nil}
}

func (n *Node) Item() int {
	return n.item
}

func (n *Node) Status() string {
	return n.status
}

func (n *Node) JobType() string {
	return n.jobType
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetItem(item int) {
	n.item = item
}

func (n *Node) SetStatus(status string) {
	n.status = status
}

func (n *Node) SetJobType(jobType string) {
	n.jobType = jobType
}

func (n *Node) SetNext(nextNode *Node) *Node {
	n.next = nextNode
	return n.next
}
