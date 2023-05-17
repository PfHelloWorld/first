package main

import "fmt"

func updateSlice(s []int){
	s[0]=100                  //切片可以修改原本参数组的原始值
	                          //相当于视窗,可以看见值并且进行修改
}

func main() {
	arr :=[...]int{0,1,2,3,4,5,6,7}
	//fmt.Println(arr)
	//fmt.Println("arr[...] =",arr[2:6])

	fmt.Println("_______视窗还可以再次打开自己_______")
	fmt.Println("Reslice")
	s2 :=arr[:5]
	fmt.Println(s2)
	s2 =s2[2:]            //重复打开,相当于是分多次定义一个视窗
	fmt.Println(s2)

	fmt.Println("__________Extending slice_________")
    fmt.Println("arr =",arr)
	s1 :=arr[2:6]
	s2 = s1[3:5]      //如果视窗读取长度不够,那就读原数组内容
	//len是目前视窗长度,而cap是视窗现长度加上可向后阅读长度
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",s1,len(s1),cap(s1))
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n",s2,len(s2),cap(s2))

	fmt.Println("_____切片可以使用尾插法进行添加元素_____")

	s1 =arr[2:6]
	s2 =s1[3:5]
	s3 :=append(s2,10)     //append是尾插法
	s4 :=append(s3,11)     //为s3提供尾插,对另一个切片使用尾插不会影响其原数组
	s5 :=append(s4,12)
	fmt.Println(s1,s2,s3,s4,s5)
	fmt.Println(arr)

}
