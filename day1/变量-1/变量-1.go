package main

import "fmt"

// 变量就如同一个事物的属性，你可以去定义
func main() {
	//1.变量的声明
	var age int
	//2.变量的赋值
	age = 18
	//3.变量的使用
	fmt.Println("age = ", age)

	//声明和赋值可以合成一句：
	var age2 int = 19
	fmt.Println("age2 = ", age2)

	// var age int = 20;
	// fmt.Println("age = ",age);
	/*变量的重复定义会报错：
	  # command-line-arguments
	  .\变量-1.go:16:6: age redeclared in this block
	                  previous declaration at .\变量-1.go:6:6
	*/

	//不可以在赋值的时候给与不匹配的类型
	//var num int = 12.56
	//fmt.Println("num = ", num)

}
