package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main()  {
	//连接redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	//通过Do函数，发送redis命令
	v, err := c.Do("set", "name", "张三")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(c.Do("get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

	//操作列表
	c.Do("lpush", "list", "james")
	c.Do("lpush", "list", "chris paul")
	c.Do("lpush", "list", "wade")

	//读取列表
	values, _ := redis.Values(c.Do("lrange", "list", "0", "3"))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

	// 或者
	var v1 string
	redis.Scan(values, &v1)
	fmt.Println(v1)
}