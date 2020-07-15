package leet_code

import (
	"math"
)

//假设Andy和Doris想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。
//
//你需要帮助他们用最少的索引和找出他们共同喜爱的餐厅。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设总是存在一个答案。
//
//示例 1:
//
//输入:
//["Shogun", "Tapioca Express", "Burger King", "KFC"]
//["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
//输出: ["Shogun"]
//解释: 他们唯一共同喜爱的餐厅是“Shogun”。
//示例 2:
//
//输入:
//["Shogun", "Tapioca Express", "Burger King", "KFC"]
//["KFC", "Shogun", "Burger King"]
//输出: ["Shogun"]
//解释: 他们共同喜爱且具有最小索引和的餐厅是“Shogun”，它有最小的索引和1(0+1)。
//提示:
//
//两个列表的长度范围都在 [1, 1000]内。
//两个列表中的字符串的长度将在[1，30]的范围内。
//下标从0开始，到列表的长度减1。
//两个列表都没有重复的元素。

func findRestaurant(list1 []string, list2 []string) []string {

	map3 := make(map[int][]string)

	for k1, v1 := range list1{
		for k2, v2 := range list2{
			if v1 == v2{
				//fmt.Println(v1)
				if map3[k2+k1] == nil{
					map3[k1+k2] = make([]string, 0)
				}
				map3[k1+k2] = append(map3[k1+k2], v1)
			}
		}
	}
	ans := math.MaxInt32
	for k := range map3{
		if ans > k{
			ans = k
		}
	}

	return map3[ans]
}


func findRestaurantV2(list1 []string, list2 []string) []string {
	map1 := make(map[string]int, len(list1))
	map2 := make(map[string]int, len(list2))

	map3 := make(map[int][]string)

	for k, v := range list1{
	    map1[v] = k
	}
	for k, v := range list2{
	    map2[v] = k
	}


	for k, v1 := range map1{
		if v2, exist := map2[k]; exist{
			if map3[v2+v1] == nil{
				map3[v2+v1] = make([]string,1)
			}
			map3[v2+v1] = append(map3[v2+v1], k)
		}
	}

	ans := math.MaxInt32
	for k := range map3{
		if ans > k{
			ans = k
		}
	}

	//fmt.Println(map3)

	return map3[ans]
}
