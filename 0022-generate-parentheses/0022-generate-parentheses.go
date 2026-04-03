func generateParenthesis(n int) []string {
	var res []string

	var backtrack func(curr string, open int, close int, depth int)
	backtrack = func(curr string, open int, close int, depth int) {
        // base condition
		if depth == n*2 {
			res = append(res, curr)
			return
		}

		if open < n {
			backtrack(curr+"(", open+1, close, depth+1)
		}

		if open > close {
			backtrack(curr+")", open, close+1, depth+1)
		}

	}

	backtrack("", 0, 0, 0)
	return res
}