package ipfs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Add(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(add(r))
}

func add(r *http.Request) interface{} {
	var results []Result
	var result Result

	type IoData struct {
		Data string
		Io   io.Reader
	}
	var ioDatas []IoData

	// if its FILE
	maxMemory := int64(32 << 20)
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	files := r.MultipartForm.File["file"]
	for _, f := range files {
		file, _ := f.Open()
		ioDatas = append(ioDatas, IoData{Data: f.Filename, Io: file})
	}

	// if its STRING
	formdata := r.MultipartForm
	if formdata.Value["arg"] != nil {
		arg := formdata.Value["arg"]
		for _, v := range arg {
			ioDatas = append(ioDatas, IoData{Data: v, Io: strings.NewReader(v)})
		}
	}

	// add to IPFS

	for _, d := range ioDatas {
		result := Result{}
		resultData := ResultData{}

		cid, err := sh.Add(d.Io)
		if err != nil {
			result.Status = "failed"
			result.ErrorMsg = fmt.Sprint(err)
			resultData.Data = d.Data
			result.Data = &resultData
		} else {
			result.Status = "success"
			resultData.Hash = cid
			resultData.Data = d.Data
			result.Data = &resultData
		}
		results = append(results, result)
	}

	return results
}
