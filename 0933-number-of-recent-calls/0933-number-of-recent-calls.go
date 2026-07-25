type RecentCounter struct {
    calls []int
}


func Constructor() RecentCounter {
    return RecentCounter{
        calls: make([]int, 0),
    }
}


func (this *RecentCounter) Ping(t int) int {
    this.calls = append(this.calls, t)

    cnt := 0
    for i := len(this.calls) - 1; i >= 0; i-- {
        if t - 3000 <= this.calls[i] {
            cnt++
        } else {
            this.calls = this.calls[i:]
            break
        }
    }

    return cnt
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */