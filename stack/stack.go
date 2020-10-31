package stack

import "fmt"

type Stack struct {
    N    int
    top  int
    data []interface{}
}

func NewStack(n int) (s *Stack) {
    s = &Stack{N: n, top: 0}
    s.data = make([]interface{}, n)
    return s
}

func (this *Stack) Push(b interface{}) {
    if this.top == this.N {
        b := make([]interface{}, this.N)
        this.data = append(this.data, b[:]...)
        this.N *= 2
    }
    this.data[this.top] = b
    this.top++
}

func (this *Stack) Pop() (b interface{}) {
    if this.top == 0 {
        return ' '
    }
    this.top--
    b = this.data[this.top]
    return
}

func (this *Stack) NotEmpty() bool {
    if this.top > 0 {
        return true
    }
    return false
}

func (this *Stack) PrintStack() {
    fmt.Printf("Stack Contents:\n")
    for i := this.top - 1; i >= 0; i-- {
        fmt.Printf("%c ", this.data[i])
    }
    fmt.Printf("End\n")
}

