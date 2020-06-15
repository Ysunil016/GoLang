package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

// HandleReceivedTrack ...
func HandleReceivedTrack(rchan <-chan *redis.Message) {
	defer wg.Done()

	for {
		info := <-rchan
		switch info.Channel {
		case "LINK_CON_POWNT":
			HandlePOWNTTrack(info.Payload)
		case "LINK_CON_PRTRK":
			HandlePRTRKTrack(info.Payload)
		case "LINK_CON_PCTRK":
			HandlePCTRKTrack(info.Payload)
		case "LINK_CON_PADSB":
			HandlePADSBTrack(info.Payload)
		default:
			log.Print("Not Recognized Channel")
		}
	}

}

// HandlePOWNTTrack ...
func HandlePOWNTTrack(track string) {
	_OWN := OWN{}
	err := json.Unmarshal([]byte(track), &_OWN)
	if err != nil {
		log.Print("Could Not UnMarshal Track Info OWN Section", err)
		return
	}
	_OWN.conv2Univ().StoreInRedisHash("POWNT").StoreInRedisHash("PUNIV").StoreOWNRedis("COP_OWN_POS", time.Minute*2)
}

// HandlePRTRKTrack ...
func HandlePRTRKTrack(track string) {
	_LINK := LINK{}
	err := json.Unmarshal([]byte(track), &_LINK)
	if err != nil {
		log.Print("Could Not UnMarshal Track Info LINK Section", err)
		return
	}
	_LINK.conv2Univ().StoreInRedisHash("PRTRK").StoreInRedisHash("PUNIV")
}

// HandlePCTRKTrack ...
func HandlePCTRKTrack(track string) {
	TrackInfo := LINK{}
	err := json.Unmarshal([]byte(track), &TrackInfo)
	if err != nil {
		log.Print("Could Not UnMarshal Track Info AIS Section", err)
		return
	}
	fmt.Println(TrackInfo)

}

// HandleReceivedTrack ...
func HandlePADSBTrack(track string) {
	TrackInfo := LINK{}
	err := json.Unmarshal([]byte(track), &TrackInfo)
	if err != nil {
		log.Print("Could Not UnMarshal Track Info ADS Section", err)
		return
	}
	fmt.Println(TrackInfo)

}
