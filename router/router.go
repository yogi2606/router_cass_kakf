package router

import (
	"practise/router_cass_kakf/dal"
	"practise/router_cass_kakf/event"
	"practise/router_cass_kakf/handler"

	"github.com/gorilla/mux"
)

//RegisterRoute RegisterRoute
func RegisterRoute(r *mux.Router) {
	bookrouter := r.PathPrefix("/crm").Subrouter()
	handler := handler.CRMHandler{
		CRM:   dal.New(),
		Kafka: event.New(),
	}
	bookrouter.HandleFunc("/customer", handler.GetAllHandler).Methods("GET")
	bookrouter.HandleFunc("/customer/{customerID}", handler.GetHandler).Methods("GET")
	bookrouter.HandleFunc("/customer", handler.PostHandler).Methods("POST")
	bookrouter.HandleFunc("/customer/{customerID}", handler.DeleteHandler).Methods("DELETE")
	bookrouter.HandleFunc("/startConsumer", handler.StartConsumer).Methods("GET")
}
