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
