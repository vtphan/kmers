/*
Author: Vinhthuy Phan, 2014
Get all kmers of a sequence
*/
package kmers

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
Store in channel "result" all kmers (represented by numbers) of "sequence" by
sliding a window of length K from "start" to "end"-1.

h := map[byte]int{'A':0, 'C':1, 'G':2, 'T':3}
id = (id<<2) - (h[sequence[i-1]] << (uint(2*K))) + h[sequence[i+K-1]]
*/
func Slide(sequence []byte, K int, start int, end int, result chan int) {
   uK := uint(K)
   var id int
   for i:=start; i<end-K+1; i++ {
      if i==start || sequence[i+K-1]=='N' {
         i = firstIndex(sequence, K, i, end)
         if i==-1 {
            return
         }
         id = Kmer2Num(sequence, K, i)
      } else {
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
      }
      result <- id
   }
}

