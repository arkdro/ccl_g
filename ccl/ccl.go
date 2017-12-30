package ccl

import (
	"github.com/asdf/ccl_g/point"
)

const connectivity = 8

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
				external_contour_tracing(color, x, y, label, &data, &dummy,
					&labels, &dummy_labels)
				label++
			} else if has_unmarked_background_below(color, x, y, &data, &labels) {
				// step 2: newly encountered internal contour
				if has_label(x, y, &labels) {
					// part of an external contour. Already labeled
				} else {
					// left neighbor must be labeled
					copy_left_label(x, y, &labels)
				}
				internal_contour_tracing(color, x, y, label, &data, &dummy,
					&labels, &dummy_labels)
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

func external_contour_tracing(color int, x int, y int, label int, data *[][]int, dummy *[]int, labels *[][]int, dummy_labels *[]int) {
	pos := 7
	contour_tracing(color, x, y, label, data, dummy, labels, dummy_labels, pos)
}

func internal_contour_tracing(color int, x int, y int, label int, data *[][]int, dummy *[]int, labels *[][]int, dummy_labels *[]int) {
	pos := 3
	contour_tracing(color, x, y, label, data, dummy, labels, dummy_labels, pos)
}

func contour_tracing(color int, x int, y int, label int, data *[][]int, dummy *[]int, labels *[][]int, dummy_labels *[]int, init_pos int) {
	initial_pair_filled := false
	initial_pair := []point.Point{
		{x, y},
		{},
	}
	prev_point := point.Point{x, y}
	for {
		new_point, pos, found := tracer(color, prev_point, label, data, dummy, labels, dummy_labels, init_pos)
		if ! found {
			// an isolated point
			return
		}
		if initial_pair_filled {
			if contour_finished(&initial_pair, new_point, prev_point) {
				return
			}
		} else {
			update_initial_pair(&initial_pair, new_point)
			initial_pair_filled = true
		}
		prev_point = new_point
	}
}

func tracer(color int, pt point.Point, label int, data *[][]int, dummy *[]int, labels *[][]int, dummy_labels *[]int, init_pos int) (point.Point, int, bool) {
	pos := init_pos
	for ; pos != init_pos; {
		point2 := get_neighbour_coord(pt, pos)
		color2 := get_color(point2, data, color)
		if same_colors(color, color2) {
			return point2, pos, true
		}
		mark_background_point(point2, data, dummy)
		pos = next_pos(pos)
	}
	return point.Point{}, 0, false
}

func update_initial_pair(initial_pair *[]point.Point, new_point point.Point) {
	(*initial_pair)[1] = new_point
}

func contour_finished(initial_pair *[]point.Point, new_point point.Point, prev_point point.Point) bool {
	return (*initial_pair)[0] == prev_point &&
		(*initial_pair)[1] == new_point
}

func get_color(pt point.Point, data *[][]int, orig_color int) int {
	if pt.Y < 0 {
		return orig_color - 1
	} else {
		return (*data)[pt.Y][pt.X]
	}
}

func same_colors(color1 int, color2 int) bool {
	return color1 == color2
}

func mark_background_point(pt point.Point, data *[][]int, dummy *[]int) {
	if pt.Y < 0 {
		(*dummy)[pt.X] = -1
	} else {
		(*data)[pt.Y][pt.X] = -1
	}
}

func next_pos(pos int) int {
	new_pos := (pos + 1) % connectivity
	return new_pos
}

func get_neighbour_coord(pt point.Point, pos int) point.Point {
	dx, dy := pos_to_delta(pos)
	point2 := point.Point{X: pt.X + dx, Y: pt.Y + dy}
	return point2
}

func pos_to_delta(pos int) (int, int) {
	var dx, dy int
	switch pos {
	case 0:
		dx, dy = 1, 0
	case 1:
		dx, dy = 1, 1
	case 2:
		dx, dy = 0, 1
	case 3:
		dx, dy = -1, 1
	case 4:
		dx, dy = -1, 0
	case 5:
		dx, dy = -1, -1
	case 6:
		dx, dy = 0, -1
	case 7:
		dx, dy = 1, -1
	}
	return dx, dy
}
func prev_point_pos(pos int) int {
	var prev_pos int
	switch pos {
	case 0:
		prev_pos = 4
	case 1:
		prev_pos = 5
	case 2:
		prev_pos = 6
	case 3:
		prev_pos = 7
	case 4:
		prev_pos = 0
	case 5:
		prev_pos = 1
	case 6:
		prev_pos = 2
	case 7:
		prev_pos = 3
	}
	return prev_pos
}

