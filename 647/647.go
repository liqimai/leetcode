// https://leetcode.com/problems/palindromic-substrings/description/
package main

import "fmt"
import "os"
import "bufio"

func countSubstrings(s string) int {
    count := 0
    for i:= 0 ; i< len(s) ; i++{
        l,r := i, i
        for l >= 0 && r < len(s){
            if s[l] == s[r]{
                count++
            }else{
                break
            }
            l--
            r++
        }
        l,r = i, i+1
        for l >= 0 && r < len(s){
            if s[l] == s[r]{
                count++
            }else{
                break
            }
            l--
            r++
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
