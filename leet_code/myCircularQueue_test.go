package leet_code

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestMyCircleQueue(t *testing.T){
	var circularQueue *CircleQueue  =  NewCircleQueue(3) // 设置长度为 3
	ok := circularQueue.EnQueue(1)  // 返回 true
	spew.Dump(ok, circularQueue, circularQueue.Queue)

	circularQueue.EnQueue(2)  // 返回 true
	spew.Dump(circularQueue.Queue)

	circularQueue.EnQueue(3)  // 返回 true
	spew.Dump(circularQueue.Queue)

	circularQueue.EnQueue(4)  // 返回 false，队列已满
	spew.Dump(circularQueue.Queue)

	circularQueue.Rear()  // 返回 3
	spew.Dump(circularQueue.Queue)

	circularQueue.IsFull()  // 返回 true
	spew.Dump(circularQueue.Queue)

	circularQueue.DeQueue()  // 返回 true
	spew.Dump(circularQueue.Queue)

	circularQueue.EnQueue(4)  // 返回 true
	spew.Dump(circularQueue.Queue)

	circularQueue.Rear()  // 返回 4
	spew.Dump(circularQueue.Queue)

}

func TestArrar(t *testing.T){
	arr := make([]int, 2, 3)
	arr[0] = 1
	spew.Dump(arr)
}