package err

import "errors"

var (
	ErrInvalidMove = errors.New("Ход должен быть числом от 1 до 5")
)
