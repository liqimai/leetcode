// 553. Optimal Division
// https://leetcode.com/problems/optimal-division/description/
import "strconv"
func optimalDivision(nums []int) string {
    var rval string
    switch len(nums) {
        case 0:
        case 1:
            rval = strconv.Itoa(nums[0])
        case 2:
            rval = strconv.Itoa(nums[0])+"/"+strconv.Itoa(nums[1])
        default:
        for i,v := range nums{
            switch{
                case i == 0:
                rval += strconv.Itoa(v)
                case i == 1:
                rval += "/("+strconv.Itoa(v)
                default:
                rval += "/"+strconv.Itoa(v)
            }
        }
        rval += ")"
    }
    return rval
}
