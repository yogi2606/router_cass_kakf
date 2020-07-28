package model

//Customer is a model
type Customer struct {
	FirstName      string      `json:"firstName" cql:"firstname"`
	LastName       string      `json:"lastName" cql:"lastname"`
	CustomerID     string      `json:"customerID" cql:"customerid"`
	ContactDetails ContactInfo `json:"contactDetails" cql:"contactdetails"`
	PersonalAddr   Address     `json:"personalAddr,omitempty" cql:"personaladdr"`
	OfficialAddr   Address     `json:"officialAddr,omitempty" cql:"officialaddr"`
}

//ContactInfo ContactInfo
type ContactInfo struct {
	MobileNumber string `json:"mobileNumber" cql:"mobilenumber"`
	EmailID      string `json:"emailID" cql:"emailid"`
}

//Address Address
type Address struct {
	Street   string `json:"street,omitempty" cql:"street"`
	District string `json:"district,omitempty" cql:"district"`
	State    string `json:"state,omitempty" cql:"state"`
	Country  string `json:"country,omitempty" cql:"country"`
}
