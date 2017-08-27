// 413. Arithmetic Slices
// https://leetcode.com/problems/arithmetic-slices/description/
package main
import "fmt"

func numberOfArithmeticSlices(A []int) int {
    longestSubarrray := func (start int) int{
        if start >= len(A) - 1{
            return 1
        }
        diff := A[start+1] - A[start]
        for i:= start + 1; i < len(A); i++{
            if A[i] != A[i-1]+diff{
                return i-start
            }
        }
        return len(A) - start
    }

    var i = 0
    var count = 0
    for i < len(A) -1{
        l := longestSubarrray(i)
        fmt.Println(i, l)
        count += (l-1)*(l-2)/2
        i = i+l-1
    }
    return count
}

func main(){
    fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
}
