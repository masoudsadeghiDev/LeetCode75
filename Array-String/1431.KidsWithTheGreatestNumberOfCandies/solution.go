package main

import "fmt"

func main() {
	candies := []int{2, 8, 7}
	extraCandies := 1

	fmt.Println(kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := candies[0]
	result := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		if maxCandy < candies[i] {
			maxCandy = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		result[i] = candies[i]+extraCandies >= maxCandy
	}
	return result
}
