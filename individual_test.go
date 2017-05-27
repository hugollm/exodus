package exodus

import "testing"
import "reflect"

func TestNewIndividual(t *testing.T) {
    individual := NewIndividual(3, newGene)
    if ! reflect.DeepEqual(individual.genome, []int{9, 9, 9}) {
        t.Fail()
    }
}

func TestEvaluate(t *testing.T) {
    individual := NewIndividual(3, newGene)
    individual.Evaluate(fitness)
    if individual.fitness != 1.25 {
        t.Fail()
    }
}

func TestCopy(t *testing.T) {
    individual := NewIndividual(3, newGene)
    individual2 := individual.Copy()
    individual.genome[2] = 8
    if ! reflect.DeepEqual(individual2.genome, []int{9, 9, 9}) {
        t.Fail()
    }
}

func newGene() int {
    return 9
}

func fitness(genome []int) float64 {
    return 1.25
}
