package result

import (
	"log"
)

type Label int

type One_color_result [][]Label

type Result []One_color_result

func Equal(r1 Result, r2 Result, color_range int) bool {
	for i := 0; i < color_range; i++ {
		res := one_color_result_equal(r1[i], r2[i])
		if !res {
			return false
		}
	}
	return true
}

func Build_result(data []*[][]int) Result {
	res := make(Result, 0)
	for _, item := range data {
		one_color_result := make([][]Label, 0)
		for _, row := range *item {
			label_row := make([]Label, 0)
			for _, point := range row {
				label := Label(point)
				label_row = append(label_row, label)
			}
			one_color_result = append(one_color_result, label_row)
		}
		res = append(res, one_color_result)
	}
	return res
}

func (r Result) Valid_data(width int, height int, color_range int) bool {
	if len(r) != color_range {
		return false
	}
	for _, item := range r {
		res := valid_item(width, height, item)
		if res == false {
			return false
		}
	}
	return true
}

func valid_item(width int, height int, item One_color_result) bool {
	if len(item) != height {
		return false
	}
	for _, row := range item {
		if len(row) != width {
			return false
		}
	}
	return true
}

func one_color_result_equal(r1 One_color_result, r2 One_color_result) bool {
	label_map := r1.map_labels(r2)
	for _, row := range r1 {
		for _, label := range row {
			if label != label_map[label] {
				return false
			}
		}
	}
	return true
}

func (r1 One_color_result) map_labels(r2 One_color_result) []Label {
	height := len(r1)
	width := len(r1[0])
	label_map := prepare_label_map(width, height)
	for y, row := range r1 {
		for x, label1 := range row {
			label2 := r2[y][x]
			if label_map[label1] == -1 {
				label_map[label1] = label2
			} else if label_map[label1] != label2 {
				log.Fatalf("multiple labels, x: %v, y: %v, label1: %v," +
					" old label2: %v, new label2: %v",
					x, y, label1, label_map[label1], label2)
			}
		}
	}
	return label_map
}

func prepare_label_map(width int, height int) []Label {
	label_map := make([]Label, width * height)
	for i := range label_map {
		label_map[i] = -1
	}
	return label_map
}
