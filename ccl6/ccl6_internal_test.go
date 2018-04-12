package ccl6

import (
	"github.com/asdf/ccl_g/point"
//	"github.com/asdf/ccl_g/color"

//	"reflect"
	"testing"
)

func Test_get_left_point1(t *testing.T) {
	color := color.Color(1)
	x := 0
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	_, status := get_left_point(color, x, y, &data)
	expected := false
	if status != expected {
		t.Error("get_left_point 1 mismatch")
	}
}

func Test_get_left_point2(t *testing.T) {
	color := color.Color(1)
	x := 1
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	actual, status := get_left_point(color, x, y, &data)
	expected := point.Point{0, 2}
	if status != true || actual != expected {
		t.Error("get_left_point 2 mismatch")
	}
}

func Test_get_left_point3(t *testing.T) {
	color := color.Color(1)
	x := 1
	y := -1
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	_, status := get_left_point(color, x, y, &data)
	expected := false
	if status != expected {
		t.Error("get_left_point 3 mismatch")
	}
}

func Test_get_left_point4(t *testing.T) {
	color := color.Color(4)
	x := 1
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	_, status := get_left_point(color, x, y, &data)
	expected := false
	if status != expected {
		t.Error("get_left_point 4 mismatch")
	}
}

func Test_get_upper_left_point1(t *testing.T) {
	color := color.Color(4)
	x := 1
	y := 0
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	_, status := get_upper_left_point(color, x, y, &data)
	expected := false
	if status != expected {
		t.Error("get_upper_left_point 1 mismatch")
	}
}

func Test_get_upper_left_point2(t *testing.T) {
	color := color.Color(4)
	x := 0
	y := 1
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	_, status := get_upper_left_point(color, x, y, &data)
	expected := false
	if status != expected {
		t.Error("get_upper_left_point 2 mismatch")
	}
}

func Test_get_upper_left_point3(t *testing.T) {
	color := color.Color(4)
	x := 1
	y := 1
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	pt, status := get_upper_left_point(color, x, y, &data)
	expected := point.Point{0, 0}
	if status != true || pt != expected {
		t.Error("get_upper_left_point 3 mismatch")
	}
}

func Test_get_upper_left_point4(t *testing.T) {
	color := color.Color(4)
	x := 1
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	pt, status := get_upper_left_point(color, x, y, &data)
	expected := point.Point{1, 1}
	if status != true || pt != expected {
		t.Error("get_upper_left_point 4 mismatch")
	}
}

func Test_get_upper_right_point1(t *testing.T) {
	color := color.Color(4)
	x := 1
	y := 0
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	width := 3
	_, status := get_upper_right_point(width, color, x, y, &data)
	expected := false
	if status != expected {
		t.Error("get_upper_right_point 1 mismatch")
	}
}

func Test_get_upper_right_point2(t *testing.T) {
	color := color.Color(4)
	x := 1
	y := 1
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	width := 3
	pt, status := get_upper_right_point(width, color, x, y, &data)
	expected := point.Point{1, 0}
	if status != true || pt != expected {
		t.Error("get_upper_right_point 2 mismatch")
	}
}

func Test_get_upper_right_point3(t *testing.T) {
	color := color.Color(3)
	x := 1
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	width := 3
	pt, status := get_upper_right_point(width, color, x, y, &data)
	expected := point.Point{2, 1}
	if status != true || pt != expected {
		t.Error("get_upper_right_point 3 mismatch")
	}
}

func Test_get_upper_right_point4(t *testing.T) {
	color := color.Color(3)
	x := 2
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	width := 3
	_, status := get_upper_right_point(width, color, x, y, &data)
	expected := false
	if status != expected {
		t.Error("get_upper_right_point 4 mismatch")
	}
}

func Test_get_upper_right_point5(t *testing.T) {
	color := color.Color(3)
	x := 2
	y := 1
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	width := 3
	pt, status := get_upper_right_point(width, color, x, y, &data)
	expected := point.Point{2, 0}
	if status != true || pt != expected {
		t.Error("get_upper_right_point 5 mismatch")
	}
}

func Test_same_color_neigbours1(t *testing.T) {
	color := 3
	x := 0
	y := 1
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	width := 3
	actual := same_color_neigbours(width, color, x, y, &data)
	expected := point.Point{0, 0}
	if len(actual) != 1 || actual[0] != expected {
		t.Error("same_color_neigbours 1 mismatch")
	}
}

func Test_same_color_neigbours2(t *testing.T) {
	color := 3
	x := 1
	y := 1
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	width := 3
	actual := same_color_neigbours(width, color, x, y, &data)
	t.Log(actual)
	expected1 := point.Point{0, 0}
	expected2 := point.Point{1, 0}
	expected0 := point.Point{0, 1}
	if len(actual) != 3 ||
		actual[0] != expected0 ||
		actual[1] != expected1 ||
		actual[2] != expected2 {
		t.Error("same_color_neigbours 2 mismatch")
	}
}

