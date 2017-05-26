package exodus

type Population struct {
    individuals []Individual
    best Individual
}

func NewPopulation(populationSize int, individualSize int, newGene NewGeneFunction) Population {
    population := Population{}
    population.individuals = make([]Individual, populationSize)
    for i := 0; i < populationSize; i++ {
        population.individuals[i] = NewIndividual(individualSize, newGene)
    }
}

func (p *Population) Evaluate(fitness FitnessFunction) {
    for i := 0; i < len(p.individuals); i++ {
        p.individuals[i].Evaluate(fitness)
        if p.best == nil || p.best.fitness < p.individuals[i].fitness {
            p.best = p.individuals[i].fitness
        }
    }
}

func (p *Population) Evolve(crossoverRate float64, mutationRate float64) {
    newPopulation := Population{}
    newPopulation.individuals = make([]Individual, populationSize)
    for i := 0; i < (populationSize/2); i++ {
        parents := p.Select()
        offspring := Crossover(parents, crossoverRate)
        offspring[0].Mutate(mutationRate)
        offspring[1].Mutate(mutationRate)
    }
}
