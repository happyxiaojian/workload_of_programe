package leet_code

import (
	"fmt"
	"sort"
	"testing"
)

func removeDuplicates(nums []int)int{
	sort.Ints(nums)
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i+1; j < l; j++ {
			if nums[i] == nums[j] {
				l--
				nums = append(nums[0:j], nums[j+1:]...)
				j--
			}else {
				break
			}
		}
	}
	return l
}

func Test2002(t *testing.T){
	//nums = [0,0,1,1,1,2,2,3,3,4],
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	le := removeDuplicates(nums)

	fmt.Printf("%#v --> %d\n", nums, le)
}


func Test2001(t *testing.T){
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("%#v\n", arr)

	arr = append(arr[0:1], arr[2:]...)

	fmt.Printf("%#v\n", arr)

	fmt.Println(len(arr), cap(arr))
}

func Test2003(t *testing.T){
	arr := []int{6, 5, 4, 4, 3, 2}

	sort.Ints(arr)

	fmt.Printf("%#v\n",arr)
}