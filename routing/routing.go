package routing

import (
	"github.com/gorilla/mux"

	"net/http"
)

// SetupRouting manages all the routes
func SetupRouting(r *mux.Router) *mux.Router {
	r.HandleFunc("/startRaft", StartRaft).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/requestVotes", RequestVotes).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/appendEntries", AppendEntries).Methods(http.MethodPost, http.MethodOptions)
	return r
}
