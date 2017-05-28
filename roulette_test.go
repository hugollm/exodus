package exodus

import "testing"

func TestSelectWithOneInvividual(t *testing.T) {
    population := NewPopulation(1, 5, NewGeneTestFunction)
    population.Evaluate(FitnessTestFunction)
    roulette := NewRoulette(population)
    individual := roulette.Select()
    if individual.Genome == nil {
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
    population.Individuals[0].Evaluate(f0)
    population.Individuals[1].Evaluate(f1)
    roulette := NewRoulette(population)
    individual := roulette.Select()
    if individual.Fitness != population.Individuals[1].Fitness {
        t.Fail()
    }
}
