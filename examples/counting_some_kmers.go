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

   numCores := runtime.NumCPU()
   runtime.GOMAXPROCS(numCores)
   var wg sync.WaitGroup

   // Here is a map of kmers that need to be counted.  Only these kmers are counted.
   freq := make(map[int]int)
   freq[158] = 0
   freq[180] = 0
   freq[39] = 0
   freq[59] = 0

   // Start a new counter that counts only kmers in freq.
   c := kmers.NewKmerCounter(K, freq)

   // Count kmers in different cores simultaneously.
   for i:=0; i<numCores; i++ {
      wg.Add(1)
      go func() {
         defer wg.Done()
         for read := range(reads){
            c.Count([]byte(read))
         }
      }()
   }

   // Finish counting
   wg.Wait()

   // Print out the result
   for kmer := range(freq) {
      fmt.Println(kmer,kmers.NumToKmer(kmer,K),freq[kmer])
   }
}


func main() {
   CountFreq(os.Args[1], 4)
}