package leet_code

type CircularQueue interface {
	MyCircularQueue(k int)     	//: 构造器，设置队列长度为 k
	Front()                    	//: 从队首获取元素。如果队列为空，返回 -1 。
	Rear()                     	//: 获取队尾元素。如果队列为空，返回 -1 。
	EnQueue(value int) 			//: 向循环队列插入一个元素。如果成功插入则返回真。
	DeQueue()                  	//: 从循环队列中删除一个元素。如果成功删除则返回真。
	IsEmpty()                  	//: 检查循环队列是否为空。
	IsFull()                   	//: 检查循环队列是否已满。
}


type CircleQueue struct{
	Queue []int
	Head int
	Tail int
	Size int
}


func NewCircleQueue(k int)*CircleQueue{
	return &CircleQueue{
		Queue:make([]int, 0, k),
		Head: -1,
		Tail: -1,
		Size: k,
	}
}

func (c *CircleQueue)IsFull()bool{
	if c == nil{
		return true
	}
	if len(c.Queue) == c.Size{
		return true
	}
	return false
}

func (c *CircleQueue)IsEmpty()bool{
	if c == nil{
		return false
	}
	if c.Head == c.Tail{
		return true
	}
	return false
}

func (c *CircleQueue)EnQueue(value int)bool{
	if c.IsFull(){
		return false
	}
	if c.IsEmpty(){
		c.Head = 0
	}
	c.Tail = (c.Tail + 1) % c.Size
	//c.Queue[c.Tail] = value
	c.Queue = append(c.Queue, value)
	return true
}

func (c *CircleQueue)DeQueue()bool{
	if c.IsEmpty(){
		return false
	}
	c.Head = (c.Head + 1) % c.Size
	return true
}

/** Get the front item from the queue. */
func (c *CircleQueue) Front() int {
	if c.IsEmpty(){
		return -1
	}
	return c.Queue[c.Head]
}

/** Get the last item from the queue. */
func (c *CircleQueue) Rear() int {
	if c.IsEmpty(){
		return -1
	}
	return c.Queue[c.Tail]
}


