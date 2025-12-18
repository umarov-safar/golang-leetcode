package midlle

func SmallestDivisor(nums []int, threshold int) int {
    return findDivisor(nums, threshold)
}

func computeSum(nums []int, divisor int) int {
    sum := 0
    for _, num := range nums {
        sum += (num + divisor - 1) / divisor
    }

    return sum
}

func findDivisor(nums []int, threshold int) int {
    left, right := 1, 0

    for _, num := range nums {
        if num > right {
            right = num
        }
    }

    for left < right {
        mid := (left + right) / 2

        if computeSum(nums, mid) > threshold {
            left = mid + 1
        } else {
            right = mid
        }
    }    

    return left
}


