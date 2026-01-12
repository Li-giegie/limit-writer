# limit-writer
Writer 实现了限制向底层Writer写入一次数据的大小，
与bufio.Writer最大的区别是bufio.Writer缓冲区没有剩余空间时，向底层Writer写入数据，
limit-writer是待写入p字节数 > 缓冲区剩余容量，先将缓冲区数据写入底层Writer，再写入p到缓存区。

## 开始
```go
go get github.com/Li-giegie/limit-writer
```
## 概述
1. 当待写入p字节数 > 缓冲区总容量，返回ErrOverflow；
2. 当待写入p字节数 > 缓冲区剩余容量，先将缓冲区数据写入底层Writer，再写入p到缓存区
3. 当待写入p字节数 <= 缓冲区剩余容量，copy到缓冲区中
