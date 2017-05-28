package exodus

func NewIndividual(size int, newGene NewGeneFunction) Individual {
    individual := NewEmptyIndividual(size)
    for i := 0; i < size; i++ {
        individual.genome[i] = newGene()
    }
    return individual
}

func NewEmptyIndividual(size int) Individual {
    individual := Individual{}
    individual.genome = make([]int, size)
    return individual
}

func (individual *Individual) Evaluate(fitness FitnessFunction) {
    individual.fitness = fitness(individual.genome)
}

func (individual *Individual) Copy() Individual {
    size := len(individual.genome)
    newIndividual := Individual{}
    newIndividual.genome = make([]int, size)
    for i := 0; i < size; i++ {
        newIndividual.genome[i] = individual.genome[i]
    }
    return newIndividual
}
