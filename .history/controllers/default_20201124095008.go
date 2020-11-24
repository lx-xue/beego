package controllers

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"myapp-beego/models"
)

type DefaultController struct {
	beego.Controller
}

func (user *DefaultController) Index() {
	numbers := make([]int, 3, 5)
	id := user.Input().Get("id")
	name := user.Input().Get("name")
	fmt.Printf("%v\n %v\n",id,name)
	fmt.Printf("sssss\n")
	user.Data["numbers"] = numbers //切片类型变量
	//输出相应
	user.Data["id"] = id
	user.Data["name"] = name
	user.TplName = "user/index.html"
}

//添加用户
func (user *DefaultController) AddUser() {
	user_name := user.GetString("user_name")
	addr_ip := user.GetString("addr_ip")
	curtime := time.Now().Unix()//当前时间戳
	user_data := models.User{
		UserName : user_name,
		AddrIp : addr_ip,
		CreateTime:int(curtime),
		UpdateTime:int(curtime),
	}
	id ,_ := models.AddUser(&user_data)
	fmt.Printf("%v\n %v\n %v\n",id,user_data)
	var slice []float64 = make([]float64, 5, 10) //创建一个切片
	var slice1 [5]int = [...]int{1, 2, 3, 4, 5}  //创建一个切片,并初始化
	arr := slice1[1:3]
	user.Data["slice"] = slice
	user.Data["slice1"] = arr
	user.Ctx.WriteString("post表单提交成功:\n")
}