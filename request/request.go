package request

import (
	"github.com/romana/rlog"

	"github.com/asdf/ccl_g/ccl"
	"github.com/asdf/ccl_g/ccl6"
	"github.com/asdf/ccl_g/plate"
	"github.com/asdf/ccl_g/result"

	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Request struct {
	Input_data plate.Plate
	Expected_data result.Result
}

func Run(file string, dir string, remove bool, connectivity int) {
	files := make([]string, 0)
	if file != "" {
		files = append(files, file)
	} else if dir != "" {
		files = get_files_in_dir(dir)
	} else {
		log.Printf("No parameters given")
		return
	}
	process_files(files, remove, connectivity)
}

func get_files_in_dir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("get_files_in_dir, read dir error: %v", err)
		files := make([]string, 0)
		return files
	}
	names := make([]string, 0)
	for _, file := range files {
		fname := path.Join(dir, file.Name())
		names = append(names, fname)
	}
	return names
}

func process_files(files []string, remove bool, connectivity int) {
	for _, file := range files {
		rlog.Info("process_files, file:", file)
		process_one_file(file, remove, connectivity)
	}
}

func process_one_file(file string, remove bool, connectivity int) {
	request, err := read_request(file)
	if err != nil {
		log.Printf("process_one_file, can't read request for file '%v': %v",
			file, err)
		return
	}
	result := run_request(request, connectivity)
	if !results_equal(result, request.Expected_data, request.Input_data.Color_range) {
		rlog.Error("process_one_file, result mismatch, file:", file)
		rlog.Warn("result:", result, "\nexpected:", request.Expected_data)
		write_result(file, result)
	} else {
		if remove {
			os.Remove(file)
		}
	}
}

func read_request(file string) (Request, error) {
	rlog.Info("read_request, file:", file)
	var request Request
	bin_data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("can't read file '%v': %v", file, err)
		return request, err
	}
	err = json.Unmarshal(bin_data, &request)
	if err != nil {
		log.Printf("can't parse file '%v': %v", file, err)
		return request, err
	}
	if !request.Input_data.Valid_data() {
		log.Printf("invalid plate data in file: %v", file)
		return request, errors.New("invalid plate data")
	}
	if !request.Expected_data.Valid_data(request.Input_data.Width, request.Input_data.Height, request.Input_data.Color_range) {
		log.Printf("invalid expected data in file: %v", file)
		return request, errors.New("invalid expected data")
	}
	return request, nil
}

func results_equal(actual result.Result, expected result.Result, color_range int) bool {
	return result.Equal(actual, expected, color_range)
}

func write_result(file string, result result.Result) {
	rlog.Info("write_result, file:", file)
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("can't encode result to json for file '%v': %v\n%+v",
			file, err, result)
		return
	}
	fname := file + "-result"
	ioutil.WriteFile(fname, data, 0644)
}

func run_request(request Request, connectivity int) result.Result {
	var raw_res []*[][]int
	if connectivity == 6 {
		raw_res = ccl6.Ccl(request.Input_data.Width, request.Input_data.Height,
			request.Input_data.Color_range, &request.Input_data.Data)
	} else {
		raw_res = ccl.Ccl(request.Input_data.Width, request.Input_data.Height,
			request.Input_data.Color_range, &request.Input_data.Data)
	}
	res := result.Build_result(raw_res)
	return res
}

