package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"goLineBot/models"
	"goLineBot/service"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineBotController struct{}

type MessagesController struct{}

func (lbc *LineBotController) RepeatHandler(bot *linebot.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.Status(http.StatusBadRequest)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			return
		}
		//For now we only handle the text message, other cases will be handled in the future
		for _, event := range events {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				newMessage := models.Message{
					ID:        message.ID,
					UserID:    event.Source.UserID,
					Type:      string(message.Type()),
					Text:      message.Text,
					CreatedAt: event.Timestamp,
				}
				log.Println(message)
				service.SaveMessage(&newMessage)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func (lbc *LineBotController) SendHandler(bot *linebot.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		message := c.PostForm("message")
		if message == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "The message in request body can't be empty",
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
			if _, err = bot.PushMessage(user.UserID, sendingMessage).Do(); err != nil {
				log.Print(err)
				return
			}
			time.Sleep(100 * time.Millisecond) // Prevent overspeeding
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "The message has been sent successfully",
		})
	}
}

func (mc *MessagesController) MessagesByUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Query("userId")
		if userId == "" {
			messagesList, err := service.GetAllMessages()
			if err != nil {
				log.Print(err)
			}
			c.JSON(http.StatusOK, gin.H{"messages list": messagesList})
			return
		}
		fmt.Println(userId)
		user, err := service.FindUserById(userId)
		if err != nil {
			log.Print(err)
			return
		}

		messagesList, err := service.GetMessagesByUser(user)
		c.JSON(http.StatusOK, gin.H{"messages list": messagesList})

	}
}