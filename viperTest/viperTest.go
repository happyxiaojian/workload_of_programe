package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)


func main01(){
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))
}


type Config struct {
	AppName  string
	LogLevel string

	MySQL    MySQLConfig
	Redis    RedisConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP   string
	Port int
}

func main02() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}

	var c Config
	_ = viper.Unmarshal(&c)

	fmt.Println(c.MySQL)
}

func main03() {
	viper.SetConfigName("config2")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.Set("app_name", "awesome web2")
	viper.Set("log_level", "DEBUG2")
	viper.Set("mysql.ip", "127.0.0.12")
	viper.Set("mysql.port", 33062)
	viper.Set("mysql.user", "root2")
	viper.Set("mysql.password", "1234562")
	viper.Set("mysql.database", "awesome2")

	viper.Set("redis.ip", "127.0.0.12")
	viper.Set("redis.port", 63812)

	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatal("write config failed: ", err)
	}
}

func main(){
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}

	viper.WatchConfig()

	go func() {
		//tt := time.NewTicker(time.Second)
		for {
			select {
				case <- time.Tick(time.Second):
					fmt.Println("redis port after sleep: ", viper.Get("redis.port"))
			}
		}

	}()

	//fmt.Println("redis port before sleep: ", viper.Get("redis.port"))
	//time.Sleep(time.Second * 10)
	//fmt.Println("redis port after sleep: ", viper.Get("redis.port"))

	select {
	}
}