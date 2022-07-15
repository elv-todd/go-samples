package main

import (
  "fmt"
  "time"
  "math"
  //"rsc.io/quote"
)

// swap : sample arg examples
func swap(x string, y string) (a string, b string) {
  a = y
  b = x
  return
}

type Vertex struct {
  Lat, Lon float64
}

func (x Vertex) dist(y Vertex) float64 {
  return math.Sqrt( (x.Lat - y.Lat) * (x.Lat - y.Lat) + (x.Lon - y.Lon) * (x.Lon - y.Lon) )
}

func describe(i interface{}) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
  fmt.Println("time: ", time.Now())
  fmt.Println(swap("a", "b"))

  defer fmt.Println("end10")
  for i := 0 ; i < 10; i++ {
    defer fmt.Print(i)
  }

  nums := []int{1, 3, 5, 7, 9}
  for i, val := range nums {
    fmt.Printf("nums[%d]=%d val has %d\n", i, nums[i], val)
  }

  m := map[string]Vertex{
    "Berkeley": {37.7, -122.2},
    "0": {0.0, 0.0},
  }
  describe(m)

  _, ok1 := m["Berkeley"]
  _, ok2 := m["SF"]
  fmt.Printf("%v val=%v\n", m, m["Berkeley"])
  fmt.Printf("exist-checks %v %v\n", ok1, ok2)
  fmt.Printf("dist %f\n", m["Berkeley"].dist(m["0"]))

}
