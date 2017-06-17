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
    http.HandleFunc("/test-connection", TestConnectionEndpoint)
    http.HandleFunc("/send-best", SendBestEndpoint)
    http.HandleFunc("/migrate", MigrateEndpoint)
    err := http.ListenAndServe(":2345", nil)
    if err != nil {
        panic(err)
    }
}

func TestConnectionEndpoint(response http.ResponseWriter, request *http.Request) {
    json.NewEncoder(response).Encode(true)
}

func SendBestEndpoint(response http.ResponseWriter, request *http.Request) {
    individual := Individual{}
    body, _ := ioutil.ReadAll(request.Body)
    json.Unmarshal(body, &individual)
    receiveBest(individual)
}

func receiveBest(individual Individual) {
    if globalSearch.Population.Best.Fitness < individual.Fitness {
        globalSearch.Population.Best = individual
        globalSearch.Best = individual
    }
}

func MigrateEndpoint(response http.ResponseWriter, request *http.Request) {
    emigrant := popImigrantFromSearch()
    json.NewEncoder(response).Encode(emigrant)
    imigrant := Individual{}
    body, _ := ioutil.ReadAll(request.Body)
    json.Unmarshal(body, &imigrant)
    globalSearch.Imigrants = append(globalSearch.Imigrants, imigrant)
}

func MigrateLocally() {
    if len(globalSearch.Imigrants) > 0 {
        imigrant := popImigrantFromSearch()
        i := rand.Intn(len(globalSearch.Population.Individuals))
        globalSearch.Population.Individuals[i] = imigrant
    }
    emigrant := globalSearch.Population.SelectIndividual()
    globalSearch.Imigrants = append(globalSearch.Imigrants, emigrant.CopyWithFitness())
}

func popImigrantFromSearch() Individual {
    imigrant := Individual{}
    if len(globalSearch.Imigrants) > 0 {
        i := rand.Intn(len(globalSearch.Imigrants))
        imigrant = globalSearch.Imigrants[i]
        globalSearch.Imigrants = append(globalSearch.Imigrants[:i], globalSearch.Imigrants[i+1:]...)
    }
    return imigrant
}
