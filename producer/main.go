package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	evt "github.com/felipeagger/go-redis-streams/packages/event"
	"github.com/felipeagger/go-redis-streams/packages/utils"
	"github.com/go-redis/redis/v7"
)

var (
	streamName string = os.Getenv("STREAM")
	client     *redis.Client
)

func init() {
	var err error
	client, err = utils.NewRedisClient()
	if err != nil {
		panic(err)
	}
}

func main() {
	for {
		status := []string{"Active", "Not Active"}[rand.Intn(2)]
		updateStatus(fmt.Sprintf("%d", rand.Intn(100)), status)
		time.Sleep( 10 * time.Second)
	}
}

func updateStatus(connectorID string, connectionsStatus string) {
	comment := fmt.Sprintf("ConnectorId=%s, ConnectorStatus=%s", connectorID, connectionsStatus)
	generateEvent(comment)
}

func generateEvent(comment string) {
	var userID uint64 = 0
	for i := 0; i < 10; i++ {

		eventType := evt.CommentType

		newID, err := produceMsg(map[string]interface{}{
			"type": string(eventType),
			"data": &evt.CommentEvent{
				Base: &evt.Base{
					Type:     eventType,
					DateTime: time.Now(),
				},
				UserID:  userID,
				Comment: comment,
			},
		})

		checkError(err, newID, string(eventType), userID, comment)

	}
}

func produceMsg(event map[string]interface{}) (string, error) {

	return client.XAdd(&redis.XAddArgs{
		Stream: streamName,
		Values: event,
	}).Result()
}

func checkError(err error, newID, eventType string, userID uint64, comment ...string) {
	if err != nil {
		fmt.Printf("produce event error:%v\n", err)
	} else {

		if len(comment) > 0 {
			fmt.Printf("produce event success Type:%v UserID:%v Comment:%v offset:%v\n",
				string(eventType), userID, comment, newID)
		} else {
			fmt.Printf("produce event success Type:%v UserID:%v offset:%v\n",
				string(eventType), userID, newID)
		}

	}
}