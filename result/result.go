package result

type Node int

type One_color_result [][]Node

type Result []One_color_result

func (r1 Result) Equal(r2 Result) bool {
	// FIXME not implemented
	return false
}

func Build_result(data []*[][]int) Result {
	res := make(Result, 0)
	for _, item := range data {
		one_color_result := make([][]Node, 0)
		for _, row := range *item {
			node_row := make([]Node, 0)
			for _, point := range row {
				node := Node(point)
				node_row = append(node_row, node)
			}
			one_color_result = append(one_color_result, node_row)
		}
		res = append(res, one_color_result)
	}
	return res
}

func (r Result) Valid_data(width int, height int, color_range int) bool {
	if len(r) != color_range {
		return false
	}
	for _, item := range r {
		res := valid_item(width, height, item)
		if res == false {
			return false
		}
	}
	return true
}

func valid_item(width int, height int, item One_color_result) bool {
	if len(item) != height {
		return false
	}
	for _, row := range item {
		if len(row) != width {
			return false
		}
	}
	return true
}
