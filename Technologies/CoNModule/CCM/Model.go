package main

import (
	"encoding/json"
	"log"
	"time"
)

type OWN struct {
	Header    string `json:"header"`
	Category  string `json:"category"`
	Source    string `json:"source"`
	Rep       string `json:"net_unit"`
	Call      string `json:"call_sign"`
	Course    string `json:"course"`
	Speed     string `json:"speed"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Height    string `json:"height_depth"`
	Time      string `json:"time"`
	Date      string `json:"date"`
}

func (t OWN) conv2Univ() UNIV {
	return UNIV{
		ID:        t.Rep,
		Header:    "OWN",
		LTN:       "0",
		Class:     "0",
		Category:  t.Category,
		Identity:  "F",
		Source:    t.Source,
		Rep:       t.Rep,
		Label:     "NIL",
		Course:    t.Course,
		Speed:     t.Speed,
		Latitude:  t.Latitude,
		Longitude: t.Longitude,
		Height:    t.Height,
		Time:      t.Time,
		Date:      t.Time,
	}
}

/*
{\"header\":\"$POWNT\",\"net_unit\":\"165\",\"category\":\"AIR\",
\"source\":\"GPS\",\"course\":\"069\",\"speed\":\"0027\",\"call_sign\":\"MOCV\",\"latitude\":\"16.85\",
\"longitude\":\"83.8\",\"height_depth\":\"2500\",\"time\":\"12:50:00\",\"date\":\"03-12-2019\"}"
*/

type LINK struct {
	Header    string `json:"header"`
	LTN       string `json:"link_track_number"`
	Class     string `json:"class"`
	Category  string `json:"category"`
	Identity  string `json:"identity"`
	Source    string `json:"source"`
	Rep       string `json:"rep_unit"`
	Label     string `json:"label"`
	Course    string `json:"course"`
	Speed     string `json:"speed"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Height    string `json:"height_depth"`
	Time      string `json:"time"`
	Date      string `json:"date"`
}

func (t LINK) conv2Univ() UNIV {
	return UNIV{
		ID:        t.Rep + t.LTN,
		Header:    "LINK",
		LTN:       t.LTN,
		Class:     t.Class,
		Category:  t.Category,
		Identity:  t.Identity,
		Source:    t.Source,
		Rep:       t.Rep,
		Label:     t.Label,
		Course:    t.Course,
		Speed:     t.Speed,
		Latitude:  t.Latitude,
		Longitude: t.Longitude,
		Height:    t.Height,
		Time:      t.Time,
		Date:      t.Time,
	}
}

/*
{\"header\":\"$PRTRK\",\"link_track_number\":\"127\",\"class\":\"MAIN BODY\",\"category\":\"SUR\",
\"identity\":\"H\",\"source\":\"CMS\",\"rep_unit\":\"165\",\"track_label\":\"TST25\",
\"course\":\"045\",\"speed\":\"16\",\"latitude\":\"16.3223\",\"longitude\":\"86.2323\",
\"height_depth\":\"0\",\"time\":\"12:21:11\",\"date\":\"03-12-2019\"}
*/

type UNIV struct {
	ID        string `json:"ID"`
	Header    string `json:"Header"`
	LTN       string `json:"LTN"`
	Class     string `json:"Class"`
	Category  string `json:"Category"`
	Identity  string `json:"Identity"`
	Source    string `json:"Source"`
	Rep       string `json:"Rep"`
	Label     string `json:"Label"`
	Course    string `json:"Course"`
	Speed     string `json:"Speed"`
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
	Height    string `json:"Height"`
	Time      string `json:"Time"`
	Date      string `json:"Date"`
}

func (t UNIV) StoreInRedisHash(Key string) UNIV {
	jsonMarshal, err := json.Marshal(t)
	if err != nil {
		log.Print("Could Marshal UNIV Data", err)
		return t
	}
	if !SaveInLocalHashRedis(Key, t.ID, string(jsonMarshal)) {
		log.Print("Could Not Save Data in Hash PUNIV", err)
	}
	return t
}

func (t UNIV) StoreOWNRedis(Key string, expiration time.Duration) UNIV {
	jsonMarshal, err := json.Marshal(t)
	if err != nil {
		log.Print("Could Marshal UNIV Data", err)
		return t
	}
	if !SaveInLocalRedis(Key, string(jsonMarshal), expiration) {
		log.Print("Could Not Save Data in Hash PUNIV", err)
	}
	return t
}

func (t UNIV) StoreInRedis(Key string) UNIV {
	jsonMarshal, err := json.Marshal(t)
	if err != nil {
		log.Print("Could Marshal UNIV Data", err)
		return t
	}
	if !SaveInLocalHashRedis(Key, t.ID, string(jsonMarshal)) {
		log.Print("Could Not Save Data in Hash PUNIV", err)
	}
	return t
}
