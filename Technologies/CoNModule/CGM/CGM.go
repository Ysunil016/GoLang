package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

// RedisHOST ...
var RedisHOST string = "127.0.0.1"

// RedisPORT ...
var RedisPORT int = 6379

// RedisClient ...
var RedisClient *redis.Client

var wg sync.WaitGroup

// CurrentEntity ... the Unit Name for the Terminal
var CurrentEntity string

// GeoRedisHashKey ...
var GeoRedisHashKey = "CON_GEO"

// OwnRedisHashKey ...
var AllRedisHashKey = "COP_ALL"

// OwnRedisHashKey ...
var OwnRedisHashKey = "COP_OWN"

// LinkRedisHashKey ...
var LinkRedisHashKey = "COP_LINK"

// AisRedisHashKey ...
var AisRedisHashKey = "COP_AIS"

// AdsRedisHashKey ...
var AdsRedisHashKey = "COP_ADS"

// IcgRedisHashKey ...
var IcgRedisHashKey = "COP_ICG"

func init() {
	// Read the File from Path
	CurrentEntity = "RANA"
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisHOST + ":" + strconv.Itoa(RedisPORT),
		Password: "",
		DB:       0,
	})
}

func main() {

	RabbitMQChannel := make(chan []byte)

	wg.Add(2)
	go Listen2RabbitMq(RabbitMQChannel)
	go HandleReceivedTrack(RabbitMQChannel)
	fmt.Println("Starting CGM")

	wg.Wait()
	defer RedisClient.Close()
}
