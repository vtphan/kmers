package kmers

/*
Store in channel "result" all kmers (represented by numbers) of "sequence" and its
reverse complement by sliding a window of length K from "start" to "end"-1.
*/
func Slide2(sequence []byte, K int, start int, end int, result chan int) {
   uK := uint(K)
   var id int
   var id_rc int
   for i:=start; i<end-K+1; i++ {
      if i==start || (sequence[i+K-1]!='A' && sequence[i+K-1]!='C' && sequence[i+K-1]!='G' && sequence[i+K-1]!='T') {
         i = firstIndex(sequence, K, i, end)
         if i==-1 {
            return
         }
         id, id_rc = Kmer2Num(sequence, K, i)
      } else {
         switch sequence[i-1] {
            case 'A':
               id = (id<<2)
               id_rc = (id_rc - 3) >> 2
            case 'C':
               id = (id<<2) - (1 << (uK<<1))
               id_rc = (id_rc - 2) >> 2
            case 'G':
               id = (id<<2) - (2 << (uK<<1))
               id_rc = (id_rc - 1) >> 2
            case 'T':
               id = (id<<2) - (3 << (uK<<1))
               id_rc = id_rc >> 2
            default:
               panic("Slide: invalid base " + string(sequence[i-1]))
         }
         switch sequence[i+K-1] {
            case 'A':
               id_rc += 3 << ((uK-1)<<1)
            case 'C':
               id += 1
               id_rc += 2 << ((uK-1)<<1)
            case 'G':
               id += 2
               id_rc += 1 << ((uK-1)<<1)
            case 'T':
               id += 3
            default:
               panic("Slide: invalid base " + string(sequence[i-1]))
         }
      }
      result <- id
      result <- id_rc
   }
}


/*
Store in channel "result" all kmers (represented by numbers) of "sequence"
by sliding a window of length K from "start" to "end"-1.
*/
func Slide1(sequence []byte, K int, start int, end int, result chan int) {
   uK := uint(K)
   var id int
   for i:=start; i<end-K+1; i++ {
      if i==start || (sequence[i+K-1]!='A' && sequence[i+K-1]!='C' && sequence[i+K-1]!='G' && sequence[i+K-1]!='T') {
         i = firstIndex(sequence, K, i, end)
         if i==-1 {
            return
         }
         id, _ = Kmer2Num(sequence, K, i)
      } else {
         switch sequence[i-1] {
            case 'A':
               id = (id<<2)
            case 'C':
               id = (id<<2) - (1 << (uK<<1))
            case 'G':
               id = (id<<2) - (2 << (uK<<1))
            case 'T':
               id = (id<<2) - (3 << (uK<<1))
            default:
               panic("Slide: invalid base " + string(sequence[i-1]))
         }
         switch sequence[i+K-1] {
            case 'A':
            case 'C':
               id += 1
            case 'G':
               id += 2
            case 'T':
               id += 3
            default:
               panic("Slide: invalid base " + string(sequence[i-1]))
         }
      }
      result <- id
   }
}
