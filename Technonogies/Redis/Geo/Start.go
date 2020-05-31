package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var wg sync.WaitGroup

// RedisParams ...
type RedisParams struct {
	SourceHost    string
	SourcePort    string
	GeoChannel    string
	StaticChannel string
}

func main() {
	// argsWithProg := os.Args
	pArgs := os.Args[1:]

	// SubHost ... Hostname for Subscription.
	var SourceHost string = "127.0.0.1"
	// SubPort ... PORT for Subscription.
	var SourcePort string = "6379"
	// SourceChannel ... Channel for GEO.
	var GeoChannel string = "GO_GEO_SOURCE"
	var StaticChannel string = "GO_STATIC_SOURCE"

	for index, V := range pArgs {
		var StringVal string = strings.ToLower(V)
		switch StringVal {
		case "-sh", "source-hostname":
			if index+1 <= len(pArgs) {
				SourceHost = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, Source-HOSTNAME after the TAG ", pArgs[index])
			}
		case "-sp", "source-port":
			if index+1 <= len(pArgs) {
				SourcePort = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, Source-PORT after the TAG ", pArgs[index])
			}
		case "-gc", "geo-channel":
			if index+1 <= len(pArgs) {
				GeoChannel = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, Geo-CHANNEL after the TAG ", pArgs[index])
			}
		case "-sc", "static-channel":
			if index+1 <= len(pArgs) {
				StaticChannel = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, Static-CHANNEL after the TAG ", pArgs[index])
			}
		}

	}

	fmt.Println("Geo Source from ", SourceHost, SourcePort, " Listening to Channel ", GeoChannel, " for GEO and ", StaticChannel, "for Its Static")
	// Making Params
	redisParam := RedisParams{
		SourceHost:    SourceHost,
		SourcePort:    SourcePort,
		GeoChannel:    GeoChannel,
		StaticChannel: StaticChannel,
	}

	// Now All Parameters are Set ... Handling Sub and Pub
	RedisHandler(redisParam)
}

const (
	// TrackLifeTime ... Defines the Expiration of Tracks in Redis
	TrackLifeTime time.Duration = time.Minute * 3 // 3 Mins
)

// RedisHandler ... Redis Connextion Handler
func RedisHandler(args RedisParams) {
	// Get Instance for 2 Clients - PUB and SUB
	client := (RedInstance{args.SourceHost, args.SourcePort}).getRedisConnection()

	wg.Add(2)
	// Get LAT-LNG and ID from Source to RedisGeo
	go handleGeoData(client, args.GeoChannel)       // Receives Geo Data and Saves Data to Redis in Geo
	go handleStaticData(client, args.StaticChannel) // Receives the Data and Save in Redis Store.
	wg.Wait()
}

// StaticData ....
type StaticData struct {
	Key   string
	Value string
}

// MarshalBinary ....
func (m StaticData) MarshalBinary() ([]byte, error) {
	fmt.Println(m)
	return json.Marshal(m)
}

func handleStaticData(client *redis.Client, channel string) {
	defer wg.Done()
	redisMsg := client.Subscribe(context.Background(), channel).Channel()
	for {
		msqRec := <-redisMsg
		msgSlice := strings.Split(msqRec.Payload, "@") // ID$<DATA>
		ID := msgSlice[0]
		StaticData := msgSlice[1]
		// Saving Directly in Redis ... And Applying Time Duration to IT
		_, err := client.Set(context.Background(), ID, StaticData, TrackLifeTime).Result()
		if err != nil {
			log.Println("Not Able to Save Data in Redis - Static Data for ID", msgSlice[1], err)
		} else {
			log.Println("Saved Data in Redis - Static Data for ID")
		}
	}
}

func handleGeoData(client *redis.Client, channel string) {
	defer wg.Done()
	redisMsg := client.Subscribe(context.Background(), channel).Channel()
	for {
		msqRec := <-redisMsg
		msgSlice := strings.Split(msqRec.Payload, "@") // KEY$ID$LAT$LONG
		Lat, _ := strconv.ParseFloat(msgSlice[2], 64)
		Lng, _ := strconv.ParseFloat(msgSlice[3], 64)
		Key := msgSlice[0]
		ID := msgSlice[1]
		_, err := client.GeoAdd(context.Background(), Key, &redis.GeoLocation{
			Name:      ID,
			Latitude:  Lat,
			Longitude: Lng,
		}).Result()
		if err != nil {
			log.Println("Not Able to Save Geo Data", err)
		} else {
			log.Println("Saved Geo Data", ID)
		}
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
