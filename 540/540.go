// 540. Single Element in a Sorted Array
// https://leetcode.com/problems/single-element-in-a-sorted-array/description/
func singleNonDuplicate(nums []int) int {
    //Our searching space is  [l, r)
    l,r := 0, len(nums)
    for {
        if l + 1== r {
            return nums[l]
        }
        var i = (r-l)/4*2 + l
        if nums[i] == nums[i+1]{
            l = i + 2
        }else{
            r = i + 1
        }
    }
}
