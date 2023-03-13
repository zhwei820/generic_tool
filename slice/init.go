package slice

import "github.com/zhwei820/generic_tool"

func Sum[T generic_tool.Number](numbers []T) T {
	var total T
	for _, x := range numbers {
		total += x
	}
	return total
}

func Avg[T generic_tool.Number](numbers []T) T {
	var total T
	for _, x := range numbers {
		total += x
	}
	return total / T(len(numbers))
}

func Max[T generic_tool.Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T generic_tool.Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MaxSlice[T generic_tool.Number](a []T) (max T) {
	if len(a) == 0 {
		panic("len ")
	}
	max = a[0]
	for _, x := range a {
		if x > max {
			max = x
		}
	}

	return max
}

func MinSlice[T generic_tool.Number](a []T) (min T) {
	if len(a) == 0 {
		panic("len ")
	}
	min = a[0]
	for _, x := range a {
		if x < min {
			min = x
		}
	}

	return min
}
