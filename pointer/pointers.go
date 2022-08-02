package main

import "fmt"

func main() {
	age := 30
	fmt.Println(age)

	// var myAge *int
	// myAge=age
	myAge := &age
	//不加*访问地址
	fmt.Println(myAge)
	//加*访问值
	fmt.Println(*myAge)
	//加*重新赋值 内存地址不变 age的值也被更改，因为age只是显示保存内存地址
	*myAge = 18
	fmt.Println(myAge)
	fmt.Println(*myAge)
	fmt.Println(age)
	//传入指针 指针对应的值能被修改
	fmt.Println(double(myAge))
	fmt.Println(*myAge)
}

//传入指针
func double(number *int) int {
	result := *number * 2
	*number = 100
	return result
}
