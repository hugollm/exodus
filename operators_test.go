package exodus

import "testing"
import "reflect"

func TestSelectParents(t *testing.T) {
    population := NewPopulation(1, 5, NewGeneTestFunction)
    population.Evaluate(FitnessTestFunction)
    parents := population.SelectParents()
    if parents[0].Genome == nil || parents[1].Genome == nil {
        t.Fail()
    }
}

func TestSelectIndividual(t *testing.T) {
    population := NewPopulation(1, 5, NewGeneTestFunction)
    population.Evaluate(FitnessTestFunction)
    individual := population.SelectIndividual()
    if individual.Genome == nil {
        t.Fail()
    }
}

func TestCrossover(t *testing.T) {
    f0 := func() float64 {
        return 0
    }
    f1 := func() float64 {
        return 1
    }
    parents := [2]Individual{NewIndividual(5, f0), NewIndividual(5, f1)}
    offspring := Crossover(parents, 0)
    if ! reflect.DeepEqual(offspring[0].Genome, parents[0].Genome) || ! reflect.DeepEqual(offspring[1].Genome, parents[1].Genome) {
        t.Fail()
    }
    offspring = Crossover(parents, 1)
    if reflect.DeepEqual(offspring[0].Genome, parents[0].Genome) || reflect.DeepEqual(offspring[1].Genome, parents[1].Genome) {
        t.Fail()
    }
}

func TestMutate(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    individual.Genome = []float64{8, 8, 8}
    individual.Mutate(1, NewGeneTestFunction)
    if ! reflect.DeepEqual(individual.Genome, []float64{9, 9, 9}) {
        t.Fail()
    }
}
