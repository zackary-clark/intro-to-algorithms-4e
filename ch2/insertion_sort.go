package ch2

func InsertionSort(a []int) []int {
	l := len(a)
	b := make([]int, l, l)
	for i, v := range a {
		if i == 0 {
			b[0] = a[0]
			continue
		}
		j := i - 1
		for j >= 0 && b[j] > v {
			b[j+1] = b[j]
			j--
		}
		b[j+1] = v
	}
	return b
}
