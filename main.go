package main

import (
	db "goLineBot/mongo"

	"goLineBot/controller"

	"goLineBot/migration"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"

	//openai "github.com/sashabaranov/go-openai"
	"goLineBot/gpt"
)

const configPath = "./configs"
const configType = "yaml"
const configName = "config"

func main() {

	app := gin.Default()

	//config
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetDefault("application.port", 8080)
	errConfig := viper.ReadInConfig()
	if errConfig != nil {
		panic("Reading configuration error because " + errConfig.Error())
	}

	//Set port address for database
	dbAdress := viper.GetString("mongo.address") + viper.GetString("mongo.port")

	//Connect to db
	db.ConnetDB(dbAdress)

	//
	migration.SaveAllCanMessages()

	//Geting linebot client through the channel secret and access token
	bot, err := linebot.New(viper.GetString("line.channel.secret"), viper.GetString("line.channel.access_token"))
	if err != nil {
		panic("linebot connect error " + err.Error())
	}

	//Connect to chatGpt
	gpt.Connect(viper.GetString("openApi.go_line_bot.token"))

	lineBotController := &controller.LineBotController{}
	messagesController := &controller.MessagesController{}
	//Api that receives message from line webhook and save the user info and message in MongoDB, then repeat the message again to user.
	app.POST("/repeat", lineBotController.RepeatHandler(bot))

	app.POST("/webhook", lineBotController.EventsHandler(bot))

	//Api that send message back to line, and then make linebot send the message to all users who has sent message to the linebot.
	app.POST("/send", lineBotController.SendHandler(bot))

	//APi that can get message list of the user by userId, and if the userId is null, it will respond with all messages in DB.
	app.GET("/messages", messagesController.MessagesByUserHandler())

	app.Run(":8080")

}
