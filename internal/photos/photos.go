package photos

import (
	tele "gopkg.in/telebot.v4"
)

var (
	MainPhoto *tele.Photo
)

func InitPhotos() {
	// photoId := os.Getenv("PHOTO_ID")

	// if photoId == "" {
	// 	log.Println("❌ PHOTO_ID не установлено. Создайте файл .env или установите переменную окружения")
	// }

	// MainPhoto = &tele.Photo{
	// 	File: tele.File{FileID: photoId},
	// }

	MainPhoto = &tele.Photo{
		File: tele.FromDisk("./internal/photos/Sparkatelogo.jpg"),
	}
}
