package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var i int;
func main() {

	http.HandleFunc("/",index);
	http.ListenAndServe("localhost:811",nil);

}

func index(rw http.ResponseWriter,req *http.Request){
	i =i+1;
	fmt.Println(i);
	t, _ := template.ParseFiles("index.html")
	t.Execute(rw,nil)
}
