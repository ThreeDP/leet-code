func firstMissingPositive(nums []int) int {
    size := len(nums) + 1
    hash := make ([]bool, size)
    hash[0] = true
    for _, num := range nums {
        if num > 0 && num < size {
            hash[num] = true
        }
    }
    for i, v := range hash {
        if v == false {
            return i 
        }
    }
    return size
}
