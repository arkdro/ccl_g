package result

import (
	"reflect"
	"testing"
)

func Test_merge_ccl(t *testing.T) {
	height := 4
	actual := Build_merge_ccl_result()
	expected_data := [][][]int {
		{{1,1}, {1,1}, {0,1}, {2,1}, {2,1}, {2,1}, {2,1}},
		{{1,1}, {1,1}, {2,1}, {0,1}, {2,1}, {2,1}, {2,1}},
		{{0,2}, {1,1}, {1,1}, {2,1}, {2,1}, {1,2}, {0,3}},
		{{0,2}, {0,2}, {0,2}, {2,1}, {2,1}, {2,1}, {0,3}},
	}
	expected := build_merged_result(expected_data)
	for y := 0; y < height; y++ {
		if !reflect.DeepEqual(actual[y], expected[y]) {
			t.Error("merge ccl result mismatch for y = ", y)
			t.Error("expected:", expected[y])
			t.Error("actual:", actual[y])
		}
	}
}
