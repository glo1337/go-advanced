package response

import (
	"encoding/json"
	"net/http"
)

func Json(writer http.ResponseWriter, data any, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(data)
}
