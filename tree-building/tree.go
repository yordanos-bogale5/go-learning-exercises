package tree

import "fmt"

// Record holds the ID of the current record and that of its nearest parent
type Record struct {
	ID     int
	Parent int
}

// Node holds the id of the current node and a reference array its nearest descendants
type Node struct {
	ID       int
	Children []*Node
}

var (
	nodeTable = map[int]*Node{}
	root      *Node
)

// Build creates node tree for an array of records
func Build(rs []Record) (*Node, error) {
	if len(rs) == 0 {
		return nil, nil
	}

	// order node id's. Since we expect contiguous id's here
`	// this will also help in error handling
`	for i, r := range rs {
		if r.ID < 0 || r.ID >= len(rs) {
			return nil, fmt.Errorf("Must have contigous ID's")
		}

		if r.ID != i {
			rs[i], rs[r.ID] = rs[r.ID], rs[i]
		}
	}

	nodes := make([]Node, len(rs))

	for i, r := range rs {
		if r.ID == 0 && r.Parent != 0 {
			return nil, fmt.Errorf("Root id cannot have parent. Found ID: %d, Parent: %d", r.ID, r.Parent)
		}

		if i != r.ID {
			return nil, fmt.Errorf("Either duplicate or non-contiguous id's")
		}

		if r.ID < r.Parent {
			return nil, fmt.Errorf("Parent id should be less than node id. ID: %d, Parent: %d", r.ID, r.Parent)
		}

		if r.ID == r.Parent && r.ID != 0 {
			return nil, fmt.Errorf("Only root ID can be equal to its parent: ID: %d, Parent: %d", r.ID, r.Parent)
		}

		nodes[i].ID = r.ID
		if i != 0 {
			p := &nodes[r.Parent]                      // pointer to parent node
			p.Children = append(p.Children, &nodes[i]) // append current not to pointer to parent
		}
	}

	return &nodes[0], nil
}
