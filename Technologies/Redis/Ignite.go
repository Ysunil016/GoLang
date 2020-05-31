package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

var wg sync.WaitGroup

func main() {
	redisInstance := RedInstance{connectRedis("127.0.0.1", 6379, "", 0)}
	Subscribe := []string{"X", "Y", "Z"}
	wg.Add(1)
	go redisInstance.subscribeToChannel(Subscribe...)

	// for i := 0; i < 10000; i++ {
	// 	redisInstance.publishOnChannel("X", strconv.Itoa(i))
	// }

	wg.Wait()
}

// RedInstance ...
type RedInstance struct {
	RedisInstance *redis.Client
}

func (r RedInstance) subscribeToChannel(channels ...string) {
	// Adding for Handling Listner
	CTX := context.Background()
	channelListener := r.RedisInstance.Subscribe(CTX, channels...).Channel()
	fmt.Println("Subscribing to ", channels)
	go handleReceivedMessage(channelListener)
}

func handleReceivedMessage(channelListener <-chan *redis.Message) {
	defer wg.Done()

	for {
		MsgReceived := <-channelListener
		switch MsgReceived.Channel {
		case "X":
			handleXMessage(MsgReceived.Payload)
			break
		case "Y":
			handleYMessage(MsgReceived.Payload)
			break
		case "Z":
			handleZMessage(MsgReceived.Payload)
			break
		}
	}
}

func handleXMessage(message string) {
	fmt.Println("X", message)
}
func handleYMessage(message string) {
	fmt.Println("Y", message)
}
func handleZMessage(message string) {
	fmt.Println("Z", message)
}

///////////////////////////////
func (r RedInstance) publishOnChannel(channel string, message string) {
	CTX := context.Background()
	_, err := r.RedisInstance.Publish(CTX, channel, message).Result()
	if err != nil {
		fmt.Println("Issue While Publishing Message")
	}
}

////// REDIS CONNEXTION //////

func connectRedis(Hostname string, Port int, Password string, DB int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Hostname + ":" + strconv.Itoa(Port),
		Password: Password,
		DB:       DB,
	})
}
