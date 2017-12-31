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

