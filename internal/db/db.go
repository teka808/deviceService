package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"test/internal/config"
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
