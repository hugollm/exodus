package exodus

import "math/rand"
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
}

func (search *Search) Start() {
    rand.Seed(time.Now().UTC().UnixNano())
    search.Population = NewPopulation(search.PopulationSize, search.IndividualSize, search.NewGene)
    if InServer() {
        go RunServer(search)
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
