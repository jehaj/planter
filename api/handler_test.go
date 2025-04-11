package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type header struct {
	Key   string
	Value string
}

func PerformRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPing(t *testing.T) {
	db := getFakeDataProvider()
	r := getRouter(db)
	w := PerformRequest(r, "GET", "/v1/ping")
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}

func TestGetPlantsForHome(t *testing.T) {
	db := getFakeDataProvider()
	defer db.shutdown()
	r := getRouter(db)
	w := PerformRequest(r, "GET", "/v1/group/h1")
	assert.Equal(t, "[{\"PlantID\":\"p1\"}]", w.Body.String())
}

func TestGetLatestWatering(t *testing.T) {
	db := getFakeDataProvider()
	defer db.shutdown()
	r := getRouter(db)
	w := PerformRequest(r, "GET", "/v1/water/h1")
	assert.Equal(t, "{\"Home\":{\"GroupID\":\"h1\"},\"Logs\":[{\"ID\":0,\"Plant\":{\"PlantID\":\"p1\"},\"WateredAt\":\"2002-07-07T00:00:00Z\"}]}", w.Body.String())
}

func TestGetHouseDB(t *testing.T) {
	db := getTestDataProvider()
	defer db.shutdown()
	r := getRouter(db)
	w := PerformRequest(r, "GET", "/v1/group/Group_A")
	expected := "[{\"PlantID\":\"Plant_001\"},{\"PlantID\":\"Plant_002\"},{\"PlantID\":\"Plant_003\"},{\"PlantID\":\"Plant_004\"},{\"PlantID\":\"Plant_005\"}]"
	assert.Equal(t, expected, w.Body.String())
}

func TestGetHouseDBforB(t *testing.T) {
	db := getTestDataProvider()
	defer db.shutdown()
	r := getRouter(db)
	w := PerformRequest(r, "GET", "/v1/group/Group_B")
	expected := "[{\"PlantID\":\"Plant_006\"},{\"PlantID\":\"Plant_007\"},{\"PlantID\":\"Plant_008\"},{\"PlantID\":\"Plant_009\"},{\"PlantID\":\"Plant_010\"}]"
	assert.Equal(t, expected, w.Body.String())
}

func TestWateringPlant1(t *testing.T) {
	db := getTestDataProvider()
	defer db.shutdown()
	r := getRouter(db)
	w := PerformRequest(r, "POST", "/v1/water/Plant_001")
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func getFakeDataProvider() DataProvider {
	return new(DataFake)
}

func getTestDataProvider() DataProvider {
	dpdb := create("test.db")
	insertFakeData(dpdb.db)
	return &dpdb
}

func insertFakeData(db *sql.DB) {
	sqlSTMT := `-- Insert sample ddpdbinto Groups table
	INSERT INTO Groups(GroupID) VALUES
		('Group_A'),
		('Group_B');

	-- Insert sample data into Plants table
	INSERT INTO Plants (PlantID, GroupID) VALUES
	('Plant_001', 'Group_A'),
	('Plant_002', 'Group_A'),
	('Plant_003', 'Group_A'),
	('Plant_004', 'Group_A'),
	('Plant_005', 'Group_A'),
	('Plant_006', 'Group_B'),
	('Plant_007', 'Group_B'),
	('Plant_008', 'Group_B'),
	('Plant_009', 'Group_B'),
	('Plant_010', 'Group_B');`

	db.Exec(sqlSTMT)
}
