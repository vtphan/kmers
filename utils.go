/*
Author: Vinhthuy Phan, 2015
*/
package kmers

/*
   compute the (base-4) numeric representations of the sequence
   and its reverse complement
*/
func Kmer2Num(sequence []byte, K int, i int) (int, int) {
   id := 0
   id_rc := 0
   for j:=i; j<K+i; j++ {
      switch sequence[j] {
         case 'A': id = id<<2
         case 'C': id = id<<2 + 1
         case 'G': id = id<<2 + 2
         case 'T': id = id<<2 + 3
         default:
            return -1, -1
      }
      switch sequence[K+i-1+i-j] {
         case 'T': id_rc = id_rc<<2
         case 'G': id_rc = id_rc<<2 + 1
         case 'C': id_rc = id_rc<<2 + 2
         case 'A': id_rc = id_rc<<2 + 3
         default:
            return -1, -1
      }
   }
   return id, id_rc
}

/*
   The first index, starting from "start" forwardly, of a kmer consisting of
   only A,C,T,G.
*/
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
   Return the K-mer (consisting of A,C,G,T) represented by x
*/
func NumToKmer(x int, K int) string {
   y := make([]byte, K)
   for i:=K-1; i>=0; i-- {
      base := x%4
      switch base {
         case 0: y[i] = 'A'
         case 1: y[i] = 'C'
         case 2: y[i] = 'G'
         case 3: y[i] = 'T'
      }
      x = (x-base)>>2
   }
   return string(y)
}
