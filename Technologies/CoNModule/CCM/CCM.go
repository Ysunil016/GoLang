package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

// RedisSADLHOST ...
var RedisSADLHOST string = "172.16.43.105"

// RedisSADLPORT ...
var RedisSADLPORT int = 6379

// RedisSADLHOST ...
var RedisLocalHOST string = "127.0.0.1"

// RedisSADLPORT ...
var RedisLocalPORT int = 6379

var RedisSADLClient *redis.Client
var RedisLocalClient *redis.Client

func init() {
	// Read the File from Path
	RedisSADLClient = redis.NewClient(&redis.Options{
		Addr:     RedisSADLHOST + ":" + strconv.Itoa(RedisSADLPORT),
		Password: "",
		DB:       0,
	})
	RedisLocalClient = redis.NewClient(&redis.Options{
		Addr:     RedisLocalHOST + ":" + strconv.Itoa(RedisLocalPORT),
		Password: "",
		DB:       0,
	})
}

var wg sync.WaitGroup

func main() {
	// Receive Tracks from SADL --
	rChannel := make(chan *redis.Message)
	wg.Add(2)
	go Listen2SADL(rChannel)
	go HandleReceivedTrack(rChannel)
	fmt.Println("Starting CCM")
	wg.Wait()
	defer RedisSADLClient.Close()
	defer RedisLocalClient.Close()
}
