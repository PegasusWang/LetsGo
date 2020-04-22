# Go goroutine 初探

!!! quote
    Before you launch a goroutine, know when it will stop. - https://the-zen-of-go.netlify.com/

## CPython 之痛 GIL

如果你用过 Python 并且使用过 CPython 解释器，应该知道”臭名昭著“的 GIL，全局解释锁。这导致 Python 在多线程下没法利用多核
CPU，对于 IO 密集程序来说可能影响还不大，但是对于 CPU 密集的程序性能可能还不如单线程。当遇到性能瓶颈的时候，你可以考虑
使用多进程或者换成 Go 来改写获得更高的性能。

## Go 杀手锏 Goroutine


## 启动一个 Goroutine


## 编写一个并发请求网址例子
