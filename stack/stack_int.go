package stack

import "fmt"

type StackInt struct {
    N    int
    top  int
    data []int
}

func NewStackInt(n int) (s *StackInt) {
    s = &StackInt{N: n, top: 0}
    s.data = make([]int, n)
    return s
}

func (this *StackInt) Push(b int) {
    if this.top == this.N {
        b := make([]int, this.N)
        this.data = append(this.data, b[:]...)
        this.N *= 2
    }
    this.data[this.top] = b
    this.top++
}

func (this *StackInt) Top() (b int) {
    if this.top == 0 {
        return 0
    }
    b = this.data[this.top-1]
    return
}

func (this *StackInt) Pop() (b int) {
    if this.top == 0 {
        return 0
    }
    this.top--
    b = this.data[this.top]
    return
}

func (this *StackInt) NotEmpty() bool {
    if this.top > 0 {
        return true
    }
    return false
}

func (this *StackInt) PrintStack() {
    fmt.Printf("StackInt Contents:\n")
    for i := this.top - 1; i >= 0; i-- {
        fmt.Printf("%d ", this.data[i])
    }
    fmt.Printf("End\n")
}

func (this *StackInt) ReturnIntArray() []int {
    s := this.top

    r := make([]int, s)
    copy(r, this.data[:s])
    return r
}

