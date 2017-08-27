// 609. Find Duplicate File in System
// https://leetcode.com/problems/find-duplicate-file-in-system/description/
import "path/filepath"
import "strings"

func findDuplicate(paths []string) [][]string {
    maps := make(map[string][]string)
    var rval [][]string
    for _, record := range paths{
        files := strings.Split(record," ")
        for _, file := range files[1:]{
            tmp := strings.Split(file[:len(file)-1], "(")
            name  := tmp[0]
            content := tmp[1]
            l, ok := maps[content]
            if ok {
                l = append(l, filepath.Join(files[0], name))
            }else{
                l = []string{filepath.Join(files[0], name)}
            }
            maps[content] = l
        }
    }
    for _, v := range maps{
        if len(v) >= 2{
            rval = append(rval, v)
        }
    }
    return rval
}
