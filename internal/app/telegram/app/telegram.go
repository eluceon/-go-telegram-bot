package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/eluceon/quizwithmebot/pkg/api/quizpb"
	"github.com/eluceon/quizwithmebot/pkg/api/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	botUrl     string
	offset     int
	parserConn quizpb.QuizServiceClient
	serverConn userpb.UserServiceClient
}

func New(apiKey string) *Client {
	return &Client{
		botUrl:     "https://api.telegram.org/bot" + apiKey,
		parserConn: quizpb.NewQuizServiceClient(clientConnection("parser:50051")),
		serverConn: userpb.NewUserServiceClient(clientConnection("sever:8080")),
	}
}

func clientConnection(endpoint string) *grpc.ClientConn {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[ERR] telegram bot could not connect gRPC: %v", err)
	}
	return conn
}

func (c Client) Start() error {
	for {
		updates, err := c.Updates()
		if err != nil {
			log.Println("[ERR] telegram getUpdates: ", err.Error())
			continue
		}

		for _, update := range updates {
			user, err := c.getUser(update.Message.From.ID, update.Message.From.Username)
			if err != nil {
				log.Println("[ERR] telegram getUser: ", err.Error())
				continue
			}
			if user.IsPassing {
				continue
			}
			user.IsPassing = true
			c.UpdateUser(user)
			c.offset = update.ID + 1
			go c.handleCmd(update.Message.Text, update.Message.Chat.ID, user)
		}
	}
}

func (c *Client) Updates() ([]Update, error) {
	resp, err := http.Get(c.botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(c.offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var updatesResponse UpdatesResponse
	err = json.Unmarshal(body, &updatesResponse)
	if err != nil {
		return nil, err
	}
	return updatesResponse.Result, nil
}

func (c *Client) SendMessage(chatID int, message string) {
	payload := fmt.Sprintf(`
		{
			"chat_id": %d,
			"text": %q
		}
	`, chatID, message)

	_, err := http.Post(c.botUrl+"/sendMessage", "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Println("[ERR] telegram http.Post error: ", err.Error())
	}
}

func (c *Client) getUser(userID int, username string) (*userpb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err := c.serverConn.ReadUser(ctx, &userpb.ID{ID: int64(userID)})
	if status.Code(err) == codes.NotFound {
		_, err = c.serverConn.CreateUser(ctx, &userpb.User{ID: int64(userID), Username: username})
		if err == nil {
			user, err = c.serverConn.ReadUser(ctx, &userpb.ID{ID: int64(userID)})
		}
	}
	return user, err
}

func (c *Client) UpdateUser(user *userpb.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := c.serverConn.UpdateUser(ctx, user)
	if err != nil {
		log.Fatalf("error while calling UpdateUser RPC: %v", err)
	}
}
