package main

const (          //iota 被重置为0
	c0 = iota    // 0
	c1 = iota    // 1
	c2 = iota    //  2
)
//表达式一样，可以省略为下列写法
const (
	a0 = iota
	a1
	a2
)

//---------------------分割线------------------------
const (
	u = iota * 2    //0
	v = iota * 2    //2
	w = iota * 2    //4
)
//表达式可以省略为下列写法
const (
	a = iota * 2
	b
	c
)
//----------------------分割线----------------------------
const x  = iota;    //0

const y  = iota    //0

// ------------------------分割线------------------------------

//以上是常量的定义
//下面是枚举的定义
/*
枚举中包含了一系列相关的常量，例如关于一个星期中每天的定义，Go语言并不支持其他语言用于
表示枚举的enum关键字，而是通过在const 后跟一对圆括号定义一组常量的方式来实现枚举
*/
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)





//下面是打印的结果，没有空格，所以是连起来的
func main() {
	print(c0,c1,c2)   //012
	print(u,v,w)     //024
    print(x,y)    //00
    print(Sunday,Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,numberOfDays)
}
