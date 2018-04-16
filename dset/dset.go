package dset

type Dset struct {
	Parent *Dset
	Rank int
	Size int
	Val int
}

func Create(x int) *Dset {
	var item Dset
	item = Dset{&item, 0, 1, x}
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
	if (*x_root).Rank == (*y_root).Rank {
		(*x_root).Rank++
	}
}
