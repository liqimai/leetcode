// 665. Non-decreasing Array
// https://leetcode.com/contest/leetcode-weekly-contest-47/problems/non-decreasing-array/
func checkPossibility(nums []int) bool {
    var modified bool

    for i,j := 0,1; i<len(nums)&&j<len(nums) ; {
        if nums[i] <= nums[j]{
        }else{
            if modified == true{
                return false
            }
            modified = true
            switch{
                case i == 0 :
                    nums[i] = nums[j]
                case nums[i-1] <= nums[j]:
                    nums[i] = nums[i-1]
                default:
                    nums[j] = nums[i]
            }
        }
        i++
        j++
    }
    return true
}
