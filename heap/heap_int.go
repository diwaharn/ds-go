package heap

import "fmt"

type HeapInt struct {
	N    int //max number of N
	top  int // current top
	data []int
}

func HeapInt_Init(N int) *HeapInt {
	heap := &HeapInt{N: N, top: 0}
	heap.data = make([]int, N+1)
	heap.top = 0
	return heap
}

func (this *HeapInt) Insert(d int) {
	this.top++
	this.data[this.top] = d
	this.Swim(this.top)
}

func (this *HeapInt) swap(i, j int) {
	this.data[i], this.data[j] = this.data[j], this.data[i]
}

func (this *HeapInt) less(i, j int) bool {
	if this.data[i] < this.data[j] {
		return true
	}
	return false
}

func (this *HeapInt) Swim(curr int) {
	parent := curr / 2
	for curr > 1 {
		if this.less(parent, curr) {
			this.swap(parent, curr)
			curr = parent
			parent = curr / 2
		} else {
			break
		}
	}
}

func (this *HeapInt) MaxItem() (bool, int) {
	if this.top < 1 {
		return false, 0
	}

	r := this.data[1]

	//copy last item hear and sink
	this.data[1] = this.data[this.top]
	this.data[this.top] = 0 //not required.. just for better print
	this.top -= 1
	this.Sink(1)
	return true, r
}

func (this *HeapInt) Sink(curr int) {
	var child1 int
	child1 = curr * 2
	for child1 <= this.top {
		if child1 < this.top && this.less(child1, child1+1) {
			child1 = child1 + 1 //this is 2 check which path to take, important
		}
		if this.less(curr, child1) {
			this.swap(curr, child1)
			curr = child1
			child1 = curr * 2
		} else {
			break
		}
	}
	return
}

func (this *HeapInt) Print() {
	fmt.Printf("%v\n", this.data)
}
