package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"runtime"
)

//创建一个操作数函数,其中eval是传入的意思,而int是返回值类型
func eval(a,b int, op string)int{
	var result int
	switch op{
	case"+":
		result=a+b
	case"-":
		result=a-b
	case"*":
		result=a*b
	case"/":
		q,_:=div(a,b)    //不想要的返回值,直接用_接收就行
		return q
	default:
		panic("unsupported operator:"+op)
	}
	return result
}

//写一个多返回值的函数
func div(a,b int)(q,r int){         //体现思想,多返回,可命名性质
	return a/b,a%b           //直接返回
}

func div2(a,b int)(q,r int){  //这里相当于是直接定义过了q,r
	q=a/b
	r=a%b
	return                    //直接return,就可以直接返回这两个返回值
}

func apply(op func(int,int)int,a,b int)int  {    //接收函数,另外再传入几个操作数
	p:=reflect.ValueOf(op).Pointer()             //创造一个指针指向函数
	opName:=runtime.FuncForPC(p).Name()          //获取函数名
	fmt.Printf("Callint function %s with args"+"(%d,%d)",opName,a,b)

	return op(a,b)
}

func sum(numbers...int)int{    //体现思想:无限传值  意思是numbers这个数组无限放入int
	s:=0
	for i:=range numbers{      //range会把数组范围内的所有数组成员遍历
		s+=numbers[i]
	}
	return s
}


func main() {
	const filename="abc.txt"                  //创建一个文件名
	if contents,err:=ioutil.ReadFile(filename);err!=nil{
		fmt.Println(err)
	}  else { //如果文件打开了,就输出文件
		fmt.Printf("%s\n", contents)
	}
	eval(1,2,"+")

    fmt.Println(div(13,3))

	//这里直接在调用处定义了一个匿名函数(随意定义性质)
	fmt.Println(apply(
		func(a int,b int)int{
			return int(math.Pow(
				float64(a),float64(b)))       //转换数据类型
		},3,4))


}
