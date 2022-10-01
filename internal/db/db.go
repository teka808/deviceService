package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"test/internal/config"
	"test/internal/utils"
)

func Open() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_DB,
		"disable")

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertOrUpdate(deviceId string, _type string, latitude float64, longitude float64, status string, timezone string) bool {
	db, err := Open()
	if err != nil {
		panic(err)
	}

	sql := fmt.Sprintf("INSERT INTO devices (device_id, type, latitude, longitude, status, timezone) "+
		"VALUES ('%s', '%s', %f, %f, '%s', '%s') "+
		"ON CONFLICT (device_id) DO UPDATE "+
		"SET type = excluded.type, "+
		"latitude = excluded.latitude, "+
		"longitude = excluded.longitude, "+
		"status = excluded.status, "+
		"timezone = excluded.timezone;", deviceId, _type, latitude, longitude, status, timezone)

	_, err = db.Exec(sql)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return true
}

type Device struct {
	device_id string
	_type     string
	latitude  float64
	longitude float64
	status    string
	timezone  string
}

func SearchByID(id string) utils.Response {
	db, err := Open()
	if err != nil {
		panic(err)
	}

	sql := fmt.Sprintf("SELECT device_id, type as _type, latitude, longitude, status, timezone"+
		" FROM devices WHERE device_id = '%s'", id)
	log.Println(sql)

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	response := utils.Response{}

	for rows.Next() {
		var device Device
		err = rows.Scan(&device.device_id, &device._type, &device.latitude, &device.longitude, &device.status, &device.timezone)
		response = BuildResponse(&device)
	}
	return response
}

func SearchByType(_type string, page int, limit int) []utils.Response {
	db, err := Open()
	if err != nil {
		panic(err)
	}

	sql := fmt.Sprintf("SELECT device_id, type as _type, latitude, longitude, status, timezone"+
		" FROM devices WHERE type = '%s'  OFFSET %d LIMIT %d", _type, page*limit, limit)
	log.Println(sql)

	return GetRows(db, sql)
}

func SearchByStatus(status string, page int, limit int) []utils.Response {
	db, err := Open()
	if err != nil {
		panic(err)
	}

	sql := fmt.Sprintf("SELECT device_id, type as _type, latitude, longitude, status, timezone"+
		" FROM devices WHERE status = '%s'  OFFSET %d LIMIT %d", status, page*limit, limit)
	log.Println(sql)

	return GetRows(db, sql)
}

func GetDevices(page int, limit int) []utils.Response {
	db, err := Open()
	if err != nil {
		panic(err)
	}

	sql := fmt.Sprintf("SELECT device_id, type as _type, latitude, longitude, status, timezone"+
		" FROM devices OFFSET %d LIMIT %d", page*limit, limit)
	log.Println(sql)

	return GetRows(db, sql)
}

func GetRows(db *sql.DB, sql string) []utils.Response {
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	var response []utils.Response

	for rows.Next() {
		var device Device
		err = rows.Scan(&device.device_id, &device._type, &device.latitude, &device.longitude, &device.status, &device.timezone)
		response = append(response, BuildResponse(&device))
	}
	return response
}

func BuildResponse(device *Device) utils.Response {
	response := utils.Response{}
	response.Device_id = device.device_id
	response.Type = device._type

	coordinates := utils.Coordinates{}
	coordinates.Latitude = device.latitude
	coordinates.Longitude = device.longitude
	response.Coordinates = coordinates

	response.Status = device.status
	response.Timezone = device.timezone
	return response
}
