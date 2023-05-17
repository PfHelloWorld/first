package main

import (
	"fmt"
	"tree"

)

type myTreeNode struct{
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode==nil || myNode.node==nil{
		return
	}

	left:=myTreeNode{myNode.node.Left}     //这里单独把左右当指针设置,因为这个函数循环是接收指针的
	right:=myTreeNode{myNode.node.Right}
	left.postOrder()                             //继续打开右边
	right.postOrder()
	myNode.node.Print()                          //输出函数
}



func main() {
    var root tree.Node
    root=tree.Node{Value:3}
    root.Left=&tree.Node{}
    root.Right=&tree.Node{5,nil,nil}
    root.Right.Left=new(tree.Node)
    root.Left.Right=tree.CreateNode(2)
    root.Right.Left.SetValue(4)

    root.Traverse()

    myRoot:=myTreeNode{&root}      //定义类
    fmt.Println()
    myRoot.postOrder()                   //调用自己写的方法
    fmt.Println()

    nodeCount := 0
    root.TraverseFunc(func(node *tree.Node){
    	nodeCount++
	})
    fmt.Println("Node count:",nodeCount)
}
