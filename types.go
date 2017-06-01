package exodus

type Population struct {
    Individuals []Individual
    Best Individual
}

type Individual struct {
    Genome []float64
    Fitness float64
}

type NewGeneFunction func() float64
type FitnessFunction func([]float64) float64
type OnGenerationFunction func(*Search)
