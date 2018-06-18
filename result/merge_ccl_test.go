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

func Test_g_to_merged_label(t *testing.T) {
	actual := G_to_merged_label("2_4")
	expected := Merged_label{2, 4}
	if actual != expected {
		t.Error("merged label mismatch, 1")
	}
	actual = G_to_merged_label("02_04")
	expected = Merged_label{2, 4}
	if actual != expected {
		t.Error("merged label mismatch, 2")
	}
	actual = G_to_merged_label("202_304")
	expected = Merged_label{202, 304}
	if actual != expected {
		t.Error("merged label mismatch, 3")
	}
}
