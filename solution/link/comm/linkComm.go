package comm

import "fmt"

/*
实现单向链表的 增、删、改、查 操作
*/

// 定义节点结构体
type Node struct {
	Data int	// 接口类型的数据data
	Next *Node			// 下一个节点
}

// 创建节点函数
func CreateNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

// 定义单向链表结构体
type LinkList struct {
	Header *Node	// 头节点，指向第一个节点
	Length int		// 链表长度
}

// 创建链表方法
func CreateLinkList() *LinkList {
	return &LinkList{
		Header: CreateNode(0),
		Length: 0,
	}
}

// 在头节点插入数据的方法
func (l *LinkList) AddInHead(data int) {
	newNode := CreateNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = newNode
	}else {
		newNode.Next = l.Header
		l.Header = newNode		// 头指针指向新加的
	}
}

// 在链表尾部添加元素
func (l *LinkList) Append(data int) {
	newNode := CreateNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = newNode
		return
	}
	// 遍历链表到最后一个节点，再进行添加​
	current := l.Header
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode	//把新节点地址给最后一个节点的Next
}

// 在链表中间 i位置 插入元素
func (l *LinkList) Insert(i int, data int) {
	newNode := CreateNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = newNode
		return
	}
	// 插入位置为首位时，直接在头部加入该元素即可
	if i == 0 {
		l.AddInHead(data)
		return
	}
	// 插入特殊位置处理，默认在尾部添加，也可校验返回错误提示信息
	if i < 0 || i >= l.Length {
		l.Append(data)
		return
	}
	// 遍历下一个节点，直到遍历的位置等于需要插入的i位置
	current := l.Header
	j := 1
	for j < i {
		current = current.Next
		j++
	}
	// 插入节点元素
	newNode.Next = current.Next
	current.Next = newNode
}

// 删除第 i 个节点的元素
func (l *LinkList) Delete(i int) {
	defer func() {
		l.Length--
	}()

	if l.Length == 0 || i < 0 || i > l.Length-1 {
		fmt.Println("i is wrongful")
		return
	}

	if i == 0 {
		l.Header = l.Header.Next
		return
	}

	j := 1
	current := l.Header
	for j < i && current.Next != nil {
		current = current.Next
		j++
	}
	if current.Next == nil {
		current = nil
		return
	}
	current.Next = current.Next.Next

}

// 遍历单向链表并打印出来
func (l *LinkList) ScanLinkList() {
	current := l.Header
	i := 1
	for current.Next != nil {
		fmt.Printf("第 %d 个节点的值为：%d\n", i, current.Data)
		current = current.Next
		i++
	}
	fmt.Printf("第 %d 个节点的值为：%d\n", i, current.Data)
}

// 单向链表翻转
func (l *LinkList) ReverseList() {
	current := l.Header
	var pre *Node
	for current != nil {
		current, current.Next, pre = current.Next, pre, current
	}
	l.Header = pre
}

// 传入头节点，打印整个链表
func PrintNodeByHead(head *Node) {
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}


