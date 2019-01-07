func indexOf(element int, data []int) (int) {
    for k, v := range data {
        if element == v {
            return k
        }
    }
    return -1    //not found.
 }

func nextGreaterElement(findNums []int, num []int) []int {
    mang:= make ([]int, len(findNums))
    // var mang []int
    for i := range mang {
        mang[i] = -1
    }
    for k, v := range findNums {
        index := indexOf(v, num)
        if index == len(num)-1 {
            continue
        }
        for i:=index+1; i<=len(num)-1;i++ {

            if num[i] > v {
                mang[k] = num[i]
                break
                }
        }


        }
        return mang
    }