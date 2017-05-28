package exodus

type Population struct {
    individuals []Individual
    best Individual
}

type Individual struct {
    genome []int
    fitness float64
}

type NewGeneFunction func() int
type FitnessFunction func([]int) float64
