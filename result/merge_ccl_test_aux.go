package result

import (
	"github.com/asdf/ccl_g/color"
)

func Build_expected_g_result_c8() G_result {
	result := map[G_label]G_item{
		"0_1": {"0_1",
			[]G_cell{
				{2, 0},
				{3, 1},
			},
			[]G_label{"1_1", "2_1"},
		},
		"0_2": {"0_2",
			[]G_cell{
				{0, 2},
				{0, 3},
				{1, 3},
				{2, 3},
			},
			[]G_label{"1_1", "2_1"},
		},
		"0_3": {"0_3",
			[]G_cell{
				{6, 2},
				{6, 3},
			},
			[]G_label{"1_2", "2_1"},
		},
		"1_1": {"1_1",
			[]G_cell{
				{0, 0},
				{1, 0},
				{0, 1},
				{1, 1},
				{0, 2},
				{1, 2},
			},
			[]G_label{"0_1", "0_2", "2_1"},
		},
		"1_2": {"1_2",
			[]G_cell{
				{5, 2},
			},
			[]G_label{"0_3", "2_1"},
		},
		"2_1": {"2_1",
			[]G_cell{
				{3, 0},
				{4, 0},
				{5, 0},
				{6, 0},
				{2, 1},
				{4, 1},
				{5, 1},
				{6, 1},
				{3, 2},
				{4, 2},
				{3, 3},
				{4, 3},
				{5, 3},
			},
			[]G_label{"0_1", "0_2", "0_3", "1_1", "1_2"},
		},
	}
	return result
}

func Build_merge_ccl_result() [][]Merged_label {
	width := 7
	height := 4
	data_color_0 := [][]int {
		{0,0,1,0,0,0,0},
		{0,0,0,1,0,0,0},
		{2,0,0,0,0,0,3},
		{2,2,2,0,0,0,3},
	}
	result_color_0 := build_one_color_result(data_color_0)
	data_color_1 := [][]int {
		{1,1,0,0,0,0,0},
		{1,1,0,0,0,0,0},
		{0,1,1,0,0,2,0},
		{0,0,0,0,0,0,0},
	}
	result_color_1 := build_one_color_result(data_color_1)
	data_color_2 := [][]int {
		{0,0,0,1,1,1,1},
		{0,0,1,0,1,1,1},
		{0,0,0,1,1,0,0},
		{0,0,0,1,1,1,0},
	}
	result_color_2 := build_one_color_result(data_color_2)
	result := []One_color_result{
		result_color_0,
		result_color_1,
		result_color_2,
	}
	actual := Merge_ccl_result(width, height, result)
	return actual
}

func Build_merge_ccl_result_2() [][]Merged_label {
	width := 7
	height := 4
	data_color_0 := [][]int {
		{0,0,1,0,0,0,0},
		{0,0,0,1,0,0,0},
		{2,0,0,0,0,0,3},
		{2,2,2,0,0,0,3},
	}
	result_color_0 := build_one_color_result(data_color_0)
	data_color_1 := [][]int {
		{1,1,0,0,0,0,0},
		{1,1,0,0,0,0,0},
		{0,1,1,0,0,2,0},
		{0,0,0,0,0,0,0},
	}
	result_color_1 := build_one_color_result(data_color_1)
	data_color_2 := [][]int {
		{0,0,0,1,1,1,1},
		{0,0,2,0,1,1,1},
		{0,0,0,1,1,0,0},
		{0,0,0,1,1,1,0},
	}
	result_color_2 := build_one_color_result(data_color_2)
	result := []One_color_result{
		result_color_0,
		result_color_1,
		result_color_2,
	}
	actual := Merge_ccl_result(width, height, result)
	return actual
}

func build_one_color_result(data [][]int) One_color_result {
	result := make([][]Label, 0)
	for _, row := range data {
		result_row := make([]Label, 0)
		for _, val := range row {
			label := Label(val)
			result_row = append(result_row, label)
		}
		result = append(result, result_row)
	}
	return result
}

func build_merged_result(data [][][]int) Merged_data {
	result := make([][]Merged_label, 0)
	for _, row := range data {
		result_row := make([]Merged_label, 0)
		for _, val := range row {
			c := color.Color(val[0])
			label := Label(val[1])
			merged_label := Merged_label{c, label}
			result_row = append(result_row, merged_label)
		}
		result = append(result, result_row)
	}
	return result
}
