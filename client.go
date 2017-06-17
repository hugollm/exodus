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
    globalSearch.mutex.Lock()
    *imigrants = append(*imigrants, imigrant)
    globalSearch.mutex.Unlock()
}

func AcceptImigrants(imigrants *[]Individual, population *Population) {
    globalSearch.mutex.Lock()
    for len(*imigrants) > 0 {
        i := rand.Intn(len(population.Individuals))
        population.Individuals[i] = (*imigrants)[0]
        *imigrants = append((*imigrants)[:0], (*imigrants)[1:]...)
    }
    globalSearch.mutex.Unlock()
}

func serverUrl(path string) string {
    url := os.Getenv("EXODUS_SERVER")
    return url + path
}
