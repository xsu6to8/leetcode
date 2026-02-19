type MyStack struct {
    queue []int
}

func Constructor() MyStack {
    return MyStack{
        queue: make([]int, 0),
    }
}

// [push to back, peek/pop from front]

func (this *MyStack) Push(x int)  {
    // push to back
    this.queue = append(this.queue, x)
    
    // len of queue
    size := len(this.queue)
    
    // {1. pop FRONT -> 2. push BACK} 위에서 구한 size 전까지
    for i := 0; i < size-1; i++ {
        removed := this.queue[0]
        this.queue = this.queue[1:]
        
        this.queue = append(this.queue, removed)
    }
}


func (this *MyStack) Pop() int {
    val := this.queue[0]
    this.queue = this.queue[1:]

    return val
}


func (this *MyStack) Top() int {
    val := this.queue[0]

    return val
}


func (this *MyStack) Empty() bool {
    if len(this.queue) == 0 {
        return true
    } else {
        return false
    }
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */