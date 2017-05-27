package exodus

import "testing"

func TestRandomInt(t *testing.T) {
    for i := 0; i < 1000; i++ {
        n := RandomInt(0, 2)
        if n < 0 || n > 2 {
            t.Fail()
        }
    }
    for i := 0; i < 1000; i++ {
        n := RandomInt(1, 3)
        if n < 1 || n > 3 {
            t.Fail()
        }
    }
}