## Usage

#### func  Slide1

```go
func Slide1(sequence []byte, K int, start int, end int, result chan int)
```
Store in channel "result" all kmers of "sequence" by
sliding a window of length K from "start" to "end"-1.

#### func  Slide2

```go
func Slide2(sequence []byte, K int, start int, end int, result chan int)
```
Store in channel "result" all kmers of "sequence" and its reverse complement by sliding a window of length K from "start" to "end"-1.

#### type Counter

```go
type Counter struct {
	K    int
	Freq []int
}
```

Counter is used to counter ALL kmers.

#### func  NewCounter

```go
func NewCounter(K int) *Counter
```

#### func (*Counter) Count1

```go
func (c *Counter) Count1(sequence []byte)
```
counting all kmers in the main strand

#### func (*Counter) Count2

```go
func (c *Counter) Count2(sequence []byte)
```
counting all kmers on both strands

#### type KmerCounter

```go
type KmerCounter struct {
	K    int
	Freq map[int]int
}
```

KmerCounter is used to counter only kmers in Freq in reads.

#### func  NewKmerCounter

```go
func NewKmerCounter(K int, freq map[int]int) *KmerCounter
```

#### func (*KmerCounter) Count1

```go
func (c *KmerCounter) Count1(sequence []byte)
```
count number of k-mers in the main strand

#### func (*KmerCounter) Count2

```go
func (c *KmerCounter) Count2(sequence []byte)
```
count number of k-mers in on both strands
