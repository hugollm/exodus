package exodus

import "math/rand"

func (p *Population) SelectParents() [2]Individual {
    roulette := NewRoulette(*p)
    a := roulette.Select()
    b := roulette.Select()
    return [2]Individual{a, b}
}

func Crossover(parents [2]Individual, probability float64) [2]Individual {
    if rand.Float64() > probability {
        return [2]Individual{parents[0].Copy(), parents[1].Copy()}
    }
    size := len(parents[0].genome)
    offspring := [2]Individual{NewEmptyIndividual(size), NewEmptyIndividual(size)}
    point := rand.Intn(size)
    for i := 0; i < size; i++ {
        if i < point {
            offspring[0].genome[i] = parents[0].genome[i]
            offspring[1].genome[i] = parents[1].genome[i]
        } else {
            offspring[0].genome[i] = parents[1].genome[i]
            offspring[1].genome[i] = parents[0].genome[i]
        }
    }
    return offspring
}

func (individual *Individual) Mutate(rate float64, newGene NewGeneFunction) {
    for i := 0; i < len(individual.genome); i++ {
        if rand.Float64() < rate {
            individual.genome[i] = newGene()
        }
    }
}
