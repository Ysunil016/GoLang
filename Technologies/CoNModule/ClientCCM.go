package main


import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

// RedisHOST ...
var RedisLocalHOST string = "127.0.0.1"

// RedisPORT ...
var RedisLocalPORT int = 6379


// RedisSADLHOST ... Read from CONFIG File
var RedisSADLHOST string = "127.0.0.1"

// RedisSADLPORT ... Read from CONFIG File
var RedisSADLPORT int = 6379


// RedisSADLClient ... Receiver from SADL 
var RedisSADLClient *redis.Client
// RedisLocalClient ... Receiver from SADL 
var RedisLocalClient *redis.Client

var wg sync.WaitGroup

func init() {
	// Read the File from Path
	CurrentEntity = "RANA"
    

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


func main(){

    RedisChannel := make(chan []byte)
	wg.Add(2)
    go ListenToSADL(RedisChannel)
    go HandleReceivedTracks(RedisChannel)
    wg.Wait()
}


func ListenToSADL(redisChannel chan<- []byte) {
    // Listening to SADL 
    
}

func HandleReceivedTracks(rChannel <-chan string){
    // Handling And Saving Data to Local Redis

}





// Saving Delta for GEOREDIS -
DELTA{ID: ID, KEY: GeoRedisHashKey, LAT: LAT, LNG: LNG}.SaveDeltaData()
// Save Static Data to Redis In-Memory
OWN{ID: ID,
    Delta:  OWNDelta{LAT: LAT, LNG: LNG, CRS: CRS, SPD: SPD, HT: HTDT, RNG: 0, BRG: 0, DTE: DTE, TME: TME},
    Header: OWNHeader{ID: ID, SRC: SRC},
    Popup:  OWNPopup{ID, 0, 0},
    Tote:   OWNTote{},
}
.SaveOWNInRedis()