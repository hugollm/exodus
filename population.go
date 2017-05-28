package exodus

func NewPopulation(populationSize int, individualSize int, newGene NewGeneFunction) Population {
    population := NewEmptyPopulation(populationSize)
    for i := 0; i < populationSize; i++ {
        population.individuals[i] = NewIndividual(individualSize, newGene)
    }
    return population
}

func NewEmptyPopulation(populationSize int) Population {
    population := Population{}
    population.individuals = make([]Individual, populationSize)
    return population
}

func (p *Population) Evaluate(fitness FitnessFunction) {
    for i := 0; i < len(p.individuals); i++ {
        p.individuals[i].Evaluate(fitness)
        if p.best.fitness < p.individuals[i].fitness {
            p.best = p.individuals[i]
        }
    }
}

func (p *Population) Evolve(crossoverRate float64, mutationRate float64, newGene NewGeneFunction) {
    populationSize := len(p.individuals)
    newPopulation := NewEmptyPopulation(populationSize)
    for i := 0; i < (populationSize/2); i++ {
        parents := p.SelectParents()
        offspring := Crossover(parents, crossoverRate)
        offspring[0].Mutate(mutationRate, newGene)
        offspring[1].Mutate(mutationRate, newGene)
    }
    for i := 0; i < populationSize; i++ {
        p.individuals[i] = newPopulation.individuals[i]
    }
    p.individuals[populationSize-1] = p.best.Copy()
}
