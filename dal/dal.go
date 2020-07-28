package dal

import (
	"practise/router_cass_kakf/cassandra"
	"practise/router_cass_kakf/model"

	"github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
)

//InsertCRM InsertCRM
func InsertCRM(customer *model.Customer) error {
	var gocqlUUIDgo gocql.UUID // FormToUser() is included in Users/processing.go
	// we will describe this later
	// attempt to save our data to Cassandra
	var err error
	gocqlUUIDgo = gocql.TimeUUID() // write data to Cassandra
	customer.CustomerID = gocqlUUIDgo.String()
	if err = cassandra.Session.Query(`
      INSERT INTO customer (customerid,firstname,lastname,contactdetails,personaladdr,officialaddr) VALUES (?, ?, ?,?,?,?)`,
		customer.CustomerID, customer.FirstName, customer.LastName, customer.ContactDetails, customer.PersonalAddr, customer.OfficialAddr).Exec(); err != nil {
		return err
	} // depending on whether we created the user, return the
	return err

}

//GetAll GetAll
func GetAll() []model.Customer {
	var customers []model.Customer
	m := map[string]interface{}{}

	iter := cassandra.Session.Query("SELECT * FROM crm.customer").Iter()
	for iter.MapScan(m) {
		customers = append(customers, model.Customer{
			CustomerID:     m["customerid"].(string),
			FirstName:      m["firstname"].(string),
			LastName:       m["lastname"].(string),
			ContactDetails: toContactDetails(m["contactdetails"]),
			PersonalAddr:   toAddr(m["personaladdr"]),
			OfficialAddr:   toAddr(m["officialaddr"]),
		})
		m = map[string]interface{}{}
	}

	return customers

}

//Get Get
func Get(customerID string) []model.Customer {
	var customers []model.Customer
	m := map[string]interface{}{}

	iter := cassandra.Session.Query("SELECT * FROM crm.customer where customerid=?", customerID).Iter()
	for iter.MapScan(m) {
		customers = append(customers, model.Customer{
			CustomerID:     m["customerid"].(string),
			FirstName:      m["firstname"].(string),
			LastName:       m["lastname"].(string),
			ContactDetails: toContactDetails(m["contactdetails"]),
			PersonalAddr:   toAddr(m["personaladdr"]),
			OfficialAddr:   toAddr(m["officialaddr"]),
		})
		m = map[string]interface{}{}
	}

	return customers

}

//Delete Delete
func Delete(customerID string) error {
	err := cassandra.Session.Query("DELETE FROM crm.customer where customerid=?", customerID).Exec()
	return err

}

var toContactDetails = func(i interface{}) model.ContactInfo {
	cdt := model.ContactInfo{}
	val, ok := i.(map[string]interface{})
	if ok {
		mapstructure.Decode(val, &cdt)
	}
	return cdt
}
var toAddr = func(i interface{}) model.Address {
	addr := model.Address{}
	val, ok := i.(map[string]interface{})
	if ok {
		mapstructure.Decode(val, &addr)
	}
	return addr
}
