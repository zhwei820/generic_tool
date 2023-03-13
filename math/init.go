package math

import "github.com/zhwei820/generic_tool"

func Abs[T generic_tool.Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
