package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func Listen2SADL(rchan chan<- *redis.Message) {
	defer wg.Done()
	rPubSub := RedisSADLClient.Subscribe(context.Background(), "LINK_CON_POWNT", "LINK_CON_PRTRK", "LINK_CON_PCTRK", "LINK_CON_PADSB")
	for {
		rchan <- <-rPubSub.Channel()
	}
}

func SaveInLocalHashRedis(Key string, ID string, Info string) bool {
	_, err := RedisLocalClient.HMSet(context.Background(), Key, ID, Info).Result()
	if err != nil {
		log.Print("Could Not Store Data in Redis Hash", err)
		return false
	}
	return true
}

func SaveInLocalRedis(Key string, Info string, expiration time.Duration) bool {
	_, err := RedisLocalClient.Set(context.Background(), Key, Info, expiration).Result()
	if err != nil {
		log.Print("Could Not Store Data in Redis Hash", err)
		return false
	}
	return true
}
