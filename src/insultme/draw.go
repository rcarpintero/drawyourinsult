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

func random(max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max)
}

func InsultEndpoint(w http.ResponseWriter, req *http.Request) {
    insult := insults[random(len(insults) - 1)]
    json.NewEncoder(w).Encode(insult)
}

func main() {
    loadInsults()
    router := mux.NewRouter()
    router.HandleFunc("/drawinsult", InsultEndpoint).Methods("GET")
    log.Fatal(http.ListenAndServe(":9090", router))
}
