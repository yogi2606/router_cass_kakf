package handler

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"practise/router_cass_kakf/dal/dalmock"
	"practise/router_cass_kakf/event/eventmock"
	"practise/router_cass_kakf/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func Test_PostHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var w *httptest.ResponseRecorder
	var r *mux.Router
	crmMock := dalmock.NewMockCRMInterface(ctrl)
	kafkaMock := eventmock.NewMockKafkaInterface(ctrl)
	btArr := []byte(`{"firstName": "yogesh", "lastName": "sharma", "contactDetails": {"mobileNumber": "882312312312", "emailID": "sharma.yogesh@connectwise.com"}, "personalAddr": {"street": "xyz", "district": "Pune", "state": "Maharashtra", "country": "India"}, "officialAddr": {"street": "xyz", "district": "Mumbai", "state": "Maharashtra", "country": "India"}}`)
	btArrInvalid := []byte(`{"firstName":}}`)
	tests := []struct {
		name         string
		expectedCode int
		payload      []byte
		executeFunc  func()
	}{
		{
			name:         "no_err",
			expectedCode: http.StatusOK,
			payload:      btArr,
			executeFunc: func() {
				crmMock.EXPECT().InsertCRM(gomock.Any()).Return(nil)
				kafkaMock.EXPECT().SendDataToKafka(gomock.Any())
			},
		},
		{
			name:         "err_InsertCRM",
			expectedCode: http.StatusBadRequest,
			payload:      btArr,
			executeFunc: func() {
				crmMock.EXPECT().InsertCRM(gomock.Any()).Return(errors.New("test"))
			},
		},
		{
			name:         "err_InsertCRM",
			expectedCode: http.StatusBadRequest,
			payload:      btArrInvalid,
			executeFunc: func() {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.executeFunc != nil {
				tt.executeFunc()
			}
			req, err := http.NewRequest(http.MethodPost, "/crm/customer", bytes.NewBuffer(tt.payload))

			if err != nil {
				t.Errorf("Failed to create request")
			}

			w = httptest.NewRecorder()
			r = mux.NewRouter()
			r.HandleFunc("/crm/customer", CRMHandler{
				CRM:   crmMock,
				Kafka: kafkaMock,
			}.PostHandler)
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, received %d", tt.expectedCode, w.Code)
				return
			}
		})
	}

}

func Test_DeleteHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var w *httptest.ResponseRecorder
	var r *mux.Router
	var url string
	crmMock := dalmock.NewMockCRMInterface(ctrl)
	customerID := "testcustomerID"
	tests := []struct {
		name         string
		expectedCode int
		customerID   string
		executeFunc  func()
	}{
		{
			name:         "no_err",
			expectedCode: http.StatusOK,
			customerID:   customerID,
			executeFunc: func() {
				crmMock.EXPECT().Delete(gomock.Any()).Return(nil)
			},
		},
		{
			name:         "err_Delete",
			expectedCode: http.StatusBadRequest,
			customerID:   customerID,
			executeFunc: func() {
				crmMock.EXPECT().Delete(gomock.Any()).Return(errors.New("test"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.executeFunc != nil {
				tt.executeFunc()
			}
			url = fmt.Sprintf("/crm/customer/%s", tt.customerID)
			req, err := http.NewRequest(http.MethodDelete, url, nil)

			if err != nil {
				t.Errorf("Failed to create request")
			}

			w = httptest.NewRecorder()
			r = mux.NewRouter()
			r.HandleFunc("/crm/customer/{customerID}", CRMHandler{
				CRM: crmMock,
			}.DeleteHandler)
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, received %d", tt.expectedCode, w.Code)
				return
			}
		})
	}

}

func Test_GetHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var w *httptest.ResponseRecorder
	var r *mux.Router
	var url string
	crmMock := dalmock.NewMockCRMInterface(ctrl)
	customerID := "testcustomerID"
	tests := []struct {
		name         string
		expectedCode int
		customerID   string
		executeFunc  func()
	}{
		{
			name:         "no_err",
			expectedCode: http.StatusOK,
			customerID:   customerID,
			executeFunc: func() {
				crmMock.EXPECT().Get(gomock.Any()).Return([]model.Customer{model.Customer{}, model.Customer{}})
			},
		},
		{
			name:         "err_Get",
			expectedCode: http.StatusNotFound,
			customerID:   customerID,
			executeFunc: func() {
				crmMock.EXPECT().Get(gomock.Any()).Return([]model.Customer{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.executeFunc != nil {
				tt.executeFunc()
			}
			url = fmt.Sprintf("/crm/customer/%s", tt.customerID)
			req, err := http.NewRequest(http.MethodGet, url, nil)

			if err != nil {
				t.Errorf("Failed to create request")
			}

			w = httptest.NewRecorder()
			r = mux.NewRouter()
			r.HandleFunc("/crm/customer/{customerID}", CRMHandler{
				CRM: crmMock,
			}.GetHandler)
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, received %d", tt.expectedCode, w.Code)
				return
			}
		})
	}

}

func Test_GetAllHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var w *httptest.ResponseRecorder
	var r *mux.Router
	crmMock := dalmock.NewMockCRMInterface(ctrl)
	tests := []struct {
		name         string
		expectedCode int
		executeFunc  func()
	}{
		{
			name:         "no_err",
			expectedCode: http.StatusOK,
			executeFunc: func() {
				crmMock.EXPECT().GetAll().Return([]model.Customer{model.Customer{}, model.Customer{}})
			},
		},
		{
			name:         "err_Get",
			expectedCode: http.StatusNotFound,
			executeFunc: func() {
				crmMock.EXPECT().GetAll().Return([]model.Customer{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.executeFunc != nil {
				tt.executeFunc()
			}
			req, err := http.NewRequest(http.MethodGet, "/crm/customer", nil)

			if err != nil {
				t.Errorf("Failed to create request")
			}

			w = httptest.NewRecorder()
			r = mux.NewRouter()
			r.HandleFunc("/crm/customer", CRMHandler{
				CRM: crmMock,
			}.GetAllHandler)
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, received %d", tt.expectedCode, w.Code)
				return
			}
		})
	}

}

func Test_StartConsumer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var w *httptest.ResponseRecorder
	var r *mux.Router
	kafkaMock := eventmock.NewMockKafkaInterface(ctrl)
	tests := []struct {
		name         string
		expectedCode int
		executeFunc  func()
	}{
		{
			name:         "no_err",
			expectedCode: http.StatusOK,
			executeFunc: func() {
				kafkaMock.EXPECT().ReceiveDataFromKafka()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.executeFunc != nil {
				tt.executeFunc()
			}
			req, err := http.NewRequest(http.MethodGet, "/crm/startConsumer", nil)

			if err != nil {
				t.Errorf("Failed to create request")
			}

			w = httptest.NewRecorder()
			r = mux.NewRouter()
			r.HandleFunc("/crm/startConsumer", CRMHandler{
				Kafka: kafkaMock,
			}.StartConsumer)
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, received %d", tt.expectedCode, w.Code)
				return
			}
		})
	}

}
