//leetcode 673 Number of Longest Increasing Subsequence
//Given an unsorted array of integers, find the number of longest increasing subsequence
// Example 1:
// Input: [1,3,5,4,7]
// Output: 2
// Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
package main
import "fmt"
import "runtime/pprof"
import "flag"
import "os"
import "log"
import "runtime"

func findNumberOfLIS(nums []int) int {
    length := make([]int, len(nums))
    count := make([]int, len(nums))
    max_length := 0
    cnt := 0
    for i := range nums {
        length[i] = 1
        count[i] = 1
        for j := 0; j < i; j++{
            if nums[i] > nums[j] {
                switch{
                case length[j] + 1 > length[i]:
                    length[i] = length[j] + 1
                    count[i] = count[j]
                case length[j] + 1 == length[i]:
                    length[i] = length[j] + 1
                    count[i] += count[j]
                }
            }
        }
        switch{
        case length[i] > max_length :
            max_length = length[i]
            cnt = count[i]
        case length[i] == max_length:
            cnt += count[i]
        }
    }
    return cnt
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }
    nums := []int{100,90,80,70,60,50,60,70,80,90,100}
    for i := 0; i < 1; i++{
        fmt.Println(findNumberOfLIS(nums))
    }

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
        f.Close()
    }
}
