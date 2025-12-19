package easy

func CheckValid(matrix [][]int) bool {
    for i, row := range matrix {
        r, c := make(map[int]bool), make(map[int]bool)
        
        for j, v := range row {
            if _, ok := r[v]; ok {
                return false
            } else {
                r[v] = true
            }

            if _, ok := c[matrix[j][i]]; ok {
                return false
            } else {
                c[matrix[j][i]] = true
            }
        }
    }

    return true
}