package main

import "time"

type DataFake struct {
	Homes  []Home
	Plants []Plant
	Logs   []PlantReport
}

func (d DataFake) wateredPlantAt(PlantID string, time time.Time) {
	//TODO implement me
	panic("implement me")
}

func (d DataFake) getLatestWaterLogsForHome(HomeID string) WaterReport {
	h1 := Home{
		GroupID: "h1",
	}
	return WaterReport{
		Home: h1,
		Logs: []PlantReport{{
			ID:        0,
			Plant:     Plant{"p1", h1},
			WateredAt: fakeDate(),
		}},
	}
}

func (d DataFake) shutdown() {
	// I am fake.
	// I do not need to shut down.
}

func (d DataFake) getHouse(HomeID string) []Plant {
	return []Plant{
		Plant{"p1", Home{"h1"}},
	}
}

func fakeDate() time.Time {
	return time.Date(2002, time.July, 7, 0, 0, 0, 0, time.UTC)
}
