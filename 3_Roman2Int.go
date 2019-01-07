package main 
import (
    "fmt"
    // S "strings"
)



func romanToInt(s string) (sum int) {
    valueMap := map[byte]int{
        'I':             1,
        'V':             5,
        'X':             10,
        'L':             50,
        'C':             100,
        'D':             500,
        'M':             1000,
    }
    var base int
    for i :=len(s)-1; i>=0; i-- {
        val := valueMap[s[i]]
        if val >= base {
            base = val
            sum += val
            continue        
        }
        sum -= val
    }
    return sum
}

func main() {
    s := "LVIII"
    fmt.Println(romanToInt(s))

}
