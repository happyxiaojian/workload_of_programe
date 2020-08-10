package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-redis/redis/v8"
	"github.com/tidwall/gjson"
	"time"
)

var redisClient *redis.Client

var ctx = context.Background()

const(
	REDIS_KEY_USER = "user:"
	REDIS_KEY_USER_LIST = "userList"
)

type User struct {
	Name string `json:"name"`
	Age  int32	`json:"age"`
}

func main() {

	InitDBForCache()

	val, err := redisClient.Get(ctx, "key").Result()
	checkErr(val, err, "redisClient.Get")

	user := User{
		Name:"landon",
		Age:24,
	}

	// 不能直接把结构体存储到redis中
	val, err = redisClient.Set(ctx, REDIS_KEY_USER + "10001", user,90-000-000-000).Result()
	checkErr(val, err, "redisClient.Set")

	// 不能直接把 map 存储到redis中
	userMap := map[string]interface{}{
		"name":"landon",
		"age":24,
	}
	val, err = redisClient.Set(ctx, REDIS_KEY_USER + "10001", userMap,90-000-000-000).Result()
	checkErr(val, err, "redisClient.Set")

	//
	userJson, err := json.Marshal(&user)
	checkErr(userJson, err, "json.Marshal(&user)")

	val, err = redisClient.Set(ctx, REDIS_KEY_USER + "10001", userJson,300 * time.Second).Result()
	checkErr(val, err, "redisClient.Set")

	val, err = redisClient.Get(ctx, REDIS_KEY_USER + "10001").Result()
	checkErr(val, err, "redisClient.Get")

	resp := redisClient.Get(ctx, REDIS_KEY_USER + "10001")
	spew.Dump(resp.Name())
	spew.Dump(resp.Args())
	spew.Dump(resp.String())
	spew.Dump(resp.Val())
	checkErr(*resp, nil, "redisClient.Get")

	result := gjson.Get(resp.Val(), "Name").String()
	spew.Dump(result)

	var user2 = &User{}
	err1 := json.Unmarshal([]byte(resp.Val()), user2)
	spew.Dump(err1)
	spew.Dump(user2)

	userList := []User{
		{"landon01", 22},
		{"landon02", 23},
		{"landon03", 24},
	}

	userListJson, err := json.Marshal(userList)
	redisClient.Set(ctx, REDIS_KEY_USER_LIST, userListJson, 300 * time.Second)

	result, err = redisClient.Get(ctx, REDIS_KEY_USER_LIST).Result()
	checkErr(result, err, "get REDIS_KEY_USER_LIST")
	userList2 := &[]User{}
	json.Unmarshal([]byte(result), userList2)
	//userList2 := gjson.Get(result, REDIS_KEY_USER_LIST)
	spew.Dump(userList2)

}

func InitDBForCache() {
	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

	pong, err := redisClient.Ping(ctx).Result()

	spew.Dump(pong, err)

}

func checkErr(resp interface{}, err error, funcName string){
	fmt.Println("==========================================")
	fmt.Printf("%s", "返回的结果是：")
	spew.Dump(resp)
	fmt.Printf("%s", "返回的错误是：")
	spew.Dump(err)
	fmt.Printf("%s", "所在的函数是：")
	spew.Dump(funcName)

}
