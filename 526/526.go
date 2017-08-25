// 526. Beautiful Arrangement
// https://leetcode.com/problems/beautiful-arrangement/description/
package main

import (
    "fmt"
    // "errors"
    // "sort"
)
import (
    "strconv"
    "os"
    "time"
    // "reflect"
    // "unsafe"
)

func main() {
    var start = time.Now()
    var n, err = strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }
    for i:=1; i<=n ; i++{
	       fmt.Printf("countArrangement(%v)= %v\n",i,countArrangement(i))
    }
    fmt.Println(time.Since(start))
}

func countArrangement(N int) int {
    visited := make([]bool, N + 1)
    count, total := 0,0
    calculate(N, 1, visited, &count, &total)
    fmt.Printf("%v/%v=%v\n", count,total,float64(count)/float64(total))
    return count
}
func calculate(N int, pos int, visited []bool, count *int, total *int) {
    *total += 1
    if pos > N {
        *count++
        return
    }
    for i := 1; i <= N; i++ {
        if (!visited[i] && (pos % i == 0 || i % pos == 0)) {
            visited[i] = true
            calculate(N, pos + 1, visited, count, total)
            visited[i] = false
        }
    }
}

// func swapWithRightMostSuccessor(nums []int, pivot int){
//     //find rightmost successor
//     var rs int
//     for i := pivot + 1; i < len(nums); i++ {
//         if nums[i] > nums[pivot] {
//             rs = i
//         }
//     }
//     //swap
//     nums[pivot], nums[rs] = nums[rs], nums[pivot]
// }
//
// func generateNextPermutation(nums []int) error{
//     //find the start of the non-increasing sufix
//     var ss = len(nums) -1  //sufix start
//     for ss > 0 && nums[ss -1] >= nums[ss] {
//         ss--
//     }
//     // If it's already the last one
//     if ss == 0 {
//         return errors.New("it's already the last one")
//     }
//     swapWithRightMostSuccessor(nums, ss-1)
//     //reverse
//     for i,j := ss, len(nums)-1; j>i ; {
//         nums[i], nums[j] = nums[j], nums[i]
//         j--
//         i++
//     }
//     return nil
// }

// func countArrangement(N int) int {
//     var count = 0
//     var total = 0
//     var nums = make([]int, N)
//     for i := range nums{
//         nums[i] = i + 1
//     }
//     for {
//         total++
//         // find the black goat
//         var pi = 0
//         for ; pi < len(nums) && ((pi+1) % nums[pi] == 0 || nums[pi] % (pi+1) == 0); {
//             pi++
//         }
//         if pi == len(nums) {
//             // if we did not find
//             count ++
//         }else{
//             // if we find it, skip all the permutations that start with it
//             sort.Sort(sort.Reverse(sort.IntSlice(nums[pi+1:])))
//         }
//         // Next Permutation
//         if err := generateNextPermutation(nums); err != nil{
//             break
//         }
//     }
//     fmt.Printf("total=%v, count/total=%v\n", total, float64(count)/float64(total))
//     return count
// }



// func countArrangement(N int) int {
//     var valid = make([][]bool, N)
//     for i := range valid {
//         valid[i] = make([]bool, N)
//         for j := range valid[i] {
//             valid[i][j] = (i + 1) % (j + 1) == 0 || (j + 1) % (i + 1) == 0
//         }
//         // fmt.Println(valid[i])
//     }
//     // p := (*reflect.SliceHeader)(unsafe.Pointer(&valid))
//     // fmt.Printf("p.len, p.cap, p.data = %v, %v, %v\n", p.Len, p.Cap, p.Data)
//
//     var nums = make([]int, N)
//     for i := range nums{
//         nums[i] = i + 1
//     }
//     var count = 0
//     var total = 0
//     for {
//         total++
//         // fmt.Println(nums)
//         var index = noRepeat(nums)
//         if index == len(nums){
//             count++
//             index--
//         }
//         if carry := addOne(nums, &valid, index); carry {
//             break
//         }
//     }
//     fmt.Printf("total=%v, count/total=%v\n", total, float64(count)/float64(total))
//     return count
// }
// func noRepeat(nums []int) int{
//     var flag = make([]bool, len(nums))
//     for i, v := range nums {
//         if flag[v-1] == true{
//             return i
//         }
//         flag[v-1] = true
//     }
//     return len(nums)
// }
// func addOne(nums []int, valid *[][]bool, i int) bool{
//     var carry = false
//     for j := nums[i]; true ;  {
//         if j >= len(nums) {
//             carry = true
//         }
//         j %= len(nums)
//         if (*valid)[i][j] {
//             nums[i] = j + 1
//             break
//         }
//         j ++
//     }
//     if carry && i>0{
//         return addOne(nums, valid, i-1)
//     }else{
//         return carry
//     }
// }
