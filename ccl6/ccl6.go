package ccl6

import (
	"github.com/romana/rlog"

	"github.com/asdf/ccl_g/point"
)

func Ccl(width int, height int, color_range int, data *[][]int) []*[][]int {
	labels := make([]*[][]int, color_range)
	for color := 0; color < color_range; color++ {
		cur_data := prepare_data(width, height, data)
		cur_labels := ccl_one_color(width, height, color, cur_data)
		labels[color] = cur_labels
	}
	return labels
}

func ccl_one_color(width int, height int, color int, data *[][]int) *[][]int {
	labels := create_empty_labels(width, height)
	linked := make(map[int]map[int]bool, (width + height)/2)
	ccl_pass1(width, height, color, data, &labels, &linked)
	rlog.Debugf("ccl_one_color, after pass1, labels: %v\nlinked: %v",
		labels, linked)
	ccl_pass2(width, height, color, data, &labels, &linked)
	rlog.Debugf("ccl_one_color, after pass2, labels: %v\nlinked: %v",
		labels, linked)
	return &labels
}

func ccl_pass1(width int, height int, color int, data *[][]int, labels *[][]int, linked *map[int]map[int]bool) {
	label := 1
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if is_background(color, x, y, data) {
				continue
			}
			rlog.Debugf("ccl_pass1, x: %v, y: %v", x, y)
			neigbours := same_color_neigbours(width, color, x, y, data)
			if no_neigbours(neigbours) {
				fresh_label_set := init_label_set(label)
				(*linked)[label] = fresh_label_set
				(*labels)[y][x] = label
				label++
			} else {
				neigbour_labels := find_neigbour_labels(neigbours, labels)
				min_label := find_minimal_item(neigbour_labels)
				(*labels)[y][x] = min_label
				set_equivalence(neigbour_labels, linked)
				rlog.Debugf("ccl_pass1, after set_equivalence\nlinked: %v",
					linked)
			}
		}
	}
}

func ccl_pass2(width int, height int, color int, data *[][]int, labels *[][]int, linked *map[int]map[int]bool) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if is_background(color, x, y, data) {
				continue
			}
			min_label, label_exists := fetch_minimal_label(x, y, labels, linked)
			if label_exists == true {
				(*labels)[y][x] = min_label
			}
		}
	}
}

func same_color_neigbours(width int, colour int, x int, y int, data *[][]int) []point.Point {
	result := make([]point.Point, 0)
	coordinates := get_neigbour_coordinates(width, x, y)
	for _, pt := range coordinates {
		if is_foreground(colour, pt.X, pt.Y, data) {
			result = append(result, pt)
		}
	}
	return result
}

func get_neigbour_coordinates(width int, x int, y int) []point.Point {
	result := make([]point.Point, 0)
	left, left_valid := get_left_coordinate(x, y)
	upper_left, upper_left_valid := get_upper_left_coordinate(x, y)
	upper_right, upper_right_valid := get_upper_right_coordinate(width, x, y)
	if left_valid == true {
		result = append(result, left)
	}
	if upper_left_valid == true {
		result = append(result, upper_left)
	}
	if upper_right_valid == true {
		result = append(result, upper_right)
	}
	return result
}

func get_left_coordinate(x int, y int) (point.Point, bool) {
	if x <= 0 {
		return point.Point{}, false
	}
	if y < 0 {
		return point.Point{}, false
	}
	new_x := x - 1
	pt := point.Point{new_x, y}
	return pt, true
}

func get_upper_left_coordinate(x int, y int) (point.Point, bool){
	if y <= 0 {
		return point.Point{}, false
	}
	var new_y = y - 1
	var new_x int
	if is_even(y) {
		new_x = x
	} else {
		if x <= 0 {
			return point.Point{}, false
		} else {
			new_x = x - 1
		}
	}
	pt := point.Point{new_x, new_y}
	return pt, true
}

func get_upper_right_coordinate(width int, x int, y int) (point.Point, bool){
	if y <= 0 {
		return point.Point{}, false
	}
	var new_y = y - 1
	var new_x int
	if is_even(y) {
		if x < width - 1 {
			new_x = x + 1
		} else {
			return point.Point{}, false
		}
	} else {
		new_x = x
	}
	pt := point.Point{new_x, new_y}
	return pt, true
}

func is_even(n int) bool {
	return n % 2 == 0
}

func no_neigbours(neigbours []point.Point) bool {
	return len(neigbours) == 0
}

func init_label_set(label int) map[int]bool {
	res := make(map[int]bool)
	res[label] = true
	return res
}

func find_neigbour_labels(neigbours []point.Point, labels *[][]int) []int {
	result := make([]int, 0)
	for _, pt := range neigbours {
		label := (*labels)[pt.Y][pt.X]
		result = append(result, label)
	}
	return result
}

func find_minimal_item(labels []int) int {
	min := labels[0]
	for _, x := range labels {
		if x < min {
			min = x
		}
	}
	return min
}

// FIXME optimize performance: move it to pass2
func set_equivalence(labels []int, linked *map[int]map[int]bool) {
	max_union := calc_max_union(labels, linked)
	for _, label := range labels {
		(*linked)[label] = max_union
	}
}

func calc_max_union(labels []int, linked *map[int]map[int]bool) map[int]bool {
	union := make(map[int]bool)
	for _, label := range labels {
		cur_set := (*linked)[label]
		add_labels_to_union(cur_set, union)
	}
	return union
}

func add_labels_to_union(cur_set map[int]bool, union map[int]bool) {
	for label := range cur_set {
		union[label] = true
	}
}

func fetch_minimal_label(x int, y int, labels *[][]int, linked *map[int]map[int]bool) (int, bool) {
	label := (*labels)[y][x]
	if label == 0 {
		return 0, false
	}
	min_label := fetch_minimal_label_by_label(label, labels, linked)
	return min_label, true
}

func fetch_minimal_label_by_label(label int, labels *[][]int, linked *map[int]map[int]bool) int {
	equiv_set := (*linked)[label]
	keys := make([]int, 0)
	for key := range equiv_set {
		keys = append(keys, key)
	}
	min_label := find_minimal_item(keys)
	if label != min_label {
		result := fetch_minimal_label_by_label(min_label, labels, linked)
		equiv_set[result] = true
		(*linked)[label] = equiv_set
		return result
	} else {
		return min_label
	}
}

func is_foreground(color int, x int, y int, data *[][]int) bool {
	return !is_background(color, x, y, data)
}

func is_background(color int, x int, y int, data *[][]int) bool {
	if (*data)[y][x] != color {
		return true
	} else {
		return false
	}
}

func create_empty_labels(width int, height int) [][]int {
	labels := make([][]int, height)
	for y := 0; y < height; y++ {
		row := make([]int, width)
		labels[y] = row
	}
	return labels
}

func prepare_data(width int, height int, orig_data *[][]int) *[][]int {
	data := make([][]int, height)
	for y := 0; y < height; y++ {
		row := make([]int, width)
		data[y] = row
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x ++ {
			data[y][x] = (*orig_data)[y][x]
		}
	}
	return &data
}
