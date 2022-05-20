package app

import (
	"context"
	"fmt"
	"github.com/eluceon/quizwithmebot/pkg/api/quizpb"
	"github.com/eluceon/quizwithmebot/pkg/api/userpb"
	"log"
	"strings"
	"time"
)

const (
	StartCmd   = "/start"
	QuizCmd    = "/quiz"
	StatsCmd   = "/stats"
	TopCmd     = "/top"
	HelpCmd    = "/help"
	COMPLEXITY = "complexity1"
)

func (c *Client) handleCmd(text string, chatID int, user *userpb.User) {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, user.Username)

	switch text {
	case StartCmd:
		c.Statistics(chatID, user.ID)
	case QuizCmd:
		c.PlayQuiz(chatID, user)
	case StatsCmd:
		c.Statistics(chatID, user.ID)
	case TopCmd:
		c.TopUsers(chatID)
	case HelpCmd:
		c.SendMessage(chatID, msgHelp)
	default:
		c.SendMessage(chatID, msgUnknownCommand)
	}
	user.IsPassing = false
	c.UpdateUser(user)
}

func (c *Client) PlayQuiz(chatID int, user *userpb.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	resp, err := c.parserConn.GetQuiz(ctx, &quizpb.QuizRequest{Complexity: COMPLEXITY})
	if err != nil {
		log.Fatalf("error while calling GetQuiz RPC: %v", err)
	}
	c.SendMessage(chatID, resp.Question)
	for start := time.Now(); time.Since(start) < 60*time.Second; {
		updates, err := c.Updates()
		if err != nil {
			log.Println("[ERR] telegram getUpdates: ", err.Error())
			continue
		}

		if len(updates) > 0 {
			c.offset = updates[0].ID + 1
			if strings.ToLower(updates[0].Message.Text) == strings.ToLower(resp.Answer) {
				c.UpdateInfo(chatID, user, correctAnswer, 1)
			} else {
				c.UpdateInfo(chatID, user, incorrectAnswer, 0)
			}
			return
		}
	}
	c.UpdateInfo(chatID, user, timeOut, 0)
}

func (c *Client) Statistics(chatID int, userID int64) {
	ctx := context.Background()

	user, err := c.serverConn.ReadUser(ctx, &userpb.ID{ID: userID})
	if err != nil {
		log.Fatalf("[ERR] telegram Statistics ReadUser error: %v", err)
	}
	c.SendStats(chatID, user.Username, user.CorrectAnswers, user.TotalAnswers)
}

func (c *Client) SendStats(chatID int, userName string, correctAnswers int32, totalAnswers int32) {
	msg := fmt.Sprintf("Имя пользователя: %s\nВерно: %d\nВсего: %d\nCоотношение: %.2f",
		userName, correctAnswers, totalAnswers, 100*float32(correctAnswers)/float32(totalAnswers))
	c.SendMessage(chatID, msg)
}

func (c *Client) TopUsers(chatID int) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	users, err := c.serverConn.GetTopUsers(ctx, &userpb.Empty{})
	if err != nil {
		log.Fatalf("[ERR] error while calling GetTopUsers RPC: %v", err)
	}
	msg := ""
	for i, user := range users.Users {
		if len([]rune(user.Username)) > 3 {
			msg += fmt.Sprintf("%d. ***%s очков: %d\n", i+1, user.Username[3:], user.CorrectAnswers)
		} else {
			msg += fmt.Sprintf("%d. *** очков: %d\n", i+1, user.CorrectAnswers)
		}
	}
	c.SendMessage(chatID, msg)
}

func (c *Client) UpdateInfo(chatID int, user *userpb.User, message string, correct int) {
	user.CorrectAnswers += int32(correct)
	user.TotalAnswers++
	user.IsPassing = false
	c.SendMessage(chatID, message)
}
