package main

// DELTA ...
type DELTA struct {
	ID  string
	KEY string
	LAT float64
	LNG float64
}

// Tracks Format ---------------------------******************---------------------------
// ************************************** POWNT *************************************

// OWN ...
type OWN struct {
	ID     string
	Popup  OWNPopup
	Header OWNHeader
	Tote   OWNTote
	Delta  OWNDelta
}

// OWNDelta ...
type OWNDelta struct {
	LAT float64
	LNG float64
	CRS float64
	SPD float64
	RNG float64
	HT  float64
	BRG float64
	TME string
	DTE string
}

// OWNHeader ...
type OWNHeader struct {
	ID  string
	SRC string
}

// OWNTote ...
type OWNTote struct {
}

// OWNPopup ...
type OWNPopup struct {
	ENTITY string
	RNG    float64
	BRG    float64
}

// ************************************** PRTRK *************************************

// LINK ...
type LINK struct {
	ID     string
	Popup  LinkPopup
	Header LinkHeader
	Tote   LinkTote
	Delta  LinkDelta
}

// LinkDelta ...
type LinkDelta struct {
	LAT float64
	LNG float64
	CRS float64
	SPD float64
	RNG float64
	HT  float64
	BRG float64
	TME string
	DTE string
}

// LinkHeader ...
type LinkHeader struct {
	ID  string
	SRC string
}

// LinkTote ...
type LinkTote struct {
}

// LinkPopup ...
type LinkPopup struct {
	ENTITY string
	RNG    float64
	BRG    float64
}

// ************************************** PCTRK *************************************

// AIS ...
type AIS struct {
	ID     string
	Popup  AISPopup
	Header AISHeader
	Tote   AISTote
	Delta  AISDelta
}

// AISDelta ...
type AISDelta struct {
	LAT float64
	LNG float64
	CRS float64
	SPD float64
	RNG float64
	HT  float64
	BRG float64
	TME string
	DTE string
}

// AISHeader ...
type AISHeader struct {
	ID  string
	SRC string
}

// AISTote ...
type AISTote struct {
}

// AISPopup ...
type AISPopup struct {
	ENTITY string
	RNG    float64
	BRG    float64
}

// ************************************** PADSB *************************************

// ADS ...
type ADS struct {
	ID     string
	Popup  ADSPopup
	Header ADSHeader
	Tote   ADSTote
	Delta  ADSDelta
}

// ADSDelta ...
type ADSDelta struct {
	LAT float64
	LNG float64
	CRS float64
	SPD float64
	RNG float64
	HT  float64
	BRG float64
	TME string
	DTE string
}

// ADSHeader ...
type ADSHeader struct {
	ID  string
	SRC string
}

// ADSTote ...
type ADSTote struct {
}

// ADSPopup ...
type ADSPopup struct {
	ENTITY string
	RNG    float64
	BRG    float64
}

// ************************************** PICG *************************************

// ICG ...
type ICG struct {
	ID     string
	Popup  ICGPopup
	Header ICGHeader
	Tote   ICGTote
	Delta  ICGDelta
}

// ICGDelta ...
type ICGDelta struct {
	LAT float64
	LNG float64
	CRS float64
	SPD float64
	RNG float64
	HT  float64
	BRG float64
	TME string
	DTE string
}

// ICGHeader ...
type ICGHeader struct {
	ID  string
	SRC string
}

// ICGTote ...
type ICGTote struct {
}

// ICGPopup ...
type ICGPopup struct {
	ENTITY string
	RNG    float64
	BRG    float64
}
