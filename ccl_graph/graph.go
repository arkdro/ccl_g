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

func create_or_fetch_node(g *graph.Graph, nodes *map[result.Merged_label]*ccl_node, label result.Merged_label, x int, y int) *ccl_node {
	c_node, found := (*nodes)[label]
	if found {
		return c_node
	}
	g_node := add_node_to_graph(g)
	*g_node.Value = label
	c_node = add_ccl_node_to_nodes(nodes, g_node, label, x, y)
	add_node_to_storage(nodes, c_node, label)
	return c_node
}

func add_ccl_node_to_nodes(nodes *map[result.Merged_label]*ccl_node, node *graph.Node, label result.Merged_label, x int, y int) *ccl_node {
	cells := make(map[cell.Ccl_cell]bool)
	current_cell := cell.Ccl_cell{X: x, Y: y}
	cells[current_cell] = true
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
			node := create_or_fetch_node(g, &nodes, label, x, y)
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

func Results_equal(graph Ccl_graph, expected result.G_result) bool {
	rlog.Warnf("results equal, graph: %+v\nexpected: %+v\n", graph, expected)
	dump_graph(graph)
	if !compare_labels(graph, expected) {
		rlog.Error("graph result labels mismatch")
		return false
	}
	for exp_key, exp_node := range expected {
		if !compare_cells(graph, exp_key, &exp_node) {
			rlog.Error("graph result cells mismatch")
			return false
		}
		if !compare_neighbours(graph, exp_key, &exp_node) {
			rlog.Error("graph result neighbours mismatch")
			return false
		}
	}
	return true
}

func compare_labels(graph Ccl_graph, expected result.G_result) bool {
	c_labels := make(map[result.Merged_label]bool)
	for label := range *graph.nodes {
		c_labels[label] = true
	}
	g_labels := make(map[result.Merged_label]bool)
	for label := range expected {
		g_label := result.G_to_merged_label(label)
		g_labels[g_label] = true
	}
	if len(c_labels) != len(g_labels) {
		rlog.Warnf("graph label lengths mismatch, ccl: %d, json: %d\n",
			len(c_labels), len(g_labels))
		return false
	}
	result := true
	for c_label := range c_labels {
		_, found := g_labels[c_label]
		if !found {
			rlog.Warnf("json label missing for ccl label: %+v\n", c_label)
			result = false
		}
	}
	for g_label := range g_labels {
		_, found := c_labels[g_label]
		if !found {
			rlog.Warnf("ccl label missing for json label: %+v\n", g_label)
			result = false
		}
	}
	return result
}

func compare_cells(graph Ccl_graph, exp_key result.G_label, exp_node *result.G_item) bool {
	result_status := true
	exp_g_cells := exp_node.Cells
	g_cells_uniq := make(map[cell.Ccl_cell]int)
	for _, g_cell := range exp_g_cells {
		c_cell := cell.Ccl_cell{g_cell.X, g_cell.Y}
		_, found := g_cells_uniq[c_cell]
		if found {
			g_cells_uniq[c_cell]++
		} else {
			g_cells_uniq[c_cell] = 1
		}
	}
	for g_cell, count := range g_cells_uniq {
		if count != 1 {
			result_status = false
			rlog.Warnf("expected cells duplicated, cell: %+v, cnt: %d\n",
				g_cell, count)
		}
	}
	g_label := result.G_to_merged_label(exp_key)
	c_node := (*graph.nodes)[g_label]
	c_cells := c_node.cells
	if len(*c_cells) != len(g_cells_uniq) {
		rlog.Warnf("lengths mismatch, ccl: %d, expected: %d\n",
			len(*c_cells), len(g_cells_uniq))
		return false
	}
	for c_cell := range *c_cells {
		_, found := g_cells_uniq[c_cell]
		if !found {
			result_status = false
			rlog.Warnf("ccl cell missing in expected cells: %+v\n", c_cell)
		}
	}
	return result_status
}

func compare_neighbours(g Ccl_graph, exp_key result.G_label, exp_node *result.G_item) bool {
	result_status := true
	g_label := result.G_to_merged_label(exp_key)
	c_node := (*g.nodes)[g_label]
	neighbors := g.g.Neighbors(*c_node.node)
	c_labels := make(map[result.Merged_label]bool)
	for _, neighbor := range neighbors {
		label := (*neighbor.Value).(result.Merged_label)
		c_labels[label] = true
	}
	exp_g_labels := exp_node.Neigbours
	exp_merged_labels := make(map[result.Merged_label]bool)
	for _, g_label := range exp_g_labels {
		label := result.G_to_merged_label(g_label)
		exp_merged_labels[label] = true
	}
	for label := range c_labels {
		_, found := exp_merged_labels[label]
		if !found {
			result_status = false
			rlog.Warnf("ccl label missing in expected neighbors: %+v\n", label)
		}
	}
	return result_status
}

func dump_nodes(nodes *map[result.Merged_label]*ccl_node) {
	for k, v := range *nodes {
		rlog.Warnf("label: %+v,\nnode: %+v, %+v, %+v\n", k, v.id, v.cells, v.node)
	}
}

func dump_graph(g Ccl_graph) {
	nodes := g.nodes
	for _, c_node := range *nodes {
		cur_label := (*c_node.node.Value).(result.Merged_label)
		rlog.Warnf("graph, cur label: %+v\n", cur_label)
		neighbors := g.g.Neighbors(*c_node.node)
		for _, g_node := range neighbors {
			label := (*g_node.Value).(result.Merged_label)
			cells := *c_node.cells
			rlog.Warnf("label: %+v\ncells: %+v\n", label, cells)
		}
	}
}
