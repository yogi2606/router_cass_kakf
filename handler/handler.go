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

//CRMHandler CRMHandler
type CRMHandler struct {
	CRM   dal.CRMInterface
	Kafka event.KafkaInterface
}

// GetAllHandler is get handler
func (c CRMHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	customers := c.CRM.GetAll()
	if len(customers) < 1 {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(customers)
}

//PostHandler PostHandler
func (c CRMHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		fmt.Println("PostHandler decode err: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// customer.CustomerID = strconv.Itoa(rand.Intn(1000000))
	err = c.CRM.InsertCRM(&customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("user not created")
		return

	}
	bArr, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("failed to marshall", err)
	}
	c.Kafka.SendDataToKafka(bArr)
	json.NewEncoder(w).Encode("user created")

}

// GetHandler is get handler
func (c CRMHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["customerID"]
	customers := c.CRM.Get(customerID)
	if len(customers) < 1 {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(customers)
}

// DeleteHandler is get handler
func (c CRMHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["customerID"]
	err := c.CRM.Delete(customerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error in deletion")
		return
	}
	json.NewEncoder(w).Encode("record deleted")
}

// StartConsumer is get handler
func (c CRMHandler) StartConsumer(w http.ResponseWriter, r *http.Request) {
	c.Kafka.ReceiveDataFromKafka()
}
