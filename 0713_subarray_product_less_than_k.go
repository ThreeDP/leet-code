func numSubarrayProductLessThanK(nums []int, k int) int {
    if k == 0 { return 0 }
    size := len(nums)
    amount := 0
    for i := 0; i < size; i++ {
        max := 1
        options := nums[i:]
        for _, num := range options {
            max *= num;
            if max >= k {
                break
            }
            amount++
        }
    } 
    return amount
}
