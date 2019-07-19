#channel

---

##缓冲channel
```go
ch := make(chan chanType，n)
n为缓冲区大小
chanType为channel的类型

假设ch := make(chan int)
那么就相当于定义了一个0缓冲区的channel
为非阻塞式的channel

当缓冲区满的时候channel会发生阻塞，除非有程序对channel进行
读操作才不会进行阻塞
读/写操作：
                        ch <- 写入的信息
       接受的对象，不写则丢弃 <- ch
```
---

#select

>linux 很早引入的函数，用来实现非阻塞的一种方式
```go
select {
    case <-chan1:
    	//如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
    	//如果成功向chan2写入数据，则进行该case处理语句
    default:
    	//假设上面都没有成功，则进入default处理流程

}
```

###协程

**协程** 
    协程业务代码会一直运行
>除非主动让出资源

>或者遇到了IO阻塞

runtime包下的Gosched()方法--主动要求让出cpu
```go
//主动要求让出cpu
runtime.Gosched()
```
