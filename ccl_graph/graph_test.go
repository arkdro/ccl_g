package ccl_graph

import (
	"github.com/twmb/algoimpl/go/graph"

	"github.com/asdf/ccl_g/cell"
	"github.com/asdf/ccl_g/result"

//	"fmt"
	"reflect"
	"testing"
)

func Test_build_graph_and_compare_labels(t *testing.T) {
	merged := result.Build_merge_ccl_result()
	width := 7
	height := 4
	connectivity := 8
	graph := Build_graph(width, height, merged, connectivity)
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

func Test_build_graph_and_compare_cells(t *testing.T) {
	merged := result.Build_merge_ccl_result()
	width := 7
	height := 4
	connectivity := 8
	graph := Build_graph(width, height, merged, connectivity)
	expected_cells := map[result.Merged_label]map[cell.Ccl_cell]bool{
		result.Make_label(0, 1): map[cell.Ccl_cell]bool{
			cell.Ccl_cell{X:2, Y:0}: true,
			cell.Ccl_cell{X:3, Y:1}: true,
		},
		result.Make_label(0, 2): map[cell.Ccl_cell]bool{
			cell.Ccl_cell{X:0, Y:2}: true,
			cell.Ccl_cell{X:0, Y:3}: true,
			cell.Ccl_cell{X:1, Y:3}: true,
			cell.Ccl_cell{X:2, Y:3}: true,
		},
		result.Make_label(0, 3): map[cell.Ccl_cell]bool{
			cell.Ccl_cell{X:6, Y:2}: true,
			cell.Ccl_cell{X:6, Y:3}: true,
		},
		result.Make_label(1, 1): map[cell.Ccl_cell]bool{
			cell.Ccl_cell{X:0, Y:0}: true,
			cell.Ccl_cell{X:1, Y:0}: true,
			cell.Ccl_cell{X:0, Y:1}: true,
			cell.Ccl_cell{X:1, Y:1}: true,
			cell.Ccl_cell{X:1, Y:2}: true,
			cell.Ccl_cell{X:2, Y:2}: true,
		},
		result.Make_label(1, 2): map[cell.Ccl_cell]bool{
			cell.Ccl_cell{X:5, Y:2}: true,
		},
		result.Make_label(2, 1): map[cell.Ccl_cell]bool{
			cell.Ccl_cell{X:3, Y:0}: true,
			cell.Ccl_cell{X:4, Y:0}: true,
			cell.Ccl_cell{X:5, Y:0}: true,
			cell.Ccl_cell{X:6, Y:0}: true,
			cell.Ccl_cell{X:2, Y:1}: true,
			cell.Ccl_cell{X:4, Y:1}: true,
			cell.Ccl_cell{X:5, Y:1}: true,
			cell.Ccl_cell{X:6, Y:1}: true,
			cell.Ccl_cell{X:3, Y:2}: true,
			cell.Ccl_cell{X:4, Y:2}: true,
			cell.Ccl_cell{X:3, Y:3}: true,
			cell.Ccl_cell{X:4, Y:3}: true,
			cell.Ccl_cell{X:5, Y:3}: true,
		},
	}
	result_cells := compare_cells(t, graph, &expected_cells)
	if !result_cells {
		t.Error("compare cells error")
	}
}

func Test_build_graph_and_compare_neighbours(t *testing.T) {
	merged := result.Build_merge_ccl_result()
	width := 7
	height := 4
	connectivity := 8
	graph := Build_graph(width, height, merged, connectivity)
	expected_neighbours := map[result.Merged_label][]result.Merged_label{
		result.Make_label(0, 1): []result.Merged_label{
			result.Make_label(1, 1),
			result.Make_label(2, 1),
		},
		result.Make_label(0, 2): []result.Merged_label{
			result.Make_label(1, 1),
			result.Make_label(2, 1),
		},
		result.Make_label(0, 3): []result.Merged_label{
			result.Make_label(1, 2),
			result.Make_label(2, 1),
		},
		result.Make_label(1, 1): []result.Merged_label{
			result.Make_label(0, 1),
			result.Make_label(2, 1),
		},
		result.Make_label(1, 2): []result.Merged_label{
			result.Make_label(0, 3),
			result.Make_label(2, 1),
		},
		result.Make_label(2, 1): []result.Merged_label{
			result.Make_label(0, 1),
			result.Make_label(0, 2),
			result.Make_label(0, 3),
			result.Make_label(1, 1),
			result.Make_label(1, 2),
		},
	}
	result_neighbours := compare_neighbours(t, graph, &expected_neighbours)
	if !result_neighbours {
		t.Error("compare neighbors error")
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

func compare_neighbours(t *testing.T, g Ccl_graph, expected_neighbours *map[result.Merged_label][]result.Merged_label) bool {
	result := true
	nodes := g.nodes
	for label, c_node := range *nodes {
		g_node := c_node.node
		neighbours := g.g.Neighbors(*g_node)
		actual_labels := get_actual_labels(t, neighbours)
		expected_labels := (*expected_neighbours)[label]
		label_result := compare_neighbour_labels(t, actual_labels, expected_labels)
		if !label_result {
			t.Errorf("wrong neighbours, label: %+v\nactual: %+v\nexpected: %+v\n",
				label, actual_labels, expected_labels)
			result = false
		}
	}
	return result
}

func get_actual_labels(t *testing.T, neighbours []graph.Node) []result.Merged_label {
	labels := []result.Merged_label{}
	for _, node := range neighbours {
		label := (*node.Value).(result.Merged_label)
		labels = append(labels, label)
	}
	return labels
}

func compare_neighbour_labels(t *testing.T, actual_labels []result.Merged_label, expected_labels []result.Merged_label) bool {
	res := true
	if len(actual_labels) != len(expected_labels) {
		t.Errorf("lengths of actual (%d) and expected labels (%d) differ\n",
			len(actual_labels), len(expected_labels))
		res = false
	}
	acc := make(map[result.Merged_label]int)
	for _, label := range actual_labels {
		_, found := acc[label]
		if found {
			acc[label]++
		} else {
			acc[label] = 1
		}
	}
	for _, label := range expected_labels {
		_, found := acc[label]
		if found {
			acc[label]--
		} else {
			acc[label] = -1
		}
	}
	for label, value := range acc {
		if value != 0 {
			t.Errorf("wrong value (%d) for label (%+v)\n", value, label)
			res = false
		}
	}
	return res
}
