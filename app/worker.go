package app

import (
	"context"
	"encoding/json"
	"fmt"
	"google-oauth/helper"
	"google-oauth/service"
	"google-oauth/web"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func ConsumeAnswers(ctx context.Context, service *service.UserAnswerService) {
	client := helper.Client

	for {
		messages, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "group-1",
			Consumer: "consumer-1",
			Streams:  []string{"quiz_answer_stream", ">"},
			Count:    5,
			Block:    5 * time.Second,
		}).Result()

		if err != nil && err != redis.Nil {
			log.Println("Stream read error:", err)
			continue
		}

		for _, message := range messages {
			for _, m := range message.Messages {
				payload := m.Values["payload"].(string)
				var data web.SubmitQuizRequest
				err := json.Unmarshal([]byte(payload), &data)
				if err != nil {
					log.Println("Unmarshal error:", err)
					continue
				}	
				
				response, err := service.SaveAllAnswers(ctx, data, data.QuizId)
				if err != nil {
					log.Println("Save error:", err)
					continue
				}

				key := fmt.Sprintf("score:%d", response.Quiz.ID)
				client.ZAdd(ctx, key, redis.Z{
					Member: response.User.Name,
					Score:  response.Score,
				})
		
				client.XAck(ctx, "quiz_answer_stream", "group-1", m.ID)
			}
		}

	}

}
