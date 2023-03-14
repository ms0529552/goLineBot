package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const configPath = "."
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
	dbAddress := viper.GetString("mongo.address") + viper.GetString("mongo.port")

	//mongodb setup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbAddress))

	// Check if successfully connected
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	///Linebot sdk testing
	bot, err := linebot.New(lineConfig.secret, lineConfig.accessToken)
	if err != nil {
		panic("linebot connect error " + err.Error())
	}

	app.POST("/callback", func(c *gin.Context) {
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
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}

	})
	app.Run(":8080")

}
