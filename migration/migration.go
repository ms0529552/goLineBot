package migration

import (
	"context"
	"goLineBot/models"
	db "goLineBot/mongo"
	"goLineBot/service"
)

func SaveAllCanMessages() {

	_, err := service.FindCanMessagesById("0")

	if err != nil {

		canMessages := []models.CanMessage{
			{ID: "0", Description: "Follow event", Content: `您好，歡迎您使用goLineBot服務，此服務是使用Line以及chatgpt開發完成，目前僅有預設與gpt3.5對話的功能，若要開啟或是關閉只要在本聊天室輸入"!gpt"便可以啟用，若不使用的時候盡量幫我保持在關閉的狀態，非常感謝您的配合其他功能仍在持續更新完善，如果有任何問題歡迎輸入"!help"取得幫助`},
			{ID: "1", Description: "help", Content: `您好，歡迎您使用goLineBot服務，此服務是使用Line以及chatgpt開發完成，目前僅有預設與gpt3.5對話的功能，若要開啟或是關閉只要在本聊天室輸入"!gpt"便可以啟用，若不使用的時候盡量幫我保持在關閉的狀態，若需要指令集可以輸入"!command"`},
			{ID: "2", Description: "command set", Content: "目前有開放的指令及如下:\n!help 可以取得說明以及幫助\n!command 可以檢視目前所有的指令集\n !gpt 可以開啟或關閉chatgpt的功能，預設為關閉\n !status 可以檢視目前chatgpt功能的開啟關閉狀態"},
			{ID: "100", Description: "command set", Content: `目前您的chatgpt功能為開啟狀態，若要關閉請輸入"!gpt"`},
			{ID: "101", Description: "command set", Content: `目前您的chatgpt功能為關閉狀態，若要開啟請輸入"!gpt"`},
			{ID: "102", Description: "command set", Content: `您已成功開啟gpt功能`},
			{ID: "103", Description: "command set", Content: `您已成功關閉gpt功能`},
		}

		collection := db.DBclient.Database("goLineBot").Collection("can_messages")
		for _, canMessage := range canMessages {
			collection.InsertOne(context.Background(), canMessage)
		}
	}
}
