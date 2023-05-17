package tree

import "fmt"

type Node struct{
	Value int
	Left,Right *Node
}

//创建一个构造函数,本质其实是一个普通函数
func CreateNode(value int)*Node {
    return &Node{Value: value} //输入一个整型,构造一个整型
}

//接收者函数
func(node Node)Print(){ //普通接收者,拷贝对象进行输出
	fmt.Print(node.Value)

}

func (node *Node) SetValue(value int) { //指针接收者,引用对象
	if node ==nil{
		fmt.Println("Setting value to nil"+"node. Ignored.")
		return
	}
	node.Value=value
	//接收者构造函数
}

//接收者--遍历树状类对象
func(node *Node) Traverse(){
    node.TraverseFunc(func(node *Node){
    	node.Print()                         //定义好这个函数,传到TraverseFunc调用
	})
    fmt.Println()                             //换行
}

func (node *Node) TraverseFunc(f func(*Node))  {   //一个方法,传入一个接收函数
	if node == nil{
		return
	}
                                   //记住,,每次成功打开(不为空)会重新向左开始遍历
	node.Left.TraverseFunc(f)      //打开分支,直到向左边走到底
	f(node)                        //输出当前位置的内容
	node.Right.TraverseFunc(f)     //打开分支,直到向右边走到底
}




