package ipfs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Resolve gets resolves the string provided to an /ipns/[name]. If asked to resolve an empty string, resolve instead resolves the node's own /ipns value.
func resolve(id string) Result {
	var result Result

	res, err := sh.Resolve(id)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	result.Status = "success"
	result.Data = res
	return result
}

func Resolve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	json.NewEncoder(w).Encode(resolve(id))
}
