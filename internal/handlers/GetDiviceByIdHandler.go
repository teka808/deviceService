package handlers

import (
	"fmt"
	"net/http"
	"test/internal/utils"
)

func GetDeviceByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := utils.GetVars(w, r)

	fmt.Fprintf(w, "Devices: %v\n", vars["id"])
}