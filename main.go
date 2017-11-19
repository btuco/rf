package main

import (
    _"encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "github.com/btuco/rest-api/rf"
    "github.com/jpillora/go433"
)

func TurnOn(w http.ResponseWriter, r *http.Request){
    fmt.Printf("Turning light on\n")
}

func TurnOff(w http.ResponseWriter, r *http.Request){
    fmt.Printf("Turning light off\n")
}


func main(){

receiving, err := go433.Receive(22, func(code uint32) {
	fmt.Printf("NEW")
})
if err != nil {
	log.Fatal(err)
}
    close(receiving)

    rf.Hello()
    router := mux.NewRouter()
    router.HandleFunc("/light/on", TurnOn).Methods("GET")
    router.HandleFunc("/light/off", TurnOff).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000",router))
}
