package ipfs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AddDir(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data Data
	if json.Unmarshal(body, &data) == nil {
		json.NewEncoder(w).Encode(addDir(data))
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}

func addDir(data Data) Result {
	var result Result
	var resultData ResultData

	cid, err := sh.AddDir(data.Dir)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
	}
	result.Status = "success"
	resultData.Hash = cid
	resultData.Data = data.Dir
	result.Data = &resultData

	return result
}
