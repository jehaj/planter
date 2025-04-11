package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

type DPDB struct {
	db *sql.DB
}

func create(dbPath string) DPDB {
	_, err := os.Open(dbPath)
	alreadyExists := !os.IsNotExist(err)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	dpdb := DPDB{db: db}
	if !alreadyExists {
		createTables(dpdb)
	}
	return dpdb
}

func (D DPDB) getHouse(GroupID string) []Plant {
	rows, err := D.db.Query("SELECT * FROM Plants WHERE GroupID = ?", GroupID)
	if err != nil {
		log.Println(err)
		return []Plant{}
	}
	defer rows.Close()
	plants := make([]Plant, 0)
	for rows.Next() {
		var plantID string
		var groupID string
		err = rows.Scan(&plantID, &groupID)
		home := Home{groupID}
		plant := Plant{plantID, home}
		if err != nil {
			log.Println(err)
		}
		plants = append(plants, plant)
	}
	return plants
}

func (D DPDB) getLatestWaterLogsForHome(HomeID string) WaterReport {
	//TODO implement me
	panic("implement me")
}

func (D DPDB) shutdown() {
	D.db.Close()
}

func (D DPDB) wateredPlantAt(PlantID string, time time.Time) {
	//TODO implement me
	panic("implement me")
}

func createTables(D DPDB) {
	sqlSTMT := `create table Groups
	(
		GroupID TEXT not null
			constraint Groups_pk
				primary key
	);
	
	create table Plants
	(
		PlantID TEXT not null
			constraint Plants_pk
				primary key,
		GroupID TEXT not null
	);`
	D.db.Exec(sqlSTMT)
}
