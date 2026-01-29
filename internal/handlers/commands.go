package handlers

import (
	kb "github.com/ilya-shymko/Sparkate/internal/keybords"
	ph "github.com/ilya-shymko/Sparkate/internal/photos"
	tele "gopkg.in/telebot.v4"
)

func ShowMainMenu(c tele.Context) error {
	// Создаем фото с текущей подписью
	photo := &tele.Photo{
		File:    ph.MainPhoto.File,
		Caption: "Main menu: ",
	}

	// Параметры отправки
	options := &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: kb.MainMenu,
	}

	// Отправляем
	return c.Send(photo, options)
}

func mainHandler(c tele.Context, caption string, menu *tele.ReplyMarkup) error {
	c.Respond()
	// Создаем фото с текущей подписью
	photo := &tele.Photo{
		File:    ph.MainPhoto.File,
		Caption: caption,
	}

	// Параметры отправки
	options := &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: menu,
	}

	// Отправляем
	return c.Edit(photo, options)
}
