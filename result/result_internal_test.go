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
	expected := []Label{5, 1, 4, 3, -1, -1, -1, -1, -1, -1, -1, -1}
	actual := res1.map_labels(res2)
	if !reflect.DeepEqual(actual, expected) {
		t.Error("map_labels mismatch")
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

