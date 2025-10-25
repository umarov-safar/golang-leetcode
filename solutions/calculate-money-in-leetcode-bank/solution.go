package calculatemoneyinleetcodebank

func totalMoney(n int) int {
    monday := 1
    ans := 0
    for n > 0 {
        for day := 0; day < min(n, 7); day++ {
            ans += monday + day
        }
        monday++
        n-=7
    }

    return ans
}
