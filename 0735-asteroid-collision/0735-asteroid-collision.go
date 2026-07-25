type Stack []int

//IsEmpty - 스택이 비어있는지 확인하는 함수
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Push - 스택에 값을 추가하는 함수.
func (s *Stack) Push(data int) {
	*s = append(*s, data) // 스택 끝(top)에 값을 추가함.
}

//Pop - 스택 top 값을 제거하는 함수.
func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	} else {
		top := len(*s) - 1
		*s = (*s)[:top]   // 스택에 마지막 데이터 제거함
	}
}

//Peep - 스택 top 값을 삭제 없이 반환만 하는 함수.
func (s *Stack) Peep() int {
    top := len(*s) - 1
	data := (*s)[top] // top 위치에 있는 값을 가져 옴
	return data
}

func asteroidCollision(asteroids []int) []int {
    var s Stack

    for i := 0; i < len(asteroids); i++ {
        // 현재 stack이 빈 경우
        if s.IsEmpty() {
            s.Push(asteroids[i])
            continue
        }

        // 양수인 경우 : 그냥 삽입
        if asteroids[i] > 0 {
            s.Push(asteroids[i])
        }

        // 음수인 경우 : top의 값 비교하며
        // 0. 원래 empty (전부 pop해서 empty) : push + break 
        // 0. 같은 음수 : push + break
        // 1. 자기 절댓값보다 작은 것들 : pop + continue
        // 2. 절댓값 같은 거 : pop + break
        // 3. 큰 것들 만나면 : break
        if asteroids[i] < 0 {
            for {
                // case 0
                if s.IsEmpty() || s.Peep() < 0 {
                    s.Push(asteroids[i])
                    break
                }

                // case 1
                currTop := s.Peep()
                if currTop < -asteroids[i] {
                    s.Pop()
                    continue
                }

                // case 2
                if currTop == -asteroids[i] {
                    s.Pop()
                    break
                }

                // case 3
                if currTop > -asteroids[i] {
                    break
                }
            }
        }
    }

    var res []int
    for !s.IsEmpty() {
        curr := s.Peep()
        res = append(res, curr)
        s.Pop()
    }

    slices.Reverse(res)

    return res
}