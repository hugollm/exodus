package exodus

type Individual struct {
    genome []int
    fitness float64
}

func NewIndividual(size int, newGene NewGeneFunction) Individual {
    individual := Individual{}
    individual.genome = make([]int, size)
    for i := 0; i < size; i++ {
        individual.genome[i] = newGene()
    }
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
        newIndividual.genome[i] = individual[i]
    }
    return newIndividual
}
