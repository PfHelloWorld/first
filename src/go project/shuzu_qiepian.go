package main

import "fmt"

func swap(a,b *int){    //传入指针,可以修改实参
	*b,*a=*a,*b
}

func swap01(a,b int)(int,int) {    //或者使用返回值来交换
	return b,a
}

func Showarr(arr* [5]int  ){      //arr[5]是一个数据类型,arr[]又是另一个数据类型
    arr[0]=100
	for i,v:=range arr{
		fmt.Println(i,v)
	}
}

//-------------------------------------------------------------

func main() {
	a,b:=10,20
	swap(&a,&b)

	a,b=swap01(a,b)     //返回值把位置换回来
	fmt.Println(a,b)

	//-------------------------------------------------------


	var arr1 [5]int     //普通创建数组
	arr2 :=[3]int{1,3,5}   //冒号创建数组
	arr3 :=[...]int{2,4,6,8,10}  //无限大小创建数组
	var grid [4][5]int

	fmt.Println(arr1,arr2,arr3)   //数组可以随意打印
	fmt.Println(grid)          //打印二维数组

    fmt.Println("输出数组一")
	Showarr(&arr1)                 //传入地址,打印数组
	fmt.Println("打印数组3")
	Showarr(&arr3)                 //此时发现,数组0号元素都变成了100

	//________________________________________

	arr01:=[...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6]",arr01[2:6])      //这些都是切片
	fmt.Println("arr[2:6]",arr01[:6])       //切片非常自由
	fmt.Println("arr[2:6]",arr01[2:])       //可以自由定义范围
	fmt.Println("arr[2:6]",arr01[:])        //也可以当作实参传入函数,可被函数修改

}
