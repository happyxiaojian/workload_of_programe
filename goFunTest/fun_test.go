package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	"os"
	"testing"
	"time"
)

type PlayState interface {
	Play()
}

type Robot struct{}

func (r *Robot)Play(){
	fmt.Println("Robot play")
}

type User struct{}

func (u *User)Play(){
	fmt.Println("Player play")
}

// Flag  0-robot 1-user
type Player struct{
	State PlayState
	Flag int
}

func NewPlayer(flag int)*Player{
	p := &Player{
		Flag:  flag,
	}
	p.changeState()
	return p
}

func (p *Player) changeState() {
	if p.Flag == 1{
		p.State = &User{}
	}else {
		p.State = &Robot{}
	}
}

func (p *Player)SetFlag(flag int){
	p.Flag = flag
	p.changeState()
}


func TestFunc(t *testing.T){

	p := NewPlayer(0)
	p.State.Play()

	p.SetFlag(1)
	p.State.Play()

	p.SetFlag(0)
	p.State.Play()
}

//================================================
func TestEmptySlice(t *testing.T){
	var arr []int

	fmt.Printf("%#v", arr)

	arr = append(arr, 1)
	arr = append(arr, 2, 3)
	arr = append(arr, []int{4, 5, 6}...)

	fmt.Printf("%#v", arr)
}

func TestEmptyMap(t *testing.T){
	var tmap map[int]string

	tmap[1] = "a"  // Assignment to entry may panic because of 'nil' map
}

// 测试 方法的接收者是否可以是 interface
type I struct {
	Name string
}

type X struct {
	I
}

type Y struct {
	I
}

func (i I)SayName(){
	fmt.Println(i.Name)
}

func (i *I)ChangeName(name string){
	i.Name = name
}

func TestSayName(t *testing.T){


	x := X{
		I{"landon001"},
	}


	x.SayName()

	x.ChangeName("landon003")

	x.SayName()

	y := Y{
		I{"landon004"},
	}

	y.SayName()

	test(y)
}

func test(para interface{}){
	switch v := para.(type) {
	case I:
		v.SayName()
	default:
		fmt.Println("unknow")
	}
}

func TestTimeParse03(t *testing.T){

	str := "2020-08-27 4:20:00" // 北京时间
	shanghai, _ := time.LoadLocation("Asia/Shanghai")
	newYork, _ := time.LoadLocation("America/New_York")

	//stt1, _ := time.Parse("2006-01-02 15:04:05", str)

	stt,_ := time.ParseInLocation("2006-01-02 15:04:05", str, shanghai)
	fmt.Println(time.Now().In(shanghai))
	fmt.Println(stt)
	fmt.Println(stt.In(newYork))

	fmt.Println(time.Now().In(newYork))
	fmt.Println(time.Now().In(shanghai))
	fmt.Println(stt)
	fmt.Println(stt.In(shanghai))


	tt,_ := time.ParseInLocation("2006-01-02 15:04:05", str, newYork)


	fmt.Println(stt.In(shanghai))
	fmt.Println(tt.In(shanghai))

	fmt.Println(stt.Add(-4 * time.Hour ))
	fmt.Println(time.Now().In(newYork))
}

const TIME_LAYOUT = "2006-01-02 15:04:05"

func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation(TIME_LAYOUT, timeStr, l)
		fmt.Println(locationName, lt)
		return lt, nil
	}
}

func TestTime(t *testing.T) {
	fmt.Println("0. now: ", time.Now())

	str := "2020-09-10 7:00:00"

	fmt.Println("1. str: ", str)
	tt, _ := time.Parse(TIME_LAYOUT, str)
	fmt.Println("2. Parse time: ", t)
	tStr := tt.Format(TIME_LAYOUT)
	fmt.Println("3. Format time str: ", tStr)
	name, offset := tt.Zone()
	name2, offset2 := tt.Local().Zone()
	fmt.Printf("4. Zone name: %v, Zone offset: %v\n", name, offset)
	fmt.Printf("5. Local Zone name: %v, Local Zone offset: %v\n", name2, offset2)
	tLocal := tt.Local()
	tUTC := tt.UTC()
	fmt.Printf("6. t: %v, Local: %v, UTC: %v\n", t, tLocal, tUTC)
	fmt.Printf("7. t: %v, Local: %v, UTC: %v\n", tt.Format(TIME_LAYOUT), tLocal.Format(TIME_LAYOUT), tUTC.Format(TIME_LAYOUT))
	fmt.Printf("8. Local.Unix: %v, UTC.Unix: %v\n", tLocal.Unix(), tUTC.Unix())
	str2 := "1969-12-31 23:59:59"
	t2, _ := time.Parse(TIME_LAYOUT, str2)
	fmt.Printf("9. str2：%v，time: %v, Unix: %v\n", str2, t2, t2.Unix())
	fmt.Printf("10. %v, %v\n", tLocal.Format(time.ANSIC), tUTC.Format(time.ANSIC))
	fmt.Printf("11. %v, %v\n", tLocal.Format(time.RFC822), tUTC.Format(time.RFC822))
	fmt.Printf("12. %v, %v\n", tLocal.Format(time.RFC822Z), tUTC.Format(time.RFC822Z))

	//指定时区
	parseWithLocation("America/Cordoba", str)
	parseWithLocation("Asia/Shanghai", str)
	parseWithLocation("Asia/Beijing", str)
}

func TestTime02(t *testing.T){

	shanghai, _ := time.LoadLocation("Asia/Shanghai") // 数据库中的时间
	fmt.Println(shanghai)
	newYork, _ := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local") // 服务器设置的时区

	malta, _ := time.LoadLocation("Europe/Malta") // 马耳他时间
	singapore, _ := time.LoadLocation("Singapore") // 新加坡时间


	now := time.Now().In(shanghai)

	FirstOpenTime := "13:04:00"
	firstOpenTimeStr := now.Format("2006-01-02") + " " +  FirstOpenTime
	firstOpenTime, _ := time.ParseInLocation("2006-01-02 15:04:05", firstOpenTimeStr, shanghai)

	firstOpenTime1 := firstOpenTime.In(local)
	firstOpenTime2 := firstOpenTime.In(malta)
	firstOpenTime3 := firstOpenTime.In(newYork)
	firstOpenTime4 := firstOpenTime.In(singapore)

	spew.Dump("local-->", firstOpenTime1)
	name1, offset1 := firstOpenTime1.Zone()
	fmt.Println(name1, offset1/3600)
	spew.Dump("malta-->", firstOpenTime2)
	name, offset := firstOpenTime2.Zone()
	fmt.Println(name, offset/3600)

	spew.Dump("newYork-->", firstOpenTime3)
	spew.Dump("singapore-->", firstOpenTime4)

}

func TestTime06(t *testing.T){
	malta, _ := time.LoadLocation("Europe/Malta") // 马耳他时间
	fmt.Println(time.Now().In(malta))
}

func TestTime03(t *testing.T){
	//tt := "2020-03-08 01:00:00"
	const timeFormat = "2006-01-02 15:04:05"

	//loc, err := time.LoadLocation("America/New_York")
	//loc, err := time.LoadLocation("UTC")
	//log.Print(loc, err)
	//
	//testz , _ := time.ParseInLocation( timeFormat, tt,  loc)
	//fmt.Println(testz)
	//

	//localNewYork := time.FixedZone("UTC", -4*60*60)
	localNewYork := time.FixedZone("UTC", -4*60*60)
	localNewYork = time.FixedZone("GMT", -4*60*60)
	//localNewYork, _ = time.LoadLocation("AST")
	fmt.Println(time.Now().Format(timeFormat))
	fmt.Println(time.Now().In(localNewYork).Format(timeFormat))

	localNewYork, _ = time.LoadLocation("America/New_York")
	fmt.Println(time.Now().Format(timeFormat))
	fmt.Println(time.Now().In(localNewYork).Format(TIME_LAYOUT))
	//fmt.Println( testz , testz.UTC())
	//testz = testz.Add( time.Minute )
	//fmt.Println( testz , testz.UTC())
	//testz = testz.Add( time.Minute )
	//fmt.Println( testz , testz.UTC())
}

func TestColor(t *testing.T){
	stdout := colorable.NewColorable(os.Stdout)
	str := "cyan"
	_, _ = fmt.Fprintf(stdout, "\x1b[36m%s\x1b[0m", str)
}


func TestPanic(t *testing.T){
	user := GetPlayer()

	fmt.Println(user.Flag)

	fmt.Println("hello")
}

func GetPlayer()*Player{
	return nil
}

func TestOR(t *testing.T){
	user := GetPlayer()
	if user != nil || user.Flag == 0 {
		fmt.Println("hello world")
	}
}


func TestPanicDefer(t *testing.T){
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("world")
			fmt.Println(err)
			fmt.Println("hello")

		}
	}()
	defer fmt.Println("11111111111111111111111")
	defer fmt.Println("22222222222222222222222")
	panic("4444444444444444444444444444")
	defer fmt.Println("33333333333333333333333")
}



