package ipfs

import (
	utl "../../utils"
	shell "github.com/ipfs/go-ipfs-api"
)

type Data struct {
	Dir       string `json:"directory,omitempty"`
	Link      string `json:"link,omitempty"`
	Path      string `json:"path,omitempty"`
	Hash      string `json:"hash,omitempty"`
	Recursive bool   `json:"recursive,omitempty"`
	Peer      bool   `json:"peer,omitempty"`
}

type ResultData struct {
	Data  string      `json:"data,omitempty"`
	Hash  string      `json:"hash,omitempty"`
	Size  int         `json:"size,omitempty"`
	Other interface{} `json:"other,omitempty"`
}

type Result struct {
	Status   string      `json:"status"`
	ErrorMsg string      `json:"error_msg,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

var sh = shell.NewShell(utl.GetEnv("IPFS_URL", "localhost:5001"))
var IHUB_URL = utl.GetEnv("IHUB_URL", "localhost:8778")
