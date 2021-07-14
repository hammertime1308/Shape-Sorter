package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hammertime1308/shape-sorter/shapes"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		http.Error(w, fmt.Sprintf("Error = %v", err.Error()), 400)
		os.Exit(2)
	}
	object := shapes.Parse(string(data))
	object.Sort()
	out := shapes.Serialize(object)
	w.Write([]byte(out))
}

func Serve(port int) {
	r := mux.NewRouter()

	fmt.Println("Starting the server on port = ", port)

	r.HandleFunc("/shape-sort", handler).Methods("GET")
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
