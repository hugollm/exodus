package exodus

type Population struct {
    Individuals []Individual
    Best Individual
}

type Individual struct {
    Genome []int
    Fitness float64
}

type NewGeneFunction func() int
type FitnessFunction func([]int) float64
type OnGenerationFunction func(*Search)
