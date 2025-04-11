package main

import (
	"time"
)

type Home struct {
	GroupID string
}

type Plant struct {
	PlantID string
	home    Home
}

type PlantReport struct {
	ID        uint
	Plant     Plant
	WateredAt time.Time
}
