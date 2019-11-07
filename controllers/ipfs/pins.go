package ipfs

import (
	"encoding/json"
	"fmt"
	"net/http"

)

// Pins returns a map of the pin hashes to their info (currently just the
// pin type, one of DirectPin, RecursivePin, or IndirectPin. A map is returned
// instead of a slice because it is easier to do existence lookup by map key
// than unordered array searching. The map is likely to be more useful to a
// client than a flat list.
func Pins(w http.ResponseWriter, r *http.Request) {
	var result Result

	data, err := sh.Pins()
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
