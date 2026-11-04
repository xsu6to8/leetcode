func reverseStr(s string, k int) string {  
    sByte := []byte(s)
	for i := 0; i < len(s); i += 2*k {
        if i + k > len(s) {
            for start, end := i, len(s)-1; start <= end; start, end = start+1, end-1 {
		        sByte[start], sByte[end] = sByte[end], sByte[start]
	        }
            break
        } else { // swaping - {default case}
            for start, end := i, i + k-1; start <= end; start, end = start+1, end-1 {
		        sByte[start], sByte[end] = sByte[end], sByte[start]
	        }
        }   
    }
    
    res := string(sByte)
    return res
}