package exodus

import "math/rand"

type Roulette struct {
    population Population
    slices []float64
}

func NewRoulette(population Population) Roulette {
    roulette := Roulette{}
    roulette.population = population
    roulette.slices = make([]float64, len(population.Individuals))
    roulette.calculateSlices()
    return roulette
}

func (roulette *Roulette) calculateSlices() {
    sumFitness := roulette.sumFitness()
    sliceOffset := 0.0
    for i := 0; i < len(roulette.population.Individuals); i++ {
        roulette.slices[i] = sliceOffset + (roulette.population.Individuals[i].Fitness / sumFitness)
        sliceOffset = roulette.slices[i]
    }
}

func (roulette *Roulette) sumFitness() float64 {
    sum := 0.0
    for i := 0; i < len(roulette.population.Individuals); i++ {
        sum += roulette.population.Individuals[i].Fitness
    }
    return sum
}

func (roulette *Roulette) Select() Individual {
    ball := rand.Float64()
    for i := 0; i < len(roulette.population.Individuals); i++ {
        if ball < roulette.slices[i] {
            return roulette.population.Individuals[i]
        }
    }
    return roulette.population.Individuals[len(roulette.population.Individuals)-1]
}
