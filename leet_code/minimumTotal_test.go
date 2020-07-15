package leet_code

import (
	"fmt"
	"testing"
)


func Test001(t *testing.T){

	ints := [][]int{{2}, {3,4}, {6,5,7,0}, {4,1,8,3}}

	got := minimumTotal(ints)

	fmt.Println(got)

	got1 := minimumTotal(ints)

	fmt.Println(got1)

}

func Test002(t *testing.T){
	list1 := []string{"Shogun","Tapioca Express","Burger King","KFC"}
	list2 := []string{"Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"}

	got1 := findRestaurant(list1, list2)

	fmt.Println(got1)

	got2 := findRestaurant(list1, list2)

	fmt.Println(got2)
}

