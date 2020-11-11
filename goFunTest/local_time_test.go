package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"
)

func GetGSTimeRange() (begin, ends time.Time) {
	localCN, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(localCN)
	startUSA := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, localCN)
	endUSA := time.Date(now.Year(), now.Month(), now.Day()+1, 8, 58, 48, 0, localCN)
	start := ParseInLocation(startUSA, localCN)
	end := ParseInLocation(endUSA, localCN)
	return start, end
}


func GetGSNextDay(t time.Time, local *time.Location, day int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+day, 0, 0, 0, 0, local)
}

func GetNextDay(t time.Time, local *time.Location) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, local)
}

//美国东部时间以位于西五区的纽约、华盛顿等城市为代表，区分夏令时（从每年的三月份的第二个周日开始 北京时间-12小时 ）和冬令时（从十一月份第一个周日开始）。
func ParseInLocation(t time.Time, loc *time.Location) time.Time {
	h := 12
	if t.Month().String() == "November" && t.Day() == 3 && t.Weekday().String() == "Sunday" {
		h = 13
	}
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour()-h, t.Minute(), t.Second(), 0, loc)
}

func TestGetGSTimeRange(t *testing.T) {
	start, end := GetGSTimeRange()
	fmt.Println(start, end)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}


/*
@name 获取当前时间和时区
*/
func GetTimeAndLocal() (time.Time, *time.Location) {
	today := time.Now()
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}

	return today, local
}

func TestTimeParse(t *testing.T){
	today, local := GetTimeAndLocal()
	spew.Dump(today)
	spew.Dump(local)

	otime := time.Date(today.Year(), today.Month()+1, 1, 10, 0, 0, 0, local)
	ctime := time.Date(today.Year(), today.Month()+1, 1, 20, 45, 0, 0, local)

	spew.Dump(otime)
	spew.Dump(ctime)
}


func TestTimeParse02(t *testing.T){
	tt := "2020-08-02 22:03:29.9999999"
	tt1, err := time.Parse("2006-01-02 15:04:05", tt)
	spew.Dump(tt1, err)
	spew.Dump(tt1.Add(time.Second * 1))
	spew.Dump(tt1.Truncate(time.Second * 2))

}


func TestTimeParse04(t *testing.T){
	date := "20201022"

	dataTime, _ := time.Parse("20060102", date)

	dayStart := dataTime.Format("2006-01-02") + " 00:00:00"

	dayEnd := dataTime.Format("2006-01-02") + " 23:59:59"

	spew.Dump(dayStart, dayEnd)
}


func TestTimeParse05(t *testing.T){
	date := "2020-11-05T06:27:00.130Z"

	dataTime, _ := time.ParseInLocation(time.RFC3339, date, time.UTC)
	fmt.Println(dataTime)
	dataTime, _ = time.ParseInLocation(time.RFC3339Nano, date, time.UTC)

	fmt.Println(dataTime)
}


func TestTimeParse06(t *testing.T){
	// 年份转换的时候 [2020]350 --> [2021]001
	var LastIssue uint64 = 2019350 + 1
	var err error
	issue := fmt.Sprintf("%d", LastIssue)
	nowYear := strconv.Itoa(time.Now().Year())
	if !strings.Contains(issue, nowYear) {
		issue = strings.ReplaceAll(issue, issue[:4], nowYear)
		issue = fmt.Sprintf("%s001", nowYear)
		if LastIssue, err = strconv.ParseUint(issue, 10, 64); err != nil{
			log.Printf("load_LHC_Impl strconv.ParseUint(v=\"%v\") err:%v\n", issue, err.Error())
		}
	}
	fmt.Println(LastIssue)
}


func TestTimeParse07(t *testing.T){
	year := "2020"
	n := 121
	ss := fmt.Sprintf("%s%03d", year, n)
	fmt.Println(ss)
}