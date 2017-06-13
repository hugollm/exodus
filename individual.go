package exodus

func NewIndividual(size int, newGene NewGeneFunction) Individual {
    individual := NewEmptyIndividual(size)
    for i := 0; i < size; i++ {
        individual.Genome[i] = newGene()
    }
    return individual
}

func NewEmptyIndividual(size int) Individual {
    individual := Individual{}
    individual.Genome = make([]float64, size)
    return individual
}

func (individual *Individual) Evaluate(fitness FitnessFunction) {
    individual.Fitness = fitness(individual.Genome)
}

func (individual *Individual) Copy() Individual {
    size := len(individual.Genome)
    newIndividual := Individual{}
    newIndividual.Genome = make([]float64, size)
    for i := 0; i < size; i++ {
        newIndividual.Genome[i] = individual.Genome[i]
    }
    return newIndividual
}

func (individual *Individual) CopyWithFitness() Individual {
    newIndividual := individual.Copy()
    newIndividual.Fitness = individual.Fitness
    return newIndividual
}
