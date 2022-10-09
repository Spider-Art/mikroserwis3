package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
    "github.com/gorilla/mux"
)

var count = 0
var awaria = false

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAwaria(w http.ResponseWriter, r *http.Request){
	awaria = true
	fmt.Fprintf(w, "Start symulacji awarii serwisu!")
}

func returnPoAwarii(w http.ResponseWriter, r *http.Request){
	awaria = false
	fmt.Fprintf(w, "Stop symulacji awarii serwisu!")
}

func returnResponse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Odpowiedź z mikroserwisu-3 wersja-2")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if awaria == false {	
		count=count+1
		fmt.Fprintf(w, "Odpowiadam z pod-a %s wersja-2, po raz %d! \n",hostname,count)
	}else{
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Symulacja awarii serwisu %s \n",hostname)
	}
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/mikroserwis3", returnResponse)
	myRouter.HandleFunc("/awaria", returnAwaria)
	myRouter.HandleFunc("/poawarii", returnPoAwarii)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    handleRequests()
}