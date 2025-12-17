package restorefinishinhorder

func RecoverOrder(order []int, friends []int) []int {
    fm := make(map[int]struct{})
    
    for _, v := range friends {
        fm[v] = struct{}{}
    }

    res := make([]int, 0)

    for _, v := range order {
        if _, ok := fm[v]; ok {
            res = append(res, v)
        }
    }
    
    return res
}