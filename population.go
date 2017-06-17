package exodus

import "math/rand"

func NewPopulation(populationSize int, individualSize int, newGene NewGeneFunction) Population {
    population := NewEmptyPopulation(populationSize)
    for i := 0; i < populationSize; i++ {
        population.Individuals[i] = NewIndividual(individualSize, newGene)
    }
    return population
}

func NewEmptyPopulation(populationSize int) Population {
    population := Population{}
    population.Individuals = make([]Individual, populationSize)
    return population
}

func (population *Population) Evaluate(fitness FitnessFunction) {
    channel := make(chan Individual)
    for i := 0; i < len(population.Individuals); i++ {
        go asyncEvaluate(population.Individuals[i], fitness, channel)
    }
    for i := 0; i < len(population.Individuals); i++ {
        population.Individuals[i] = <- channel
        if population.Best.Fitness < population.Individuals[i].Fitness {
            population.Best = population.Individuals[i]
        }
    }
}

func asyncEvaluate(individual Individual, fitness FitnessFunction, channel chan Individual) {
    individual.Evaluate(fitness)
    channel <- individual
}

func (population *Population) Migrate(migrationRate float64, imigrants *[]Individual) {
    if rand.Float64() < migrationRate {
        if InServer() {
            MigrateLocally()
        }
        if InClient() {
            AcceptImigrants(imigrants, population)
            go SendBestToServer(population.Best)
            go MigrateToServer(population.SelectIndividual(), imigrants)
        }
    }
}

func (population *Population) Evolve(crossoverRate float64, mutationRate float64, newGene NewGeneFunction) {
    populationSize := len(population.Individuals)
    newPopulation := NewEmptyPopulation(populationSize)
    for i := 0; i < (populationSize/2); i++ {
        parents := population.SelectParents()
        offspring := Crossover(parents, crossoverRate)
        offspring[0].Mutate(mutationRate, newGene)
        offspring[1].Mutate(mutationRate, newGene)
        newPopulation.Individuals[i+i] = offspring[0].Copy()
        newPopulation.Individuals[i+i+1] = offspring[1].Copy()
    }
    for i := 0; i < populationSize; i++ {
        population.Individuals[i] = newPopulation.Individuals[i]
    }
    population.Individuals[populationSize-1] = population.Best.Copy()
}
