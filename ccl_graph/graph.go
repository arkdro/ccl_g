package ccl_graph

import (
	"github.com/twmb/algoimpl/go/graph"
	"github.com/romana/rlog"

	"github.com/asdf/ccl_g/cell"
	"github.com/asdf/ccl_g/coordinates"
	"github.com/asdf/ccl_g/result"
)

type Ccl_graph struct {
	g *graph.Graph
	nodes *map[result.Merged_label]*ccl_node
}

type ccl_node struct {
	id result.Merged_label
	cells *map[cell.Ccl_cell]bool
	node *graph.Node
}

func create_or_fetch_node(g *graph.Graph, nodes *map[result.Merged_label]*ccl_node, label result.Merged_label) *ccl_node {
	c_node, found := (*nodes)[label]
	if found {
		return c_node
	}
	g_node := add_node_to_graph(g)
	*g_node.Value = label
	c_node = add_ccl_node_to_nodes(nodes, g_node, label)
	add_node_to_storage(nodes, c_node, label)
	return c_node
}

func add_ccl_node_to_nodes(nodes *map[result.Merged_label]*ccl_node, node *graph.Node, label result.Merged_label) *ccl_node {
	cells := make(map[cell.Ccl_cell]bool)
	cnode := &ccl_node{
		id: label,
		node: node,
		cells: &cells,
	}
	(*nodes)[label] = cnode
	return cnode
}

func add_node_to_storage(nodes *map[result.Merged_label]*ccl_node, node *ccl_node, label result.Merged_label) {
	(*nodes)[label] = node
}

func add_node_to_graph(g *graph.Graph) *graph.Node {
	node := g.MakeNode()
	return &node
}

func get_same_label_neighbour_cells(width int, height int, connectivity int, label result.Merged_label, x int, y int, merged [][]result.Merged_label) []cell.Ccl_cell {
	coords := coordinates.Get_neighbour_coordinates(connectivity, width, height, x, y)
	result := []cell.Ccl_cell{}
	for _, cell := range coords {
		if same_label(label, cell, merged) {
			result = append(result, cell)
		}
	}
	return result
}

func same_label(label result.Merged_label, cell cell.Ccl_cell, merged [][]result.Merged_label) bool {
	return label == merged[cell.Y][cell.X]
}

func add_cells_to_node(nodes *map[result.Merged_label]*ccl_node, label result.Merged_label, neighbour_cells []cell.Ccl_cell) {
	rlog.Warnf("label: %+v\n", label)
	rlog.Warnf("nodes1: %+v\n", nodes)
	c_node := (*nodes)[label]
	cells := c_node.cells
	rlog.Warnf("neighbour_cells: %+v\n", neighbour_cells)
	rlog.Warnf("cells1: %+v\n", cells)
	for _, cell := range neighbour_cells {
		(*cells)[cell] = true
	}
	rlog.Warnf("p, t: %T, v: %+v\n", c_node, c_node)
	rlog.Warnf("cells2: %+v\n", cells)
	rlog.Warnf("nodes2: %+v\n", nodes)
}

func get_neighbour_labels(width int, height int, connectivity int, label result.Merged_label, x int, y int, merged [][]result.Merged_label) []result.Merged_label {
	result := []result.Merged_label{}
	coords := coordinates.Get_neighbour_coordinates(connectivity, width, height, x, y)
	for _, coord := range coords {
		current_label := merged[coord.Y][coord.X]
		if label != current_label {
			result = append(result, current_label)
		}
	}
	return result
}

func create_neighbours_in_graph(g *graph.Graph, nodes *map[result.Merged_label]*ccl_node, c_node *ccl_node, neighbour_labels []result.Merged_label) {
	g_node := (*c_node).node
	for _, label := range neighbour_labels {
		current_c_node, found := (*nodes)[label]
		if found {
			current_g_node := (*current_c_node).node
			g.MakeEdge(*g_node, *current_g_node)
		}
	}
}

func Build_graph(width int, height int, merged [][]result.Merged_label, connectivity int) Ccl_graph {
	g := graph.New(graph.Undirected)
	nodes := make(map[result.Merged_label]*ccl_node)
	res := Ccl_graph{g: g, nodes: &nodes}
	for y, row := range merged {
		rlog.Warn("y: ", y)
		for x, label := range row {
			rlog.Warn("x: ", x)
			node := create_or_fetch_node(g, &nodes, label)
			rlog.Warnf("node: %+v\n", *node)
			neighbour_labels := get_neighbour_labels(width, height, connectivity, label, x, y, merged)
			neighbour_cells := get_same_label_neighbour_cells(width, height, connectivity, label, x, y, merged)
			add_cells_to_node(&nodes, label, neighbour_cells)
			dump_nodes(&nodes)
			create_neighbours_in_graph(g, &nodes, node, neighbour_labels)
		}
	}
	return res
}

func Results_equal(graph Ccl_graph, expected interface{}) bool {
	return false
}

func dump_nodes(nodes *map[result.Merged_label]*ccl_node) {
	for k, v := range *nodes {
		rlog.Warnf("label: %+v,\nnode: %+v, %+v, %+v\n", k, v.id, v.cells, v.node)
	}
}
