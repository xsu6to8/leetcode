type Stack struct {
    items []int
}

func (s *Stack) Push(n int) {
    s.items = append(s.items, n)
}
func (s *Stack) Pop() int{
    val := s.items[len(s.items) - 1]
    s.items = s.items[: len(s.items) - 1]
    return val
}
func (s *Stack) Peek() int{
    val := s.items[len(s.items) - 1]
    return val
}

func (s *Stack) Len() int {
    return len(s.items)
}
func (s *Stack) IsEmpty() bool {
    if s.Len() == 0 {
        return true
    }
    return false
}

func largestRectangleArea(heights []int) int {
    n := len(heights)

    // NSE: 오른쪽에서 처음으로 나보다 작은 막대의 인덱스 (+a 없으면 n)
    nse := make([]int, n)
    s1 := new(Stack)
    for i := n - 1; i >= 0; i-- {
        // 나보다 크거나 같은 막대는 내 왼쪽 막대들의 NSE가 될 수 없으니 버림
        for !s1.IsEmpty() && heights[s1.Peek()] >= heights[i] {
            s1.Pop()
        }
        if s1.IsEmpty() {
            nse[i] = n
        } else {
            nse[i] = s1.Peek()
        }
        s1.Push(i)
    }

    // PSE: 왼쪽에서 처음으로 나보다 작은 막대의 인덱스 (+a 없으면 -1)
    pse := make([]int, n)
    s2 := new(Stack)
    for i := 0; i < n; i++ {
        // 나보다 크거나 같은 막대는 내 오른쪽 막대들의 PSE가 될 수 없으니 버림
        for !s2.IsEmpty() && heights[s2.Peek()] >= heights[i] {
            s2.Pop()
        }
        if s2.IsEmpty() {
            pse[i] = -1
        } else {
            pse[i] = s2.Peek()
        }
        s2.Push(i)
    }

    // 각 막대를 "가장 낮은 막대"로 하는 직사각형 넓이 중 최댓값
    maxArea := 0
    for i := 0; i < n; i++ {
        currArea := heights[i] * (nse[i] - pse[i] - 1)
        maxArea = max(maxArea, currArea)
    }
    return maxArea
}