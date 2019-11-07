package ipfs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PublishData struct {
	Node  string `json:"node"`
	Value string `json:"value"`
}

// Publish updates a mutable name to point to a given value.
func publish(data PublishData) Result {
	var result Result

	err := sh.Publish(data.Node, data.Value)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	result.Status = "success"
	return result
}

func Publish(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data PublishData
	if json.Unmarshal(body, &data) == nil {
		json.NewEncoder(w).Encode(publish(data))
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}
