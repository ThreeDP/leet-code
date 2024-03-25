func findDuplicates(nums []int) []int {
    var hash [1000000]bool
    var res []int
    for _, num := range nums {
        if hash[num] == true {
            res = append(res, num)
        } else {
            hash[num] = true
        }
    }
    return res
}
