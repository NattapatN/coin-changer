package coinchanger

import "sort"

var coin = []int{1, 2, 5, 10}

func coinchanger(money int) []int {
	//implement logic
	result := []int{}
	//reverse coin
	reveseCoin := sortArrayLargeToSmall(coin)

	currentIndex := 0
	for money > 0 && currentIndex < len(reveseCoin) {
		if money >= reveseCoin[currentIndex] {
			result = append(result, reveseCoin[currentIndex])
			money = money - reveseCoin[currentIndex]
		} else {
			currentIndex++
		}
	}
	return result
}

func sortArrayLargeToSmall(input []int) []int {
	output := input
	sort.SliceStable(output, func(i, j int) bool {
		return output[i] > output[j]
	})
	return output
}
