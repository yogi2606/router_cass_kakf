package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practise/router_cass_kakf/dal"
	"practise/router_cass_kakf/event"
	"practise/router_cass_kakf/model"

	"github.com/gorilla/mux"
)

// GetAllHandler is get handler
func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	customers := dal.GetAll()
	event.ReceiveDataFromKafka()
	json.NewEncoder(w).Encode(customers)
}

//PostHandler PostHandler
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		fmt.Println("PostHandler decode err: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// customer.CustomerID = strconv.Itoa(rand.Intn(1000000))
	err = dal.InsertCRM(&customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("user not created")
		return

	}
	bArr, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("failed to marshall", err)
	}
	event.SendDataToKafka(bArr)
	json.NewEncoder(w).Encode("user created")

}

// GetHandler is get handler
func GetHandler(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["customerID"]
	customers := dal.Get(customerID)
	json.NewEncoder(w).Encode(customers)
}

// DeleteHandler is get handler
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["customerID"]
	err := dal.Delete(customerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error in deletion")
		return
	}
	json.NewEncoder(w).Encode("record deleted")
}
