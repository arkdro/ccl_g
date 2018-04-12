package ccl6

import (
	"github.com/asdf/ccl_g/point"
	"github.com/asdf/ccl_g/color"

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
	if status != true && actual != expected {
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

