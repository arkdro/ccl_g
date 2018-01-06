package result

type Label int

type One_color_result [][]Label

type Result []One_color_result

func (r1 Result) Equal(r2 Result) bool {
	// FIXME not implemented
	return false
}

func Build_result(data []*[][]int) Result {
	res := make(Result, 0)
	for _, item := range data {
		one_color_result := make([][]Label, 0)
		for _, row := range *item {
			label_row := make([]Label, 0)
			for _, point := range row {
				label := Label(point)
				label_row = append(label_row, label)
			}
			one_color_result = append(one_color_result, label_row)
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

func prepare_label_map(width int, height int) []Label {
	label_map := make([]Label, width * height)
	for i := range label_map {
		label_map[i] = -1
	}
	return label_map
}
