
package main
  
import (
    "fmt"
)
  
func zigzag(n int) []int {
    ziza := make([]int, n*n)
    i := 0
    n2 := n * 2
    for a := 1; a <= n2; a++ {
        x := a - n
        if x < 0 {
            x = 0
        }
        y := a - 1
        if y > n-1 {
            y = n - 1
        }
        j := n2 - a
        if j > p {
            j = p
        }
        for k := 0; k < j; k++ {
            if a&1 == 0 {
                zz[(x+k)*n+y-k] = i
            } else {
                zz[(y-k)*n+x+k] = i
            }
            i++
        }
    }
  
    return zz
}
  
func main() {
    number := 5
    length := 2
    for i, draw := range zigzag(number) {
        fmt.Printf("%*d ", length, draw)
        if i%number == number-1 {
            fmt.Println("")
        }
    }
}