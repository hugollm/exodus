package exodus

import "math/rand"

func (population *Population) SelectParents() [2]Individual {
    roulette := NewRoulette(*population)
    a := roulette.Select()
    b := roulette.Select()
    return [2]Individual{a, b}
}

func (population *Population) SelectIndividual() Individual {
    roulette := NewRoulette(*population)
    return roulette.Select()
}

func Crossover(parents [2]Individual, probability float64) [2]Individual {
    if rand.Float64() > probability {
        return [2]Individual{parents[0].Copy(), parents[1].Copy()}
    }
    size := len(parents[0].Genome)
    offspring := [2]Individual{NewEmptyIndividual(size), NewEmptyIndividual(size)}
    point := rand.Intn(size)
    for i := 0; i < size; i++ {
        if i < point {
            offspring[0].Genome[i] = parents[0].Genome[i]
            offspring[1].Genome[i] = parents[1].Genome[i]
        } else {
            offspring[0].Genome[i] = parents[1].Genome[i]
            offspring[1].Genome[i] = parents[0].Genome[i]
        }
    }
    return offspring
}

func (individual *Individual) Mutate(rate float64, newGene NewGeneFunction) {
    for i := 0; i < len(individual.Genome); i++ {
        if rand.Float64() < rate {
            individual.Genome[i] = newGene()
        }
    }
}
