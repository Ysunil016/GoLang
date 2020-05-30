package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/go-redis/redis"
)

var wg sync.WaitGroup

// RedisParams ...
type RedisParams struct {
	SubHost    string
	PubHost    string
	SubPort    string
	PubPort    string
	SubChannel string
	PubChannel string
}

func main() {
	// argsWithProg := os.Args
	pArgs := os.Args[1:]

	// SubHost ... Hostname for Subscription.
	var SubHost string = "127.0.0.1"
	// PubHost ... Hostname for Publishing.
	var PubHost string = "127.0.0.1"
	// SubPort ... PORT for Subscription.
	var SubPort string = "6379"
	// PubPort ... PORT for Publishing.
	var PubPort string = "6379"

	// ChannelTo ... Channel for Subscription.
	var fromChannel string
	// ChannelTo ... Channel for Publishing.
	var toChannel string

	for index, V := range pArgs {
		var StringVal string = strings.ToLower(V)
		switch StringVal {
		case "-sh", "sub-hostname":
			if index+1 <= len(pArgs) {
				SubHost = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, SUB-HOSTNAME after the TAG ", pArgs[index])
			}
		case "-ph", "pub-hostname":
			if index+1 <= len(pArgs) {
				PubHost = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, PUB-HOSTNAME after the TAG ", pArgs[index])
			}
		case "-sp", "sub-port":
			if index+1 <= len(pArgs) {
				SubPort = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, SUB-PORT after the TAG ", pArgs[index])
			}
		case "-pp", "pub-port":
			if index+1 <= len(pArgs) {
				PubPort = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, PUB-PORT after the TAG ", pArgs[index])
			}
		case "-sc", "sub-channel":
			if index+1 <= len(pArgs) {
				fromChannel = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, SUB-CHANNEL after the TAG ", pArgs[index])
			}
		case "-pc", "pub-channel":
			if index+1 <= len(pArgs) {
				toChannel = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, PUB-CHANNEL after the TAG ", pArgs[index])
			}
		}
	}

	fmt.Println("Subscribing to ", SubHost, " at Port", SubPort)
	fmt.Println("Publishing to ", PubHost, " at Port", PubPort)
	if fromChannel == "" || toChannel == "" {
		log.Fatal("Please Provide Channels for PUB and SUB")
	}
	fmt.Println("Subscribing to Channel ", fromChannel)
	fmt.Println("Publishing to Channel ", toChannel)

	// Making Params
	redisParam := RedisParams{
		SubHost:    SubHost,
		SubPort:    SubPort,
		PubHost:    PubHost,
		PubPort:    PubPort,
		SubChannel: fromChannel,
		PubChannel: toChannel,
	}

	// Now All Parameters are Set ... Handling Sub and Pub
	RedisHandler(redisParam)
}

// RedisHandler ... Redis Connextion Handler
func RedisHandler(args RedisParams) {
	// Get Instance for 2 Clients - PUB and SUB
	subClient := (RedInstance{args.SubHost, args.SubPort}).getRedisConnection()
	pubClient := (RedInstance{args.PubHost, args.PubPort}).getRedisConnection()

	redisMessage := make(chan *redis.Message)
	wg.Add(2)
	go handleSubClient(subClient, args.SubChannel, redisMessage)
	go handlePubClient(pubClient, args.PubChannel, redisMessage)
	wg.Wait()
}

func handleSubClient(subClient *redis.Client, channel string, redisChannel chan<- *redis.Message) {
	defer wg.Done()
	redisMsg := subClient.Subscribe(context.Background(), channel).Channel()
	for {
		redisChannel <- <-redisMsg
	}
}
func handlePubClient(pubClient *redis.Client, channel string, redisChannel <-chan *redis.Message) {
	defer wg.Done()
	for {
		Message := <-redisChannel
		pubClient.Publish(context.Background(), channel, Message.Payload)
	}
}

// RedInstance ... Redis Instance, form Multi Connections
type RedInstance struct {
	HostName string
	Port     string
}

func (r RedInstance) getRedisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.HostName + ":" + r.Port,
		DB:       0,
		Password: "",
	})
}
