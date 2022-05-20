package app

import (
	"context"
	"encoding/xml"
	pb "github.com/eluceon/quizwithmebot/pkg/api/quizpb"
	"io/ioutil"
	"net/http"
	"strings"
)

type Quiz struct {
	Question string `xml:"question>Question""`
	Answer   string `xml:"question>Answer"`
}

func (s *server) GetQuiz(ctx context.Context, req *pb.QuizRequest) (*pb.QuizResponse, error) {
	resp, err := http.Get("https://db.chgk.info/xml/random/answers/" + req.Complexity + "/limit1")
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quiz Quiz
	xml.Unmarshal(bytes, &quiz)
	answer := strings.ToLower(quiz.Answer)
	if strings.HasSuffix(answer, ".") {
		answer = answer[:len(answer)-1]
	}
	return &pb.QuizResponse{
		Question: quiz.Question,
		Answer:   answer,
	}, nil
}
