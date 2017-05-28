package exodus

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
    for i := 0; i < len(population.Individuals); i++ {
        population.Individuals[i].Evaluate(fitness)
        if population.Best.Fitness < population.Individuals[i].Fitness {
            population.Best = population.Individuals[i]
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
