package plate

type Plate struct {
	Width int
	Height int
	Color_range int
	Data [][]int
}

func (plt Plate) Valid_data() bool {
	if plt.Width <= 0 {
		return false
	}
	if plt.Height <= 0 {
		return false
	}
	if len(plt.Data) != plt.Height {
		return false
	}
	for _, row := range plt.Data {
		if len(row) != plt.Width {
			return false
		}
	}
	return true
}

