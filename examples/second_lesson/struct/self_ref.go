package main

func main() {

}

type Node struct {
	//自引用只能使用指针  ，因为编译时需要确定结构体的内存大小，而指针的大小固定（不讨论压缩指针的话）
	//left Node
	//right Node

	left  *Node
	right *Node

	// 这个也会报错
	//nn NodeNode
}

type NodeNode struct {
	node Node
}
