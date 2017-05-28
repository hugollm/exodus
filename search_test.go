package exodus

import "testing"

func TestSearch(t *testing.T) {
    search := Search{
        IndividualSize: 10,
        PopulationSize: 30,
        CrossoverRate: 0.6,
        MutationRate: 0.01,
        NewGene: newGene,
        Fitness: fitness,
        OnGeneration: onGeneration,
    }
    search.Start()
}

func newGene() int {
    return RandomInt(0, 1)
}

func fitness(genome []int) float64 {
    fit := 0
    for i := 0; i < len(genome); i++ {
        fit += genome[i]
    }
    return float64(fit)
}

func onGeneration(search *Search) {
    if search.Generation > 100 {
        panic("Too many generations")
    }
    if (search.Best.Fitness == 10) {
        search.Stop()
    }
}
