package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func main(){
	//i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		Print1()

		Print2()

		Print3()
	})

	//c.AddFunc(spec, func() {
	//	Print3()
	//})

	c.AddFunc(spec, func() {
		Print4()
	})

	c.Start()

	select{}
}


func Print1(){
	fmt.Println("111111111111111")
}

func Print2(){
	panic("22")

	fmt.Println("222222222222222")
}

func Print3(){
	fmt.Println("33333333333333")
}

func Print4(){
	fmt.Println("444444444444444444")
}


func (cs *Crontable) TaskFunc(rule string, task func()) {
	if err := cs.cron.AddFunc(rule, task); err != nil {
		log.Fatal("start new task is error: ", err)
	}
}


// Crontable 定时任务服务
type Crontable struct {
	cron *cron.Cron
}

// 新建定时服务
func NewCron() *Crontable {
	cronS := new(Crontable)
	cronS.cron = cron.New()
	return cronS
}


// 启动定时
func (cs *Crontable) Start() {
	cs.cron.Start()
}
