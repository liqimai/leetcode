// https://leetcode.com/problems/palindromic-substrings/description/
package main

import "fmt"
import "os"
import "bufio"

func countSubstrings(s string) int {
    l := len(s)
    count := 0
    for n := 0; n <= 2*l-2 ; n++{
        var k int
        if n%2 == 1{
            k = 1
        }else{
            k = 0
        }
        for ; k <= 2*l-2 - n && k <= n; k += 2{
            j := (n+k)/2
            i := (n-k)/2
            if s[i] == s[j]{
				fmt.Println(s[i:j+1])
                count++
            }else{
                break
            }
        }
    }
    return count
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        s := scanner.Text()
        fmt.Println(countSubstrings(s))
    }
}
