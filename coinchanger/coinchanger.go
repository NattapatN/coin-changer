package coinchanger

import (
	"sort"
)

var coin = []int{1, 5, 7}

func coinchanger(money int) []int {
	//reverse coin
	reveseCoin := sortArrayLargeToSmall(coin)

	//implement logic
	result := recursiveSummaryValue(money, 0, reveseCoin, []int{})

	return result
}

func sortArrayLargeToSmall(input []int) []int {
	output := input
	sort.SliceStable(output, func(i, j int) bool {
		return output[i] > output[j]
	})
	return output
}

func recursiveSummaryValue(target int, sum int, restArray []int, resultArray []int) []int {
	tempResult := []int{}
	if len(restArray) == 0 {
		return []int{}
	} else if sum+restArray[0] == target {
		return append(resultArray, restArray[0])
	} else if sum+restArray[0] < target {
		tempSum := sum + restArray[0]
		tempResult = recursiveSummaryValue(target, tempSum, restArray, append(resultArray, restArray[0]))
	}
	result := recursiveSummaryValue(target, sum, restArray[1:], resultArray)
	switch {
	case len(result) == 0:
		return tempResult
	case len(tempResult) == 0:
		return result
	case len(result) < len(tempResult):
		return result
	default:
		return tempResult
	}
}
