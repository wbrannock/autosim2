package main

import (
  "fmt"
  "os"
  "math/rand"
  "time"
)

var garray [12]float64

func input(){
  f, err := os.Open("test.txt")
   if err != nil {
      fmt.Println(err)
   }
   var is int = 0
   for {
      is++
      var flt float64
      var n int
      n, err = fmt.Fscanln(f, &flt)
      if n == 0 || err != nil {
        break
      }
      //fmt.Println("Probability:", flt)
      garray[is- 1] = flt
   }
}
var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func uniform() float64 {
  return r1.Float64()
}

func simulate(sims int) {
  var reps = sims
  var wins = 0
  var bowl = 0
  var tenWins = 0
  var undefeated = 0
  for i := 1; i <= reps; i++ {
    var internalWins = 0
    for z := 0; z < 12; z++ {
      if uniform() < garray[z]{
        wins++
        internalWins++
      }

    }
   
    //Event Checker
    if internalWins >= 6 {
      bowl++
    }
    if internalWins >= 10 {
      tenWins++
    }
    if internalWins == 12 {
      undefeated++
    }

  }
  fmt.Printf("Season Simulated %d times\n", reps)
  var outNumber = float64(wins) / float64(reps)
  var w = wins / reps
  var l = 12 - w
  fmt.Printf("Median Record: ")
  fmt.Printf("%d-%d\n", w, l)
  fmt.Printf("Average Games Won Per Season: ")
  fmt.Printf("%f\n", outNumber)
  fmt.Printf("Bowl Chance: ")
  var bowlPercent = (float64(bowl) / float64(reps)) * 100
  fmt.Printf("%f\n", bowlPercent)
  fmt.Printf("10 Wins Chance: ")
  var tenWinsPercent = (float64(tenWins) / float64(reps)) * 100
  fmt.Printf("%f\n", tenWinsPercent)
  fmt.Printf("Undefeated Chance: ")
  var undefeatedPercent = (float64(undefeated) / float64(reps)) * 100
  fmt.Printf("%f\n", undefeatedPercent)

}
func main() {
   input()
   simulate(10000000)
}