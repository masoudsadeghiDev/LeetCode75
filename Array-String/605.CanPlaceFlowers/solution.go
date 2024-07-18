package array_string

func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)

	if n == 0 {
		return true
	} else if n == size && size > 1 {
		return false
	} else if size == 1 && flowerbed[0] == 0 && n == 1 {
		return true
	} else if size == 1 && flowerbed[0] == 1 {
		return false
	}
	counter := n
	if size >= 2 {

		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			flowerbed[0] = 1
			counter--
		}

		if flowerbed[size-2] == 0 && flowerbed[size-1] == 0 {
			flowerbed[size-1] = 1
			counter--
		}
	}

	for i := 1; i < size-2; i++ {
		if counter == 0 {
			return true
		}
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			counter--
		}
	}

	
	return counter <= 0
}
