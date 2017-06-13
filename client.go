package exodus

import "bytes"
import "encoding/json"
import "io/ioutil"
import "math/rand"
import "net/http"
import "os"

func InClient() bool {
    return os.Getenv("EXODUS_SERVER") != ""
}

func TestConnectionWithServer() {
    _, e := http.Get(serverUrl("/test-connection"))
    if e != nil {
        panic(e)
    }
}

func SendBestToServer(individual Individual) {
    jsonStr, _ := json.Marshal(individual)
    request, _ := http.NewRequest("POST", serverUrl("/send-best"), bytes.NewBuffer(jsonStr))
    request.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    response, _ := client.Do(request)
    response.Body.Close()
}

func MigrateToServer(emigrant Individual, imigrants *[]Individual) {
    jsonStr, _ := json.Marshal(emigrant)
    request, _ := http.NewRequest("POST", serverUrl("/migrate"), bytes.NewBuffer(jsonStr))
    request.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    response, e := client.Do(request)
    if e != nil {
        return
    }
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)
    imigrant := Individual{}
    json.Unmarshal(body, &imigrant)
    if imigrant.Genome == nil {
        return
    }
    *imigrants = append(*imigrants, imigrant)
}

func AcceptImigrant(imigrants *[]Individual, population *Population) {
    if len(*imigrants) > 0 {
        imigrantsIndex := rand.Intn(len(*imigrants))
        populationIndex := rand.Intn(len(population.Individuals))
        population.Individuals[populationIndex] = (*imigrants)[imigrantsIndex]
        *imigrants = append((*imigrants)[:imigrantsIndex], (*imigrants)[imigrantsIndex+1:]...)
    }
}

func serverUrl(path string) string {
    url := os.Getenv("EXODUS_SERVER")
    return url + path
}
