package ipfs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Pin the given path.
func Pin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	var result Result

	err := sh.Pin(hash)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Status = "success"
	json.NewEncoder(w).Encode(result)
}
