package exodus

import "testing"

func TestNewPopulation(t *testing.T) {
    population := NewPopulation(5, 3, NewGeneTestFunction)
    if len(population.Individuals) != 5 {
        t.Fail()
    }
    for i := 0; i < 5; i++ {
        if population.Individuals[i].Genome == nil {
            t.Fail()
        }
    }
}

func TestPopulationEvaluate(t *testing.T) {
    population := NewPopulation(5, 3, NewGeneTestFunction)
    population.Evaluate(FitnessTestFunction)
    for i := 0; i < 5; i++ {
        if population.Individuals[i].Fitness != 1.25 {
            t.Fail()
        }
    }
}

func TestPopulationEvolve(t *testing.T) {
    population := NewPopulation(5, 3, NewGeneTestFunction)
    population.Evaluate(FitnessTestFunction)
    population.Evolve(0.7, 0.01, NewGeneTestFunction)
    for i := 0; i < len(population.Individuals); i++ {
        if population.Individuals[i].Fitness != 0 {
            t.Fail()
        }
    }
}
