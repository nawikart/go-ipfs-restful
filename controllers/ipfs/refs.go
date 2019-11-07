package ipfs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

// Lists links (references) from an object.
func Refs(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data Data
	if json.Unmarshal(body, &data) == nil {
		var result Result

			refs, err := sh.Refs(data.Hash, data.Recursive)
		if err != nil {
			result.Status = "failed"
			result.ErrorMsg = fmt.Sprint(err)
			json.NewEncoder(w).Encode(result)
			return
		}

		var actual []string
		for r := range refs {
			actual = append(actual, r)
		}

		result.Status = "success"
		result.Data = actual

		json.NewEncoder(w).Encode(result)
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}
