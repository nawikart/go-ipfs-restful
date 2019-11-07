package routes

import (
	c "../controllers/ipfs"
	"github.com/gorilla/mux"
)

func Ipfs(r *mux.Router) {
	r.HandleFunc("/ipfs/add", c.Add).Methods("POST")
	r.HandleFunc("/ipfs/adddir", c.AddDir).Methods("POST")
	r.HandleFunc("/ipfs/addlink", c.AddLink).Methods("POST")
	r.HandleFunc("/ipfs/cat/{hash}", c.Cat).Methods("GET")
	r.HandleFunc("/ipfs/list/{hash}", c.List).Methods("GET")
	r.HandleFunc("/ipfs/refs", c.Refs).Methods("POST")
	r.HandleFunc("/ipfs/pin/{hash}", c.Pin).Methods("GET")
	r.HandleFunc("/ipfs/unpin/{hash}", c.Unpin).Methods("GET")
	r.HandleFunc("/ipfs/pins", c.Pins).Methods("GET")
	// r.HandleFunc("/ipfs/name/publish", c.Publish).Methods("POST")
	r.HandleFunc("/ipfs/name/publish", c.PublishWithDetails).Methods("POST")
	r.HandleFunc("/ipfs/name/resolve/{id}", c.Resolve).Methods("GET")
}
