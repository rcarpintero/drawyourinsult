package main

import (
    "bufio"
    "math/rand"
    "time"
    "os"
    "encoding/json"
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)

var insults []string

func loadInsults() {
    file, _ := os.Open("insult_list")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        insults = append(insults, scanner.Text())
    }
}

func aRandomInsult() string {
    position := rand.Intn(len(insults) - 1)
    return insults[position]
}

func InsultEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(aRandomInsult())
}

func main() {
    rand.Seed(time.Now().Unix())
    loadInsults()
    router := mux.NewRouter()
    router.HandleFunc("/drawinsult", InsultEndpoint).Methods("GET")
    log.Fatal(http.ListenAndServe(":9090", router))
}
