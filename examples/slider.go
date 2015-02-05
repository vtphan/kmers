// example of how to slide a K-window across a read
// and get all kmers and store them into a channel

package main

import (
   "bufio"
   "os"
   "sync"
   "fmt"
   "github.com/vtphan/kmers"
)

func FreqFromReads(readFile string) {
   // Get all reads into a channel
   reads := make(chan string)
   go func() {
      f, err := os.Open(readFile)
      if err != nil {
         panic("Error opening " + readFile)
      }
      defer f.Close()
      scanner := bufio.NewScanner(f)
      for scanner.Scan() {
         reads <- scanner.Text()
      }
      close(reads)
   }()

   // Spread the processing of reads into different cores
   result := make(chan int)
   numCores := 4
   K := 5
   var wg sync.WaitGroup
   for i:=0; i<numCores; i++ {
      wg.Add(1)
      go func() {
         defer wg.Done()
         for read := range(reads){
            kmers.Slide([]byte(read), K, 0, len(read), result)
         }
      }()
   }

   go func() {
      wg.Wait()
      close(result)
   }()
   for r := range(result) {
      fmt.Println(r, kmers.NumToKmer(r,K))
   }
}

func main() {
   FreqFromReads("reads1.txt")
}