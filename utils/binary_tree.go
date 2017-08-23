package utils
import "fmt"
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func Slice2Tree(s []interface{}) *TreeNode{
    nodes := make([]*TreeNode, len(s))
    for i,v := range s{
        if v != nil{
            nodes[i] = new(TreeNode)
            nodes[i].Val = v.(int)
        }
    }
    // fmt.Println(nodes)
    for i, node := range nodes{
        if node != nil{
            if 2*i + 1 < len(nodes){
                node.Left = nodes[2*i+1]
            }
            if 2*i + 2 < len(nodes){
                node.Right = nodes[2*i+2]
            }
        }
    }
    return nodes[0]
}

func (root *TreeNode)Tree2Slice() (s []interface{}){
    if root == nil{
        s = append(s, nil)
        return
    }
    var queue_of_node []*TreeNode
    var queue_of_index []int
    queue_of_node = append(queue_of_node, root)
    queue_of_index = append(queue_of_index, 0)

    for len(queue_of_index) > 0{
        node := queue_of_node[0]
        i := queue_of_index[0]
        queue_of_index = queue_of_index[1:]
        queue_of_node = queue_of_node[1:]
        for len(s) <= i{
            s = append(s, nil)
        }
        s[i] = node.Val
        if node.Left != nil{
            queue_of_node = append(queue_of_node, node.Left)
            queue_of_index = append(queue_of_index, 2*i+1)
        }
        if node.Right != nil{
            queue_of_node = append(queue_of_node, node.Right)
            queue_of_index = append(queue_of_index, 2*i+2)
        }
    }
    return
}

func (t *TreeNode) String() string{
    return fmt.Sprintf("{%v %v %v}", t.Val, t.Left, t.Right)
}
