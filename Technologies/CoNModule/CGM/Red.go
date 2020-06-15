package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

// Redis Geo-Save   ---------------------------******************---------------------------

// SaveDeltaData ...
func (delta DELTA) SaveDeltaData() {
	fmt.Println("Store Data In Redis")
	_, err := RedisClient.GeoAdd(context.Background(),
		delta.KEY, &redis.GeoLocation{
			Name:      delta.ID,
			Latitude:  delta.LAT,
			Longitude: delta.LNG,
		}).Result()
	if err != nil {
		log.Println(err)
	}
}

// SaveOWNInRedis ...
func (t OWN) SaveOWNInRedis() {
	JM, err := json.Marshal(t)
	if err != nil {
		log.Println("UNMARSHAL SAVE OWN", err)
	}
	done, err := RedisClient.HMSet(context.Background(), OwnRedisHashKey, t.ID, string(JM)).Result()
	if done {
		log.Println("SAVE OWN")
	}
}

// SaveLINKInRedis ...
func (t LINK) SaveLINKInRedis() {
	JM, err := json.Marshal(t)
	if err != nil {
		log.Println("UNMARSHAL SAVE LINK", err)
	}
	done, err := RedisClient.HMSet(context.Background(), LinkRedisHashKey, t.ID, string(JM)).Result()
	if done {
		log.Println("SAVE LINK")
	}
}

// SaveAISInRedis ...
func (t AIS) SaveAISInRedis() {
	JM, err := json.Marshal(t)
	if err != nil {
		log.Println("UNMARSHAL SAVE AIS", err)
	}
	done, err := RedisClient.HMSet(context.Background(), AisRedisHashKey, t.ID, string(JM)).Result()
	if done {
		log.Println("SAVE AIS")
	}
}

// SaveADSInRedis ...
func (t ADS) SaveADSInRedis() {
	JM, err := json.Marshal(t)
	if err != nil {
		log.Println("UNMARSHAL SAVE AIS", err)
	}
	done, err := RedisClient.HMSet(context.Background(), AdsRedisHashKey, t.ID, string(JM)).Result()
	if done {
		log.Println("SAVE ADS ", t.ID)
	}
}

// SaveICGInRedis ...
func (t ICG) SaveICGInRedis() {
	JM, err := json.Marshal(t)
	if err != nil {
		log.Println("UNMARSHAL SAVE ICG", err)
	}
	done, err := RedisClient.HMSet(context.Background(), IcgRedisHashKey, t.ID, string(JM)).Result()
	if done {
		log.Println("SAVE ICG")
	}
}
