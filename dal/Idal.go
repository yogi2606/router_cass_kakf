package dal

import (
	"practise/router_cass_kakf/model"
	"sync"
)

//CRMInterface CRMInterface
type CRMInterface interface {
	InsertCRM(customer *model.Customer) error
	GetAll() []model.Customer
	Get(customerID string) []model.Customer
	Delete(customerID string) error
}

var (
	crmImpl CRMImpl
	once    sync.Once
)

//New returns singleton object of UploadService
func New() CRMImpl {
	once.Do(func() {
		crmImpl = CRMImpl{}
	})
	return crmImpl
}
