package ipfs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// List entries at the given path.
func List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	var result Result

	data, err := sh.List(hash)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Status = "success"
	result.Data = data
	json.NewEncoder(w).Encode(result)
}
