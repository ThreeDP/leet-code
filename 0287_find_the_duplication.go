func findDuplicate(nums []int) int {
    size := len(nums)
    hash := [1000000]bool{}
    for i := 0; i < size; i++ {
        if hash[nums[i]] {
            return nums[i];
        } else {
            hash[nums[i]] = true
        }
    }
    return -1
}
