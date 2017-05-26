package exodus

import "math/rand"

type Search struct {

    individualSize int
    populationSize int
    crossoverRate float64
    mutationRate float64
    newGene NewGeneFunction
    fitness FitnessFunction
    onGeneration OnGenerationFunction

    generation int
    best Individual
    stop bool
}

type NewGeneFunction func() int
type FitnessFunction func([]int) float64
type OnGenerationFunction func(*Search)

func (s *Search) Start() {
    rand.Seed(time.Now().UTC().UnixNano())
    population := NewPopulation(s.populationSize, s.individualSize, s.newGene)
    for {
        s.generation++
        population.Evaluate(fitness)
        s.best = population.best
        s.onGeneration(&s)
        if s.stop {
            break
        }
        population.Evolve()
    }
}

func (s *Search) Stop() {
    s.stop = true
}
