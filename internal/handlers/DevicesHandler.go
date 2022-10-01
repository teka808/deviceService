package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test/internal/db"
	"test/internal/utils"
)

func GetDeviceByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := utils.GetVars(w, r)
	if vars == nil || vars["id"] == "" {
		fmt.Fprintf(w, "{'error':'ID required'}")
	}
	res := db.SearchByID(vars["id"])
	b, err := json.Marshal(&res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(b))
}

func GetDeviceByTypeHandler(w http.ResponseWriter, r *http.Request) {
	vars := utils.GetVars(w, r)
	if vars == nil || vars["type"] == "" {
		fmt.Fprintf(w, "{'error':'Type required'}")
	}

	_page := r.URL.Query().Get("page")
	_limit := r.URL.Query().Get("limit")

	page := 0
	limit := 50

	p, err1 := strconv.Atoi(_page)
	if err1 == nil {
		page = p
	}

	l, err2 := strconv.Atoi(_limit)
	if err2 == nil {
		limit = l
	}

	res := db.SearchByType(vars["type"], page, limit)
	b, err := json.Marshal(&res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(b))
}

func GetDeviceByStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := utils.GetVars(w, r)
	if vars == nil || vars["status"] == "" {
		fmt.Fprintf(w, "{'error':'status required'}")
	}

	_page := r.URL.Query().Get("page")
	_limit := r.URL.Query().Get("limit")

	page := 0
	limit := 50

	p, err1 := strconv.Atoi(_page)
	if err1 == nil {
		page = p
	}

	l, err2 := strconv.Atoi(_limit)
	if err2 == nil {
		limit = l
	}

	res := db.SearchByStatus(vars["type"], page, limit)
	b, err := json.Marshal(&res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(b))
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	_page := r.URL.Query().Get("page")
	_limit := r.URL.Query().Get("limit")

	page := 0
	limit := 50

	p, err1 := strconv.Atoi(_page)
	if err1 == nil {
		page = p
	}

	l, err2 := strconv.Atoi(_limit)
	if err2 == nil {
		limit = l
	}

	res := db.GetDevices(page, limit)
	b, err := json.Marshal(&res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(b))
}
