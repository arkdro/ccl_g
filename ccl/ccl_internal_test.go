package ccl

import (
	"github.com/asdf/ccl_g/point"

	"reflect"
	"testing"
)

func Test_has_unmarked_background_below1(t *testing.T) {
	color := 1
	x := 1
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	labels := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	actual := has_unmarked_background_below(color, x, y, &data, &labels)
	expected := true
	if actual != expected {
		t.Error("has_unmarked_background_below 1 mismatch")
	}
}

func Test_has_unmarked_background_below2(t *testing.T) {
	color := 1
	x := 1
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 1, 3},
	}
	labels := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	actual := has_unmarked_background_below(color, x, y, &data, &labels)
	expected := false
	if actual != expected {
		t.Error("has_unmarked_background_below 2 mismatch")
	}
}

func Test_has_unmarked_background_below3(t *testing.T) {
	color := 1
	x := 1
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	labels := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 1, 0},
	}
	actual := has_unmarked_background_below(color, x, y, &data, &labels)
	expected := false
	if actual != expected {
		t.Error("has_unmarked_background_below 3 mismatch")
	}
}

func Test_has_background_above1(t *testing.T) {
	color := 1
	x := 1
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	dummy := []int{0, 0, 0}
	actual := has_background_above(color, x, y, &data, &dummy)
	expected := true
	if actual != expected {
		t.Error("has_background_above 1 mismatch")
	}
}

func Test_has_background_above2(t *testing.T) {
	color := 1
	x := 1
	y := 2
	data := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	dummy := make([]int, 3)
	actual := has_background_above(color, x, y, &data, &dummy)
	expected := false
	if actual != expected {
		t.Error("has_background_above 2 mismatch")
	}
}

func Test_has_background_above3(t *testing.T) {
	color := 1
	x := 1
	y := 0
	data := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	dummy := make([]int, 3)
	actual := has_background_above(color, x, y, &data, &dummy)
	expected := true
	if actual != expected {
		t.Error("has_background_above 3 mismatch")
	}
}

func Test_has_background_above4(t *testing.T) {
	color := 1
	x := 1
	y := 0
	data := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	dummy := []int{1, 1, 1} // should not happen
	actual := has_background_above(color, x, y, &data, &dummy)
	expected := false
	if actual != expected {
		t.Error("has_background_above 4 mismatch")
	}
}

func Test_is_background1(t *testing.T) {
	color := 1
	x := 1
	y := 0
	data := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	actual := is_background(color, x, y, &data)
	expected := true
	if actual != expected {
		t.Error("is_background 1 mismatch")
	}
}

func Test_is_background2(t *testing.T) {
	color := 1
	x := 1
	y := 1
	data := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	actual := is_background(color, x, y, &data)
	expected := false
	if actual != expected {
		t.Error("is_background 2 mismatch")
	}
}

func Test_has_no_label1(t *testing.T) {
	x := 1
	y := 1
	labels := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	actual := has_no_label(x, y, &labels)
	expected := false
	if actual != expected {
		t.Error("has_no_label 1 mismatch")
	}
}

func Test_has_no_label2(t *testing.T) {
	x := 1
	y := 1
	labels := [][]int{
		{0, 0, 0},
		{1, 0, 0},
		{1, 2, 3},
	}
	actual := has_no_label(x, y, &labels)
	expected := true
	if actual != expected {
		t.Error("has_no_label 2 mismatch")
	}
}

func Test_has_label1(t *testing.T) {
	x := 1
	y := 1
	labels := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	actual := has_label(x, y, &labels)
	expected := true
	if actual != expected {
		t.Error("has_label 1 mismatch")
	}
}

func Test_has_label2(t *testing.T) {
	x := 1
	y := 1
	labels := [][]int{
		{1, 2, 3},
		{0, 0, 3},
		{1, 2, 3},
	}
	actual := has_label(x, y, &labels)
	expected := false
	if actual != expected {
		t.Error("has_label 2 mismatch")
	}
}

func Test_copy_left_label(t *testing.T) {
	x := 1
	y := 2
	labels := [][]int{
		{1, 2, 3},
		{0, 0, 0},
		{4, 0, 0},
	}
	expected := [][]int{
		{1, 2, 3},
		{0, 0, 0},
		{4, 4, 0},
	}
	copy_left_label(x, y, &labels)
	if !reflect.DeepEqual(labels, expected) {
		t.Error("copy_left_label mismatch")
	}
}

func Test_create_empty_labels(t *testing.T) {
	width := 2
	height := 3
	expected := [][]int{
		{0, 0},
		{0, 0},
		{0, 0},
	}
	actual := create_empty_labels(width, height)
	if !reflect.DeepEqual(actual, expected) {
		t.Error("create_empty_labels mismatch")
	}
}

func Test_update_initial_pair(t *testing.T) {
	initial_pair := []point.Point{
		{2, 3},
		{0, 0},
	}
	new_point := point.Point{X: 5, Y: 3}
	expected := []point.Point{
		{2, 3},
		{5, 3},
	}
	update_initial_pair(&initial_pair, new_point)
	if !reflect.DeepEqual(initial_pair, expected) {
		t.Error("update_initial_pair mismatch")
	}
}

func Test_contour_finished1(t *testing.T) {
	initial_pair := []point.Point{
		{2, 3},
		{0, 0},
	}
	new_point := point.Point{X: 5, Y: 3}
	prev_point := point.Point{X: 4, Y: 4}
	expected := false
	actual := contour_finished(&initial_pair, new_point, prev_point)
	if actual != expected {
		t.Error("contour_finished 1 mismatch")
	}
}

func Test_contour_finished2(t *testing.T) {
	initial_pair := []point.Point{
		{2, 3},
		{1, 1},
	}
	new_point := point.Point{X: 1, Y: 1}
	prev_point := point.Point{X: 2, Y: 3}
	expected := true
	actual := contour_finished(&initial_pair, new_point, prev_point)
	if actual != expected {
		t.Error("contour_finished 2 mismatch")
	}
}

func Test_get_color1(t *testing.T) {
	pt := point.Point{X: 1, Y: 1}
	color := 1
	data := [][]int{
		{1, 2, 3},
		{1, 4, 3},
		{1, 2, 3},
	}
	actual := get_color(pt, &data, color)
	expected := 4
	if actual != expected {
		t.Error("get_color 1 mismatch")
	}
}

func Test_get_color2(t *testing.T) {
	pt := point.Point{X: 1, Y: -1}
	color := 4
	data := [][]int{
		{1, 2, 3},
		{1, 4, 3},
		{1, 2, 3},
	}
	actual := get_color(pt, &data, color)
	expected := 3
	if actual != expected {
		t.Error("get_color 2 mismatch")
	}
}

func Test_same_colors1(t *testing.T) {
	color1 := 4
	color2 := 4
	actual := same_colors(color1, color2)
	expected := true
	if actual != expected {
		t.Error("same_colors 1 mismatch")
	}
}

func Test_same_colors2(t *testing.T) {
	color1 := 4
	color2 := 2
	actual := same_colors(color1, color2)
	expected := false
	if actual != expected {
		t.Error("same_colors 2 mismatch")
	}
}

func Test_mark_background_point1(t *testing.T) {
	pt := point.Point{X: 1, Y: 0}
	data := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	dummy := make([]int, 3)
	expected_data := [][]int{
		{1, -1, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	expected_dummy := make([]int, 3)
	mark_background_point(pt, &data, &dummy)
	if !reflect.DeepEqual(data, expected_data) ||
		!reflect.DeepEqual(dummy, expected_dummy) {
		t.Error("mark_background_point 1 mismatch")
	}
}

func Test_mark_background_point2(t *testing.T) {
	pt := point.Point{X: 1, Y: -1}
	data := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	dummy := make([]int, 3)
	expected_data := [][]int{
		{1, 2, 3},
		{1, 1, 3},
		{1, 2, 3},
	}
	expected_dummy := []int{0, -1, 0}
	mark_background_point(pt, &data, &dummy)
	if !reflect.DeepEqual(data, expected_data) ||
		!reflect.DeepEqual(dummy, expected_dummy) {
		t.Error("mark_background_point 2 mismatch")
	}
}

func Test_next_pos1(t *testing.T) {
	pos := 4
	actual := next_pos(pos)
	expected := 5
	if actual != expected {
		t.Error("next_pos 1 mismatch")
	}
}

func Test_next_pos2(t *testing.T) {
	pos := 7
	actual := next_pos(pos)
	expected := 0
	if actual != expected {
		t.Error("next_pos 2 mismatch")
	}
}

func Test_pos_to_delta(t *testing.T) {
	actual := make([]int, connectivity * 2)
	for i := 0; i < connectivity; i++ {
		dx, dy := pos_to_delta(i)
		actual[i * 2] = dy
		actual[i * 2 + 1] = dx
	}
	expected := []int{
		0, 1,
		1, 1,
		1, 0,
		1, -1,
		0, -1,
		-1, -1,
		-1, 0,
		-1, 1,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("pos_to_delta mismatch")
	}
}

func Test_get_neighbour_coord(t *testing.T) {
	actual := make([]point.Point, connectivity)
	for i := 0; i < connectivity; i++ {
		pt := point.Point{
			X: i * 3 - 1,
			Y: i * 5 + 1}
		point2 := get_neighbour_coord(pt, i)
		actual[i] = point2
	}
	expected := []point.Point{
		{X: 0, Y: 1},
		{X: 3, Y: 7},
		{X: 5, Y: 12},
		{X: 7, Y: 17},
		{X: 10, Y: 21},
		{X: 13, Y: 25},
		{X: 17, Y: 30},
		{X: 21, Y: 35},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("get_neighbour_coord mismatch")
	}
}

func Test_prev_point_pos(t *testing.T) {
	actual := make([]int, connectivity)
	for i := 0; i < connectivity; i++ {
		pos := prev_point_pos(i)
		actual[i] = pos
	}
	expected := []int{4, 5, 6, 7, 0, 1, 2, 3}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("prev_point_pos mismatch")
	}
}

func Test_calc_next_pos(t *testing.T) {
	actual := make([]int, connectivity)
	for i := 0; i < connectivity; i++ {
		pos := calc_next_pos(i)
		actual[i] = pos
	}
	expected := []int{6, 7, 0, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("calc_next_pos mismatch")
	}
}

func Test_tracer1(t *testing.T) {
	pt := point.Point{X: 1, Y: 1}
	color := 4
	label := 1
	init_pos := 7
	data := [][]int{
		{1, 2, 3},
		{1, 4, 3},
		{1, 2, 3},
	}
	labels := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	dummy := make([]int, 3)
	dummy_labels := make([]int, 3)
	expected_data := [][]int{
		{-1, -1, -1},
		{-1, 4, -1},
		{-1, -1, -1},
	}
	expected_dummy := []int{0, 0, 0}
	expected_labels := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	expected_dummy_labels := []int{0, 0, 0}
	expected_point := point.Point{}
	point2, pos2, status := tracer(color, pt, label, &data, &dummy, &labels, &dummy_labels, init_pos)
	if point2 != expected_point {
		t.Error("tracer 1 point mismatch")
	} else if pos2 != 0 {
		t.Error("tracer 1 pos mismatch")
	} else if status != false {
		t.Error("tracer 1 status mismatch")
	} else if !reflect.DeepEqual(data, expected_data) {
		t.Error("tracer 1 data mismatch")
	} else if !reflect.DeepEqual(dummy, expected_dummy) {
		t.Error("tracer 1 dummy mismatch")
	} else if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("tracer 1 labels mismatch")
	} else if !reflect.DeepEqual(dummy_labels, expected_dummy_labels) {
		t.Error("tracer 1 dummy labels mismatch")
	}
}

func Test_tracer2(t *testing.T) {
	pt := point.Point{X: 1, Y: 1}
	color := 4
	label := 1
	init_pos := 7
	data := [][]int{
		{1, 2, 3},
		{1, 4, 3},
		{4, 2, 3},
	}
	labels := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	dummy := make([]int, 3)
	dummy_labels := make([]int, 3)
	expected_data := [][]int{
		{1, 2, -1},
		{1, 4, -1},
		{4, -1, -1},
	}
	expected_dummy := []int{0, 0, 0}
	expected_labels := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
	}
	expected_dummy_labels := []int{0, 0, 0}
	expected_point := point.Point{X: 0, Y: 2}
	point2, pos2, status := tracer(color, pt, label, &data, &dummy, &labels, &dummy_labels, init_pos)
	if point2 != expected_point {
		t.Error("tracer 2 point mismatch")
	} else if pos2 != 3 {
		t.Error("tracer 2 pos mismatch")
	} else if status != true {
		t.Error("tracer 2 status mismatch")
	} else if !reflect.DeepEqual(data, expected_data) {
		t.Error("tracer 2 data mismatch")
	} else if !reflect.DeepEqual(dummy, expected_dummy) {
		t.Error("tracer 2 dummy mismatch")
	} else if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("tracer 2 labels mismatch")
	} else if !reflect.DeepEqual(dummy_labels, expected_dummy_labels) {
		t.Error("tracer 2 dummy labels mismatch")
	}
}

func Test_external_contour_tracing(t *testing.T) {
	x := 1
	y := 0
	color := 1
	label := 3
	dummy := []int{0, 0, 0, 0, 0, 0, 0, 0}
	data := [][]int{
		{0, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	dummy_labels := []int{0, 0, 0, 0, 0, 0, 0, 0}
	labels := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected_dummy := []int{-1, -1, -1, -1, -1, -1, -1, 0}
	expected_data := [][]int{
		{-1, 1, 1, 1, 1, 1, -1, -1},
		{-1, 1, 0, 0, 0, 1, 1, -1},
		{-1, 1, 1, 0, 0, 1, 1, -1},
		{-1, 1, 1, 1, 1, 1, 1, -1},
		{-1, 1, 1, 1, 1, 1, 1, -1},
		{-1, 1, 1, 1, 1, 1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, 0},
	}
	expected_dummy_labels := []int{0, 0, 0, 0, 0, 0, 0, 0}
	expected_labels := [][]int{
		{0, 3, 3, 3, 3, 3, 0, 0},
		{0, 3, 0, 0, 0, 0, 3, 0},
		{0, 3, 0, 0, 0, 0, 3, 0},
		{0, 3, 0, 0, 0, 0, 3, 0},
		{0, 3, 0, 0, 0, 0, 3, 0},
		{0, 3, 3, 3, 3, 3, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	external_contour_tracing(color, x, y, label, &data, &dummy, &labels, &dummy_labels)
	if !reflect.DeepEqual(data, expected_data) {
		t.Error("external_contour_tracing data mismatch")
	} else if !reflect.DeepEqual(dummy, expected_dummy) {
		t.Error("external_contour_tracing dummy mismatch")
	} else if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("external_contour_tracing labels mismatch")
	} else if !reflect.DeepEqual(dummy_labels, expected_dummy_labels) {
		t.Error("external_contour_tracing dummy labels mismatch")
	}
}

func Test_external_contour_tracing2(t *testing.T) {
	x := 1
	y := 0
	color := 1
	label := 3
	dummy := []int{0, 0, 0, 0, 0, 0, 0, 0}
	data := [][]int{
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	dummy_labels := []int{0, 0, 0, 0, 0, 0, 0, 0}
	labels := [][]int{
		{0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected_dummy := []int{-1, -1, -1, 0, 0, 0, 0, 0}
	expected_data := [][]int{
		{-1, 1, -1, 0, 0, 0, 0, 0},
		{-1, -1, -1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected_dummy_labels := []int{0, 0, 0, 0, 0, 0, 0, 0}
	expected_labels := [][]int{
		{0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	external_contour_tracing(color, x, y, label, &data, &dummy, &labels, &dummy_labels)
	if !reflect.DeepEqual(data, expected_data) {
		t.Error("external_contour_tracing 2 data mismatch")
	} else if !reflect.DeepEqual(dummy, expected_dummy) {
		t.Error("external_contour_tracing 2 dummy mismatch")
	} else if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("external_contour_tracing 2 labels mismatch")
	} else if !reflect.DeepEqual(dummy_labels, expected_dummy_labels) {
		t.Error("external_contour_tracing 2 dummy labels mismatch")
	}
}

