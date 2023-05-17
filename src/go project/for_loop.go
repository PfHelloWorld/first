package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(n int)string  {
	result:=""
	for;n>0;n/=2{
		lsb:=n%2
		result=strconv.Itoa(lsb)+result   //先算小位再算大位
	}
	return result
}

func printFile(filename string)  {
	file,err:=os.Open(filename)      //要一行行读,需要先打开文件
	if err!=nil{
		panic(err)             //这里,如果文件没有成功打开,那就显示报错位置
	}

	scanner:=bufio.NewScanner(file)    //这里,先创建一个扫描器,指向文件开头的位置

	for scanner.Scan(){                //只写循环条件,可以直接当while语句使用
		fmt.Println(scanner.Text())    //输出扫描器的第一行
	}
}

func main() {
	fmt.Printf(converToBin(269))

	printFile("abc.txt")
}
