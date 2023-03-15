package main

import (
	"fmt"
	"goLineBot/models"
	db "goLineBot/mongo"
	"goLineBot/service"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const configPath = "./configs"
const configType = "yaml"
const configName = "config"

type lineconfig struct {
	secret      string
	accessToken string
}

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

	lineConfig := lineconfig{viper.GetString("line.channel.secret"), viper.GetString("line.channel.access_token")}
	dbAdress := viper.GetString("mongo.address") + viper.GetString("mongo.port")
	db.ConnetDB(dbAdress)

	///Linebot sdk testing
	bot, err := linebot.New(lineConfig.secret, lineConfig.accessToken)
	if err != nil {
		panic("linebot connect error " + err.Error())
	}
	app.POST("/repeat", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.Status(http.StatusBadRequest)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			return
		}

		for _, event := range events {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				newMessage := models.Message{ID: message.ID, UserID: event.Source.UserID, Type: string(message.Type()), Text: message.Text, CreatedAt: event.Timestamp}
				fmt.Println(message)
				service.SaveMessage(&newMessage)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	})

	app.POST("/send", func(c *gin.Context) {
		message := c.PostForm("message")
		if message == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "請提供訊息內容",
			})
			return
		}

		sendingMessage := linebot.NewTextMessage(message)
		usersList, err := service.GetUsersList()
		if err != nil {
			log.Print(err)
			return
		}

		for _, user := range usersList {
			_, err = bot.PushMessage(user.UserID, sendingMessage).Do()
			if err != nil {
				log.Print(err)
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "The message has been sent successfully",
		})

	})

	app.Run(":8080")

}
