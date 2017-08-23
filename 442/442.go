// https://leetcode.com/problems/find-all-duplicates-in-an-array/
// Firstly, we put each element x in nums[x - 1].
// Secondly, we check through the array. If a number x doesn't present in nums[x - 1], then x is absent.
// The swap here should be down carefully
package main
import "fmt"

func findDuplicates(nums []int) []int {
    var res []int
    for i := 0; i < len(nums); {
        if nums[i] != nums[nums[i] - 1]{
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }else{
            i++
        }
        // fmt.Println(nums)
    }
    for i, v := range nums {
        if v != i+1 {
            res = append(res, v)
        }
    }
    return res
}

func main(){
    fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}
