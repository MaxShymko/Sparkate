package blindgame

import (
	"errors"

	err "github.com/ilya-shymko/Sparkate/internal/errors"
)

func CalculateResults(player1, player2 int8) (ResultOfBGRound, error) {
	result := ResultOfBGRound{}
	if player1 > 5 || player1 < 1 || player2 > 5 || player2 < 1 {
		return result, err.ErrInvalidMove
	}

	// Если числа равны - ничья
	if player1 == player2 {
		result.IsDraw = true
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

	return ResultOfBGRound{}, errors.New("неожиданная ошибка в логике игры")
}
