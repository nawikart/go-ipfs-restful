package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	utl "../../utils"
)

// Displays the data contained by an IPFS or IPNS object(s) at the given path.
func Cat(w http.ResponseWriter, r *http.Request) {
	defer utl.Catch()
	vars := mux.Vars(r)
	hash := vars["hash"]

	var result Result

	cid, err := sh.Cat(hash)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		json.NewEncoder(w).Encode(result)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(cid)
	newStr := buf.String()

	w.Write([]byte(newStr))
}
