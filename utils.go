package exodus

import "math/rand"

func RandomInt(min int, max int) int {
    return rand.Intn(max - min + 1) + min
}

func RandomFloat64(min float64, max float64) float64 {
    return rand.Float64() * (max - min) + min
}
