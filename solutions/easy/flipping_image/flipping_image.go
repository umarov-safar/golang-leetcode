package flippingimage

// Task: 832. Flipping an Image
// Link: https://leetcode.com/problems/flipping-an-image/

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image[0])
	for i := range image {
		for j := 0; j < ((n + 1) / 2); j++ {
			image[i][j], image[i][n-j-1] = image[i][n-j-1], image[i][j]

			if image[i][j] == 1 {
				image[i][j] = 0
			} else {
				image[i][j] = 1
			}

			if j != n-j-1 {
				if image[i][n-j-1] == 1 {
					image[i][n-j-1] = 0
				} else {
					image[i][n-j-1] = 1
				}
			}
		}
	}

	return image
}
