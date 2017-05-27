package exodus

import "math/rand"

func RandomInt(min int, max int) int {
    return (rand.Int() % max) + min
}
