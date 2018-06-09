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
	expected_cells := make(map[result.Merged_label]cell.Ccl_cell)
	expected_neighbours := make(map[result.Merged_label]result.Merged_label)
	t.Errorf("graph: %v\n", graph)
	expected_labels := []result.Merged_label{
		result.Make_label(0, 1),
		result.Make_label(0, 2),
		result.Make_label(0, 3),
		result.Make_label(1, 1),
		result.Make_label(1, 2),
		result.Make_label(2, 1),
	}
	result1 := compare_labels(t, graph, expected_labels)
	if !result1 {
		t.Error("compare labels error")
	}
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

func compare_cells(t *testing.T, g Ccl_graph, expected_cells *map[result.Merged_label]map[cell.Ccl_cell]bool) bool {
	result := true
	nodes := g.nodes
	for label, node := range *nodes {
		actual_cells := (*node).cells
		expected_node_cells := (*expected_cells)[label]
		node_result := compare_node_cells(t, actual_cells, &expected_node_cells)
		if !node_result {
			t.Errorf("cells mismatch, label: %+v\nactual: %+v\nexpected: %+v\n",
				label, *actual_cells, expected_node_cells)
			result = false
		}
	}
	return result
}

func compare_node_cells(t *testing.T, actual_cells *map[cell.Ccl_cell]bool, expected_node_cells *map[cell.Ccl_cell]bool) bool {
	if !reflect.DeepEqual(*actual_cells, *expected_node_cells) {
		acc := make(map[cell.Ccl_cell]int)
		for k := range *actual_cells {
			_, found := acc[k]
			if found {
				acc[k]++
			} else {
				acc[k] = 1
			}
		}
		for k := range *expected_node_cells {
			_, found := acc[k]
			if found {
				acc[k] += 10
			} else {
				acc[k] = 10
			}
		}
		for k, v := range acc {
			if v != 11 {
				t.Errorf("wrong cell: %+v, %d\n", k, v)
			}
		}
		return false
	}
	return true
}
