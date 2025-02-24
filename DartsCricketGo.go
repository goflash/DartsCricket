package main

import (
	"fmt"
	"strconv"
)

func getValidThrow(sector int) int {
	for {
		var input string
		fmt.Print("Первый бросок: ")
		fmt.Scanln(&input)

		throwValue, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Пожалуйста, введите целое число.")
			continue
		}

		if throwValue == sector || throwValue == sector*2 || throwValue == sector*3 {
			return throwValue
		} else {
			fmt.Println("______________________")
			fmt.Println("Ошибка, вводите заново: ")
			fmt.Println("______________________")
		}
	}
}

func calculateScore(playerNumber int) int {
	totalScore := 0
	for sector := 1; sector <= 3; sector++ {
		fmt.Printf("Сектор %d для Игорька %d\n", sector, playerNumber)
		for i := 0; i < 3; i++ {
			totalScore += getValidThrow(sector)
		}
	}
	return totalScore
}

func main() {
	sum1 := calculateScore(1)
	sum2 := calculateScore(2)

	fmt.Println("Итоговый результат Игорька 1:", sum1)
	fmt.Println("Итоговый результат Игорька 2:", sum2)

	if sum1 > sum2 {
		fmt.Println("Выиграл Игорёк 1")
	} else if sum1 < sum2 {
		fmt.Println("Выиграл Игорёк 2")
	} else {
		fmt.Println("Ничья")
	}
}
