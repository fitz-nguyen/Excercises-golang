package main
import "fmt"

func twoSum(nums []int, target int) []int {
    sumMap := make(map[int]int)
    for index, key := range nums{
        _, ok := sumMap[target - key]
        if ok{
            return []int{sumMap[target - key], index}
        } else {
            sumMap[key] = index
        }
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println("Result = ", twoSum(nums, target))
}