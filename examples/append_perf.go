package main

import (
   "fmt"
)

func main() {
   N := 200000000
   a := make([]int, N)
   for i:=0; i<N; i++ {
      a[i] = i
   }

   // var b []int
   // for i:=0; i<N; i++ {
   //    b = append(b,i)
   // }
   fmt.Println("done")
}