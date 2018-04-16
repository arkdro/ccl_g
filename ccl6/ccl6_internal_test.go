package ccl6

import (
	"github.com/asdf/ccl_g/point"
//	"github.com/asdf/ccl_g/color"

	"reflect"
	"testing"
)

func Test_get_left_coordinate1(t *testing.T) {
	x := 0
	y := 2
	_, status := get_left_coordinate(x, y)
	expected := false
	if status != expected {
		t.Error("get_left_coordinate 1 mismatch")
	}
}

func Test_get_left_coordinate2(t *testing.T) {
	x := 1
	y := 2
	actual, status := get_left_coordinate(x, y)
	expected := point.Point{0, 2}
	if status != true || actual != expected {
		t.Error("get_left_coordinate 2 mismatch")
	}
}

func Test_get_left_coordinate3(t *testing.T) {
	x := 1
	y := -1
	_, status := get_left_coordinate(x, y)
	expected := false
	if status != expected {
		t.Error("get_left_coordinate 3 mismatch")
	}
}

func Test_get_upper_left_coordinate1(t *testing.T) {
	x := 1
	y := 0
	_, status := get_upper_left_coordinate(x, y)
	expected := false
	if status != expected {
		t.Error("get_upper_left_coordinate 1 mismatch")
	}
}

func Test_get_upper_left_coordinate2(t *testing.T) {
	x := 0
	y := 1
	_, status := get_upper_left_coordinate(x, y)
	expected := false
	if status != expected {
		t.Error("get_upper_left_coordinate 2 mismatch")
	}
}

func Test_get_upper_left_coordinate3(t *testing.T) {
	x := 1
	y := 1
	pt, status := get_upper_left_coordinate(x, y)
	expected := point.Point{0, 0}
	if status != true || pt != expected {
		t.Error("get_upper_left_coordinate 3 mismatch")
	}
}

func Test_get_upper_left_coordinate4(t *testing.T) {
	x := 1
	y := 2
	pt, status := get_upper_left_coordinate(x, y)
	expected := point.Point{1, 1}
	if status != true || pt != expected {
		t.Error("get_upper_left_coordinate 4 mismatch")
	}
}

func Test_get_upper_right_coordinate1(t *testing.T) {
	x := 1
	y := 0
	width := 3
	_, status := get_upper_right_coordinate(width, x, y)
	expected := false
	if status != expected {
		t.Error("get_upper_right_coordinate 1 mismatch")
	}
}

func Test_get_upper_right_coordinate2(t *testing.T) {
	x := 1
	y := 1
	width := 3
	pt, status := get_upper_right_coordinate(width, x, y)
	expected := point.Point{1, 0}
	if status != true || pt != expected {
		t.Error("get_upper_right_coordinate 2 mismatch")
	}
}

func Test_get_upper_right_coordinate3(t *testing.T) {
	x := 1
	y := 2
	width := 3
	pt, status := get_upper_right_coordinate(width, x, y)
	expected := point.Point{2, 1}
	if status != true || pt != expected {
		t.Error("get_upper_right_coordinate 3 mismatch")
	}
}

func Test_get_upper_right_coordinate4(t *testing.T) {
	x := 2
	y := 2
	width := 3
	_, status := get_upper_right_coordinate(width, x, y)
	expected := false
	if status != expected {
		t.Error("get_upper_right_coordinate 4 mismatch")
	}
}

func Test_get_upper_right_coordinate5(t *testing.T) {
	x := 2
	y := 1
	width := 3
	pt, status := get_upper_right_coordinate(width, x, y)
	expected := point.Point{2, 0}
	if status != true || pt != expected {
		t.Error("get_upper_right_coordinate 5 mismatch")
	}
}

func Test_same_color_neigbours1(t *testing.T) {
	color := 1
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
	color := 1
	x := 1
	y := 1
	data := [][]int{
		{1, 1, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	width := 3
	actual := same_color_neigbours(width, color, x, y, &data)
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

func Test_no_neigbours1(t *testing.T) {
	actual := no_neigbours([]point.Point{})
	expected := true
	if actual != expected {
		t.Error("no_neigbours 1 mismatch")
	}
}

func Test_no_neigbours2(t *testing.T) {
	actual := no_neigbours([]point.Point{{1,1}})
	expected := false
	if actual != expected {
		t.Error("no_neigbours 2 mismatch")
	}
}

func Test_find_minimal_label(t *testing.T) {
	actual := find_minimal_item([]int{10, 5, 2, 2, 1, 3})
	expected := 1
	if actual != expected {
		t.Error("find_minimal_label mismatch")
	}
}

func Test_find_neigbour_labels(t *testing.T) {
	neigbours := []point.Point{
		{1, 2},
		{2, 3},
	}
	data := [][]int{
		{ 1,  2,  3},
		{11, 21, 31},
		{12, 22, 32},
		{13, 23, 33},
	}
	actual := find_neigbour_labels(neigbours, &data)
	if len(actual) != 2 || actual[0] != 22 || actual[1] != 33 {
		t.Error("find_neigbour_labels mismatch")
	}
}

func Test_set_equivalence(t *testing.T) {
	label1 := 1
	label2 := 2
	labels := []int{label1, label2}
	width := 3
	height := 3
	linked := create_linked_storage(width, height)
	set1 := init_label_set(label1)
	set2 := init_label_set(label2)
	(*linked)[label1] = set1
	(*linked)[label2] = set2
	set_equivalence(labels, linked)
	set1 = (*linked)[1]
	if set1.Val != 1 {
		t.Error("set_equivalence set1 mismatch")
	}
	set2 = (*linked)[2]
	if set2.Val != 1 {
		t.Error("set_equivalence set2 mismatch")
	}
}

func Test_fetch_minimal_label(t *testing.T) {
	label1 := 1
	label2 := 4
	labels := [][]int{
		{0, 1, 1},
		{2, 1, 0},
		{2, 1, 1},
		}
	linked := make(map[int]map[int]bool)
	set1 := init_label_set(label1)
	set2 := init_label_set(label2)
	set1[label2] = true
	linked[label1] = set1
	linked[label2] = set2
	x := 1
	y := 1
	label, status := fetch_minimal_label(x, y, &labels, &linked)
	if status != true {
		t.Error("fetch_minimal_label status mismatch")
	}
	if label != 1 {
		t.Error("fetch_minimal_label label mismatch")
	}
}

func Test_fetch_minimal_label2(t *testing.T) {
	label1 := 1
	label2 := 4
	labels := [][]int{
		{0, 1, 1},
		{2, 1, 0},
		{2, 1, 1},
		}
	linked := make(map[int]map[int]bool)
	set1 := init_label_set(label1)
	set2 := init_label_set(label2)
	set1[label2] = true
	linked[label1] = set1
	linked[label2] = set2
	x := 2
	y := 1
	_, status := fetch_minimal_label(x, y, &labels, &linked)
	if status != false {
		t.Error("fetch_minimal_label, 2, status mismatch")
	}
}

func Test_ccl_pass1_1(t *testing.T) {
	width := 3
	height := 3
	color := 1
	data := [][]int{
		{1, 2, 1},
		{4, 1, 3},
		{1, 2, 1},
	}
	labels := create_empty_labels(width, height)
	linked := make(map[int]map[int]bool)
	expected_labels := [][]int{
		{1, 0, 2},
		{0, 1, 0},
		{1, 0, 3},
	}
	expected_linked := make(map[int]map[int]bool)
	set1 := make(map[int]bool)
	set1[1] = true
	set2 := make(map[int]bool)
	set2[2] = true
	set3 := make(map[int]bool)
	set3[3] = true
	expected_linked[1] = set1
	expected_linked[2] = set2
	expected_linked[3] = set3
	ccl_pass1(width, height, color, &data, &labels, &linked)
	if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("ccl_pass1, 1, labels mismatch")
	}
}

func Test_ccl_pass1_2(t *testing.T) {
	width := 3
	height := 3
	color := 1
	data := [][]int{
		{1, 2, 1},
		{4, 1, 1},
		{1, 2, 1},
	}
	labels := create_empty_labels(width, height)
	linked := make(map[int]map[int]bool)
	expected_labels := [][]int{
		{1, 0, 2},
		{0, 1, 1},
		{1, 0, 1},
	}
	expected_linked := make(map[int]map[int]bool)
	set1 := make(map[int]bool)
	set1[1] = true
	set1[2] = true
	set1[3] = true
	set2 := make(map[int]bool)
	set2[1] = true
	set2[2] = true
	set2[3] = true
	set3 := make(map[int]bool)
	set3[1] = true
	set3[2] = true
	set3[3] = true
	expected_linked[1] = set1
	expected_linked[2] = set2
	expected_linked[3] = set3
	ccl_pass1(width, height, color, &data, &labels, &linked)
	if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("ccl_pass1, 2, labels mismatch")
	}
}

func Test_ccl_pass2_1(t *testing.T) {
	width := 3
	height := 3
	color := 1
	data := [][]int{
		{1, 2, 1},
		{4, 1, 1},
		{1, 2, 1},
	}
	labels := [][]int{
		{1, 0, 2},
		{0, 1, 1},
		{1, 0, 1},
	}
	linked := make(map[int]map[int]bool)
	set1 := make(map[int]bool)
	set1[1] = true
	set1[2] = true
	set1[3] = true
	set2 := make(map[int]bool)
	set2[1] = true
	set2[2] = true
	set2[3] = true
	set3 := make(map[int]bool)
	set3[1] = true
	set3[2] = true
	set3[3] = true
	linked[1] = set1
	linked[2] = set2
	linked[3] = set3
	expected_labels := [][]int{
		{1, 0, 1},
		{0, 1, 1},
		{1, 0, 1},
	}
	ccl_pass2(width, height, color, &data, &labels, &linked)
	if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("ccl_pass2, 1, labels mismatch")
	}
}

func Test_ccl_one_color1(t *testing.T) {
	width := 8
	height := 7
	color := 1
	data := [][]int{
		{ 0, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{ 0, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{ 0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0, 0},
		{ 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected_labels := [][]int{
		{ 0, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{ 0, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{ 0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0, 0},
		{ 0, 0, 0, 0, 0, 0, 0, 0},
	}
	labels := ccl_one_color(width, height, color, &data)
	if !reflect.DeepEqual(*labels, expected_labels) {
		t.Error("ccl_one_color 1 labels mismatch")
	}
}

