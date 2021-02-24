package main

import (
	"fmt"
	"os"
	"time"
	"encoding"

	"github.com/felipeagger/go-redis-streams/packages/utils"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack/v4"
)

var (
	streamName string = os.Getenv("STREAM")
	client     *redis.Client
)

////////////////////////////////////////////////////
type UpdateEvent struct {
	*Base
	UserID  uint64
	ConnectorId string
	ConnectorStatus string
}

func (o *UpdateEvent) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

func (o *UpdateEvent) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}

type Type string

const (
	UpdateType Type = "Update"
)

type Base struct {
	ID       string
	ConnectorId string
	ConnectorStatus string
	Type     Type
	DateTime time.Time
	Retry    bool
}

// Event ...
type Event interface {
	GetID() string
	GetConnectorId() string
	GetConnectorStatus() string
	GetType() Type
	GetDateTime() time.Time
	SetID(id string)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

func New(t Type) (Event, error) {
	b := &Base{
		Type: t,
	}

	switch t {

	case UpdateType:
		return &UpdateEvent{
			Base: b,
		}, nil

	}

	return nil, fmt.Errorf("type %v not supported", t)
}

func (o *Base) GetID() string {
	return o.ID
}

func (o *Base) SetID(id string) {
	o.ID = id
}

func (o *Base) GetConnectorId() string {
	return o.ConnectorId
}

func (o *Base) GetConnectorStatus() string {
	return o.ConnectorStatus
}

func (o *Base) GetType() Type {
	return o.Type
}

func (o *Base) GetDateTime() time.Time {
	return o.DateTime
}

func (o *Base) String() string {

	return fmt.Sprintf("id:%s type:%s", o.ID, o.Type)
}
////////////////////////////////////////////////////

func init() {
	var err error
	client, err = utils.NewRedisClient()
	if err != nil {
		panic(err)
	}
}

func main() {
	for {
		generateEvent()
		time.Sleep(10 * time.Second)
	}
}

func generateEvent() {
	var userID uint64 = 0
	for i := 0; i < 10; i++ {

		userID++

		newID, err := produceMsg(map[string]interface{}{
			"type": string(UpdateType),
			"data": &UpdateEvent{
				Base: &Base{
					Type:     UpdateType,
					DateTime: time.Now(),
				},
				UserID: userID,
			},
		})

		checkError(err, newID, string(UpdateType), userID)

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
