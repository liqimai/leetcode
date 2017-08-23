/** https://leetcode.com/problems/find-bottom-left-tree-value/description/
 * 513. Find Bottom Left Tree Value
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "github.com/liqimai/leetcode/utils"
import "fmt"

type TreeNode utils.TreeNode

func findBottomLeftValue(root *TreeNode) int {
    v, _ := _findBottomLeftValue((*TreeNode)(root))
    return v
}

func _findBottomLeftValue(root *TreeNode) (v, d int){
    if root == nil{
        return 0, -1
    }
    lv, ld := _findBottomLeftValue((*TreeNode)(root.Left))
    rv, rd := _findBottomLeftValue((*TreeNode)(root.Right))
    switch{
        case ld < 0 && rd < 0:
            return root.Val, 0
        case ld < rd:
            return rv, rd+1
        default:
            return lv, ld+1
    }
}

func main(){
    // fmt.Printf("%v\n",utils.Slice2Tree([]int{0, 1, 2, 3}))
    fmt.Println(_findBottomLeftValue((*TreeNode)(utils.Slice2Tree([]interface{}{0, 1, 2, nil, nil, 5}))))
    fmt.Println(utils.Slice2Tree([]interface{}{0, 1, 2, nil, nil, 5}).Tree2Slice())
}
