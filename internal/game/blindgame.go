package game

import (
	"fmt"
)

type ResultOfGameRound struct {
	Player1Points int8
	Player2Points int8
}

func CalculateResults(player1, player2 int8) (ResultOfGameRound, error) {
	result := ResultOfGameRound{}
	if player1 > 5 || player1 < 1 || player2 > 5 || player2 < 1 {
		return result, fmt.Errorf("Ход должен быть числом от 1 до 5")
	}

	// Если числа равны - ничья
	if player1 == player2 {
		result.Player1Points = 0
		result.Player2Points = 0
		return result, nil
	}

	// Определяем большее и меньшее число
	var higher, lower int8
	var isFirstPlayerHigher bool

	if player1 > player2 {
		higher, lower = player1, player2
		isFirstPlayerHigher = true
	} else {
		higher, lower = player2, player1
		isFirstPlayerHigher = false
	}

	diff := higher - lower

	// Если разница равна 1
	if diff == 1 {
		points := player1 + player2 // сумма обоих чисел

		if isFirstPlayerHigher { // Если первый игрок показал большее число
			result.Player2Points = points // то второй игрок (с меньшим) получает очки
		} else {
			result.Player1Points = points // и наоборот
		}

		return result, nil
	}

	// Если разница больше 1
	if diff > 1 {

		if isFirstPlayerHigher {
			result.Player1Points = diff
		} else {
			result.Player2Points = diff
		}

		return result, nil
	}

	return ResultOfGameRound{}, fmt.Errorf("неожиданная ошибка в логике игры")
}
