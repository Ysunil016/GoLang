package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/go-redis/redis"
)

var wg sync.WaitGroup

// RedisParams ...
type RedisParams struct {
	SourceHost    string
	SourcePort    string
	SourceChannel string
}

func main() {
	// argsWithProg := os.Args
	pArgs := os.Args[1:]

	// SubHost ... Hostname for Subscription.
	var SourceHost string = "127.0.0.1"
	// SubPort ... PORT for Subscription.
	var SourcePort string = "6379"
	// SourceChannel ... Channel for GEO.
	var SourceChannel string = ""

	for index, V := range pArgs {
		var StringVal string = strings.ToLower(V)
		switch StringVal {
		case "-sh", "source-hostname":
			if index+1 <= len(pArgs) {
				SourceHost = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, Source-HOSTNAME after the TAG ", pArgs[index])
			}
		case "-ph", "source-port":
			if index+1 <= len(pArgs) {
				SourcePort = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, Source-PORT after the TAG ", pArgs[index])
			}
		case "-sc", "source-channel":
			if index+1 <= len(pArgs) {
				SourceChannel = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, Source-CHANNEL after the TAG ", pArgs[index])
			}
		}
	}

	if SourceChannel == "" {
		log.Fatal("Please Provide Source Channel")
	}

	fmt.Println("Geo Source from ", SourceHost, SourcePort, " at Channel ", SourceChannel)
	// Making Params
	redisParam := RedisParams{
		SourceHost:    SourceHost,
		SourcePort:    SourcePort,
		SourceChannel: SourceChannel,
	}

	// Now All Parameters are Set ... Handling Sub and Pub
	RedisHandler(redisParam)
}

// RedisHandler ... Redis Connextion Handler
func RedisHandler(args RedisParams) {
	// Get Instance for 2 Clients - PUB and SUB
	client := (RedInstance{args.SourceHost, args.SourcePort}).getRedisConnection()
	redisMessage := make(chan *redis.GeoLocation)

	wg.Add(1)
	// Get LAT-LNG from Source
	go handleSourceClient(client, args.SourceChannel, redisMessage)
	wg.Wait()
}

func storeGeoLocation(client *redis.Client, HashKey string, ID string) {
	_, err := client.GeoAdd(context.Background(), HashKey, &redis.GeoLocation{Name: ID, Longitude: -12.32, Latitude: 34}).Result()
	if err != nil {
		log.Fatal(err)
	}
}

func handleSourceClient(client *redis.Client, channel string, redisChannel chan<- *redis.GeoLocation) {
	defer wg.Done()
	redisMsg := client.Subscribe(context.Background(), channel).Channel()
	for {
		msqRec := <-redisMsg
		msgSlice := strings.Split(msqRec.Payload, " ") // KEY ID LAT LONG
		Lat, _ := strconv.ParseFloat(msgSlice[2], 64)
		Lng, _ := strconv.ParseFloat(msgSlice[3], 64)
		Key := msgSlice[0]
		ID := msgSlice[1]
		client.GeoAdd(context.Background(), Key, &redis.GeoLocation{
			Name:      ID,
			Latitude:  Lat,
			Longitude: Lng,
		})
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
