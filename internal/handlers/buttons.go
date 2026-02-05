package handlers

import (
	"fmt"
	"log"

	kb "github.com/ilya-shymko/Sparkate/internal/keybords"
	sys "github.com/ilya-shymko/Sparkate/internal/system"
	txt "github.com/ilya-shymko/Sparkate/internal/texts"
	tele "gopkg.in/telebot.v4"
)

// Main Menu

func ShowRulesOfGame(c tele.Context) error {
	defer c.Respond()
	id := c.Sender().ID
	state, err := sys.UserMap.GetUserState(id)
	if err != nil {
		log.Printf("Ошибка при выполнении функции(ShowRulesOfGame): %v", err)
		return err
	}

	if state == sys.StateMainMenu {
		sys.UserMap.UpdateUserState(id, sys.StateRulesMenu)

		return mainHandler(c, txt.RulesOfBG, kb.BackToMainMenuMenu)
	} else {
		err = fmt.Errorf("Пользователь находясь не в MainMenu хотел посмотреть rules")
		log.Printf("Ошибка при выполнении функции(ShowRulesOfGame): %v", err)
		return err
	}
}

func BackToMainMenu(c tele.Context) error {
	defer c.Respond()
	id := c.Sender().ID
	state, err := sys.UserMap.GetUserState(id)
	if err != nil {
		log.Printf("Ошибка при выполнении функции(BackToMainMenu): %v", err)
		return err
	}

	if state == sys.StateRulesMenu || state == sys.StateQueue || state == sys.StateEndOfGame {
		sys.UserMap.UpdateUserState(id, sys.StateMainMenu)
		return mainHandler(c, "Main menu: ", kb.MainMenu)
	} else {
		err = fmt.Errorf("Пользователь находясь не в RulesMenu хотел вернуться в MainMenu")
		log.Printf("Ошибка при выполнении функции(BackToMainMenu): %v", err)
		return err
	}
}

// Queue
func SelectionOfRivals(c tele.Context) error {
	defer c.Respond()
	id := c.Sender().ID
	state, err := sys.UserMap.GetUserState(id)
	if err != nil {
		log.Printf("Ошибка при выполнении функции(SelectionOfRivals): %v", err)
		return err
	}

	if state == sys.StateMainMenu {
		exist, idOfUserInQueue := sys.IsInQueueSmb()
		if exist {
			return StartGame(c, idOfUserInQueue, id)

		} else {
			sys.UserMap.UpdateUserState(id, sys.StateQueue)

			sys.UserMap.UpdateUserLastCtx(id, c)
			return mainHandler(c, "Going selection of rivals: ", kb.BackToMainMenuMenu)
		}
	} else {
		err = fmt.Errorf("Пользователь находясь не в MainMenu хотел зайти в очередь")
		log.Printf("Ошибка при выполнении функции(SelectionOfRivals): %v", err)
		return err
	}
}

func Helper1(c tele.Context) error {
	return CalculateResultsHandler(c, "GameBtn10")
}
func Helper2(c tele.Context) error {
	return CalculateResultsHandler(c, "GameBtn20")
}
func Helper3(c tele.Context) error {
	return CalculateResultsHandler(c, "GameBtn30")
}
func Helper4(c tele.Context) error {
	return CalculateResultsHandler(c, "GameBtn40")
}
func Helper5(c tele.Context) error {
	return CalculateResultsHandler(c, "GameBtn50")
}

// func CalculateResultsHandler(c tele.Context, data string) error {
// 	defer c.Respond()
// 	id := c.Sender().ID
// 	state, err := sys.UserMap.GetUserState(id)
// 	if err != nil {
// 		log.Printf("Ошибка при выполнении функции(BackToMainMenu): %v", err)
// 		return err
// 	}

// 	if state == sys.StateChoiceDigit {

// 		var choice int8
// 		switch data {
// 		case "GameBtn10":
// 			choice = 1
// 		case "GameBtn20":
// 			choice = 2
// 		case "GameBtn30":
// 			choice = 3
// 		case "GameBtn40":
// 			choice = 4
// 		case "GameBtn50":
// 			choice = 5
// 		default:
// 			choice = 0
// 			return fmt.Errorf("пользователь %d отправил фигню а нe цифру(1 2 3 4 5)", id)
// 		}
// 		log.Printf("choice равно - %d", choice)

// 		gameId, _ := sys.UserMap.GetUserGameID(id)
// 		is, idOfPlayerWaiting := sys.GameSessionMap.IsWaitingSmb(gameId)
// 		if is {
// 			ctx2, _ := sys.UserMap.GetUserLastCtx(idOfPlayerWaiting)
// 			player2choice, _ := sys.GameSessionMap.GetPlayerLastChoice(gameId, idOfPlayerWaiting)

// 			if sys.GameSessionMap.IsFirstPlayer(gameId, id) {
// 				StartRound(c, ctx2, id, idOfPlayerWaiting, gameId, choice, player2choice)
// 			} else {
// 				StartRound(ctx2, c, idOfPlayerWaiting, id, gameId, player2choice, choice)
// 			}

// 		} else {
// 			sys.GameSessionMap.UdatePlayerLastChoice(gameId, id, choice)
// 			sys.UserMap.UpdateUserLastCtx(id, c)
// 			sys.UserMap.UpdateUserState(id, sys.StateWaitingRival)
// 			return mainHandler(c, "Waitng rival's choice...", kb.BackToMainMenuMenu)
// 		}
// 	} else {
// 		err = fmt.Errorf("Пользователь находясь не в StateChoiceDigit хотел нажать gamebtn")
// 		log.Printf("Ошибка при выполнении функции (CalculateResultsHandler): %v", err)
// 		return err
// 	}

// 	return nil
// }

func CalculateResultsHandler(c tele.Context, data string) error {
	log.Printf("[INFO] CalculateResultsHandler запущен для пользователя %d, data: %s", c.Sender().ID, data)
	defer log.Println("[INFO] CalculateResultsHandler завершен")

	defer c.Respond()
	id := c.Sender().ID

	// Получение состояния пользователя
	state, err := sys.UserMap.GetUserState(id)
	if err != nil {
		log.Printf("[ERROR] Ошибка получения состояния пользователя %d: %v", id, err)
		return err
	}
	log.Printf("[DEBUG] Пользователь %d, состояние: %s", id, state)

	if state == sys.StateChoiceDigit {
		log.Printf("[DEBUG] Обработка выбора цифры для пользователя %d", id)

		var choice int8
		switch data {
		case "GameBtn10":
			choice = 1
		case "GameBtn20":
			choice = 2
		case "GameBtn30":
			choice = 3
		case "GameBtn40":
			choice = 4
		case "GameBtn50":
			choice = 5
		default:
			log.Printf("[WARN] Пользователь %d отправил невалидные данные: %s", id, data)
			choice = 0
			return fmt.Errorf("пользователь %d отправил фигню а не цифру(1 2 3 4 5)", id)
		}
		log.Printf("[INFO] Пользователь %d выбрал: %d", id, choice)

		// Получение ID игры
		gameId, err := sys.UserMap.GetUserGameID(id)
		if err != nil {
			log.Printf("[ERROR] Ошибка получения gameId для пользователя %d: %v", id, err)
			return err
		}
		log.Printf("[DEBUG] Пользователь %d, gameId: %d", id, gameId)

		// Проверка ожидания соперника
		is, idOfPlayerWaiting := sys.GameSessionMap.IsWaitingSmb(gameId)
		log.Printf("[DEBUG] Ожидание соперника: %v, ID ожидающего: %d", is, idOfPlayerWaiting)

		if is {
			log.Printf("[INFO] Соперник найден, начинаем раунд")

			ctx2, err := sys.UserMap.GetUserLastCtx(idOfPlayerWaiting)
			if err != nil {
				log.Printf("[ERROR] Ошибка получения контекста для пользователя %d: %v", idOfPlayerWaiting, err)
				return err
			}

			player2choice, err := sys.GameSessionMap.GetPlayerLastChoice(gameId, idOfPlayerWaiting)
			if err != nil {
				log.Printf("[ERROR] Ошибка получения выбора соперника %d: %v", idOfPlayerWaiting, err)
				return err
			}
			log.Printf("[DEBUG] Выбор соперника %d: %d", idOfPlayerWaiting, player2choice)

			// Определение кто первый игрок
			isFirstPlayer := sys.GameSessionMap.IsFirstPlayer(gameId, id)
			log.Printf("[DEBUG] Пользователь %d первый игрок: %v", id, isFirstPlayer)

			if isFirstPlayer {
				log.Printf("[INFO] Запуск StartRound (первый игрок)")
				StartRound(c, ctx2, id, idOfPlayerWaiting, gameId, choice, player2choice)
			} else {
				log.Printf("[INFO] Запуск StartRound (второй игрок)")
				StartRound(ctx2, c, idOfPlayerWaiting, id, gameId, player2choice, choice)
			}

		} else {
			log.Printf("[INFO] Соперник не найден, обновляем состояние ожидания")

			// Сохраняем выбор пользователя
			err := sys.GameSessionMap.UdatePlayerLastChoice(gameId, id, choice)
			if err != nil {
				log.Printf("[ERROR] Ошибка обновления выбора игрока %d: %v", id, err)
				return err
			}

			err = sys.UserMap.UpdateUserLastCtx(id, c)
			if err != nil {
				log.Printf("[ERROR] Ошибка обновления контекста пользователя %d: %v", id, err)
				return err
			}

			err = sys.UserMap.UpdateUserState(id, sys.StateWaitingRival)
			if err != nil {
				log.Printf("[ERROR] Ошибка обновления состояния пользователя %d: %v", id, err)
				return err
			}

			log.Printf("[INFO] Пользователь %d переведен в состояние ожидания", id)
			return mainHandler(c, "Waitng rival's choice...", kb.BackToMainMenuMenu)
		}
	} else {
		log.Printf("[WARN] Неправильное состояние! Пользователь %d в состоянии %s, но пытается нажать gamebtn",
			id, state)
		err = fmt.Errorf("Пользователь находясь не в StateChoiceDigit хотел нажать gamebtn")
		return err
	}

	return nil
}
