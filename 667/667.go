//667. Beautiful Arrangement II
//https://leetcode.com/contest/leetcode-weekly-contest-47/problems/beautiful-arrangement-ii/
//The only correct answer for constructArray(5 ,4) is [1, 5, 2, 4, 3]
//This is the key to this problem 
package main

import "fmt"
func main(){
    fmt.Println(constructArray(5 ,1))
}

func constructArray(n int, k int) []int {
    nums := make([]int, n)
    // l,r :=
    last := 0
    inc  := false
    for i:= 0; i < n; i++{
        if i < k{
            switch(i%2){
                case 0:
                    nums[i] = i/2+1
                    inc = true
                    last = nums[i]
                case 1:
                    nums[i] = n-(i/2)
                    inc = false
                    last = nums[i]
            }
        }else{
            if inc {
                last ++
                nums[i] = last
            }else{
                last --
                nums[i] = last
            }
        }
    }
    return nums
}
