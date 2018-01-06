package result

type Node int

type One_color_result [][]Node

type Result []One_color_result

func (r1 Result) Equal(r2 Result) bool {
	// FIXME not implemented
	return false
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
