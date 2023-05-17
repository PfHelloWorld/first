package main

import "fmt"

func printSlice(s []int){
	fmt.Printf("len=%d,cap=%d\n",len(s),cap(s))
}

func main() {

	var s[]int //不定义的切片是空长度

	for i:=0;i<100;i++{
		printSlice(s)
		s=append(s,2*i+1)     //尾插一些基数进去,返回的
	}
    fmt.Println(s)

	//__________自建切片长度___________

	s1 :=[]int{2,4,6,8}
	printSlice(s1)

	s2:=make([]int,16)       //自定义长度len

	s3:=make([]int,10,32)    //同时定义len,和cap
    printSlice(s2)
	printSlice(s3)

	fmt.Println("___________复制切片___________")
	copy(s2,s1)          //利用copy函数把s1复制到s2中
	printSlice(s2)

	fmt.Println("___________删除切片内元素___________")
	fmt.Println(s2)
	s2=append(s2[:3],s2[4:]...)      //利用...可以把一整段切片尾插进另一段切片中
	fmt.Println(s2)

    fmt.Println("头插法")
	front :=s2[0]       //可以看作是一个指向头的指针(视窗)
	s2=s2[1:]           //把截去头元素的数组重新复制给自己,相当于头删法

	fmt.Println("尾插法")
	tail :=s2[len(s2)-1]      //计算尾部长度-1,就是尾删要删的元素
    s2 = s2[:len(s2)-1]
    fmt.Println(front,tail)
    printSlice(s2)            //现在的s2删去头和尾了
}
