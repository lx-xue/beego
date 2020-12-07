package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var QuanJuBianLiang bool = true
var a, b int = 1, 2
var c string = "cc"
var index *string

const (
	Unknow = 0
	Male   = 2
)

//声明结构体
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

//定义结构体Books的指针变量struct_pointer
var struct_pointer *Books
var balance = [...]float32{1000.0, 2, 3.4, 7.0, 50.0}

//var balance [10] float32
//初始化函数,go 自动加载函数
// func init() {

// 	book3 := Books{
// 		title:   "ssss",
// 		author:  "xxxxx",
// 		subject: "string",
// 		book_id: 22,
// 	}
// 	fmt.Printf("%v\n %v\n", book3.title, &book3.title)
// 	var Book1 Books /* 声明 Book1 为 Books 类型 */
// 	var Book2 Books /* 声明 Book2 为 Books 类型 */
// 	/* book 1 描述 */
// 	Book1.title = "Go 语言"
// 	Book1.author = "www.runoob.com"
// 	Book1.subject = "Go 语言教程"
// 	Book1.book_id = 6495407

// 	/* book 2 描述 */
// 	Book2.title = "Python 教程"
// 	Book2.author = "www.runoob.com"
// 	Book2.subject = "Python 语言教程"
// 	Book2.book_id = 6495700
// 	fmt.Println(Book1.title)
// 	printBook(Book1)

// 	fmt.Printf("%v", balance[0])
// 	var ff int = 1
// 	for ff < 6 {
// 		ff++
// 		fmt.Printf("666\n")
// 	}
// 	f := index
// 	d := "12"
// 	index = &c
// 	g := index
// 	h := *(index)
// 	fmt.Printf("%v %v %v %v %v", f, d, index, g, h)
// }

//参数为结构体
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
func test(u int, o string) (int, int, string) {
	fmt.Printf("%v %v", u, o)
	l, m, n := 1, 195, "aaaa"
	return l, m, n
}

// 调用器接口
type Invoker interface {
	// 需要实现一个Call方法
	Call(interface{})
}

// 结构体类型
type Struct struct {
}

//实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 函数定义为类型
type FuncCaller func(interface{})

// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	// 调用f函数本体
	f(p)
}

// //http包中包含有Handler接口定义
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// // Handler 用于定义每个 HTTP 的请求和响应的处理过程
// // 同时，也可以使用处理函数实现接口
// type HandlerFunc func(ResponseWriter, *Request)

// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
// 	f(w, f) //函数变量自调用
// }

// // 要使用闭包实现默认的http请求处理,可以使用http.HandleFunc()函数
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
// 	DefaultServeMux.HandleFunc(pattern, handler)
// }

// // 而 DefaultServeMux 是 ServeMux 结构，拥有 HandleFunc() 方法，定义如下
// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
// 	mux.Handle(pattern, HandlerFunc(handler))
// }

// 闭包即使匿名函数
//累加器函数
func Accumulate(value int) func() int {
	//返回一个闭包
	return func() int {
		value++      //累加
		return value //返回一个int累加值
	}
}

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

//可变参数
func myfunc(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

var (
	valueByKey = make(map[interface{}]interface{})
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

//对共享资源加锁解锁处理
func readValue(key interface{}) interface{} {
	//对共享资源加锁
	valueByKeyGuard.Lock()
	//取值
	v := valueByKey[key]
	// fmt.Printf("%T\n", v)
	// 对共享资源解锁
	valueByKeyGuard.Unlock()
	// 返回值
	return v
}

//使用延迟并发解锁
func readValue1(key string) interface{} {
	valueByKeyGuard.Lock()
	valueByKey["aa"] = 888
	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

// 根据文件名查询其大小
func fileSize(filename string) int64 {
	// 根据文件名打开文件, 返回文件句柄和错误
	f, err := os.Open(filename)
	// 如果打开时发生错误, 返回文件大小为0
	if err != nil {
		return 0
	}
	//文件打开成功后,使用延迟调用
	defer f.Close()
	// 取文件状态信息
	info, err := f.Stat()

	if err != nil {
		return 0
	}
	size := info.Size()
	return size
}

//递归函数
func fibonacci(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

// 当编译正则表达式发生错误时，使用 panic 触发宕机，该函数适用于直接使用正则表达式而无须处理正则表达式错误的情况。
// func MustCompile(str string) *Regexp {
// regexp, error := Compile(str)
// if error != nil {
// 	panic(`regexp: Compile(` + quote(str) + `): ` + error.Error())
// }
// return regexp
// }
//使用struct定义一个单项链表
// type Node struct {
// 	data int
// 	next *node
// }

// func Shownode(p *Node) { //遍历
// 	for p != nil {
// 		fmt.Println(*p)
// 		p = p.next //移动指针
// 	}
// }

// type error interface{
// 	Error() string
// }
// func Sqrt(f float64) (float64,error){
// 	if f < 0 {
// 		return 0,errors.new("math: square root of negative number")
// 	}
// }
// //定义一个divideError结构
// type DivideError struct{
// 	dividee int
// 	divider int
// }
// func (de *DivideError) Error() string{
// 	strFormat :=`
// 	Cannot proceed,the divider is zero.
// 	dividee: %d
// 	divider: 0
// 	`
// 	return fmt.Sprintf(strFormat,de.dividee)
// }
//并发
// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
//测试下协程阻塞
func sum(s []int, c chan int) {
	for _, v := range s {
		c <- v //把sum发送到通道c
	}

}
func HelloMakeChanSize() {
	size := 0
	c1 := make(chan int, size)
	go func() {
		for i := 0; i < 4; i++ {
			val := i*10 + 7
			fmt.Println(time.Now(), "<- ", val, "at", i)
			c1 <- i*10 + 7
		}
		c1 <- 0
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("After Sleep")
	for val := range c1 {
		fmt.Println(time.Now(), "receive:", val)
		if val == 0 {
			break
		}
	}
}

//go 区分定义类型和定义变量的区别
func main() {
	HelloMakeChanSize()
	// s := []int{7, 2, 8, -9, 4, 0}
	// c := make(chan int)
	// go sum(s, c) //切出s前一半
	// // go sum(s[len(s)/2:], c) //切出s后一半
	// x, y, z, m := <-c, <-c, <-c, <-c //从通道c中接受
	// fmt.Println(x, y, z, m)
	//输出是无序的,应为是两个线程
	// go say("world")
	// say("hello")
	// var head = new(Node)
	// fmt.Printf("%v\n",head)
	// head.data = 1
	// var node1 = new(Node)
	// node1.data = 2
	// panic("crash")
	// var e error
	// // 创建一个错误实例，包含文件名和行号
	// e = newParseError("main.go", 1)
	// // 通过error接口查看错误描述
	// fmt.Println(e.Error())
	// // 根据错误接口具体的类型，获取详细错误信息
	// switch detail := e.(type) {
	// case *ParseError: // 这是一个解析错误
	// 	fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	// default: // 其他类型的错误
	// 	fmt.Println("other error")
	// }
	// var i int = 4
	// fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
	// // 递归
	// result := 0
	// for i := 1; i <= 10; i++ {
	// 	result = fibonacci(i)
	// 	fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	// }

	// rv := readValue("aa")
	// rv1 := readValue1("aa")
	// fmt.Printf("%v\n %v\n", rv, rv1)
	// //延迟执行语句,类似栈储存先进后出
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)

	// var v1 int = 1
	// var v2 int64 = 234
	// var v3 string = "hello"
	// var v4 float32 = 1.234
	// myfunc(1, 234, "hello", 1.234)
	// // 创建一个玩家生成器
	// generator := playerGen("high noon")
	// // 返回玩家的名字和血量
	// name, hp := generator()
	// // 打印值
	// fmt.Println(name, hp)
	// //赋值
	// accumulator := Accumulate(1)

	// //第一次调用
	// fmt.Println(accumulator())
	// //第二次调用
	// fmt.Println(accumulator())
	// //第三次调用
	// fmt.Println(accumulator())
	// // 打印累加器的函数地址
	// fmt.Printf("%v\n", &accumulator)
	// // 创建一个累加器, 初始值为1
	// accumulator2 := Accumulate(10)
	// // 累加1并打印
	// fmt.Println(accumulator2())
	// // 打印累加器的函数地址
	// fmt.Printf("%v\n", &accumulator2)
	// fmt.Println(accumulator())
	// // 声明接口变量
	// var invoker Invoker
	// // 实例化结构体
	// s := new(Struct)
	// // 将实例化的结构体赋值到接口
	// invoker = s
	// // 使用接口调用实例化结构体的方法Struct.Call
	// invoker.Call("hello")
	// // 将匿名函数转为FuncCaller类型，再赋值给接口
	// invoker = FuncCaller(func(v interface{}) {
	// 	fmt.Println("from function", v)
	// })
	// // 使用接口调用FuncCaller.Call，内部会调用函数本体
	// invoker.Call("hello")
}
