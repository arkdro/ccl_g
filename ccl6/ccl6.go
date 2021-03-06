package ccl6

import (
	"github.com/romana/rlog"

	"github.com/asdf/ccl_g/dset"
	"github.com/asdf/ccl_g/point"

	"fmt"
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
	linked := create_linked_storage(width, height)
	ccl_pass1(width, height, color, data, &labels, linked)
	rlog.Debugf("ccl_one_color, after pass1, labels: %v\nlinked: %v",
		labels, linked_as_string(linked))
	ccl_pass2(width, height, color, data, &labels, linked)
	rlog.Debugf("ccl_one_color, after pass2, labels: %v\nlinked: %v",
		labels, linked_as_string(linked))
	return &labels
}

func ccl_pass1(width int, height int, color int, data *[][]int, labels *[][]int, linked *[]*dset.Dset) {
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
					linked_as_string(linked))
			}
		}
	}
}

func ccl_pass2(width int, height int, color int, data *[][]int, labels *[][]int, linked *[]*dset.Dset) {
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

func init_label_set(label int) *dset.Dset {
	return dset.Create(label)
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

func set_equivalence(labels []int, linked *[]*dset.Dset) {
	union := calc_max_union(labels, linked)
	for _, label := range labels {
		(*linked)[label] = union
	}
}

func calc_max_union(labels []int, linked *[]*dset.Dset) *dset.Dset {
	base_label := labels[0]
	base_item := (*linked)[base_label]
	for _, label := range labels {
		cur_set := (*linked)[label]
		var new_item *dset.Dset
		if cur_set == nil {
			new_item = dset.Create(label)
		} else {
			new_item = cur_set
		}
		dset.Union(new_item, base_item)
	}
	return base_item
}

func fetch_minimal_label(x int, y int, labels *[][]int, linked *[]*dset.Dset) (int, bool) {
	label := (*labels)[y][x]
	if label == 0 {
		return 0, false
	}
	equiv_set := (*linked)[label]
	rlog.Debugf("fetch_minimal_label, x: %v, y: %v, label: %v, val: %v, min: %v",
		x, y, label, equiv_set.Val, equiv_set.Min)
	root := equiv_set.Find()
	return root.Min, true
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

func create_linked_storage(width int, height int) *[]*dset.Dset {
	res := make([]*dset.Dset, width * height)
	return &res
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

func linked_as_string(linked *[]*dset.Dset) string {
	str := ""
	for idx, item := range *linked {
		if item != nil {
			str += fmt.Sprintf("%v: %v(%v), ", idx, (*item).Val, (*item).Min)
		}
	}
	return str
}
