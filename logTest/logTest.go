package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogInfo()*log.Logger{
	//创建io对象，日志的格式为当前时间.log;2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	file := "./" + time.Now().Format("2006-01-02") + ".log"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}

	//创建一个Logger：参数1：日志写入目的地 ； 参数2：每条日志的前缀 ；参数3：日志属性
	return  log.New(logFile, "自定义的前缀",log.Lshortfile)
}

var myLoger *log.Logger

func main() {
	//调用日志初始化方法
	 myLoger= LogInfo()

	//Prefix返回前缀，Flags返回Logger的输出选项属性值
	fmt.Printf("创建时前缀为:%s\n创建时输出项属性值为:%d\n", myLoger.Prefix(), myLoger.Flags())

	//SetFlags 重新设置输出选项  myLoger.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	//重新设置输出前缀  myLoger.SetPrefix("test_")

	//获取修改后的前缀和输出项属性值
	fmt.Printf("修改后前缀为:%s\n修改后输出项属性值为:%d\n", myLoger.Prefix(), myLoger.Flags())

	//输出一条日志
	_ = myLoger.Output(2, "使用Output进行日志输出")

	//格式化输出日志  myLoger.Printf("我是%v方法在%d行内容为:%s","Printf", 40, "其实我底层是以fmt.Printf的方式处理的，相当于Java里的Info级别")

	//开启这个注释，下面代码就不会继续走，并且程序停止
	// myLoger.Fatal("我是Fatal方法，我会停止程序，但不会抛出异常")

	//调用业务层代码
	serviceCode()
	 myLoger.Printf("业务代码里的Panicln不会影响到我，因为他已经被处理干掉了，程序目前正常")
}

func serviceCode()  {
	defer func() {
		if r := recover(); r != nil {
			//用以捕捉Panicln抛出的异常
			fmt.Printf("使用recover()捕获到的错误：%s\n", r)
		}
	}()
	// 模拟错误业务逻辑，使用抛出异常和捕捉的方式记录日志
	if 1 == 1 { myLoger.Panicln("我是Panicln方法，我会抛异常信息的，相当于Error级别")
	}
}