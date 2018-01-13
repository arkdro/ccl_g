package result

import (
	"reflect"
	"testing"
)

func Test_map_labels(t *testing.T) {
	res1 := One_color_result{
		{0, 0, 1, 2},
		{3, 0, 1, 2},
		{0, 3, 3, 2},
	}
	res2 := One_color_result{
		{5, 5, 1, 4},
		{3, 5, 1, 4},
		{5, 3, 3, 4},
	}
	expected := map[Label]Label{0:5, 1:1, 2:4, 3:3}
	actual, status, _ := res1.map_labels(res2)
	if status != true {
		t.Error("map_labels status mismatch")
	} else if !reflect.DeepEqual(actual, expected) {
		t.Error("map_labels data mismatch")
	}
}

func Test_one_color_result_equal(t *testing.T) {
	res1 := One_color_result{
		{4, 4, 1, 2},
		{3, 4, 1, 2},
		{4, 3, 3, 2},
	}
	res2 := One_color_result{
		{5, 5, 1, 4},
		{3, 5, 1, 4},
		{5, 3, 3, 4},
	}
	expected := true
	actual := one_color_result_equal(res1, res2)
	if !reflect.DeepEqual(actual, expected) {
		t.Error("one_color_result_equal mismatch")
	}
}

func Test_one_color_result_equal2(t *testing.T) {
	res1 := One_color_result{
		{4, 4, 1, 2},
		{3, 4, 1, 2},
		{4, 3, 3, 2},
	}
	res2 := One_color_result{
		{5, 5, 1, 4},
		{5, 5, 1, 4},
		{5, 5, 5, 4},
	}
	expected := false
	actual := one_color_result_equal(res1, res2)
	if actual != expected {
		t.Error("one_color_result_equal 2 mismatch")
	}
}

func Test_one_color_result_equal3(t *testing.T) {
	res1 := One_color_result{
		{4, 4, 1, 2},
		{3, 4, 1, 2},
		{4, 3, 3, 2},
	}
	res2 := One_color_result{
		{5, 5, 1, 4},
		{5, 5, 1, 4},
		{5, 5, 5, 4},
	}
	expected := false
	actual := one_color_result_equal(res2, res1)
	if actual != expected {
		t.Error("one_color_result_equal 3 mismatch")
	}
}

func Test_one_color_result_equal4(t *testing.T) {
	res1 := One_color_result{
		{4, 4, 1, 2},
		{3, 4, 1, 2},
		{4, 3, 3, 2},
	}
	res2 := One_color_result{
		{5, 5, 1, 4},
		{3, 5, 1, 4},
		{5, 3, 3, 4},
	}
	expected := true
	actual := one_color_result_equal(res1, res2)
	if !reflect.DeepEqual(actual, expected) {
		t.Error("one_color_result_equal 4 mismatch")
	}
}

