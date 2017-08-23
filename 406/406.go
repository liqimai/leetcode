package main
import "fmt"

func reconstructQueue(people [][]int) [][]int {
    rval := make([][]int, len(people))

    for j := 0; j < len(people); j++ {
        var min_v int = 0x7fffffff
        //find first element
        for _, v := range people {
            if v[1] == 0{
                if v[0] < min_v{
                    min_v = v[0]
                }
            }
        }

        //reduce count
        for _, v := range people {
            if v[0] <= min_v{
                v[1]--
            }
        }

        //add to rval
        rval[j] = make([]int, 2)
        rval[j][0] = min_v
    }

    //reconstruct the count
    for i := 0; i < len(people); i++ {
        for j := i + 1; j < len(people); j++ {
            if rval[j][0] <= rval[i][0] {
                rval[j][1]++
            }
        }
    }
    return rval
}

func main(){
    fmt.Println(reconstructQueue([][]int{
        {7,0},
        {4,4},
        {7,1},
        {5,0},
        {6,1},
        {5,2}}))
}
