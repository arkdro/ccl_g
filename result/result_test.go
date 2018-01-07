package result

import (
	"testing"
)

func Test_valid_data1(t *testing.T) {
	width := 4
	height := 3
	color_range := 2
	res1 := One_color_result{
		{0, 0, 1, 1},
		{3, 0, 0, 0},
		{0, 3, 3, 0},
	}
	res2 := One_color_result{
		{1, 1, 0, 0},
		{0, 1, 1, 1},
		{1, 0, 0, 1},
	}
	res := Result{
		res1,
		res2,
	}
	expected := true
	actual := res.Valid_data(width, height, color_range)
	if actual != expected {
		t.Error("valid data 1 mismatch")
	}
}

func Test_valid_data2(t *testing.T) {
	width := 4
	height := 3
	color_range := 2
	res1 := One_color_result{
		{0, 0, 1, 1},
		{3, 0, 0, 0},
		{0, 3, 3, 0},
	}
	res := Result{
		res1,
	}
	expected := false
	actual := res.Valid_data(width, height, color_range)
	if actual != expected {
		t.Error("valid data 2 mismatch")
	}
}

func Test_valid_data3(t *testing.T) {
	width := 4
	height := 3
	color_range := 2
	res1 := One_color_result{
		{0, 0, 1, 1},
		{3, 0, 0, 0},
		{0, 3, 3, 0},
	}
	res2 := One_color_result{
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}
	res := Result{
		res1,
		res2,
	}
	expected := false
	actual := res.Valid_data(width, height, color_range)
	if actual != expected {
		t.Error("valid data 3 mismatch")
	}
}

func Test_valid_data4(t *testing.T) {
	width := 4
	height := 3
	color_range := 2
	res1 := One_color_result{
		{0, 0, 1, 1, 1},
		{3, 0, 0, 0, 1},
		{0, 3, 3, 0, 0},
	}
	res2 := One_color_result{
		{1, 1, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 1, 1},
	}
	res := Result{
		res1,
		res2,
	}
	expected := false
	actual := res.Valid_data(width, height, color_range)
	if actual != expected {
		t.Error("valid data 4 mismatch")
	}
}

func Test_equal1(t *testing.T) {
	color_range := 2
	res1a := One_color_result{
		{0, 0, 1, 1},
		{2, 0, 0, 0},
		{0, 2, 2, 0},
	}
	res1b := One_color_result{
		{4, 4, 0, 0},
		{0, 4, 4, 4},
		{4, 0, 0, 4},
	}
	r1 := Result{
		res1a,
		res1b,
	}
	res2a := One_color_result{
		{0, 0, 4, 4},
		{3, 0, 0, 0},
		{0, 3, 3, 0},
	}
	res2b := One_color_result{
		{1, 1, 0, 0},
		{0, 1, 1, 1},
		{1, 0, 0, 1},
	}
	r2 := Result{
		res2a,
		res2b,
	}
	expected := true
	actual := Equal(r1, r2, color_range)
	if actual != expected {
		t.Error("equal 1 mismatch")
	}
}

func Test_equal2(t *testing.T) {
	color_range := 2
	res1a := One_color_result{
		{0, 0, 1, 1},
		{2, 0, 0, 0},
		{0, 2, 2, 0},
	}
	res1b := One_color_result{
		{1, 1, 0, 0},
		{0, 1, 1, 1},
		{1, 0, 0, 1},
	}
	r1 := Result{
		res1a,
		res1b,
	}
	res2a := One_color_result{
		{0, 0, 2, 2},
		{3, 0, 0, 0},
		{0, 3, 3, 0},
	}
	res2b := One_color_result{
		{4, 4, 0, 0},
		{4, 4, 4, 4},
		{4, 0, 0, 4},
	}
	r2 := Result{
		res2a,
		res2b,
	}
	expected := false
	actual := Equal(r1, r2, color_range)
	if actual != expected {
		t.Error("equal 2 mismatch")
	}
}

