package exodus

import "math/rand"

func RandInt(min int, max int) {
    return (rand.Int() % max) + min
}
