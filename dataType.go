package main

import "fmt"
/*
Go语言内置支持的基本数据类型： 布尔类型：bool
                               整型：int8、byte、int16、int、uint、uintptr等
                               浮点类型：float32、float64
                               复数类型：complex64、complex128
                               字符串：string
                               字符类型：rune
                               错误类型：error
Go语言支持复合类型
                    指针（pointer）
                    数组（array）
                    切片（slice）
                    字典（map）
                    通道（chan）
                    结构体（struct）
                    接口（interface）

*/
//Go与PHP类型比较（下面5点区别）
/*
1、多了字符类型（单个字符）、错误类型和复数类型。
2、PHP通过系统级配置函数error_reporting定义应用的错误报告级别，不区分单独的字符与字符串类型，
3、Go还对整型的精度及是否有符号（正数还是负数）做了区分
4、php只有一个int类型标识整型数据，
5、另外php通过float和double来区分浮点精度

PHP并不支持指针类型，对于数组，切片，字典，php则通过数组类型一网打尽，
通道，结构体，接口，php也不支持，
    通道类型主要用于并发编程
    结构体类似php中的类（class）
   之后要详细了解结构体和接口类型的使用

*/
//-------------------------布尔类型---------------------------------
/*
与PHP不同的是，Go是强类型语言，变量类型一旦确定，就不能将其他类型的值赋值给该变量。因此，布尔类型不能接受其他类型的赋值
也不支持自动或强制的类型转换。下面示范列子
*/

func main() {
	//Go语言的错误示例，但是PHP则可以这么使用
	var b bool
	b = 1  //编译错误
	b = bool(1)  //编译错误

	//Go正确使用
	var c  bool
	c = (1 != 0)   //编译正确
	fmt.Println("结果：",c)


	/*
	由于强类型的缘故，Go 语言在进行布尔值真假判断时，对值的类型有严格限制，
	在 PHP 这种弱类型语言中，以下这些值在进行布尔值判断的时候（使用非严格的 == 比较符）都会被认为是 false：
					布尔值 FALSE 本身
					整型值 0（零）
					浮点型值 0.0（零）
					空字符串，以及字符串 "0"
					不包括任何元素的数组
					特殊类型 NULL（包括尚未赋值的变量）
					从空标记生成的 SimpleXML 对象
	而在 Go 语言中则不然，甚至不同类型的值直接不能使用 == 或 != 运算符进行比较，在编译期就会报错，比如下面这段代码：
	*/
	d := (false == 0)    //报错如下
	/*
	    cannot convert 0 (type untyped number) to type bool
		invalid operation: false == 0 (mismatched types bool and int)
	*/

//同样，！运算符也不能作用与非布尔类型值

}
