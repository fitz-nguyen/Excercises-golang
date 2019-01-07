func Remove(slice []int, index int) []int{
    slice = append(slice[:index], slice[index+1:]...)
    return slice
}
func removeElement(nums []int, val int) int {
    i := 0
    
    for i < len(nums)  {
        if nums[i]==val {
            nums = Remove(nums, i)
            i--
        } 
        i++
        
    }
    return len(nums)
}