package controllers

import (
	"flag"
	"fmt"

	"github.com/astaxie/beego"
)
var skillParam = flag.String("skill", "", "skill to perform")
type Ccc struct {
	Id         int    `orm:"column(user_id);auto"`
	UserName   string `orm:"column(user_name);size(255);not null"`
	AddrIp     string `orm:"column(addr_ip);size(255);not null"`
	CreateTime int    `orm:"column(create_time);size(10);null"`
	UpdateTime int    `orm:"column(update_time);size(10);null"`
}
type DefaultController struct {
	beego.Controller
}

func (user *DefaultController) Index() {
	// current_time := time.Now().Unix()
	// b := *new(Ccc)
	// a := &Ccc{
	// 	Id:         1,
	// 	UserName:   "name",
	// 	AddrIp:     "127.0.0.1",
	// 	CreateTime: int(current_time),
	// 	UpdateTime: int(current_time),
	// }
	// a_id := a.Id
	// fmt.Print("%v\n %v\n %v\n", b, a, a_id)
	//一维数组
	// var arr [2]int
	// arr[0] = 1
	// arr[1] = 2
	// arr1 := [3]int{1, 2, 3}
	// fmt.Printf("%v\n %v\n", arr1, arr)
	//多维数组
	// var arrarr [4][2]int //声明一个二维数组
	// arrarr[0][0] = 1
	// arrarr[0][1] = 2
	// arrarr[1] = arrarr[0]
	// arrarr1 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}} //声明一个二维数组,同时初始化它
	// fmt.Printf("%v\n %v\n", arrarr, arrarr1)

	array := [30]int{ //数组
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	}
	slice := []int{ //切片声明是[],不指定数量
		1, 2, 3, 4, 5,
	}
	//切片1
	slice1 := array[20:]
	fmt.Printf("%v\n %v\n", slice, slice1)
	numbers := make([]int, 3, 5)
	id := user.Input().Get("id")
	name := user.Input().Get("name")
	user.Data["numbers"] = numbers //切片类型变量
	//输出相应
	user.Data["id"] = id
	user.Data["name"] = name
	user.TplName = "user/index.html"
}

//添加用户
func (user *DefaultController) AddUser() {
	// user_name := user.GetString("user_name")
	// addr_ip := user.GetString("addr_ip")
	// curtime := time.Now().Unix() //当前时间戳
	// user_data := models.User{
	// 	UserName:   user_name,
	// 	AddrIp:     addr_ip,
	// 	CreateTime: int(curtime),
	// 	UpdateTime: int(curtime),
	// }
	// id, _ := models.AddUser(&user_data)
	// fmt.Printf("%v\n %v\n %v\n", id, user_data)
	// var slice []float64         //创建一个切片
	// var slice1 = []int{1, 2, 3} //创建一个切片,并初始化
	// arr := slice1[1:3]
	// var a [5]string
	// var b []int
	// c := a
	// d := b
	// fmt.Println("%v\n %v\n %v\n %v\n", c, d)
	// user.Data["slice"] = slice
	// user.Data["slice1"] = arr
	// var slice []int
	// a := append(slice, 1)
	// b := append(slice, 1, 2)
	// c := append(slice, []int{1, 2, 3}...)
	// fmt.Println("%v\n %v\n %v\n %v\n %v\n", a, b, c, slice)
	// var numbers []int
	// for i := 0; i < 10; i++ {
	// 	numbers = append(numbers, i)
	// 	fmt.Printf("len: %d  cap: %d pointer: %v\n", len(numbers), cap(numbers), numbers)
	// }

	// // 设置元素数量为1000
	// const elementCount = 1000
	// // 预分配足够多的元素切片
	// srcData := make([]int, elementCount)
	// //将切片赋值
	// for i := 0; i < elementCount; i++ {
	// 	srcData[i] = i
	// }
	// // 引用切片数据
	// refData := srcData
	// // 预分配足够多的元素切片
	// copyData := make([]int, elementCount)
	// // 将数据复制到新的切片空间中
	// copy(copyData, srcData)
	// // 修改原始数据的第一个元素
	// srcData[0] = 999
	// // 打印引用切片的第一个元素
	// fmt.Println(refData[0])
	// // 打印复制切片的第一个和最后一个元素
	// fmt.Println(copyData[0], copyData[elementCount-1])
	// // 复制原始数据从4到6(不包含)
	// copy(copyData, srcData[4:6])
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("%d ", copyData[i])
	// }

	// fmt.Printf("%p\n %v\n",srcData,srcData)

	//map(集合)
	// var countryCapitalMap map[string]string /*创建集合 */
	// countryCapitalMap = make(map[string]string)

	// /* map插入key - value对,各个国家对应的首都 */
	// countryCapitalMap["France"] = "巴黎"
	// countryCapitalMap["Italy"] = "罗马"
	// countryCapitalMap["Japan"] = "东京"
	// countryCapitalMap["India "] = "新德里"
	// fmt.Printf("%v\n", countryCapitalMap)
	// for k,v := range countryCapitalMap {
	// 	fmt.Printf("%v %v\n",k,v)
	// }
	// map1 := make(map[string]int)
	// map1["route"] = 66
	// map1["brazil"] = 4
	// map1["china"] = 960
	// fmt.Printf("%v\n", map1)
	// delete(map1, "china")
	// fmt.Printf("%v\n", map1)
	//切片
	// slice1 := []int{1, 2, 3}
	// slice2 := slice1[1:2]
	// slice3 := slice2
	// slice3[0] = 5
	// slice2[0] = 4
	// var a = []int{1, 2, 3} //切片
	// m := a[:]              //切片的切片
	// c := append(a, 1, 2)
	// d := append(a[:0], a[1:]...) // 在开头添加1个元素
	// a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	// for k1, v1 := range slice2 {
	// 	fmt.Printf("%v\n %v\n", k1, v1)
	// }
	// for k, v := range slice3 {
	// 	fmt.Printf("%v\n %v\n", k, v)
	// }
	// var map1 = map[string]int{"a": 1}
	// // map1 = map[string]int{"a": 1}
	// map1["c"] = 2
	// map1["d"] = 3
	// fmt.Printf("%v\n %v\n %v\n %v\n %v\n %v\n", map1)
	// 流程控制语句
	// if 1==1 && false {
	// 	fmt.Printf("%v\n",12)
	// }else{
	// 	fmt.Printf("%v\n",11)
	// }
	// count :=10
	// for i := 0; i < count; i++ {
	// 	fmt.Printf("%v\n",i)
	// }
	// map1 := map[int]int{
	// 	1:1,
	// 	2:2,
	// 	3:3,
	// }
	// for k,v := range map1{
	// 	fmt.Printf("%v %v %v\n",k,"=>",v)
	// }
	// case1 := "3"
	// switch case1{
	// case "1":
	// 	fmt.Printf("%v\n",case1)
	// case "2":
	// 	fmt.Printf("%v\n",case1)	
	// default :
	// 	fmt.Printf("%v\n",0)

	// }
	
    // for x := 0; x < 10; x++ {
    //     for y := 0; y < 10; y++ {
    //         if y == 2 {
    //             // 跳转到标签
    //             goto breakHere
    //         }
    //     }
    // }
    // // 手动返回, 避免执行进入标签
    // return
    // // 标签
	// breakHere:
	// fmt.Println("done")
	// Go语言中 break 语句可以结束 for、switch 和 select 的代码块，另外 break 语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的 for、switch 和 select 的代码块上。

	// OuterLoop:
    // for i := 0; i < 2; i++ {
    //     for j := 0; j < 5; j++ {
    //         switch j {
    //         case 2:
    //             fmt.Println(i, j)
    //             break OuterLoop
    //         case 3:
    //             fmt.Println(i, j)
    //             break OuterLoop
    //         }
    //     }
	// }
	// Go语言中 continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用，在 continue 语句后添加标签时，表示开始标签对应的循环
	// OuterLoop:
    // for i := 0; i < 2; i++ {
    //     for j := 0; j < 5; j++ {
    //         switch j {
    //         case 2:
    //             fmt.Println(i, j)
    //             continue OuterLoop
    //         }
    //     }
    // }
	user.Ctx.WriteString("post表单提交成功:\n")
}

//普通函数,用户传入不同的匿名函数体可以实现对元素不同的遍历操作
func test(a int, b int, c string) (int, string) {
	//匿名函数
	func(data int) {
		fmt.Println("hello", data)
	}(100) //(100)是对匿名函数进行调用,参数100
	f := func(data1 int) {
		fmt.Printf("%v\n", data1)
	}
	f(90) //调用
	return b, c
}

//普通函数+匿名函数用作回调,(第一个参数是切片,第二个参数是匿名函数,匿名函数有一个int类型参数)
func test1(slice1 []int, f1 func(int)) {
	for _, v := range slice1 {
		f1(v)
	}
}

func huidiao(cs int, nmhs func(int, int)) {
	nmhs(cs, cs)
}
