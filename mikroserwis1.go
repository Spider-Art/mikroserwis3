package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"net/http"
    "github.com/gorilla/mux"
)

var count = 0
var awaria = false

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}


func returnResponse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Odpowiedź z mikroserwisu-1")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := &http.Client{}
 	req, err := http.NewRequest("GET", "http://mikroserwis2:8080/mikroserwis2/", nil)
	if err != nil {
		fmt.Print(err.Error())
		fmt.Println("Brak serwisu 2")
 	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Brak serwisu 2")
		fmt.Fprintf(w,"Brak mikroserwisu 2 !!!")
		return
 	}
	defer resp.Body.Close()
 	bodyBytes, err := ioutil.ReadAll(resp.Body)
 		if err != nil {
			log.Fatal(err)  
		}
	fmt.Println(string(bodyBytes)) 	
	fmt.Fprintf(w, "Przekazuje z pod-a %s, odpowiedź: <<-- %s !",hostname,string(bodyBytes))	

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/mikroserwis1", returnResponse)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    handleRequests()
}