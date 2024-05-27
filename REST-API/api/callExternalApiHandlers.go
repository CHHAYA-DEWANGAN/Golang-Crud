package api

import (
	"REST-API/REST-API/client"
	"encoding/json"
	"fmt"
	"net/http"
)

func CallExternalApi(w http.ResponseWriter, r *http.Request) {

	result, err := client.CallExternal3rdPatryApi()

	if err != nil {
		return
	}
	fmt.Printf("Data get from externalApi %#v", result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
