// 668. Kth largest Number in Multiplication Table
// https://leetcode.com/contest/leetcode-weekly-contest-47/problems/kth-largest-number-in-multiplication-table/
// The area of range 0<x<m, 0<y<n, xy < v is noted as area(v), it can be calculated by integration.
// Out goal is to find a v satisfy area(v-1) < k <= area(v) by binary search
func min(a int, b int) int{
    if a<b{
        return a
    }else{
        return b
    }
}
func numLessThan(v int, m int, n int) int{
    area := 0;
    for i := 1; i <= min(n, v); i++{
        area += min(m, v/i);
    }
    return area;
}

func findKthNumber(m int, n int, k int) int {
    l,r := 1, m*n;
    for{
        mid := (l+r)/2;
        mid_area := numLessThan(mid, m, n);
        switch{
            case k <= mid_area && k > numLessThan(mid - 1, m, n):
                return mid;
            case k > mid_area:
                l = mid + 1;
            case k <= numLessThan(mid - 1, m, n):
                r = mid;
        }
    }
}
