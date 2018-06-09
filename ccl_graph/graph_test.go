package ccl_graph

import (
//	"github.com/twmb/algoimpl/go/graph"

	"github.com/asdf/ccl_g/cell"
	"github.com/asdf/ccl_g/result"

//	"fmt"
	"testing"
)

func Test_build_graph(t *testing.T) {
	merged := result.Build_merge_ccl_result()
	t.Errorf("merged: %v\n", merged)
	width := 7
	height := 4
	connectivity := 6
	graph := Build_graph(width, height, merged, connectivity)
	expected_labels := []result.Merged_label{}
	expected_cells := make(map[result.Merged_label]cell.Ccl_cell)
	expected_neighbours := make(map[result.Merged_label]result.Merged_label)
	t.Errorf("graph: %v\n", graph)
}

func compare_labels(t *testing.T, g Ccl_graph, expected_labels []result.Merged_label) bool {
	nodes := g.nodes
	if len(*nodes) != len(expected_labels) {
		t.Errorf("lengths of graph (%d) and expected labels array (%d) mismatch\n",
			len(*nodes), len(expected_labels))
		return false
	}
	expected := make(map[result.Merged_label]int)
	for _, key := range expected_labels {
		expected[key] = 0
	}
	if len(*nodes) != len(expected) {
		t.Errorf("lengths of graph (%d) and expected labels map (%d) mismatch\n",
			len(*nodes), len(expected))
		return false
	}

	for label := range *nodes {
		elem, found := expected[label]
		if !found {
			t.Errorf("label (%+v) exists in graph, missing in expected labels\n",
				label)
			return false
		}
		if elem != 0 {
			t.Errorf("repeated labels (%+v) in graph\n", label)
			return false
		}
		expected[label]++
	}
	result := true
	for key, val := range expected {
		if val != 1 {
			t.Errorf("label (%+v) not equal to 1 (%v)\n", key, val)
			result = false
		}
	}
	return result
}
