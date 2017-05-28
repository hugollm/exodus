package exodus

import "math/rand"
import "time"

type Search struct {

    individualSize int
    populationSize int
    crossoverRate float64
    mutationRate float64
    newGene NewGeneFunction
    fitness FitnessFunction
    onGeneration OnGenerationFunction

    generation int
    population Population
    best Individual
    stop bool
}

func (s *Search) Start() {
    rand.Seed(time.Now().UTC().UnixNano())
    s.population = NewPopulation(s.populationSize, s.individualSize, s.newGene)
    for {
        s.generation++
        s.population.Evaluate(fitness)
        s.best = s.population.best
        s.onGeneration(s)
        if s.stop {
            break
        }
        s.population.Evolve(s.crossoverRate, s.mutationRate, s.newGene)
    }
}

func (s *Search) Stop() {
    s.stop = true
}
