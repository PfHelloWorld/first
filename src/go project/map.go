package main

import "fmt"

func main() {
	m:=map[string]string{   //定义map容器,输入key和value
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}

	m2:=make(map[string]int)     //map容器可以实现空定义

	var m3 map[string]int       //可以用var,也可以使用:=来创建map

	fmt.Println(m,m2,m3)


	//______________遍历______________
	for k,v:=range m{            //map容器是乱序的,所以输出时是乱序的
		fmt.Println(k,v)         //k接收key值,而v接收value值
	}
    //______________利用key值来寻找value_______________
	fmt.Println("Getting values")
    courseName:=m["course"]        //这里是用key值来寻找value值,把这个value值赋值给courseName
    fmt.Println(courseName)        //输出刚刚赋值的value
}
