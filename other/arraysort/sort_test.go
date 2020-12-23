package arraysort

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Sort(t *testing.T) {
	sortArr := generateRandomData()
	fmt.Println("排序前", sortArr)
	quickSort(sortArr)
	fmt.Println("排序后", sortArr)
}

// return []int array which length is 10 and element between [1-100]
func generateRandomData() []int {
	res := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		res = append(res, randZeroToHundred())
	}
	return res
}

func randZeroToHundred() int {
	return rand.Intn(100) + 1
}
