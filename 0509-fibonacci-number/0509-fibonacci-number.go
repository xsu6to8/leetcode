func fib(n int) int {
    var recursive func(curr int) int
    recursive = func(curr int) int {
        if curr == 0 {
            return 0
        } else if curr == 1 {
            return 1
        }

        return recursive(curr - 1) + recursive(curr - 2)
    }

    return recursive(n)
}