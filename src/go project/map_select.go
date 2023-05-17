package main

import "fmt"

func lengthOf_NoRepeat_Str(s string)int{
	lastOccured:=make(map[rune]int)         //创建一个"最后出现位置"空map
	start :=0                             //查找重复元素的起始位置
	maxLength:=0                          //最后找到的最长不重复为

	for i,ch:=range []rune(s){            //计划让i到x结束,并且把S强转换成byte
		if lastOccured[ch]>=start{        //lastOccured[ch]返回的是int,相当于用了key来查找位置
			start =lastOccured[ch]+1      //起始位置改变,改到重复元素前面的位置
		}
		if i-start+1>maxLength{           //如果现在新的长度大于上次寻找的长度,起始位置更改,最大长度也更改
			maxLength=i-start+1
		}
        lastOccured[ch]=i              //把重复位置设为当前遍历到的位置,下次循环继续查找是否重复
	}
	return maxLength                      //输出最大长度
}

func main() {
	s:=lengthOf_NoRepeat_Str("abcba")
    fmt.Println(s)
}
