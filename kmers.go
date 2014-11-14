/*
Author: Vinhthuy Phan, 2014
Get all kmers of a sequence
*/
package main

import (
   "fmt"
   "sync"
)

func Kmer2Num(sequence []byte, K int, i int) int {
   id := 0
   for j:=i; j<K+i; j++ {
      switch sequence[j] {
         case 'A': id = id<<2
         case 'C': id = id<<2 + 1
         case 'G': id = id<<2 + 2
         case 'T': id = id<<2 + 3
         default:
            return -1
      }
   }
   return id
}

func firstIndex(sequence []byte, K int, start int, end int) int {
   var j int
   for i:=start; i<end-K+1; i++ {
      for j=i; j<i+K; j++ {
         if sequence[j]!='A' && sequence[j]!='C' && sequence[j]!='G' && sequence[j]!='T' {
            break
         }
      }
      if j==i+K {
         return i
      }
   }
   return -1
}


/*
This function gets all kmers of "sequence" from index "start" to index "end"-1.
The kmers are represented by numbers, stored in "result".

enumerate all k-mers of sequence and compute their ids
h := map[byte]int{'A':0, 'C':1, 'G':2, 'T':3}
id = (id<<2) - (h[sequence[i-1]] << (uint(2*K))) + h[sequence[i+K-1]]
*/
func Kmers(sequence []byte, K int, start int, end int, result chan int) {
   i := firstIndex(sequence, K, start, end)
   if i == -1 {
      return
   }
   uK := uint(K)
   id := Kmer2Num(sequence, K, i)
   result <- id
   for i:=i+1; i<end-K+1; i++ {
      if sequence[i+K-1]=='A' || sequence[i+K-1]=='C' || sequence[i+K-1]=='G' || sequence[i+K-1]=='T' {
         switch sequence[i-1] {
            case 'A': id = (id<<2)
            case 'C': id = (id<<2) - (1 << (uK<<1))
            case 'G': id = (id<<2) - (2 << (uK<<1))
            case 'T': id = (id<<2) - (3 << (uK<<1))
         }
         switch sequence[i+K-1] {
            case 'A':
            case 'C': id += 1
            case 'G': id += 2
            case 'T': id += 3
         }
         result <- id
      } else {
         i = firstIndex(sequence, K, i, end)
         if i==-1 {
            return
         }
         id = Kmer2Num(sequence, K, i)
         result <- id
      }
   }
}

func main(){
   seq := []byte("CCAAGNGT")
   var K int = 3

   var wg sync.WaitGroup
   result := make(chan int)

   wg.Add(1)
   go func() {
      defer wg.Done()
      Kmers(seq,K,0,len(seq),result)
   }()

   go func() {
      wg.Wait()
      close(result)
   }()


   for r := range(result) {
      fmt.Println(r)
   }
}
