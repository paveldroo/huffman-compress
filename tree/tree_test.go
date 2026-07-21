package tree_test

import (
	"testing"

	"github.com/paveldroo/huffman-compress/tree"
)

func TestTree(t *testing.T) {
	charsCount := map[string]int{
		"C": 32,
		"D": 42,
		"E": 120,
		"K": 7,
		"L": 42,
		"M": 24,
		"U": 37,
		"Z": 2,
	}
	want := tree.Node{
		Value:     "",
		Frequency: 306,
		Left: &tree.Node{
			Value:     "E",
			Frequency: 120,
			Left:      nil,
			Right:     nil,
		},
		Right: &tree.Node{
			Value:     "",
			Frequency: 186,
			Left: &tree.Node{
				Value:     "",
				Frequency: 79,
				Left: &tree.Node{
					Value:     "U",
					Frequency: 37,
					Left:      nil,
					Right:     nil,
				},
				Right: &tree.Node{
					Value:     "D",
					Frequency: 42,
					Left:      nil,
					Right:     nil,
				},
			},
			Right: &tree.Node{
				Value:     "",
				Frequency: 107,
				Left: &tree.Node{
					Value:     "L",
					Frequency: 42,
					Left:      nil,
					Right:     nil,
				},
				Right: &tree.Node{
					Value:     "",
					Frequency: 65,
					Left: &tree.Node{
						Value:     "C",
						Frequency: 32,
						Left:      nil,
						Right:     nil,
					},
					Right: &tree.Node{
						Value:     "",
						Frequency: 33,
						Left: &tree.Node{
							Value:     "",
							Frequency: 9,
							Left: &tree.Node{
								Value:     "Z",
								Frequency: 2,
								Left:      nil,
								Right:     nil,
							},
							Right: &tree.Node{
								Value:     "K",
								Frequency: 7,
								Left:      nil,
								Right:     nil,
							},
						},
						Right: &tree.Node{
							Value:     "M",
							Frequency: 24,
							Left:      nil,
							Right:     nil,
						},
					},
				},
			},
		},
	}
	root := tree.Tree(charsCount)

	if !treesEqual(&want, &root) {
		t.Fatal("trees are not equal")
	}

}

func treesEqual(root1, root2 *tree.Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	// fmt.Printf("node1.value: %s, node1.freq: %d\nnode2.value: %s, node2.freq: %d\n", root1.Value, root1.Frequency, root2.Value, root2.Frequency)
	if !(root1 != nil && root2 != nil) || root1.Value != root2.Value || root1.Frequency != root2.Frequency {
		return false
	}

	return treesEqual(root1.Left, root2.Left) && treesEqual(root1.Right, root2.Right)
}
