package leet_code

type CircularQueue interface {
	MyCircularQueue(k int)     //: 构造器，设置队列长度为 k
	Front()                    //: 从队首获取元素。如果队列为空，返回 -1 。
	Rear()                     //: 获取队尾元素。如果队列为空，返回 -1 。
	EnQueue(value int) //: 向循环队列插入一个元素。如果成功插入则返回真。
	DeQueue()                  //: 从循环队列中删除一个元素。如果成功删除则返回真。
	IsEmpty()                  //: 检查循环队列是否为空。
	IsFull()                   //: 检查循环队列是否已满。
}


type CircleQueue struct{
	Queue []int
	Head int
	Tail int
	Len int
	Cap int

}


func Constructor(k int)CircleQueue{
	return CircleQueue{
		Queue:make([]int, 0, k),
		Head: 0,
		Tail: 0,
		Len: 0,
		Cap: k,
	}
}

func (c *CircleQueue)IsFull()bool{
	if c == nil{
		return true
	}
	if (c.Tail + 1) % c.Cap == c.Head{
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
	c.Queue[c.Tail] = value
	c.Tail = (c.Tail + 1) % c.Cap
	c.Len += 1
	return true
}

func (c *CircleQueue)DeQueue()bool{
	if c.IsEmpty(){
		return false
	}
	c.Queue[c.Head] = 0
	return true
}

/** Get the front item from the queue. */
func (c *CircleQueue) Front() int {
	if c.IsEmpty(){
		return 0
	}
	return c.Queue[c.Head]
}

/** Get the last item from the queue. */
func (c *CircleQueue) Rear() int {
	if c.IsEmpty(){
		return 0
	}
	return c.Queue[c.Tail]
}



