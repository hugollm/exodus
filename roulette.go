package exodus

import "math/rand"

type Roulette struct {
    population Population
    slices []float64
}

func NewRoulette(population Population) Roulette {
    roulette := Roulette{}
    roulette.population = population
    roulette.slices = make([]float64, len(population.individuals))
    roulette.calculateSlices()
    return roulette
}

func (r *Roulette) calculateSlices() {
    sumFitness := r.sumFitness()
    sliceOffset := 0.0
    for i := 0; i < len(r.population.individuals); i++ {
        r.slices[i] = sliceOffset + (r.population.individuals[i].fitness / sumFitness)
        sliceOffset = r.slices[i]
    }
}

func (r *Roulette) sumFitness() float64 {
    sum := 0.0
    for i := 0; i < len(r.population.individuals); i++ {
        sum += r.population.individuals[i].fitness
    }
    return sum
}

func (r *Roulette) Select() Individual {
    ball := rand.Float64()
    for i := 0; i < len(r.population.individuals); i++ {
        if ball < r.slices[i] {
            return r.population.individuals[i]
        }
    }
    return r.population.individuals[len(r.population.individuals)-1]
}
