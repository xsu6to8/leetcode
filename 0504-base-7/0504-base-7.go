func convertToBase7(num int) string {
    num64 := int64(num)
    return strconv.FormatInt(num64, 7)
}