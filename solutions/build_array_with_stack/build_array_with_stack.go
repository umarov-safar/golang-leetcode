package buildarraywithstackgo

func buildArray(target []int, n int) []string {
    res := []string{}
    j := 0

    for i := 1; i <= n && j < len(target); i++ {
        res = append(res, "Push")

        if i == target[j] {
            j++
        } else {
            res = append(res, "Pop")
        }
    }

    return res
}