package ulog

import (
	"io"
	"sync"
	"time"
)

const (

	// 缓冲队列长度
	asyncBufferSize = 100

	// 开启内存池范围的写入大小
	maxTextBytes = 1024
)

// 异步写入器, 此输出器在有多级输出时, 应作为中间缓冲用
// 例如: AsyncOutput -> RollingOutput 异步输出到滚动文件
// 		AsyncOutput -> NetworkOutput 异步输出到网络
type AsyncOutput struct {
	output io.Writer

	queue     chan interface{}
	bytesPool *sync.Pool
}

func (self *AsyncOutput) writeLoop() {
	for {

		// 从队列中获取一个要写入的日志
		raw := <-self.queue

		switch d := raw.(type) {
		case func(): // Flush操作
			d()
		case []byte:
			// 写入目标
			self.output.Write(d)

			// 必须是由pool分配的，才能用池释放
			if cap(d) < maxTextBytes {
				self.bytesPool.Put(d)
			}
		}

	}
}

type SyncWrite interface {
	Sync() error
}

// 保证在调用本函数结束时, 之前的内容已经完全写入
func (self *AsyncOutput) Flush(timeout time.Duration) {
	ch := make(chan struct{})

	self.queue <- func() {

		// 确保文件已经写入, Output需要实现这个接口
		if f, ok := self.output.(SyncWrite); ok {
			f.Sync()
		}

		ch <- struct{}{}
	}

	select {
	case <-ch:
	case <-time.After(timeout):
	}
}

func (self *AsyncOutput) Write(b []byte) (n int, err error) {
	var newb []byte

	// 超大时, 直接分配
	if len(b) >= maxTextBytes {
		newb = make([]byte, len(b))
	} else { // 从池中分配
		newb = self.bytesPool.Get().([]byte)[:len(b)]
	}

	copy(newb, b)
	self.queue <- newb

	return len(b), nil
}

// 异步写入的目标
func NewAsyncOutput(output io.Writer) *AsyncOutput {

	self := &AsyncOutput{
		output:    output,
		queue:     make(chan interface{}, asyncBufferSize),
		bytesPool: new(sync.Pool),
	}

	// 拷贝数据的池
	self.bytesPool.New = func() interface{} {
		return make([]byte, maxTextBytes)
	}

	go self.writeLoop()

	return self
}
