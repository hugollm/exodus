package exodus

import "testing"
import "reflect"

func TestNewIndividual(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    if ! reflect.DeepEqual(individual.genome, []int{9, 9, 9}) {
        t.Fail()
    }
}

func TestIndividualEvaluate(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    individual.Evaluate(FitnessTestFunction)
    if individual.fitness != 1.25 {
        t.Fail()
    }
}

func TestIndividualMutate(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    individual.genome = []int{8, 8, 8}
    individual.Mutate(1, NewGeneTestFunction)
    if ! reflect.DeepEqual(individual.genome, []int{9, 9, 9}) {
        t.Fail()
    }
}

func TestIndividualCopy(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    individual2 := individual.Copy()
    individual.genome[2] = 8
    if ! reflect.DeepEqual(individual2.genome, []int{9, 9, 9}) {
        t.Fail()
    }
}
