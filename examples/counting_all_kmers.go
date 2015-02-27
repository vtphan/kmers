// example how to count k-mer frequencies in a set of reads, using multiple goroutines.

package main

import (
   "github.com/vtphan/kmers"
   "os"
   "bufio"
   "fmt"
   "runtime"
   "sync"
)


func CountFreq(readFile string, K int) {

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
   numCores := runtime.NumCPU()
   runtime.GOMAXPROCS(numCores)
   var wg sync.WaitGroup

   // Define a new counter that count *all* kmers.
   c := kmers.NewCounter(K)

   // Start counting all k-mers in reads, using different cores simultaneously.
   for i:=0; i<numCores; i++ {
      wg.Add(1)
      go func() {
         defer wg.Done()
         for read := range(reads){
            c.Count2([]byte(read))
         }
      }()
   }

   // Finish counting
   wg.Wait()


   // Print out the results
   for i:=0; i<len(c.Freq); i++ {
      if c.Freq[i] > 0 {
         fmt.Println(i,kmers.NumToKmer(i,K),c.Freq[i])
      }
   }
}


func main() {
   CountFreq(os.Args[1], 4)
}