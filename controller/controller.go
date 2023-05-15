package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"goLineBot/models"
	"goLineBot/service"

	"goLineBot/gpt"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sashabaranov/go-openai"
)

type LineBotController struct{}

type MessagesController struct{}

func (lbc *LineBotController) EventsHandler(bot *linebot.Client) gin.HandlerFunc {
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
		for _, event := range events {

			switch event.Type {
			case linebot.EventTypeFollow:
				FollowHandler(event, bot)
			case linebot.EventTypeUnfollow:
				UnFollowHandler(event, bot)
			case linebot.EventTypeMessage:
				MessageHandler(event, bot)
			}

		}
	}

}

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
				service.SaveMessage(&newMessage, bot)
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
			"success": "The message has been sent successfully",
			"message": sendingMessage,
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

func FollowHandler(event *linebot.Event, bot *linebot.Client) {

}

func UnFollowHandler(event *linebot.Event, bot *linebot.Client) {

}

func MessageHandler(event *linebot.Event, bot *linebot.Client) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:

		userId := event.Source.UserID
		messageId := message.ID
		newMessage := models.Message{
			ID:        message.ID,
			UserID:    userId,
			Type:      string(message.Type()),
			Text:      message.Text,
			CreatedAt: event.Timestamp,
		}
		log.Println(message)

		newSystemMessage := models.SystemMessageLog{
			ReplyID:     messageId,
			ReplyUserID: userId,
			CreatedAt:   event.Timestamp,
		}

		user, getUserErr := service.FindUserById(userId)
		if getUserErr != nil {
			log.Print(getUserErr)
			canMessage, err := service.FindCanMessagesById("0")
			if err != nil {
				log.Print(err)
			}
			_, replyErr := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(canMessage.Content)).Do()
			if replyErr != nil {
				log.Print(replyErr)
			}

			newSystemMessage.Text = canMessage.Content
			newSystemMessage.Type = "can"
			service.SaveSystemMessage(newSystemMessage)
		}

		service.SaveMessage(&newMessage, bot)

		if strings.HasPrefix(message.Text, "!") {
			command := strings.TrimPrefix(message.Text, "!")
			commandId := CommandHandler(command, userId)
			if commandId != " " {
				canMessage, err := service.FindCanMessagesById(commandId)
				if err != nil {
					log.Print(err)
				}
				_, replyErr := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(canMessage.Content)).Do()
				if replyErr != nil {
					log.Print(replyErr)
				}
				newSystemMessage.Text = canMessage.Content
				newSystemMessage.Type = "can"
				service.SaveSystemMessage(newSystemMessage)
			}
			return
		}

		if user.ChatGptSwitch {
			gptReplyMessage := gptMessageHandler(message.Text)

			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(gptReplyMessage)).Do()

			if err != nil {
				log.Print(err)
			}
			newSystemMessage.Text = gptReplyMessage
			newSystemMessage.Type = "gpt3.5"
			service.SaveSystemMessage(newSystemMessage)
			return
		}

		_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
		if err != nil {
			log.Print(err)
		}
		newSystemMessage.Text = message.Text
		newSystemMessage.Type = "repeat"
		service.SaveSystemMessage(newSystemMessage)
	}

}

func CommandHandler(command, userId string) string {
	switch command {
	case "gpt":
		service.ChangeGptSwitch(userId)
		return " "
	case "help":
		return "1"
	case "command":
		return "2"
	case "status":
		user, err := service.FindUserById(userId)

		if err != nil {
			log.Print(err)
		}
		if user.ChatGptSwitch {
			return "100"
		} else {
			return "101"
		}

	}
	return " "
}

func gptMessageHandler(sendingMessage string) string {

	resp, err := gpt.GptClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleUser, Content: sendingMessage},
			},
		},
	)

	if err != nil {
		log.Print(err)
		return ""
	}

	return resp.Choices[0].Message.Content

}
