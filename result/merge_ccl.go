package result

import (
	"github.com/asdf/ccl_g/color"
)

type Merged_label struct {
	c color.Color
	l Label
}

type Merged_data [][]Merged_label

func init_2d_array(width int, height int) Merged_data {
	data := make([][]Merged_label, height)
	for y := 0; y < height; y++ {
		row := make([]Merged_label, width)
		for x := 0; x < width; x++ {
			label := Merged_label{color.Color(-1), Label(-1)}
			row[x] = label
		}
		data[y] = row
	}
	return data
}

func Merge_ccl_result(width int, height int, result Result) [][]Merged_label {
	merged := init_2d_array(width, height)
	for c, color_result := range result {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				v := color_result[y][x]
				if v != 0 {
					label := Merged_label{color.Color(c), Label(v)}
					merged[y][x] = label
				}
			}
		}
	}
	return merged
}

func Make_label(c int, l int) Merged_label {
	return Merged_label{color.Color(c), Label(l)}
}
