// example how to count k-mer frequencies in a set of reads, using multiple goroutines.

package main

import (
   "github.com/vtphan/kmers"
   "math"
   "os"
   "sync"
   "bufio"
   "fmt"
   "runtime"
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

   // frequncies and locks tables.  Each entry in freq has a lock in lock.
   freq := make([]int, int(math.Pow(4,float64(K))))
   lock := make([]sync.RWMutex, len(freq))

   for i:=0; i<numCores; i++ {
      wg.Add(1)
      go func() {
         defer wg.Done()
         for read := range(reads){
            kmers.Count([]byte(read), K, 0, len(read), freq, lock)
         }
      }()
   }

   wg.Wait()

   for i:=0; i<len(freq); i++ {
      fmt.Println(i,kmers.NumToKmer(i,K),freq[i])
   }
}


func main() {
   CountFreq(os.Args[1], 4)
}