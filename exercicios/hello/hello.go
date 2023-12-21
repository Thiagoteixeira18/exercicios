package hello

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleHello(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "word"
	}

	rw.Header().Add("Context-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]string{
		"message": fmt.Sprintf("Hello %s", name),
	})
}
