package ccl6

import (
	"reflect"
	"testing"
)

func Test_ccl1(t *testing.T) {
	width := 7
	height := 4
	color_range := 3
	data := [][]int{
		{ 1,1,0,2,2,2,2},
		{1,1,2,0,2,2,2},
		{ 0,1,1,2,2,1,0},
		{0,0,0,2,2,2,0},
	}
	expected_labels0 := [][]int{
		{0,0,1,0,0,0,0},
		{0,0,0,1,0,0,0},
		{2,0,0,0,0,0,3},
		{2,2,2,0,0,0,3},
	}
	expected_labels1 := [][]int{
		{1,1,0,0,0,0,0},
		{1,1,0,0,0,0,0},
		{0,1,1,0,0,2,0},
		{0,0,0,0,0,0,0},
	}
	expected_labels2 := [][]int{
		{0,0,0,1,1,1,1},
		{0,0,2,0,1,1,1},
		{0,0,0,1,1,0,0},
		{0,0,0,1,1,1,0},
	}
	expected_labels := []*[][]int{
		&expected_labels0,
		&expected_labels1,
		&expected_labels2,
	}
	labels := Ccl(width, height, color_range, &data)
	if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("ccl 1 labels mismatch")
	}
}

func Test_ccl2(t *testing.T) {
	width := 7
	height := 4
	color_range := 1
	data := [][]int{
		{ 7,7,0,7,7,7,7},
		{7,7,7,7,0,7,0},
		{ 7,7,0,0,0,7,0},
		{0,0,0,0,7,0,7},
	}
	expected_labels0 := [][]int{
		{0,0,1,0,0,0,0},
		{0,0,0,0,2,0,3},
		{0,0,2,2,2,0,3},
		{2,2,2,2,0,2,0},
	}
	expected_labels := []*[][]int{
		&expected_labels0,
	}
	labels := Ccl(width, height, color_range, &data)
	if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("ccl 2, labels mismatch")
		t.Error("ccl 2, labels 0: ", labels[0])
	}
}

func Test_ccl3(t *testing.T) {
	width := 5
	height := 4
	color_range := 1
	data := [][]int{
		{ 7,7,7,7,0},
		{0,7,0,0,0},
		{ 0,0,0,7,7},
		{0,0,0,0,7},
	}
	expected_labels0 := [][]int{
		{0,0,0,0,1},
		{1,0,1,1,1},
		{1,1,1,0,0},
		{1,1,1,1,0},
	}
	expected_labels := []*[][]int{
		&expected_labels0,
	}
	labels := Ccl(width, height, color_range, &data)
	if !reflect.DeepEqual(labels, expected_labels) {
		t.Error("ccl 3, labels mismatch")
		t.Error("ccl 3, labels 0: ", labels[0])
	}
}
