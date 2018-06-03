package result

import (
	"github.com/asdf/ccl_g/color"
)

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
