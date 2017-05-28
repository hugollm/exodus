package exodus

import "testing"

func TestNewPopulation(t *testing.T) {
    population := NewPopulation(5, 3, NewGeneTestFunction)
    if len(population.individuals) != 5 {
        t.Fail()
    }
    for i := 0; i < 5; i++ {
        if population.individuals[i].genome == nil {
            t.Fail()
        }
    }
}

func TestPopulationEvaluate(t *testing.T) {
    population := NewPopulation(5, 3, NewGeneTestFunction)
    population.Evaluate(FitnessTestFunction)
    for i := 0; i < 5; i++ {
        if population.individuals[i].fitness != 1.25 {
            t.Fail()
        }
    }
}

func TestPopulationEvolve(t *testing.T) {
    population := NewPopulation(5, 3, NewGeneTestFunction)
    population.Evaluate(FitnessTestFunction)
    population.Evolve(0.7, 0.01, NewGeneTestFunction)
    for i := 0; i < len(population.individuals); i++ {
        if population.individuals[i].fitness != 0 {
            t.Fail()
        }
    }
}
