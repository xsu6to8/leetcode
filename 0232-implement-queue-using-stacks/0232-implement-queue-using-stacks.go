type MyQueue struct {
    inStack  []int
    outStack []int
}

func Constructor() MyQueue {
    return MyQueue{
        inStack:  make([]int, 0),
        outStack: make([]int, 0),
    }
}

func (this *MyQueue) Push(x int)  {
    this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
    for i := len(this.inStack) - 1; i > 0; i-- {
        this.outStack = append(this.outStack, this.inStack[i])
    }

    tmp := this.inStack[0]
    this.inStack = make([]int, 0)
    
    for i := len(this.outStack) - 1; i >= 0; i-- {
        this.inStack = append(this.inStack, this.outStack[i])
    }
    
    this.outStack = make([]int, 0) 

    return tmp
}

func (this *MyQueue) Peek() int {
    return this.inStack[0]
}

func (this *MyQueue) Empty() bool {
    if len(this.inStack) == 0 {
        return true
    } else {
        return false
    }
}