package exodus

import "math/rand"

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

func (individual *Individual) Mutate(rate float64, newGene NewGeneFunction) {
    for i := 0; i < len(individual.genome); i++ {
        if rand.Float64() < rate {
            individual.genome[i] = newGene()
        }
    }
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
