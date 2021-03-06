package leet_code

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"time"
)


var  empthStruct = struct{}{}

func Test1001(t *testing.T){
	map1 := make(map[string]struct{})

	map1["landon"] = empthStruct

	a, ok := map1["a"]
	t.Log(ok)
	t.Log(a)
}


// 测试匿字段的重写
type A struct{
	C
	Name string

}

type A1 struct{
	B
 	Name string

}

type B struct{
	Age int
}

type D struct{
	Addr string
}



type C interface {
	Hello()
}

func (b B)Hello(){
	fmt.Printf("hello my age is %d\n", b.Age)
}

func (d D)Hello(){
	fmt.Printf("hello my address is %s\n", d.Addr)
}

//func (a A)Hello(){
//	fmt.Printf("hello my name is %s, age is %d\n", a.Name, a.Age)
//}


func Test1002(t *testing.T){
	aa := &A{
		B{18},
		"landon",
	}

	aa.Hello()

	aa.C = D{"beijing"}

	aa.Hello()

}

//=======================================
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Array []int

func (arr Array) Len() int {
	return len(arr)
}

func (arr Array) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Array) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 匿名struct
type reverse struct {
	Array
}

// 重写
func (r reverse) Less(i, j int) bool {
	return r.Array.Less(j, i)
}

// 构造reverse Interface
func Reverse(data Array) Interface {
	return &reverse{data}
}

func Test1003(t *testing.T) {
	arr := Array{1, 2, 3}
	rarr := Reverse(arr)
	sort.Sort(arr)
	fmt.Println(arr)

	sort.Sort(rarr)
	fmt.Println(rarr)

	fmt.Println(arr.Less(0, 1))
	fmt.Println(rarr.Less(0, 1))
}


//===========================================

type Array1 []int

func (arr Array1) Len() int {
	return len(arr)
}

func (arr Array1) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Array1) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type Array2 []int

func (arr Array2) Len() int {
	return len(arr)
}

func (arr Array2) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Array2) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type Sortable struct {
	sort.Interface
	// other field
	Type string
}

func NewSortable(i sort.Interface) Sortable {
	t := reflect.TypeOf(i).String()

	return Sortable{
		Interface: i,
		Type:      t,
	}
}

func DoSomething(s Sortable) {
	fmt.Println(s.Type)
	fmt.Println(s.Len())
	fmt.Println(s.Less(0, 1))
}

func Test1005(t *testing.T) {

	arr1 := Array1{1, 2, 3}
	arr2 := Array2{3, 2, 1, 0}

	DoSomething(NewSortable(arr1))
	DoSomething(NewSortable(arr2))
}


//==============================================
// 遍历一个结构体的字段以及对应的值

type Person struct {
	Name     string
	Sex      string
	Age      int
	PhoneNum string
	School   string
	City     string
}

func Test10005(t *testing.T) {
	p1 := Person{
		Name:     "tom",
		Sex:      "male",
		Age:      10,
		PhoneNum: "1000000",
		School:   "spb-kindergarden",
		City:     "cq",
	}

	rv := reflect.ValueOf(p1)
	rt := reflect.TypeOf(p1)

	//if rv.Kind() == reflect.Struct {
	if rt.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			//按顺序遍历
			fmt.Printf("field:%+v\tvalue:%+v\n", rt.Field(i).Name, rv.Field(i))
		}
	}


	//可以直接取想要的字段
	//reflect的type interface，FieldByName方法会返回字段信息以及是否有该字段；
	if f, ok := rt.FieldByName("Age"); ok {
		fmt.Printf("\n\nfield:%+v\tvalue:%+v\n", f.Name, rv.FieldByName("Age"))
	}
}


//===========================================================
// 	//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//	//输出: [3,3,5,5,6,7]

func TestMaxSliceWindow(t *testing.T){
	//nums := []int{1,3,-1,-3,5,3,6,7}

	nums := []int{1, -1}

	k := 2
	slice1 := maxSlidingWindow(nums, k)
	t.Log(slice1)
}



func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}


func TestFibonacci(t *testing.T)  {
	f := Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci: %d\n", f())
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci: %d\n", f())
	}
}

func fib(n int, helper map[int]int)int {
	if n == 1 || n == 2 {
		return 1
	}
	if v1, ok := helper[n]; ok{
		return v1
	}
	helper[n] = fib(n-1, helper) + fib(n-2, helper)
	return helper[n]
}

func fib3(n int, helper map[int]int) int{
	helper[1], helper[2] = 1, 1
	for i := 3; i <= n; i++ {
		helper[i] = helper[i-1] + helper[i-2]
	}
	return helper[n]
}


func fib2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib2(n - 1) + fib2(n - 2)
}

func TestFib(t *testing.T){
	start := time.Now()
	helper := map[int]int{}
	fmt.Println(fib(200, helper))
	cc := time.Since(start)
	t.Log(cc)

}

func TestFib2(t *testing.T){
	start2 := time.Now()
	fmt.Println(fib2(40))
	t.Log(time.Since(start2))
}

func TestSlice(t *testing.T){
	arr := []int{4, 3, 2, 1}

	fmt.Println(len(arr))

	arr = append(arr[:1], arr[2:]...)

	fmt.Printf("arr: len: %#v --> %#v\n", len(arr), arr)


}

func TestSlice2(t *testing.T){
	arr := []int{1, 1}

	resp := removeDuplicates2(&arr)

	fmt.Printf("arr: %#v --> %#v\n", resp, arr)
}

func removeDuplicates2(nums *[]int) int {
	leng := len(*nums)
	var n int
	for i := 0; i < len(*nums) -1; i++ {
		for j := i + 1; j < len(*nums); j ++ {
			if (*nums)[i] == (*nums)[j] {
				*nums = append((*nums)[:j], (*nums)[j+1:]...)
				n++
				j--
			}
		}
	}
	return leng - n
}

func TestModifySlice(t *testing.T){
	arr := []int{1, 3, 4}

	modifySlice(arr)

	fmt.Printf("arr: %#v\n addr:%p --> len:%d, cap:%d\n", arr, &arr, len(arr), cap(arr))

}


func modifySlice(nums []int){
	fmt.Printf("inner nums: %p\n", &nums)
	nums[0] = 1000

	nums = append(nums[:1], nums[2:]...)

	fmt.Printf("inner nums: %#v --> len:%d, cap:%d\n", nums, len(nums), cap(nums))
}
