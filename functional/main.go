package main

import (
	"strconv"

	lop "github.com/samber/lo/parallel"
)

func main() {
	resultArr := lop.Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	lop.ForEach(resultArr, func(x string, _ int) { println(x) })
}
