package main

import "time"

type DataProvider interface {
	getHouse(HomeID string) []Plant
	getLatestWaterLogsForHome(HomeID string) WaterReport
	shutdown()
	wateredPlantAt(PlantID string, time time.Time)
}

type WaterReport struct {
	Home Home
	Logs []PlantReport
}
