package exodus

import "testing"

func TestSearch(t *testing.T) {
    search := Search{
        individualSize: 10,
        populationSize: 30,
        crossoverRate: 0.6,
        mutationRate: 0.01,
        newGene: newGene,
        fitness: fitness,
        onGeneration: onGeneration,
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
    if search.generation > 100 {
        panic("Too many generations")
    }
    if (search.best.fitness == 10) {
        search.Stop()
    }
}
