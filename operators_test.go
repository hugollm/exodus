package exodus

import "testing"
import "reflect"

func TestSelectParents(t *testing.T) {
    population := NewPopulation(1, 5, NewGeneTestFunction)
    population.Evaluate(FitnessTestFunction)
    parents := population.SelectParents()
    if parents[0].genome == nil || parents[1].genome == nil {
        t.Fail()
    }
}

func TestCrossover(t *testing.T) {
    f0 := func() int {
        return 0
    }
    f1 := func() int {
        return 1
    }
    parents := [2]Individual{NewIndividual(5, f0), NewIndividual(5, f1)}
    offspring := Crossover(parents, 0)
    if ! reflect.DeepEqual(offspring[0].genome, parents[0].genome) || ! reflect.DeepEqual(offspring[1].genome, parents[1].genome) {
        t.Fail()
    }
    offspring = Crossover(parents, 1)
    if reflect.DeepEqual(offspring[0].genome, parents[0].genome) || reflect.DeepEqual(offspring[1].genome, parents[1].genome) {
        t.Fail()
    }
}

func TestMutate(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    individual.genome = []int{8, 8, 8}
    individual.Mutate(1, NewGeneTestFunction)
    if ! reflect.DeepEqual(individual.genome, []int{9, 9, 9}) {
        t.Fail()
    }
}
