package exodus

import "testing"

func TestSelectWithOneInvividual(t *testing.T) {
    population := NewPopulation(1, 5, NewGeneTestFunction)
    population.Evaluate(FitnessTestFunction)
    roulette := NewRoulette(population)
    individual := roulette.Select()
    if individual.genome == nil {
        t.Fail()
    }
}

func TestSelectWithTwoIndividuals(t *testing.T) {
    f0 := func(genome []int) float64 {
        return 0.0
    }
    f1 := func(genome []int) float64 {
        return 1.0
    }
    population := NewPopulation(2, 5, NewGeneTestFunction)
    population.individuals[0].Evaluate(f0)
    population.individuals[1].Evaluate(f1)
    roulette := NewRoulette(population)
    individual := roulette.Select()
    if individual.fitness != population.individuals[1].fitness {
        t.Fail()
    }
}
