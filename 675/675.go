package main
import "fmt"

import "sort"
import "container/heap"
type MyPoint struct {
    X,Y int
}

type Cell struct {
    MyPoint
    Height int
    dist int
}

type ByHeight []Cell
func (a ByHeight) Len() int           { return len(a) }
func (a ByHeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHeight) Less(i, j int) bool { return a[i].Height < a[j].Height }


// An IntHeap is a min-heap of ints.
func (h *ByHeight) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Cell))
}

func (h *ByHeight) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func cutOffTree(forest [][]int) int {
    trees := make([]Cell, 0, 50*50)
    for i, row := range forest {
        for j, height := range row {
            if height > 1 {
                trees = append(trees, Cell{ MyPoint{i,j}, height, 0})
            }
        }
    }
    sort.Sort(ByHeight(trees))
    start := Cell{ MyPoint{0,0}, 0, 0}
    // fmt.Println(trees)
    rval := 0
    for i := range trees {
        inc := walk(start.MyPoint, trees[i].MyPoint, forest)
        // fmt.Println(start.MyPoint, trees[i].MyPoint, inc, l1dis(start.MyPoint, trees[i].MyPoint))
        if inc < 0 {
            return -1
        }
        rval += inc
        start = trees[i]
    }
    return rval
}
func walk(start, end MyPoint, forest [][]int ) int {
    rows, cols := len(forest), len(forest[0])
    distance := make([][]int, rows)
    for i := range distance {
        distance[i] = make([]int, cols)
        for j := range distance[i] {
            distance[i][j] = 50*50 + 1000
        }
    }
    // _walk(start, end, forest, 0, distance)
    open := &ByHeight{Cell{start, 0+l1dis(MyPoint{0,0}, end), 0}}
    heap.Init(open)
    closed := make(map[MyPoint]bool)
    for {
        if len(*open) == 0 {
            return -1
        }
        now := heap.Pop(open).(Cell)
        closed[now.MyPoint] = true
        distance[now.X][now.Y] = now.dist
        if now.MyPoint == end {
            return now.dist
        }
        neighbors := []MyPoint{
            MyPoint{now.X+1, now.Y},
            MyPoint{now.X-1, now.Y},
            MyPoint{now.X, now.Y+1},
            MyPoint{now.X, now.Y-1},
        }
        for _, v := range(neighbors) {
            if v.X < rows && v.X >= 0 && v.Y < cols && v.Y >= 0 && forest[v.X][v.Y] != 0 && !closed[v]{
                heap.Push(open, Cell{ v, now.dist+1+l1dis(v, end), now.dist + 1 })
            }
        }
    }
}

func l1dis(a,b MyPoint) int {
    abs := func (a int) int {
        if a > 0 {
            return a
        } else{
            return -a
        }
    }
    return abs(a.X - b.X) + abs(a.Y - b.Y)
}

// func _walk(now, end MyPoint, forest [][]int, pathLength int, distance [][]int) {
//     rows, cols := len(forest), len(forest[0])
//     if now.X >= rows || now.X < 0 || now.Y >= cols || now.Y < 0 || pathLength >= distance[now.X][now.Y] || forest[now.X][now.Y] == 0 {
//         return
//     }
//     distance[now.X][now.Y] = pathLength
//     if now == end {
//         fmt.Println(distance)
//     }
//     g := pathLength+1
//     neighbors := []Cell {
//         Cell{ MyPoint{now.X+1, now.Y}, g+l1dis(MyPoint{now.X+1, now.Y}, end)},
//         Cell{ MyPoint{now.X-1, now.Y}, g+l1dis(MyPoint{now.X-1, now.Y}, end)},
//         Cell{ MyPoint{now.X, now.Y+1}, g+l1dis(MyPoint{now.X, now.Y+1}, end)},
//         Cell{ MyPoint{now.X, now.Y-1}, g+l1dis(MyPoint{now.X, now.Y-1}, end)},
//     }
//     sort.Sort(ByHeight(neighbors))
//     for i := range neighbors {
//         if distance[end.X][end.Y] > 50*50 {
//             _walk(neighbors[i].MyPoint, end, forest, g, distance)
//         }
//     }
//     return
// }

func main() {
    // forest := [][]int{
    //     []int {4557,6199,7461,2554,6132,7471,7103,4290},
    //     []int {2034,2301,6733,6040,2603,5697,9541,6678},
    //     []int {7308,7368,9618,4386,6944,3923,4754,4294},
    //     []int {0000,3016,7242,5284,6631,1897,1767,7603},
    //     []int {2228,0000,3625,7713,2956,3264,3371,6124},
    //     []int {9195,7804,2787,0000,4919,9304,5161,502},
    // }
    forest := [][]int {
        []int {1,2,3},
        []int {0,0,0},
        []int {7,6,5},
    }
    fmt.Println(cutOffTree(forest))
}
