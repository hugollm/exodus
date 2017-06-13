package exodus

import "testing"
import "reflect"

func TestNewIndividual(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    if ! reflect.DeepEqual(individual.Genome, []float64{9, 9, 9}) {
        t.Fail()
    }
}

func TestIndividualEvaluate(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    individual.Evaluate(FitnessTestFunction)
    if individual.Fitness != 1.25 {
        t.Fail()
    }
}

func TestIndividualCopy(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    individual2 := individual.Copy()
    individual.Genome[2] = 8
    if ! reflect.DeepEqual(individual2.Genome, []float64{9, 9, 9}) {
        t.Fail()
    }
}

func TestIndividualCopyWithFitness(t *testing.T) {
    individual := NewIndividual(3, NewGeneTestFunction)
    individual.Fitness = 5
    individual2 := individual.CopyWithFitness()
    individual.Genome[2] = 8
    if ! reflect.DeepEqual(individual2.Genome, []float64{9, 9, 9}) || individual2.Fitness != 5 {
        t.Fail()
    }
}
