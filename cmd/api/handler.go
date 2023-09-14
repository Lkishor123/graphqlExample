package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
func HomeHandler() handles request on "/" to verify if server is up and running.
*/
func (serve *GQServer) HomeHandler(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "GQ Test Server Up and Running",
		Version: "1.0.0",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

/*
func GQHandler() handles request on "/graphql", it will be used to handle all graphql queries.
*/
func (s *GQServer) GQHandler(w http.ResponseWriter, r *http.Request) {
	// Mock DB
	var dbinfo = DBinfo{
		ID:   1,
		Name: "Max",
		Complex: ComplexType{
			C1param: "C1",
			C2param: "C2",
		},
	}

	var dbinfoList []DBinfo
	dbinfoList = append(dbinfoList, dbinfo)

	// get query from Request:
	q, _ := io.ReadAll(r.Body)
	query := string(q)

	// create a new variable of type *Graph
	g := New(&dbinfoList)

	// set the query string on New variable
	g.QueryString = query

	// perform the query
	resp, err := g.Query()

	// handle error
	if err != nil {
		var payload = struct {
			Error   bool        `json:"error"`
			Message string      `json:"message"`
			Data    interface{} `json:"data,omitempty"`
		}{
			Error:   true,
			Message: err.Error(),
		}
		out, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("error while Marshal")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(out)
		return
	}
	// send response
	j, _ := json.MarshalIndent(resp, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
