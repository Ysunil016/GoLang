package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	LocalPort     string
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
	var LocalPort string = "6379"

	// SourceChannel ... Channel for GEO.
	var GeoChannel string = "GO_GEO_SOURCE"
	// SourceChannel ... Channel for Static Data.
	var StaticChannel string = "GO_STATIC_SOURCE"

	// ListeningPort ... PORT that is Listening for Accepting HTTP Request.
	var ListeningPort string = ":8080"

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
		case "-lp", "local-port":
			if index+1 <= len(pArgs) {
				LocalPort = pArgs[index+1]
			} else {
				log.Fatal("Please Provide, Local-Port after the TAG ", pArgs[index])
			}
		case "-port", "port":
			if index+1 <= len(pArgs) {
				ListeningPort = ":" + pArgs[index+1]
			} else {
				log.Fatal("Please Provide, Listening at Port ", pArgs[index])
			}
		}

	}

	fmt.Println("Geo Source from ", SourceHost, SourcePort, " Listening to Channel ", GeoChannel, " for GEO and ", StaticChannel, "for Its Static")
	// Making Params
	redisParam := RedisParams{
		SourceHost:    SourceHost,
		SourcePort:    SourcePort,
		LocalPort:     LocalPort,
		GeoChannel:    GeoChannel,
		StaticChannel: StaticChannel,
	}

	// Now All Parameters are Set ... Handling Sub and Pub
	go RedisHandler(redisParam) // On Seperate Core

	// Listening to Expose API
	fmt.Println("Server Listening at Port", ListeningPort)
	http.HandleFunc("/geoRadius", getGeoRadius)
	log.Fatal(http.ListenAndServe(ListeningPort, nil))
}

// GeoRadiusRequest ...
type GeoRadiusRequest struct {
	KEY string
	LAT float64
	LNG float64
	RAD float64
}

func getGeoRadius(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, "Please Use POST METHOD, with Required Parameters")
		return
	}
	P := GeoRadiusRequest{}
	err := json.NewDecoder(r.Body).Decode(&P)
	if err != nil {
		fmt.Fprint(w, "Could Not Parse Your Received Parameter")
	}
	findIDsInRad(w, P)
}

func findIDsInRad(w http.ResponseWriter, params GeoRadiusRequest) {
	client := (RedInstance{"127.0.0.1", "6379"}).getRedisConnection()
	defer client.Close()
	geoPos, _ := client.GeoRadius(context.Background(), params.KEY, params.LNG, params.LAT, &redis.GeoRadiusQuery{
		Radius:      params.RAD,
		WithCoord:   true,
		WithDist:    true,
		WithGeoHash: true,
		Unit:        "km",
	}).Result()

	// if err.Error() != "nil" {
	// 	fmt.Fprint(w, "Could Not Fetch GeoPos")
	// 	return
	// }
	json.NewEncoder(w).Encode(geoPos)
}

const (
	// TrackLifeTime ... Defines the Expiration of Tracks in Redis
	TrackLifeTime time.Duration = time.Minute * 3 // 3 Mins
)

// RedisHandler ... Redis Connextion Handler
func RedisHandler(args RedisParams) {
	// Get Instance for 2 Clients - PUB and SUB
	client := (RedInstance{args.SourceHost, args.SourcePort}).getRedisConnection()
	defer client.Close()
	wg.Add(2)
	// Get LAT-LNG and ID from Source to RedisGeo
	go handleGeoData(client, args.GeoChannel, args) // Receives Geo Data and Saves Data to Redis in Geo
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

func handleGeoData(client *redis.Client, channel string, args RedisParams) {
	defer wg.Done()
	localClient := (RedInstance{"127.0.0.1", args.LocalPort}).getRedisConnection()
	defer client.Close()

	redisMsg := client.Subscribe(context.Background(), channel).Channel()
	for {
		msqRec := <-redisMsg
		msgSlice := strings.Split(msqRec.Payload, "@") // KEY$ID$LAT$LONG
		Lat, _ := strconv.ParseFloat(msgSlice[2], 64)
		Lng, _ := strconv.ParseFloat(msgSlice[3], 64)
		Key := msgSlice[0]
		ID := msgSlice[1]
		// Saving Geo Data in Local Redis
		_, err := localClient.GeoAdd(context.Background(), Key, &redis.GeoLocation{
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
