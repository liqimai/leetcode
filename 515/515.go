// 515. Find Largest Value in Each Tree Row
// https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"

func largestValues(root *TreeNode) []int {
    qInt := list.New()
    qNode := list.New()
    qInt.PushBack(0)
    qNode.PushBack(root)
    var rval []int
    for nodeE, depthE := qNode.Front(), qInt.Front(); nodeE != nil; nodeE, depthE = nodeE.Next(), depthE.Next() {
        node := nodeE.Value.(*TreeNode)
        depth := depthE.Value.(int)
        if node == nil{
            continue
        }
        if depth == len(rval){
            rval = append(rval, node.Val)
        }else if node.Val > rval[depth]{
            rval[depth] = node.Val
        }
        if node.Left != nil{
            qNode.PushBack(node.Left)
            qInt.PushBack(depth+1)
        }
        if node.Right != nil{
            qNode.PushBack(node.Right)
            qInt.PushBack(depth+1)
        }
    }
    return rval
}
