package ipfs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PublishWithDetailsData struct {
	ContentHash string        `json:"contenthash"`
	Key         string        `json:"key"`
	Lifetime    time.Duration `json:"lifetime"`
	Ttl         time.Duration `json:"ttl"`
	Resolve     bool          `json:"resolve"`
}

// Publish IPNS names.
func publishWithDetails(data PublishWithDetailsData) Result {
	var result Result

	res, err := sh.PublishWithDetails(data.ContentHash, data.Key, data.Lifetime, data.Ttl, data.Resolve)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	result.Status = "success"
	result.Data = &res
	return result
}

func PublishWithDetails(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data PublishWithDetailsData
	if json.Unmarshal(body, &data) == nil {
		json.NewEncoder(w).Encode(publishWithDetails(data))
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}
