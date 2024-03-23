func reverseBase(x, rev int) int {
    if x == 0 {
        if rev > math.MaxInt32 { return 0 }
        if rev < math.MinInt32 { return 0 }
        return rev
    }
    return reverseBase(x / 10, (rev * 10) + x % 10)
}

func reverse(x int) int {
    return reverseBase(x, 0)
}
