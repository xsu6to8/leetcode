func reverseString(s []byte)  {
    l := len(s)
    for i := 0; i < l/2; i++ {
        temp := s[i]
        s[i] = s[l - 1 - i]
        s[l - 1 - i] = temp
    }
}