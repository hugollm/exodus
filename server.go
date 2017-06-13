package exodus

import "encoding/json"
import "io/ioutil"
import "math/rand"
import "net/http"
import "os"

var globalSearch *Search

func InServer() bool {
    return os.Getenv("EXODUS_SERVER") == ""
}

func RunServer(search *Search) {
    globalSearch = search
    http.HandleFunc("/test-connection", testConnection)
    http.HandleFunc("/send-best", sendBest)
    http.HandleFunc("/migrate", migrate)
    err := http.ListenAndServe(":2345", nil)
    if err != nil {
        panic(err)
    }
}

func testConnection(response http.ResponseWriter, request *http.Request) {
    json.NewEncoder(response).Encode(true)
}

func sendBest(response http.ResponseWriter, request *http.Request) {
    individual := Individual{}
    body, _ := ioutil.ReadAll(request.Body)
    json.Unmarshal(body, &individual)
    if globalSearch.Population.Best.Fitness < individual.Fitness {
        globalSearch.Population.Best = individual
        globalSearch.Best = individual
    }
}

func migrate(response http.ResponseWriter, request *http.Request) {
    emigrant := Individual{}
    if len(globalSearch.Imigrants) > 0 {
        i := rand.Intn(len(globalSearch.Imigrants))
        emigrant = globalSearch.Imigrants[i]
        globalSearch.Imigrants = append(globalSearch.Imigrants[:i], globalSearch.Imigrants[i+1:]...)
    }
    json.NewEncoder(response).Encode(emigrant)
    imigrant := Individual{}
    body, _ := ioutil.ReadAll(request.Body)
    json.Unmarshal(body, &imigrant)
    globalSearch.Imigrants = append(globalSearch.Imigrants, imigrant)
}

func MigrateLocally() {
    if len(globalSearch.Imigrants) > 0 {
        imigrantsIndex := rand.Intn(len(globalSearch.Imigrants))
        emigrant := globalSearch.Imigrants[imigrantsIndex]
        globalSearch.Imigrants = append(globalSearch.Imigrants[:imigrantsIndex], globalSearch.Imigrants[imigrantsIndex+1:]...)
        populationIndex := rand.Intn(len(globalSearch.Population.Individuals))
        globalSearch.Population.Individuals[populationIndex] = emigrant
    }
    imigrant := globalSearch.Population.SelectIndividual()
    globalSearch.Imigrants = append(globalSearch.Imigrants, imigrant.CopyWithFitness())
}
