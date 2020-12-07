package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

const PASSWORD string = ""
const OK string = "ok"
const NO string = "no"

type RedisConnController struct {
	beego.Controller
}

func newPool(idc, prekey string) *redis.Pool {
	var redisConn map[string]string = map[string]string{
		"beijing": "127.0.0.1:6379",
		"tianjin": "10.10.10.10:6379"}

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConn[idc])
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", PASSWORD); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func (c *RedisConnController) GetDictData() string {
	userid, err := c.GetInt("userid")
	idc := c.GetString("idc")
	prekey := c.GetString("prekey")
	if err != nil {
		fmt.Printf("用户id[%d]参数出错", userid)
		return NO
	}

	var pool *redis.Pool = newPool(idc, prekey)
	conn := pool.Get()
	value, err := conn.Do("lrange", prekey, 0, -1)
	if err != nil {
		fmt.Printf("用户id[%d]从redis读取数据出错", userid)
		return NO
	}
	type ids []int
	data, _ := redis.Ints(value, err)
	for _, v := range data {
		if v == userid {
			fmt.Printf("用户id[%d]在黑名单中", userid)
			return OK
		}
	}
	fmt.Printf("用户id[%d]不在黑名单中", userid)
	return NO
}
