# Connected components labeling

## Parameters

	--infile - input file. Mutually exclusive with --indir
	--indir - input dir. Mutually exclusive with --infile
	--remove - remove input file after successful processing

## Run

### run with a json request:

	ccl_g --infile correct.json

### process all the files in a directory

	ccl_g --indir some_input_dir

## Requests

	A request contains an input data and an expected data.

	The input data contains width, height, color range and a 2d array of integers.

	The expected data is the list of 2d arrays, one array for each color in the color range.
	Every 2d array contains positive integers as labels and negative integers as
	a background.
