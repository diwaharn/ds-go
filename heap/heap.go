package heap

import "fmt"
import . "ds-go/util"

type Less func(i, j interface{}) bool

type Heap struct {
	N    int  //max number of N
	top  int  // current top
	less Less //function pointer
	data []interface{}
}

func Init_Heap(N int, func_cb Less) *Heap {
	heap := &Heap{N: N, top: 0}
	heap.data = make([]interface{}, N+1)
	heap.less = func_cb
	return heap
}

func (this *Heap) Count() int {
	return this.top
}

func (this *Heap) Insert(d interface{}) {
	Printf("Insert Enter %v\n", d)
	if this.top+1 == this.N {
		//array full increase the size
		fmt.Printf("Realloc happening %d to %d\n", this.N, this.N*2)
		this.N *= 2
		n := make([]interface{}, this.N+1)
		copy(n, this.data)
		this.data = n
	}
	this.top++
	this.data[this.top] = d
	this.Swim(this.top)
}

func (this *Heap) swap(i, j int) {
	this.data[i], this.data[j] = this.data[j], this.data[i]
}

// func (this *Heap) less(i, j int) bool {
// 	if this.data[i] < this.data[j] {
// 		return true
// 	}
// 	return false
// }

func (this *Heap) Swim(curr int) {
	parent := curr / 2
	Printf("Swim Enter : curr:%d, top:%d, parent:%d\n", curr, this.top, parent)
	for curr > 1 {
		Printf("\t{this.data[%d]=%d < this.data[%d]=%d}\n", parent, this.data[parent], curr, this.data[curr])
		if this.less(this.data[parent], this.data[curr]) {
			Printf("\tless so swap %d %d in array \t\t%v\n", parent, curr, this.data)
			this.swap(parent, curr)
			Printf("\tafter swap %d %d in array \t\t%v\n", parent, curr, this.data)
			curr = parent
			parent = curr / 2
		} else {
			break
		}
	}
}

func (this *Heap) MaxItem() (bool, interface{}) {
	if this.top < 1 {
		return false, 0
	}

	r := this.data[1]
	Printf("MaxItem :%d\n", r)
	//copy last item hear and sink
	this.data[1] = this.data[this.top]
	this.data[this.top] = 0 //not required.. just for better print
	this.top -= 1
	this.Sink(1)
	return true, r
}

func (this *Heap) Sink(curr int) {
	Printf("Sink Enter : curr:%d, top:%d\t\t%v\n", curr, this.top, this.data)
	var child1 int
	child1 = curr * 2
	Printf("\tChild : %d=%d, %d=%d\n", child1, this.data[child1], child1+1, this.data[child1+1])
	for child1 <= this.top {
		Printf("\tsink{this.data[%d]=%d < this.data[%d]=%d}\n", child1, this.data[child1], child1+1, this.data[child1+1])
		if child1+1 <= this.top && this.less(this.data[child1], this.data[child1+1]) {
			child1 = child1 + 1 //this is 2 check which path to take, important
		}
		Printf("\tsink swap before %d %d\n", curr, child1)
		if this.less(this.data[curr], this.data[child1]) {
			Printf("\tsink swap to do %d %d\n", curr, child1)
			this.swap(curr, child1)
			curr = child1
			child1 = curr * 2
		} else {
			break
		}
	}
	return
}

func (this *Heap) Print(s string) {
	fmt.Printf("%s : %v\n", s, this.data)
}
