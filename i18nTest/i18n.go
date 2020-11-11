package main

import (
	"encoding/json"
	"fmt"
	"github.com/Chain-Zhang/pinyin"
	"github.com/Xuanwo/go-locale"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"strings"
	"time"

	"log"
)

type user struct {
	name string
	Age int
}

func main01() {
	log.SetPrefix("[debug]-->")
	tag, err := locale.Detect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", tag)
	spew.Dump(tag)

	spew.Dump(user{})
	// Have fun with language.Tag!

}

func main02()  {

	str, err := pinyin.New("大小、单双").Split("").Mode(pinyin.InitialsInCapitals).Convert()
	str = strings.ReplaceAll(str, "、", "-")
	stru := strings.ToUpper(str)
	if err != nil {
		// 错误处理
	}else{
		fmt.Println(stru)
	}

	str, err = pinyin.New("我是中国人").Split(" ").Mode(pinyin.WithoutTone).Convert()

	if err != nil {
		// 错误处理
	}else{
		fmt.Println(str)
	}

	str, err = pinyin.New("我是中国人").Split("-").Mode(pinyin.Tone).Convert()
	if err != nil {
		// 错误处理
	}else{
		fmt.Println(str)
	}

	str, err = pinyin.New("我是中国人").Convert()
	if err != nil {
		// 错误处理
	}else{
		fmt.Println(str)
	}
}

func main03() {
	filePath := "D:/torch_work/lottery_odds"
	files, _ := ioutil.ReadDir(filePath)

	for _, file := range files {
		spew.Dump(file.Name())
		contens, _ := ioutil.ReadFile(filePath + "/" + file.Name())

		var contensJson interface{}

		json.Unmarshal(contens, &contensJson)

		//spew.Dump(contensJson)

		if v, ok := contensJson.(map[string]interface{}); ok {
			for k, value := range v {
				if k == "zh-CN" {
					spew.Dump(value)
					if v1, ok := value.(map[string]interface{}); ok {
						for k, v := range v1 {
							fmt.Printf("%v  %v\n", k, v)
						}

					}

				}
			}

		}
	}
}

func main(){
	tt := "2020-10-31 10:00:00"
	t1, _ := time.Parse(TIME_FORMAT, tt)
	fmt.Println(t1.Location().String())
	fmt.Println(t1)

	shanghai, _, _ := GetTimeZone(TimeZoneShangHai)
	t2, _ := time.ParseInLocation(TIME_FORMAT, tt, shanghai)
	fmt.Println(t2.Location().String())
	fmt.Println(t2)


	tz, offset, _ := GetTimeZone(0)
	fmt.Println(offset, time.Now())
	fmt.Printf("tz:%3v--time.now:%v\n", tz, time.Now().In(tz))
}

const (
	TimeZoneNewYork   = 0 // 美东时区
	TimeZoneShangHai  = 1 // 北京时间
	TimeZoneSingapore = 2 // 新加坡
	TimeZoneMalta     = 3 // 马耳他
	TIME_FORMAT = "2006-01-02 15:04:05"
)

// offsetHour 是和上海时间相比的时间差
func GetTimeZone(t int) (*time.Location, time.Duration, error) {
	var timeZone *time.Location
	var offsetHour time.Duration
	var err error
	switch t {
	case TimeZoneNewYork:
		//timeZone, err = time.LoadLocation("America/New_York") // 美东时区
		timeZone = time.FixedZone("UTC", -4*60*60) // 美东时区
		offsetHour = time.Duration(-12*time.Hour)			   // 和北京时差固定为12个小时
	case TimeZoneShangHai:
		timeZone, err = time.LoadLocation("Asia/Shanghai") // 上海时区
		offsetHour = time.Duration(0*time.Hour)
	case TimeZoneSingapore:
		timeZone, err = time.LoadLocation("Singapore") // 新加坡
		offsetHour = time.Duration(0*time.Hour)
	case TimeZoneMalta:
		//timeZone, err = time.LoadLocation("Europe/Malta") // 马耳他
		timeZone = time.FixedZone("UTC", 1*60*60) // 马耳他时区
		offsetHour = time.Duration(-7*time.Hour)  // 和北京时差固定为7个小时
		//timeZone = time.FixedZone("UTC", 2*60*60) // 美东时区

	}
	return timeZone, offsetHour, err
}