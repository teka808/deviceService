package handlers

import (
	"fmt"
	"net/http"
	"test/internal/utils"
)

func UpdatesHandler(w http.ResponseWriter, r *http.Request) {
	vars := utils.GetVars(w, r)

	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
