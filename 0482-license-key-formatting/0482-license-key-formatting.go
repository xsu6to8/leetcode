import (
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	undashed := strings.ReplaceAll(s, "-", "")
	uppered := strings.ToUpper(undashed)

	if len(uppered) == 0 {
		return ""
	}

	resbyte := []byte{}
	cnt := 0
	for i := len(uppered) - 1; i >= 0; i-- {
		resbyte = append(resbyte, uppered[i])
		cnt++
		if cnt == k {
			resbyte = append(resbyte, '-')
			cnt = 0
		}
	}

    // 마지막에 '-'로 끝나는 case
	if len(resbyte) > 0 && resbyte[len(resbyte)-1] == '-' {
		resbyte = resbyte[:len(resbyte)-1]
	}

    // reversing
	left, right := 0, len(resbyte)-1
	for left < right {
		resbyte[left], resbyte[right] = resbyte[right], resbyte[left]
		left++
		right--
	}

	return string(resbyte)
}
