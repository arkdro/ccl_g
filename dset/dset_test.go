package dset

import (
	"testing"
)

func Test_find_1(t *testing.T) {
	x := Create(10)
	actual := x.Find()
	if (*actual).Rank != 0 {
		t.Error("find rank mismatch")
	}
	if (*actual).Size != 1 {
		t.Error("find size mismatch")
	}
	if (*actual).Val != 10 {
		t.Error("find value mismatch")
	}
}

func Test_union_1(t *testing.T) {
	x := Create(10)
	y := Create(15)
	Union(x, y)
	if x.Rank != 1 {
		t.Error("find x rank mismatch")
	}
	if y.Rank != 0 {
		t.Error("find y rank mismatch")
	}
	if x.Size != 2 {
		t.Error("find x size mismatch")
	}
	if y.Size != 1 {
		t.Error("find y size mismatch")
	}
	if x.Val != 10 {
		t.Error("find x value mismatch")
	}
	if y.Val != 15 {
		t.Error("find y value mismatch")
	}
}

func Test_union_2(t *testing.T) {
	x := Create(10)
	y := Create(15)
	y2 := Create(22)
	Union(x, y)
	Union(x, y2)
	if x.Rank != 1 {
		t.Error("find x3 rank mismatch")
	}
	if y.Rank != 0 {
		t.Error("find y rank mismatch")
	}
	if (*x).Size != 3 {
		t.Error("find x3 size mismatch")
	}
	if y.Size != 1 {
		t.Error("find y size mismatch")
	}
	if x.Val != 10 {
		t.Error("find x3 value mismatch")
	}
	if y.Val != 15 {
		t.Error("find y value mismatch")
	}
}

func Test_find_and_union_1(t *testing.T) {
	x := Create(10)
	y := Create(15)
	y2 := Create(22)
	Union(x, y)
	Union(y, y2)
	res := y2.Find()
	if res.Rank != 1 {
		t.Error("find and union, rank mismatch")
	}
	if res.Size != 3 {
		t.Error("find and union, size mismatch")
	}
	if res.Val != 10 {
		t.Error("find and union, value mismatch")
	}
}

func Test_find_and_union_2(t *testing.T) {
	x := Create(10)
	y := Create(15)
	y2 := Create(22)
	Union(x, y)
	Union(y2, y)
	res := y2.Find()
	if res.Rank != 1 {
		t.Error("find and union, 2, rank mismatch")
	}
	if res.Size != 3 {
		t.Error("find and union, 2, size mismatch")
	}
	if res.Val != 10 {
		t.Error("find and union, 2, value mismatch")
	}
}

func Test_find_and_union_3(t *testing.T) {
	x := Create(10)
	y := Create(15)
	y2 := Create(22)
	Union(y, x)
	Union(y2, y)
	Union(x, y2)
	res := x.Find()
	if res.Rank != 1 {
		t.Error("find and union, 3, rank mismatch")
	}
	if res.Size != 3 {
		t.Error("find and union, 3, size mismatch")
	}
	if res.Val != 15 {
		t.Error("find and union, 3, value mismatch")
	}
	if res.Min != 10 {
		t.Error("find and union, 3, min value mismatch")
	}
}

func Test_random_1(t *testing.T) {
	even := Create(0)
	odd := Create(1)
	N := 3000
	data := make([]*Dset, N * 2)
	for i := 0; i < N; i+=2 {
		x_even := Create(i * 2)
		data[i] = x_even
		x_odd := Create(i * 2 + 1)
		data[i+1] = x_odd
		Union(even, x_even)
		Union(odd, x_odd)
		if even.Find().Val != x_even.Find().Val {
			t.Error("random test, init set, item not found", i, even, x_even)
		}
		if odd.Find().Val != x_odd.Find().Val {
			t.Error("random test, init set, item not found", i, odd, x_odd)
		}
	}
	for i := 0; i < N; i+=2 {
		x_even := data[i]
		x_odd := data[i+1]
		if even.Find().Val != x_even.Find().Val {
			t.Error("random test, even item not found", i, even, x_even)
		}
		if odd.Find().Val != x_odd.Find().Val {
			t.Error("random test, odd item not found", i, odd, x_odd)
		}
		if even.Find().Min != 0 {
			t.Error("random test, even min item wrong", i, even, x_even)
		}
		if odd.Find().Min != 1 {
			t.Error("random test, odd min item wrong", i, odd, x_odd)
		}
	}
}
