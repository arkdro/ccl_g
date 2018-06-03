package coordinates

import (
	"github.com/asdf/ccl_g/cell"
)

func extract_valid_results(width int, height int, coord_results []cell.Ccl_cell) []cell.Ccl_cell {
	result := []cell.Ccl_cell{}
	for _, cell := range coord_results {
		if is_valid(width, height, cell.X, cell.Y) {
			result = append(result, cell)
		}
	}
	return result
}

func is_even(n int) bool {
	return n % 2 == 0
}

func is_valid(width, height, x, y int) bool {
	if y < 0 {
		return false
	}
	if y >= height {
		return false
	}
	if x < 0 {
		return false
	}
	if x >= width {
		return false
	}
	return true
}

func get_left_coordinate(width, height, x, y int) cell.Ccl_cell {
	var x2, y2 int
	y2 = y
	x2 = x - 1
	return cell.Ccl_cell{X: x2, Y: y2}
}

func get_right_coordinate(width, height, x, y int) cell.Ccl_cell {
	var x2, y2 int
	y2 = y
	x2 = x + 1
	return cell.Ccl_cell{X: x2, Y: y2}
}

func Get_neighbour_coordinates(connectivity, width, height, x, y int) []cell.Ccl_cell {
	switch connectivity {
	case 6:
		return get_neighbour_coordinates_6(width, height, x, y)
	case 8:
		return get_neighbour_coordinates_8(width, height, x, y)
	default:
		return []cell.Ccl_cell{}
	}
}

