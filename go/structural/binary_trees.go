// Another very common approac to the composite pattern
// is when working with binary tree structures. In a
// binary tree, you need to store instances of itself in a field

package main

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

func main() {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
			Left:      nil,
		},
		Left: &Tree{4, nil, nil},
	}
}
