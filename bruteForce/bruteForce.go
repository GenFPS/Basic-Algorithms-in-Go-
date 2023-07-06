package main

import "fmt"

func BruteForceChange(amount int, coins []int) []int {
	numCoins := len(coins)
	minCoins := make([]int, amount+1)
	coinsUsed := make([]int, amount+1)

	for cents := 1; cents <= amount; cents++ {
		coinCount := cents
		newCoin := 1

		for j := 0; j < numCoins; j++ {
			if coins[j] <= cents {
				if minCoins[cents-coins[j]]+1 < coinCount {
					coinCount = minCoins[cents-coins[j]] + 1
					newCoin = coins[j]
				}
			}
		}

		minCoins[cents] = coinCount
		coinsUsed[cents] = newCoin
	}

	result := make([]int, 0)
	start := amount
	for start > 0 {
		coin := coinsUsed[start]
		result = append(result, coin)
		start -= coin
	}

	return result
}

func main() {
	amount := 40
	coins := []int{25, 20, 10, 5, 1}

	// Вызов функции BruteForceChange
	change := BruteForceChange(amount, coins)

	fmt.Println("Amount:", amount)
	fmt.Println("Coins:", coins)
	fmt.Println("Change:", change)
	fmt.Println("AmountCoins", len(change))
}
