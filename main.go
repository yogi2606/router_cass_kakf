package main

import (
	"fmt"
	"net/http"
	"practise/router_cass_kakf/cassandra"
	"practise/router_cass_kakf/router"

	"github.com/gorilla/mux"
)

func main() {
	// mapDemo()
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()
	r := mux.NewRouter()
	router.RegisterRoute(r)

	err := http.ListenAndServe(":9090", r)
	if err != nil {
		fmt.Println(err)
	}
}

/* type temp struct {
	EmailID string `json:"emailid"`
}

func mapDemo() {
	mapt := map[string]interface{}{}
	mapt["emailid"] = "10"
	var a temp
	mapstructure.Decode(mapt, &a)
	fmt.Println(a)
} */
