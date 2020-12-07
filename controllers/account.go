package controllers

import (
	"database/sql"
	"fmt"
	"myapp-beego/models"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

//定义控制器(结构体)
type AccountController struct {
	beego.Controller
}

//登录页面
func (this *AccountController) LoginIndex() {
	this.Data["name"] = "name"
	this.TplName = "account/login.html"
}

//登录
func (this *AccountController) LoginAction() {
	addr_ip := this.GetString("addr_ip")
	password := this.GetString("password")
	//登录成功后,将账户信息存入redis(并设置缓存时间)redis获取账户信息
	redis, err := cache.NewCache("redis", `{"key":"user","conn":":6379","dbNum":"0"}`)
	// fmt.Printf("%v\n %v\n", redis, err)
	if err == nil {
		get := redis.Get(addr_ip)
		if get != nil { //redis存在此账户信息,直接返回结果
			fmt.Printf("%v\n", string(get.([]byte)))
			this.Data["json"] = map[string]interface{}{"status": 1, "msg": "success", "data": string(get.([]byte))}
			this.ServeJSON()
		} else { //redis不存在此账户信息,先保存redis,再返回结果
			//框架orm获取账户信息
			res, err := models.GetUserByMap(addr_ip, password)
			if res == nil {
				this.Data["json"] = map[string]interface{}{"status": 0, "msg": err}
				this.ServeJSON()
			} else {
				user_data := map[string]interface{}{}

				user_data["Id"] = res.Id
				user_data["UserName"] = res.UserName
				user_data["AddrIp"] = res.AddrIp
				user_data["Password"] = res.Password
				user_data["CreateTime"] = res.CreateTime
				user_data["UpdateTime"] = res.UpdateTime
				fmt.Printf("%v\n", user_data)
				redis.Put(addr_ip, user_data, 3600*time.Second)
				this.Data["json"] = map[string]interface{}{"status": 1, "msg": "success", "data": res}
				this.ServeJSON()
			}
		}

	}
	//db库接口获取账户信息
	db, err := sql.Open("mysql", "root:root@/beego?charset=utf8")
	// fmt.Printf("%v\n", &err)
	rows, err := db.Query("SELECT * FROM user")
	// fmt.Printf("%v\n", rows)
	user_arr := make([]interface{}, 10) //声明指定长度的字典类型变量
	user_arr_key := 0
	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns)) //声明指定长度的字典类型变量
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i] //将values的指针循环付给scanArgs
		}

		//将数据保存到 record 字典scanArgs元素必须是指针元素
		err = rows.Scan(scanArgs...) //此操作也相当于把数据保存到values
		// fmt.Printf("%v\n", err)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				// fmt.Println(string(col.([]byte)))
				record[columns[i]] = string(col.([]byte))
			}
		}
		// fmt.Println(record)
		user_arr[user_arr_key] = record
		user_arr_key++
	}
	// fmt.Printf("%v\n", user_arr)
	var ss models.User
	//查询一条
	err = db.QueryRow("SELECT * FROM user WHERE user_id =?", 1).Scan(&ss.Id, &ss.UserName, &ss.AddrIp, &ss.Password, &ss.CreateTime, &ss.UpdateTime)
	this.Data["aa"] = map[string]interface{}{"status": 1, "msg": "success", "data": "ss"}
	this.ServeJSON()
	// fmt.Printf("%v\n %v\n", err, ss)
}

//nil 检查函数
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
