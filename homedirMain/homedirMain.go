package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os/user"
	"sort"
)

func main01() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Home dir:", u.HomeDir)
}

func main02() {
	//dir, err := homedir.Dir()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Home dir:", dir)
	//
	//dir = "~/golang/src"
	//expandedDir, err := homedir.Expand(dir)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("Expand of %s is: %s\n", dir, expandedDir)
	//
	m := map[int]string{
		1:"a",
		2:"b",
	}

	a, exist := m[3]
	spew.Dump(exist, a)


	var aa []*int

	i := 1
	aa = append(aa, &i)

	spew.Dump(aa)
}

func main03(){
	fArr := []float64{1.1, 5.5, 2.2, 7.7, 4.4}

	sort.Float64s(fArr)

	spew.Dump(fArr)
}

type Book struct {
	Name string
	Price float64
}

type User struct {
	Id int
	Books map[string]string
}

func main04(){
	//users := map[int]*User{}

	user1 := User{
		Id: 1,
		Books: map[string]string{
			"a":"1111",
			"b":"2222",
		},
	}

	user2 := User{
		Id: 2,
		Books: map[string]string{
			"a":"333",
			"b":"4444",
		},
	}

	users := map[int]*User{
		user1.Id:&user1,
		user2.Id:&user2,
	}

	var users2 map[int]*User

	users2 = users

	user3 := users[2]
	//spew.Dump(user3)

	delete(users, user2.Id)
	spew.Dump(users)

	fmt.Println("================")
	spew.Dump(user3)

	spew.Dump(users2)
}

func main(){
	aMap := map[int]int{
		1:100,
		2:2,
	}

	a, exist := aMap[1]
	fmt.Println(a, exist)
}
