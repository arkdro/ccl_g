package coordinates

import (
	"github.com/asdf/ccl_g/cell"
)

func get_neighbour_coordinates_6(width, height, x, y int) []cell.Ccl_cell {
	coord_results := []cell.Ccl_cell{
		get_upper_left_coordinate_6(width, height, x, y),
		get_upper_right_coordinate_6(width, height, x, y),
		get_left_coordinate(width, height, x, y),
		get_right_coordinate(width, height, x, y),
		get_lower_left_coordinate_6(width, height, x, y),
		get_lower_right_coordinate_6(width, height, x, y),
	}
	result := extract_valid_results(width, height, coord_results)
	return result
}

func get_upper_left_coordinate_6(width, height, x, y int) cell.Ccl_cell {
	var x2, y2 int
	y2 = y - 1
	if is_even(y) {
		x2 = x
	} else {
		x2 = x - 1
	}
	return cell.Ccl_cell{X: x2, Y: y2}
}

func get_upper_right_coordinate_6(width, height, x, y int) cell.Ccl_cell {
	var x2, y2 int
	y2 = y - 1
	if is_even(y) {
		x2 = x + 1
	} else {
		x2 = x
	}
	return cell.Ccl_cell{X: x2, Y: y2}
}

func get_lower_left_coordinate_6(width, height, x, y int) cell.Ccl_cell {
	var x2, y2 int
	y2 = y + 1
	if is_even(y) {
		x2 = x
	} else {
		x2 = x - 1
	}
	return cell.Ccl_cell{X: x2, Y: y2}
}

func get_lower_right_coordinate_6(width, height, x, y int) cell.Ccl_cell {
	var x2, y2 int
	y2 = y + 1
	if is_even(y) {
		x2 = x + 1
	} else {
		x2 = x
	}
	return cell.Ccl_cell{X: x2, Y: y2}
}
