package exodus

import "math/rand"
import "sync"
import "time"

type Search struct {

    IndividualSize int
    PopulationSize int
    CrossoverRate float64
    MutationRate float64
    MigrationRate float64
    NewGene NewGeneFunction
    Fitness FitnessFunction
    OnGeneration OnGenerationFunction

    Generation int
    Population Population
    Best Individual
    Imigrants []Individual

    stop bool
    mutex sync.Mutex
}

var globalSearch *Search

func (search *Search) Start() {
    globalSearch = search
    rand.Seed(time.Now().UTC().UnixNano())
    search.Population = NewPopulation(search.PopulationSize, search.IndividualSize, search.NewGene)
    if InServer() {
        go RunServer()
    } else {
        TestConnectionWithServer()
    }
    for {
        search.Generation++
        search.Population.Evaluate(search.Fitness)
        search.Population.Migrate(search.MigrationRate, &search.Imigrants)
        search.Best = search.Population.Best
        search.OnGeneration(search)
        if search.stop {
            break
        }
        search.Population.Evolve(search.CrossoverRate, search.MutationRate, search.NewGene)
    }
}

func (search *Search) Stop() {
    search.stop = true
}
