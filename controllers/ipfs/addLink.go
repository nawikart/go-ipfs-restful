package ipfs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AddLink(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data Data
	if json.Unmarshal(body, &data) == nil {
		json.NewEncoder(w).Encode(addLink(data))
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}

func addLink(data Data) Result {
	var result Result
	var resultData ResultData

	cid, err := sh.AddLink(data.Link)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
	}
	result.Status = "success"
	resultData.Hash = cid
	resultData.Data = data.Link
	result.Data = &resultData

	return result
}
