package ccl

func Ccl(width int, height int, color_range int, data [][]int) {
}

func ccl_one_color(width int, height int, color int, data [][]int) [][]int {
	labels := create_empty_labels(width, height)
	dummy := make([]int, width)
	dummy_labels := make([]int, width)
	label := 1
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if is_background(color, x, y, &data) {
				continue
			}
			if has_no_label(x, y, &labels) &&
				has_background_above(color, x, y, &data, &dummy) {
				// step 1: external contour of a new component
				labels[y][x] = label
				countour_tracing(color, x, y, label, &data, &dummy, &labels, &dummy_labels)
				label++
			} else if has_unmarked_background_below(color, x, y, &data, &labels) {
				// step 2: newly encountered internal contour
				if has_label(x, y, &labels) {
					// part of an external contour. Already labeled
				} else {
					// left neighbor must be labeled
					copy_left_label(x, y, &labels)
				}
				countour_tracing(color, x, y, label, &data, &dummy, &labels, &dummy_labels)
			} else {
				// step 3: left neighbor must be a labeled pixel
				copy_left_label(x, y, &labels)
			}
		}
	}
	return labels
}

func has_unmarked_background_below(color int, x int, y int, data *[][]int, labels *[][]int) bool {
	if (*data)[y+1][x] != color &&
		(*labels)[y+1][x] == 0 {
		return true
	} else {
		return false
	}
}

func has_background_above(color int, x int, y int, data *[][]int, dummy *[]int) bool {
	if y > 0 {
		if (*data)[y][x] != color {
			return true
		} else {
			return false
		}
	} else {
		if (*dummy)[x] != color {
			return true
		} else {
			return false
		}
	}
}

func is_background(color int, x int, y int, data *[][]int) bool {
	if (*data)[y][x] != color {
		return true
	} else {
		return false
	}
}

func has_no_label(x int, y int, labels *[][]int) bool {
	return !has_label(x, y, labels)
}

func has_label(x int, y int, labels *[][]int) bool {
	if (*labels)[y][x] == 0 {
		return false
	} else {
		return true
	}
}

func copy_left_label(x int, y int, labels *[][]int) {
	left_label := get_left_label(x, y, labels)
	(*labels)[y][x] = left_label
}

func get_left_label(x int, y int, labels *[][]int) int {
	return (*labels)[y][x-1]
}

func create_empty_labels(width int, height int) [][]int {
	labels := make([][]int, height)
	for y := 0; y < height; y++ {
		row := make([]int, width)
		labels[y] = row
	}
	return labels
}

func countour_tracing(color int, x int, y int, label int, data *[][]int, dummy *[]int, labels *[][]int, dummy_labels *[]int) {
}

