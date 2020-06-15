package main

import (
	"fmt"
	"strconv"
	"strings"
)

// HandleReceivedTrack ...
func HandleReceivedTrack(track <-chan []byte) {
	defer wg.Done()
	for {
		Track := strings.Split(string(<-track), ",")
		switch Track[0] {
		case "$POWNT":
			HandlePOWNT(Track)
		case "$PRTRK":
			HandlePRTRK(Track)
		case "$PCTRK":
			HandlePCTRK(Track)
		case "$PADSB":
			HandlePADSB(Track)
		case "$PCGTRK":
			HandlePICG(Track)
		}
	}
}

// HandlePOWNT ...
func HandlePOWNT(track []string) {
	fmt.Println("Received POWNT")
	fmt.Println(track)
	// Save Tracks to Redis Geo
	LAT, err := strconv.ParseFloat(track[6], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LAT")
	}
	LNG, err := strconv.ParseFloat(track[7], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LNG")
	}
	CRS, err := strconv.ParseFloat(track[4], 64)
	if err != nil {
		fmt.Println("Could'nt Convert CRS")
	}
	SPD, err := strconv.ParseFloat(track[5], 64)
	if err != nil {
		fmt.Println("Could'nt Convert SPD")
	}
	HTDT, err := strconv.ParseFloat(track[8], 64)
	if err != nil {
		fmt.Println("Could'nt Convert Height")
	}
	ID := track[1]
	// CAT := track[2]
	SRC := track[3]
	// CAL := track[9]
	DTE := track[10]
	TME := track[11]

	// Saving Delta for GEOREDIS -
	DELTA{ID: ID, KEY: GeoRedisHashKey, LAT: LAT, LNG: LNG}.SaveDeltaData()
	// Save Static Data to Redis In-Memory
	OWN{ID: ID,
		Delta:  OWNDelta{LAT: LAT, LNG: LNG, CRS: CRS, SPD: SPD, HT: HTDT, RNG: 0, BRG: 0, DTE: DTE, TME: TME},
		Header: OWNHeader{ID: ID, SRC: SRC},
		Popup:  OWNPopup{ID, 0, 0},
		Tote:   OWNTote{},
	}.SaveOWNInRedis()
}

// HandlePRTRK ...
func HandlePRTRK(track []string) {
	fmt.Println("Received PRTRK")
	fmt.Println(track)
	// Save Tracks to Redis Geo
	LAT, err := strconv.ParseFloat(track[10], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LAT")
	}
	LNG, err := strconv.ParseFloat(track[11], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LNG")
	}
	CRS, err := strconv.ParseFloat(track[8], 64)
	if err != nil {
		fmt.Println("Could'nt Convert CRS")
	}
	SPD, err := strconv.ParseFloat(track[9], 64)
	if err != nil {
		fmt.Println("Could'nt Convert SPD")
	}
	HTDT, err := strconv.ParseFloat(track[12], 64)
	if err != nil {
		fmt.Println("Could'nt Convert Height")
	}
	ID := track[2]
	// REP := track[3]
	// LTN := track[4]
	// IDN := track[5]
	SRC := track[6]
	// LBL := track[7]
	DTE := track[13]
	TME := track[14]

	DELTA{ID: ID, KEY: GeoRedisHashKey, LAT: LAT, LNG: LNG}.SaveDeltaData()
	// Save Static Data to Redis In-Memory
	LINK{
		ID:     ID,
		Delta:  LinkDelta{LAT: LAT, LNG: LNG, CRS: CRS, SPD: SPD, RNG: 0, BRG: 0, DTE: DTE, TME: TME, HT: HTDT},
		Header: LinkHeader{ID: ID, SRC: SRC},
		Popup:  LinkPopup{ID, 0, 0},
		Tote:   LinkTote{},
	}.SaveLINKInRedis()
}

// HandlePCTRK ...
func HandlePCTRK(track []string) {
	fmt.Println("Received PCTRK")

	fmt.Println(track)
	// Save Tracks to Redis Geo
	LAT, err := strconv.ParseFloat(track[8], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LAT")
	}
	LNG, err := strconv.ParseFloat(track[9], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LNG")
	}
	CRS, err := strconv.ParseFloat(track[6], 64)
	if err != nil {
		fmt.Println("Could'nt Convert CRS")
	}
	SPD, err := strconv.ParseFloat(track[7], 64)
	if err != nil {
		fmt.Println("Could'nt Convert SPD")
	}
	SRC := track[3]
	ID := track[1]
	DTE := track[10]
	TME := track[11]

	// Saving in Geo Redis
	DELTA{ID: track[1], KEY: GeoRedisHashKey, LAT: LAT, LNG: LNG}.SaveDeltaData()

	// Save Static Data to Redis In-Memory
	AIS{
		ID:     ID,
		Delta:  AISDelta{BRG: 0, RNG: 0, CRS: CRS, DTE: DTE, HT: 0, LAT: LAT, LNG: LNG, SPD: SPD, TME: TME},
		Header: AISHeader{ID: ID, SRC: SRC},
		Popup:  AISPopup{ID, 0, 0},
		Tote:   AISTote{},
	}.SaveAISInRedis()
}

// HandlePADSB ...
func HandlePADSB(track []string) {
	fmt.Println("Received ADSB")
	fmt.Println(track)
	// Save Tracks to Redis Geo
	LAT, err := strconv.ParseFloat(track[11], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LAT")
	}
	LNG, err := strconv.ParseFloat(track[12], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LNG")
	}

	CRS, err := strconv.ParseFloat(track[9], 64)
	if err != nil {
		fmt.Println("Could'nt Convert CRS")
	}
	SPD, err := strconv.ParseFloat(track[10], 64)
	if err != nil {
		fmt.Println("Could'nt Convert SPD")
	}
	HT, err := strconv.ParseFloat(track[13], 64)
	if err != nil {
		fmt.Println("Could'nt Convert SPD")
	}
	ID := track[1]
	TSRC := "ADSB"
	DTE := track[14]
	TME := track[15]

	// Saving in Geo Redis
	DELTA{ID: ID, KEY: GeoRedisHashKey, LAT: LAT, LNG: LNG}.SaveDeltaData()
	// Save Static Data to Redis In-Memory
	ADS{
		ID:     ID,
		Delta:  ADSDelta{BRG: 0, RNG: 0, CRS: CRS, DTE: DTE, HT: HT, LAT: LAT, LNG: LAT, SPD: SPD, TME: TME},
		Header: ADSHeader{ID: ID, SRC: TSRC},
		Popup:  ADSPopup{ID, 0, 0},
		Tote:   ADSTote{},
	}.SaveADSInRedis()
}

// HandlePICG ...
func HandlePICG(track []string) {
	fmt.Println("Received ICG")
	fmt.Println(track)
	// Save Tracks to Redis Geo
	LAT, err := strconv.ParseFloat(track[10], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LAT")
	}
	LNG, err := strconv.ParseFloat(track[11], 64)
	if err != nil {
		fmt.Println("Could'nt Convert LNG")
	}

	DELTA{ID: track[1], KEY: GeoRedisHashKey, LAT: LAT, LNG: LNG}.SaveDeltaData()

	// Save Static Data to Redis In-Memory
}
