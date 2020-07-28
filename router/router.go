package router

import (
	"practise/router_cass_kakf/handler"

	"github.com/gorilla/mux"
)

//RegisterRoute RegisterRoute
func RegisterRoute(r *mux.Router) {
	bookrouter := r.PathPrefix("/crm").Subrouter()
	bookrouter.HandleFunc("/customer", handler.GetAllHandler).Methods("GET")
	bookrouter.HandleFunc("/customer/{customerID}", handler.GetHandler).Methods("GET")
	bookrouter.HandleFunc("/customer", handler.PostHandler).Methods("POST")
	bookrouter.HandleFunc("/customer/{customerID}", handler.DeleteHandler).Methods("DELETE")
}
