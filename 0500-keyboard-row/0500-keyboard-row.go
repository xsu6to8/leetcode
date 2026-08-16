func findWords(words []string) []string {
    firRow := make(map[byte]struct{})
    for _, b := range []byte("qwertyuiop") {
        firRow[b] = struct{}{}
    }

    secRow := make(map[byte]struct{})
    for _, b := range []byte("asdfghjkl") {
        secRow[b] = struct{}{}
    }

    thiRow := make(map[byte]struct{})
    for _, b := range []byte("zxcvbnm") {
        thiRow[b] = struct{}{}
    }

    var res []string
    for _, origin := range words {
        s := strings.ToLower(origin)
        sByte := []byte(s)
        firByte := sByte[0]

        var targetRow map[byte]struct{}
        if _, ok := firRow[firByte]; ok {
            targetRow = firRow
        } else if _, ok := secRow[firByte]; ok {
            targetRow = secRow
        } else {
            targetRow = thiRow
        }

        inRow := true
        for i := 1; i < len(sByte); i++ {
            if _, ok := targetRow[sByte[i]]; !ok {
                inRow = false
                break
            }
            inRow = true
        }

        if inRow {
            res = append(res, origin)
        }
    }

    return res
}