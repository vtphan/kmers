
package kmers

import (
   "testing"
   "fmt"
   "sync"
)


func Test1(t *testing.T){
   seq := []byte("CCAAGGT")
   var K int = 3

   var wg sync.WaitGroup
   result := make(chan int)
   wg.Add(1)
   go func() {
      wg.Wait()
      close(result)
   }()
   go func() {
      defer wg.Done()
      Kmers(seq,K,0,len(seq),result)
   }()
   for r := range(result) {
      fmt.Println(r)
   }
   fmt.Println("Done Test1")
}

func Test2(t *testing.T){
   seq := []byte("CGNTCAG")
   var K int = 2

   var wg sync.WaitGroup
   result := make(chan int)
   wg.Add(1)
   go func() {
      wg.Wait()
      close(result)
   }()
   go func() {
      defer wg.Done()
      Kmers(seq,K,0,len(seq),result)
   }()
   for r := range(result) {
      fmt.Println(r)
   }
   fmt.Println("Done Test2")
}