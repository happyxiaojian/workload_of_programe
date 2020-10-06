package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)


//判断是否为数字
func IsNo(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		// 符号和空格
		str = strings.Trim(str, " \\t\\n\\r\\v\\f")
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// 十六进制
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // 小数点
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' {
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}

var d string

func main(){
	Init()
	flag.Parse()

	now := time.Now().Format("2006-01-02")


	fmt.Println(now)
	fmt.Println(d)

	if d == "" {
		d = now
		tt, err := time.Parse("2006-01-02", d)
		if err != nil{
			d = now
		}
		fmt.Println(err)
		printTime(tt)
	}

	tt2, _ := time.Parse("2006-01-02", "2020-09-26")
	printTime(tt2)
}

func printTime(t time.Time){
	fmt.Println(t)
}

var (
	intflag int
	boolflag bool
	stringflag string
)

func Init() {
	flag.StringVar(&d, "day", "", "input date like -day=2020-03-04")
	flag.IntVar(&intflag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflag, "stringflag", "default", "string flag value")
}

func main02() {
	flag.Parse()

	fmt.Println("int flag:", intflag)
	fmt.Println("bool flag:", boolflag)
	fmt.Println("string flag:", stringflag)
}
