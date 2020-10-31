package queue

import "fmt"

type Queue struct {
	N     int
	front int
	tail  int
	data  []int
}

func NewQueue(n int) (q *Queue) {
	q = &Queue{N: n, front: 0, tail: 0}
	q.data = make([]int, n)
	return q
}

func (this *Queue) IsEmpty() bool {
	if this.front == this.tail {
		return true
	}
	return false
}

func (this *Queue) Count() int {
	if this.front == this.tail {
		return 0
	}
	return (this.tail - this.front)
}

func (this *Queue) ResizeQueue() {
	//end of the array, resize
	c := this.Count()
	if c < this.N/2 {
		j := this.tail
		for i := 0; i < c; i++ {
			this.data[i] = this.data[j]
		}
		this.tail = 0
		this.front = c
	} else {
		d := make([]int, this.N*2)
		copy(d, this.data)
		this.data = d
		this.N = this.N * 2
	}
}

func (this *Queue) Add(n int) {
	if this.front+1 == this.N {
		this.ResizeQueue()
	}

	this.data[this.front+1] = n
	this.front++
}

func (this *Queue) Tail() int {
	if this.front == this.tail {
		return 0
	}
	this.front++
	return this.data[this.front-1]
}

func (this *Queue) Print() {
	fmt.Printf("Queue Print : N:%d, tail:%d, front:%d\n%v\n", this.N, this.tail, this.front, this.data)
}
