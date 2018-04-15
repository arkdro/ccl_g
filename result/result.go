package result

import (
	"github.com/romana/rlog"

	"log"
)

type Label int

type One_color_result [][]Label

type Result []One_color_result

type Label_error struct {
	x int
	y int
	key Label
	val_old Label
	val_new Label
}

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
	label_map, status, err1 := r1.map_labels(r2)
	if status == false {
		log.Printf("multiple labels, x: %v, y: %v, label1: %v," +
			" old label2: %v, new label2: %v",
			err1.x, err1.y, err1.key, err1.val_old, err1.val_new)
		return false
	}
	label_map2, status2, err2 := r2.map_labels(r1)
	if status2 == false {
		log.Printf("multiple labels, reverse, x: %v, y: %v, label1: %v," +
			" old label2: %v, new label2: %v",
			err2.x, err2.y, err2.key, err2.val_old, err2.val_new)
		return false
	}
	for y, row := range r1 {
		for x, label := range row {
			label2 := r2[y][x]
			if label_map[label] != label2 ||
				label_map2[label2] != label {
				return false
			}
		}
	}
	return true
}

func (r1 One_color_result) map_labels(r2 One_color_result) (map[Label]Label, bool, Label_error) {
	rlog.Debug("map_labels,\nr1: %v\nr2: %v", r1, r2)
	label_map := prepare_label_map()
	for y, row := range r1 {
		for x, label1 := range row {
			label2 := r2[y][x]
			val, found := label_map[label1]
			if !found {
				label_map[label1] = label2
			} else if val != label2 {
				label_error := Label_error{
					x: x,
					y: y,
					key: label1,
					val_old: val,
					val_new: label2,
				}
				return prepare_label_map(), false, label_error
			}
		}
	}
	return label_map, true, Label_error{}
}

func prepare_label_map() map[Label]Label {
	label_map := make(map[Label]Label)
	return label_map
}
