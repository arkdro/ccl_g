package dset

type Dset struct {
	Parent *Dset
	Min int
	Rank int
	Size int
	Val int
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Create(x int) *Dset {
	var item Dset
	item = Dset{&item, x, 0, 1, x}
	return &item
}

// path compression
func (x *Dset) Find() *Dset {
	if x.Parent != x {
		x.Parent = x.Parent.Find()
	}
	return x.Parent
}

// union by size
func Union(x *Dset, y *Dset) {
	x_root := x.Find()
	y_root := y.Find()
	if *x_root == *y_root {
		return
	}
	if (*x_root).Size < (*y_root).Size {
		x_root, y_root = y_root, x_root
	}
	(*y_root).Parent = x_root
	(*x_root).Size += (*y_root).Size
	minimum := min((*x_root).Min, (*y_root).Min)
	(*x).Min = minimum
	(*y).Min = minimum
	(*x_root).Min = minimum
	(*y_root).Min = minimum
	if (*x_root).Rank == (*y_root).Rank {
		(*x_root).Rank++
	}
}
